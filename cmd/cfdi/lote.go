package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/OrtaMarco/cfdi-go"
)

// trabajo es un CFDI por leer: o una ruta en disco (la lee el worker) o los
// bytes ya extraídos de un zip. El resultado vuelve por res.
type trabajo struct {
	origen string
	ruta   string
	datos  []byte
	err    error // error al obtener los bytes (zip dañado, permisos…)
	res    chan resultado
}

// resultado es lo que queda de un CFDI al salir del worker: la línea JSONL ya
// serializada y lo poco que pide el resumen. El Comprobante no sale del
// worker, así que un resultado en espera pesa lo que su línea.
type resultado struct {
	origen       string
	linea        []byte
	errLinea     error // no se pudo serializar
	err          error // error fatal de lectura
	version      string
	tipo         string
	complementos []string
	advertencias []string // el código de cada advertencia
	huella       string   // ejemplo anónimo para el resumen; solo si hubo advertencias, hallazgos o error
	sello        *cfdi.ResultadoSello
	tfd          *cfdi.ResultadoSello
	validacion   string   // estado de la validación; "" si no se pidió
	hallazgos    []string // el código de cada hallazgo de la validación
	retenciones  bool     // es un CFDI de Retenciones: version y complementos van aparte en el resumen
}

// Conteo de una clase (código de advertencia o clase de error fatal).
type Conteo struct {
	CFDI        int      `json:"cfdi"`                  // documentos con al menos una
	Ocurrencias int      `json:"ocurrencias,omitempty"` // advertencias en total (no aplica a fatales)
	Ejemplos    []string `json:"ejemplos"`              // hasta maxEjemplos huellas (u orígenes con -rutas)
}

// Resumen del lote, pensado para compartirse: solo conteos y, como ejemplos,
// huellas anónimas (ver huella). Nada de RFC, nombres, montos ni rutas, salvo
// que se pidan las rutas con -rutas.
type Resumen struct {
	Total                 int                `json:"total"`
	Leidos                int                `json:"leidos"`
	Fatales               int                `json:"fatales"`
	Advertidos            int                `json:"advertidos"`
	AdvertenciasPorCodigo map[string]*Conteo `json:"advertencias_por_codigo"`
	FatalesPorClase       map[string]*Conteo `json:"fatales_por_clase"`
	PorVersion            map[string]int     `json:"por_version"`
	PorTipo               map[string]int     `json:"por_tipo"`
	PorComplemento        map[string]int     `json:"por_complemento"`
	// Sellos solo está si se verificaron (-sello o -verificar).
	Sellos *ConteoSellos `json:"sellos,omitempty"`
	// SellosTFD, el SelloSAT de los timbres, solo si se pidió -certs-sat.
	SellosTFD *ConteoSellos `json:"sellos_tfd,omitempty"`
	// Validacion solo está si se pidió -validar.
	Validacion *ConteoValidacion `json:"validacion,omitempty"`
	// Retenciones solo está si el lote trae CFDI de Retenciones. Los de
	// arriba (leidos, advertidos, advertencias, sellos, validación) los
	// cuentan junto con los CFDI; por_version, por_tipo y por_complemento
	// son solo de CFDI, y los de Retenciones van aquí.
	Retenciones *ResumenRetenciones `json:"retenciones,omitempty"`

	rutas bool // ejemplos con el origen completo en vez de la huella
}

// ResumenRetenciones cuenta los CFDI de Retenciones leídos, por versión
// ("1.0", "2.0") y por complemento ("ServiciosPlataformasTecnologicas 1.0").
type ResumenRetenciones struct {
	Leidos         int            `json:"leidos"`
	PorVersion     map[string]int `json:"por_version"`
	PorComplemento map[string]int `json:"por_complemento"`
}

// ConteoSellos cuenta los CFDI leídos por resultado de la verificación del
// sello del emisor; los inválidos y no verificables, por motivo.
type ConteoSellos struct {
	Valido        int            `json:"valido"`
	Invalido      map[string]int `json:"invalido"`
	NoVerificable map[string]int `json:"no_verificable"`
}

// ConteoValidacion cuenta los CFDI leídos por estado de la validación
// (valido, invalido, incompleto) y los hallazgos por código, como las
// advertencias.
type ConteoValidacion struct {
	PorEstado          map[string]int     `json:"por_estado"`
	HallazgosPorCodigo map[string]*Conteo `json:"hallazgos_por_codigo"`
}

func (c *ConteoValidacion) String() string {
	return fmt.Sprintf("validación: %d válidos, %d inválidos, %d incompletos",
		c.PorEstado[cfdi.ValidacionValida], c.PorEstado[cfdi.ValidacionInvalida], c.PorEstado[cfdi.ValidacionIncompleta])
}

func nuevoConteoSellos() *ConteoSellos {
	return &ConteoSellos{Invalido: map[string]int{}, NoVerificable: map[string]int{}}
}

func (c *ConteoSellos) agregar(r *cfdi.ResultadoSello) {
	switch r.Estado {
	case cfdi.SelloValido:
		c.Valido++
	case cfdi.SelloInvalido:
		c.Invalido[r.Motivo]++
	default:
		c.NoVerificable[r.Motivo]++
	}
}

func (c *ConteoSellos) String() string {
	suma := func(m map[string]int) (n int) {
		for _, v := range m {
			n += v
		}
		return n
	}
	return fmt.Sprintf("sellos: %d válidos, %d inválidos, %d no verificables", c.Valido, suma(c.Invalido), suma(c.NoVerificable))
}

const maxEjemplos = 5

func nuevoResumen(rutas bool, o opcionesLote) *Resumen {
	r := &Resumen{
		rutas:                 rutas,
		AdvertenciasPorCodigo: map[string]*Conteo{},
		FatalesPorClase:       map[string]*Conteo{},
		PorVersion:            map[string]int{},
		PorTipo:               map[string]int{},
		PorComplemento:        map[string]int{},
	}
	if o.verificar {
		r.Sellos = nuevoConteoSellos()
	}
	if o.certs != nil {
		r.SellosTFD = nuevoConteoSellos()
	}
	if o.validar {
		r.Validacion = &ConteoValidacion{PorEstado: map[string]int{}, HallazgosPorCodigo: map[string]*Conteo{}}
	}
	return r
}

func contar(m map[string]*Conteo, clave, origen string, ocurrencias int) {
	c := m[clave]
	if c == nil {
		c = &Conteo{Ejemplos: []string{}}
		m[clave] = c
	}
	c.CFDI++
	c.Ocurrencias += ocurrencias
	if len(c.Ejemplos) < maxEjemplos {
		c.Ejemplos = append(c.Ejemplos, origen)
	}
}

func valorOVacio(s string) string {
	if s == "" {
		return "(vacío)"
	}
	return s
}

// huella identifica un documento sin decir de dónde salió: los primeros 12
// hex del sha256 de su contenido (del origen si no se pudo leer) y la
// extensión del archivo de entrada o del contenedor más externo
// ("3f2a9c1b4d5e.zip"). Con el sha256 del contenido no se puede confirmar si
// una ruta adivinada (con un RFC, p. ej.) está en el lote; para ubicar el
// archivo, `sha256sum` o volver a correr con -rutas.
func huella(datos []byte, origen string) string {
	if datos == nil {
		datos = []byte(origen)
	}
	h := sha256.Sum256(datos)
	contenedor, _, _ := strings.Cut(origen, "!")
	return hex.EncodeToString(h[:6]) + extension(contenedor)
}

func (r *Resumen) agregar(res resultado) {
	r.Total++
	if r.rutas {
		res.huella = res.origen
	}
	if res.err != nil {
		r.Fatales++
		clase := cfdi.ClaseDe(res.err)
		if clase == "" {
			clase = cfdi.ClaseLectura
		}
		contar(r.FatalesPorClase, clase, res.huella, 0)
		return
	}
	r.Leidos++
	if r.Sellos != nil && res.sello != nil {
		r.Sellos.agregar(res.sello)
	}
	if r.SellosTFD != nil && res.tfd != nil {
		r.SellosTFD.agregar(res.tfd)
	}
	if r.Validacion != nil && res.validacion != "" {
		r.Validacion.PorEstado[res.validacion]++
		contarCodigos(r.Validacion.HallazgosPorCodigo, res.hallazgos, res.huella)
	}
	if res.retenciones {
		if r.Retenciones == nil {
			r.Retenciones = &ResumenRetenciones{PorVersion: map[string]int{}, PorComplemento: map[string]int{}}
		}
		r.Retenciones.Leidos++
		r.Retenciones.PorVersion[valorOVacio(res.version)]++
		for _, comp := range res.complementos {
			r.Retenciones.PorComplemento[comp]++
		}
	} else {
		r.PorVersion[valorOVacio(res.version)]++
		r.PorTipo[valorOVacio(res.tipo)]++
		for _, comp := range res.complementos {
			r.PorComplemento[comp]++
		}
	}
	if len(res.advertencias) == 0 {
		return
	}
	r.Advertidos++
	contarCodigos(r.AdvertenciasPorCodigo, res.advertencias, res.huella)
}

// contarCodigos suma un documento con esos códigos (con repeticiones) a m.
func contarCodigos(m map[string]*Conteo, cods []string, huella string) {
	porCodigo := map[string]int{}
	for _, a := range cods {
		porCodigo[a]++
	}
	codigos := make([]string, 0, len(porCodigo))
	for k := range porCodigo {
		codigos = append(codigos, k)
	}
	sort.Strings(codigos)
	for _, k := range codigos {
		contar(m, k, huella, porCodigo[k])
	}
}

type lineaError struct {
	Origen string `json:"origen"`
	Error  struct {
		Clase   string `json:"clase"`
		Mensaje string `json:"mensaje"`
	} `json:"error"`
}

func cmdLote(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("lote", flag.ContinueOnError)
	fs.SetOutput(stderr)
	salida := fs.String("o", "", "archivo JSONL de salida (por omisión stdout)")
	resumenPath := fs.String("resumen", "", "archivo JSON con el resumen")
	workers := fs.Int("j", runtime.GOMAXPROCS(0), "workers concurrentes")
	sello := fs.Bool("sello", false, "incluir sellos, certificados y cadenas, y verificar el sello del emisor")
	verificar := fs.Bool("verificar", false, "verificar el sello del emisor sin incluir los sellos")
	validar := fs.Bool("validar", false, "validar contra los XSD del SAT y las reglas aritméticas")
	certsDir := fs.String("certs-sat", "", "directorio con los certificados (.cer) del SAT para verificar el timbre")
	conRutas := fs.Bool("rutas", false, "ejemplos del resumen con la ruta completa en vez de la huella")
	rutas, err := parsearIntercalado(fs, args)
	if err != nil {
		return 2
	}
	if len(rutas) == 0 {
		fmt.Fprint(stderr, "uso: "+usoLote+"\n")
		return 2
	}
	if *workers < 1 {
		*workers = 1
	}

	out := stdout
	if *salida != "" {
		f, err := os.Create(*salida)
		if err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
		defer f.Close()
		out = f
	}
	w := bufio.NewWriter(out)

	var erroresArgs []string
	producir := func(emitir func(trabajo)) {
		for _, r := range rutas {
			if err := recorrer(r, emitir); err != nil {
				erroresArgs = append(erroresArgs, err.Error())
			}
		}
	}
	opciones := opcionesLote{json: cfdi.OpcionesJSON{IncluirSello: *sello}, verificar: *sello || *verificar, validar: *validar}
	if *certsDir != "" {
		if opciones.certs, err = cargarCertificados(*certsDir, stderr); err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
	}
	res := nuevoResumen(*conRutas, opciones)
	var errEscritura error
	canalizar(producir, *workers, 256**workers,
		func(t trabajo) resultado { return procesar(t, opciones) },
		func(r resultado) {
			res.agregar(r)
			if errEscritura == nil {
				errEscritura = r.errLinea
			}
			if errEscritura == nil {
				_, errEscritura = w.Write(r.linea)
			}
		})
	if errEscritura == nil {
		errEscritura = w.Flush()
	}

	codigo := 0
	for _, e := range erroresArgs {
		fmt.Fprintln(stderr, e)
		codigo = 1
	}
	if errEscritura != nil {
		fmt.Fprintln(stderr, "escribiendo la salida:", errEscritura)
		codigo = 1
	}
	if *resumenPath != "" {
		b, _ := json.MarshalIndent(res, "", "  ")
		if err := os.WriteFile(*resumenPath, append(b, '\n'), 0o644); err != nil {
			fmt.Fprintln(stderr, err)
			codigo = 1
		}
	}
	fmt.Fprintf(stderr, "%d documentos: %d leídos (%d con advertencias), %d fatales\n",
		res.Total, res.Leidos, res.Advertidos, res.Fatales)
	if res.Retenciones != nil {
		fmt.Fprintf(stderr, "de los leídos, %d son CFDI de Retenciones\n", res.Retenciones.Leidos)
	}
	if res.Sellos != nil {
		fmt.Fprintln(stderr, res.Sellos)
	}
	if res.SellosTFD != nil {
		fmt.Fprintln(stderr, "timbres: "+strings.TrimPrefix(res.SellosTFD.String(), "sellos: "))
	}
	if res.Validacion != nil {
		fmt.Fprintln(stderr, res.Validacion)
	}
	return codigo
}

// canalizar reparte los trabajos entre workers y entrega los resultados a
// consumir en el orden en que producir los emitió, no en el que terminan: la
// salida es reproducible.
//
// La memoria no crece con el número de archivos: producir encola el canal de
// respuesta de cada trabajo en orden, que tiene capacidad ventana, y consumir
// los espera uno por uno. Si un documento lento retiene la cabeza, producir se
// detiene en vez de acumular resultados detrás de él; a lo más hay ventana+2
// trabajos emitidos sin consumir.
func canalizar(producir func(emitir func(trabajo)), workers, ventana int, procesar func(trabajo) resultado, consumir func(resultado)) {
	trabajos := make(chan trabajo, workers)
	orden := make(chan chan resultado, ventana)
	go func() {
		defer close(trabajos)
		defer close(orden)
		producir(func(t trabajo) {
			t.res = make(chan resultado, 1)
			orden <- t.res
			trabajos <- t
		})
	}()
	for range workers {
		go func() {
			for t := range trabajos {
				t.res <- procesar(t)
			}
		}()
	}
	for res := range orden {
		consumir(<-res)
	}
}

type opcionesLote struct {
	json      cfdi.OpcionesJSON
	verificar bool                 // verificar el sello del emisor
	certs     cfdi.CertificadosSAT // verificar el timbre (nil = no)
	validar   bool                 // ValidarEsquema y ValidarReglas
}

// procesar lee el CFDI, arma su línea JSONL y se queda solo con lo que pide
// el resumen.
func procesar(t trabajo, o opcionesLote) resultado {
	r := resultado{origen: t.origen, err: t.err}
	if r.err == nil && t.ruta != "" {
		t.datos, r.err = os.ReadFile(t.ruta)
	}
	if r.err != nil {
		r.err = &cfdi.Error{Clase: cfdi.ClaseLectura, Err: r.err}
	}
	if r.err == nil && cfdi.Detectar(t.datos) == cfdi.TipoRetenciones {
		return procesarRetenciones(t, r, o)
	}
	var c *cfdi.Comprobante
	if r.err == nil {
		c, r.err = cfdi.ParseBytes(t.datos)
	}
	if r.err != nil {
		return conError(t, r)
	}
	d := c.Documento(o.json)
	d.Origen = r.origen
	if o.verificar || o.certs != nil {
		v := cfdi.Verificar(t.datos, o.certs)
		if o.verificar {
			d.VerificacionSello, r.sello = &v.Sello, &v.Sello
		}
		d.VerificacionTFD, r.tfd = v.TFD, v.TFD
	}
	if o.validar {
		v := cfdi.Validar(t.datos, c)
		d.Validacion, r.validacion = v, v.Estado
		for _, h := range v.Hallazgos() {
			r.hallazgos = append(r.hallazgos, h.Codigo)
		}
	}
	r.linea, r.errLinea = lineaJSON(d)
	r.version, r.tipo, r.complementos = c.Version, c.TipoDeComprobante, c.Complementos()
	for _, a := range c.Advertencias {
		r.advertencias = append(r.advertencias, a.Codigo)
	}
	if len(r.advertencias) > 0 || len(r.hallazgos) > 0 {
		r.huella = huella(t.datos, t.origen)
	}
	return r
}

// conError arma la línea de un documento que no se pudo leer.
func conError(t trabajo, r resultado) resultado {
	r.huella = huella(t.datos, t.origen)
	var l lineaError
	l.Origen = r.origen
	l.Error.Clase = cfdi.ClaseDe(r.err)
	l.Error.Mensaje = r.err.Error()
	r.linea, r.errLinea = lineaJSON(l)
	return r
}

// procesarRetenciones es procesar para un CFDI de Retenciones: su línea
// lleva "tipo_documento":"retenciones" y la clave "retenciones" en vez de
// "cfdi". -validar es el esquema y las reglas de Retenciones.
func procesarRetenciones(t trabajo, r resultado, o opcionesLote) resultado {
	rt, err := cfdi.ParseRetencionesBytes(t.datos)
	if err != nil {
		r.err = err
		return conError(t, r)
	}
	r.retenciones = true
	d := rt.Documento(o.json)
	d.Origen = r.origen
	if o.verificar || o.certs != nil {
		v := cfdi.VerificarRetenciones(t.datos, o.certs)
		if o.verificar {
			d.VerificacionSello, r.sello = &v.Sello, &v.Sello
		}
		d.VerificacionTFD, r.tfd = v.TFD, v.TFD
	}
	if o.validar {
		v := cfdi.ValidarRetenciones(t.datos)
		d.Validacion, r.validacion = v, v.Estado
		for _, h := range v.Hallazgos() {
			r.hallazgos = append(r.hallazgos, h.Codigo)
		}
	}
	r.linea, r.errLinea = lineaJSON(d)
	r.version, r.complementos = rt.Version, rt.Complementos()
	for _, a := range rt.Advertencias {
		r.advertencias = append(r.advertencias, a.Codigo)
	}
	if len(r.advertencias) > 0 || len(r.hallazgos) > 0 {
		r.huella = huella(t.datos, t.origen)
	}
	return r
}

// lineaJSON serializa v en una línea terminada en \n, sin escapar &, < ni >
// (igual que cfdi.Documento.Escribir).
func lineaJSON(v any) ([]byte, error) {
	var b bytes.Buffer
	e := json.NewEncoder(&b)
	e.SetEscapeHTML(false)
	err := e.Encode(v)
	return b.Bytes(), err
}

func extension(nombre string) string { return strings.ToLower(path.Ext(nombre)) }

// ignorado descarta la basura que dejan los compresores de macOS.
func ignorado(nombre string) bool {
	base := path.Base(nombre)
	return strings.HasPrefix(base, "._") || strings.HasPrefix(nombre, "__MACOSX/") || strings.Contains(nombre, "/__MACOSX/")
}

// recorrer emite un trabajo por cada .xml que encuentre bajo ruta. Devuelve
// error solo si la ruta misma no se puede abrir.
func recorrer(ruta string, emitir func(trabajo)) error {
	info, err := os.Stat(ruta)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		archivo(ruta, emitir)
		return nil
	}
	return filepath.WalkDir(ruta, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			emitir(trabajo{origen: p, err: err})
			if d != nil && d.IsDir() {
				return fs.SkipDir
			}
			return nil
		}
		if !d.IsDir() && !ignorado(filepath.ToSlash(p)) {
			switch extension(p) {
			case ".xml", ".zip":
				archivo(p, emitir)
			}
		}
		return nil
	})
}

// archivo emite un .xml (el worker lo lee) o abre un .zip. Una ruta explícita
// con otra extensión se trata como XML.
func archivo(p string, emitir func(trabajo)) {
	if extension(p) != ".zip" {
		emitir(trabajo{origen: p, ruta: p})
		return
	}
	z, err := zip.OpenReader(p)
	if err != nil {
		emitir(trabajo{origen: p, err: err})
		return
	}
	defer z.Close()
	entradasZip(&z.Reader, p, emitir)
}

func entradasZip(z *zip.Reader, origen string, emitir func(trabajo)) {
	for _, f := range z.File {
		if f.FileInfo().IsDir() || ignorado(f.Name) {
			continue
		}
		ext := extension(f.Name)
		if ext != ".xml" && ext != ".zip" {
			continue
		}
		o := origen + "!" + f.Name
		b, err := leerEntrada(f)
		if err != nil {
			emitir(trabajo{origen: o, err: err})
			continue
		}
		if ext == ".xml" {
			emitir(trabajo{origen: o, datos: b})
			continue
		}
		interno, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
		if err != nil {
			emitir(trabajo{origen: o, err: err})
			continue
		}
		entradasZip(interno, o, emitir)
	}
}

func leerEntrada(f *zip.File) ([]byte, error) {
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	// Con el tamaño declarado se reserva el búfer una vez, en vez de ir
	// duplicándolo como io.ReadAll. archive/zip ya exige que el tamaño real
	// coincida con el declarado (y el CRC) al llegar a EOF.
	if n := f.UncompressedSize64; n > 0 && n < 1<<30 {
		b := make([]byte, n)
		if _, err := io.ReadFull(rc, b); err != nil {
			return nil, err
		}
		if _, err := rc.Read(make([]byte, 1)); err != io.EOF {
			if err == nil {
				err = zip.ErrFormat
			}
			return nil, err
		}
		return b, nil
	}
	return io.ReadAll(rc)
}
