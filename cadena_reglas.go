// Código generado por `go run ./internal/gencadena/generar` a partir de los
// XSLT del SAT en internal/gencadena/xslt. NO EDITAR: cambiar el generador.

package cfdi

// hojaCFDI32 sale de cfd/3/cadenaoriginal_3_2/cadenaoriginal_3_2.xslt y de lo que incluye.
//
// Archivos (sha256):
//   - cfd/3/cadenaoriginal_3_2/cadenaoriginal_3_2.xslt 17bb9f2fa3bd585bc5aab37256856690922da1ecb3e28b49f961e1fed46cd1cd
//   - cfd/2/cadenaoriginal_2_0/utilerias.xslt d9ef68267747e0f6725295d4012684a4c051cb3f2265dcb2035d32bf40895016
//   - cfd/ecc/ecc.xslt 8fca4c74f2b2d66946da84eac9e5422c92b54990d44be45324d3e66d4d5538c8
//   - cfd/donat/donat11.xslt 1c72d07d9945fd7824471caaebaee47dafc3b56785a037549fcbee5e123c6d36
//   - cfd/divisas/divisas.xslt be7167e1d421b66ba17a6f1788c964c5bb1f207ed7ad9803166d60966a332b1d
//   - cfd/ecb/ecb.xslt c61979e58b4971d68c7ebf76d77215a7c0fb7149da1763c0fc99cfe7e020e655
//   - cfd/implocal/implocal.xslt 3c2ee6c3ecce9b2e74b2a225f3c17250259751a1307f58a96799ff119e2fafe6
//   - cfd/terceros/terceros11.xslt 62b8986466db9b1032462bfbe0cf3514caa8d5e1b8b5906caeab44179853996a
//   - cfd/iedu/iedu.xslt 2e9ddb05e5cd222cded57a5306b43f954f15a2ac9f8ddf3e4818156cd703a49d
//   - cfd/ventavehiculos/ventavehiculos11.xslt 95fd6522700a5125dacc363e07fac8ff2f470b2e23aec5bbe3d3f684c8a4152f
//   - cfd/pfic/pfic.xslt 0e9e992e4e0fa6e13ee7b18ce5ddc51c55f3615236c7c51b80426811f9779bab
//   - cfd/TuristaPasajeroExtranjero/TuristaPasajeroExtranjero.xslt 06aee48f06ed53ab4fad9f2ddc950bca51840cfd4fb137ec8d14632e2c2caeb3
//   - cfd/leyendasFiscales/leyendasFisc.xslt 68b0016e44c1829a5c7d5da340b69945100d8fda1243e4e279c2cdc8c745d1fb
//   - cfd/spei/spei.xslt c1607fd73a8ea890b3c5018234b500bbec7eb0cfba14afead4e03038d48e11d3
//   - cfd/nomina/nomina11.xslt 3527ad6f93a0f10b289c996c6846da208a9ea54e7b647dd245aeb2c44cd4590c
//   - cfd/cfdiregistrofiscal/cfdiregistrofiscal.xslt d824c4993e1f5940ff0605bdd3ea040669bef69355b551ab9da916f3a3746e7d
//   - cfd/pagoenespecie/pagoenespecie.xslt ed33ce278525a43db377d82fe641d42391d74b731f82015b1fedfa84654f1798
//   - cfd/consumodecombustibles/consumodecombustibles.xslt e1141cfa02dc67de719f9e31c6abc997c19831715544a2d525ad14dadde4b623
//   - cfd/valesdedespensa/valesdedespensa.xslt e22cf591f5669b78b3306c33e2434acbf7f40c3671e318c3fb2b736a2bf6ef8c
//   - cfd/aerolineas/aerolineas.xslt d000d0ed64a45a05a7f59b1eb2399d55d1588f18056ed37e8bc642ca2fe655c0
//   - cfd/notariospublicos/notariospublicos.xslt b3aeebdbfa1858b60d9f8b437f9f07a95f600e712b52aa0a8c08a7f924fe9ce7
//   - cfd/vehiculousado/vehiculousado.xslt 226664fcdcbc51e57abb260104b59a06d3bba62a62ae92fbbfeee6f16791ae2c
//   - cfd/servicioparcialconstruccion/servicioparcialconstruccion.xslt b3b9a8983a673a352ac291c1e0a9605e983dc0af8a4ecb4f6ec0cb13e258fe85
//   - cfd/certificadodestruccion/certificadodedestruccion.xslt 0b306d8142eb118f47d3ab6c967a6be6c9c31436678e776e571ce496018b68d5
//   - cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt 347bf6dfdae5c0b18d6a107c8c339f212763496fa4ae1fa748d0354a7c3338d2
//   - cfd/arteantiguedades/obrasarteantiguedades.xslt 2d11a854ff914b229ef7d2d43cdfd6296315d040c36e0ad6c721ace08ba89d1d
//   - cfd/acreditamiento/AcreditamientoIEPS10.xslt 83caef3dfe64d2f69d7994c902d583da3025db8a31ad3eba97a05804b2f2717d
//   - cfd/EstadoDeCuentaCombustible/ecc11.xslt d9bb8d26d511136e68ca8b986c7fedcb946aaffe0a29d986c30c30970ef219d6
//   - cfd/ComercioExterior/ComercioExterior10.xslt 3fe537833ae5cdc7990239276c3b2569b6fa1d1c57f3ebcc7d87e2ab220759ae
//   - cfd/ine/ine10.xslt 02781b96a0da332693f719f893f7a72117196b073c5d0df48e7fa6a4ed444882
//   - cfd/ine/ine11.xslt 02781b96a0da332693f719f893f7a72117196b073c5d0df48e7fa6a4ed444882
//   - cfd/nomina/nomina12.xslt a14cdb44995526aa5888036bedf7fac9f84ad0041f9393371b0bbbea0572649e
//   - cfd/ComercioExterior11/ComercioExterior11.xslt bb9e7c3d2a2a3e8af9c65498fd1422b0f400e0787d2418b460d51ece6c988f9d
//
// Includes omitidos (complementos que la librería no lee):
//   - cfd/psgecfd/psgecfd.xslt
//   - cfd/detallista/detallista.xslt
//
// Plantillas match="/" de hojas incluidas, que la de la raíz sustituye:
//   - cfd/vehiculousado/vehiculousado.xslt
var hojaCFDI32 = &hojaCadena{
	nombre: "cadenaoriginal_3_2",
	raiz: []pasoCadena{
		txt("|"),
		aplicar(abs(hijo("http://www.sat.gob.mx/cfd/3", "Comprobante"))),
		txt("||"),
	},
	plantillas: map[nombreXML][]pasoCadena{
		// cfd/ecc/ecc.xslt
		{"http://www.sat.gob.mx/ecc", "EstadoDeCuentaCombustible"}: {
			req("tipoOperacion"),
			req("numeroDeCuenta"),
			opc("subTotal"),
			req("total"),
			aplicar(rel(hijo("http://www.sat.gob.mx/ecc", "Conceptos"))),
		},
		// cfd/ecc/ecc.xslt
		{"http://www.sat.gob.mx/ecc", "Conceptos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/ecc", "ConceptoEstadoDeCuentaCombustible")),
				aplicar(rel()),
			),
		},
		// cfd/ecc/ecc.xslt
		{"http://www.sat.gob.mx/ecc", "Traslados"}: {
			cada(rel(hijo("http://www.sat.gob.mx/ecc", "Traslado")),
				aplicar(rel()),
			),
		},
		// cfd/ecc/ecc.xslt
		{"http://www.sat.gob.mx/ecc", "ConceptoEstadoDeCuentaCombustible"}: {
			req("identificador"),
			req("fecha"),
			req("rfc"),
			req("claveEstacion"),
			req("cantidad"),
			req("nombreCombustible"),
			req("folioOperacion"),
			req("valorUnitario"),
			req("importe"),
			cada(rel(hijo("http://www.sat.gob.mx/ecc", "Traslados")),
				aplicar(rel()),
			),
		},
		// cfd/ecc/ecc.xslt
		{"http://www.sat.gob.mx/ecc", "Traslado"}: {
			req("impuesto"),
			req("tasa"),
			req("importe"),
		},
		// cfd/donat/donat11.xslt
		{"http://www.sat.gob.mx/donat", "Donatarias"}: {
			req("version"),
			req("noAutorizacion"),
			req("fechaAutorizacion"),
			req("leyenda"),
		},
		// cfd/divisas/divisas.xslt
		{"http://www.sat.gob.mx/divisas", "Divisas"}: {
			req("version"),
			req("tipoOperacion"),
		},
		// cfd/ecb/ecb.xslt
		{"http://www.sat.gob.mx/ecb", "EstadoDeCuentaBancario"}: {
			req("version"),
			req("numeroCuenta"),
			req("nombreCliente"),
			cada(rel(hijo("http://www.sat.gob.mx/ecb", "Movimientos"), hijo("http://www.sat.gob.mx/ecb", "MovimientoECBFiscal")),
				req("fecha"),
				req("RFCenajenante"),
				req("Importe"),
			),
		},
		// cfd/implocal/implocal.xslt
		{"http://www.sat.gob.mx/implocal", "ImpuestosLocales"}: {
			req("version"),
			req("TotaldeRetenciones"),
			req("TotaldeTraslados"),
			cada(rel(hijo("http://www.sat.gob.mx/implocal", "RetencionesLocales")),
				req("ImpLocRetenido"),
				req("TasadeRetencion"),
				req("Importe"),
			),
			cada(rel(hijo("http://www.sat.gob.mx/implocal", "TrasladosLocales")),
				req("ImpLocTrasladado"),
				req("TasadeTraslado"),
				req("Importe"),
			),
		},
		// cfd/terceros/terceros11.xslt
		{"http://www.sat.gob.mx/terceros", "PorCuentadeTerceros"}: {
			req("version"),
			req("rfc"),
			opc("nombre"),
			aplicar(rel(desc("http://www.sat.gob.mx/terceros", "InformacionFiscalTercero"))),
			cada(rel(desc("http://www.sat.gob.mx/terceros", "InformacionAduanera")),
				aplicar(rel()),
			),
			si(rel(hijo("http://www.sat.gob.mx/terceros", "CuentaPredial")),
				aplicar(rel(hijo("http://www.sat.gob.mx/terceros", "CuentaPredial"))),
			),
			cada(rel(desc("http://www.sat.gob.mx/terceros", "Retenciones"), hijo("http://www.sat.gob.mx/terceros", "Retencion")),
				aplicar(rel()),
			),
			cada(rel(desc("http://www.sat.gob.mx/terceros", "Traslados"), hijo("http://www.sat.gob.mx/terceros", "Traslado")),
				aplicar(rel()),
			),
		},
		// cfd/terceros/terceros11.xslt
		{"http://www.sat.gob.mx/terceros", "Retencion"}: {
			req("impuesto"),
			req("importe"),
		},
		// cfd/terceros/terceros11.xslt
		{"http://www.sat.gob.mx/terceros", "Traslado"}: {
			req("impuesto"),
			req("tasa"),
			req("importe"),
		},
		// cfd/terceros/terceros11.xslt
		{"http://www.sat.gob.mx/terceros", "InformacionAduanera"}: {
			req("numero"),
			req("fecha"),
			opc("aduana"),
		},
		// cfd/terceros/terceros11.xslt
		{"http://www.sat.gob.mx/terceros", "CuentaPredial"}: {
			req("numero"),
		},
		// cfd/terceros/terceros11.xslt
		{"http://www.sat.gob.mx/terceros", "InformacionFiscalTercero"}: {
			req("calle"),
			opc("noExterior"),
			opc("noInterior"),
			opc("colonia"),
			opc("localidad"),
			opc("referencia"),
			req("municipio"),
			req("estado"),
			req("pais"),
			req("codigoPostal"),
		},
		// cfd/iedu/iedu.xslt
		{"http://www.sat.gob.mx/iedu", "instEducativas"}: {
			req("version"),
			req("nombreAlumno"),
			req("CURP"),
			req("nivelEducativo"),
			req("autRVOE"),
			opc("rfcPago"),
		},
		// cfd/ventavehiculos/ventavehiculos11.xslt
		{"http://www.sat.gob.mx/ventavehiculos", "VentaVehiculos"}: {
			req("version"),
			req("ClaveVehicular"),
			siIgual("version", "1.1",
				req("Niv"),
			),
			cada(rel(desc("http://www.sat.gob.mx/ventavehiculos", "InformacionAduanera")),
				aplicar(rel()),
			),
		},
		// cfd/ventavehiculos/ventavehiculos11.xslt
		{"http://www.sat.gob.mx/ventavehiculos", "InformacionAduanera"}: {
			req("numero"),
			req("fecha"),
			opc("aduana"),
		},
		// cfd/pfic/pfic.xslt
		{"http://www.sat.gob.mx/pfic", "PFintegranteCoordinado"}: {
			req("version"),
			req("ClaveVehicular"),
			req("Placa"),
			opc("RFCPF"),
		},
		// cfd/TuristaPasajeroExtranjero/TuristaPasajeroExtranjero.xslt
		{"http://www.sat.gob.mx/TuristaPasajeroExtranjero", "TuristaPasajeroExtranjero"}: {
			req("version"),
			req("fechadeTransito"),
			req("tipoTransito"),
			aplicar(rel(hijo("http://www.sat.gob.mx/TuristaPasajeroExtranjero", "datosTransito"))),
		},
		// cfd/TuristaPasajeroExtranjero/TuristaPasajeroExtranjero.xslt
		{"http://www.sat.gob.mx/TuristaPasajeroExtranjero", "datosTransito"}: {
			req("Via"),
			req("TipoId"),
			req("NumeroId"),
			req("Nacionalidad"),
			req("EmpresaTransporte"),
			opc("IdTransporte"),
		},
		// cfd/leyendasFiscales/leyendasFisc.xslt
		{"http://www.sat.gob.mx/leyendasFiscales", "LeyendasFiscales"}: {
			req("version"),
			cada(rel(hijo("http://www.sat.gob.mx/leyendasFiscales", "Leyenda")),
				aplicar(rel()),
			),
		},
		// cfd/leyendasFiscales/leyendasFisc.xslt
		{"http://www.sat.gob.mx/leyendasFiscales", "Leyenda"}: {
			opc("disposicionFiscal"),
			opc("norma"),
			req("textoLeyenda"),
		},
		// cfd/spei/spei.xslt
		{"http://www.sat.gob.mx/spei", "Complemento_SPEI"}: {
			cada(rel(hijo("http://www.sat.gob.mx/spei", "SPEI_Tercero")),
				aplicar(rel()),
			),
		},
		// cfd/spei/spei.xslt
		{"http://www.sat.gob.mx/spei", "SPEI_Tercero"}: {
			req("FechaOperacion"),
			req("Hora"),
			req("ClaveSPEI"),
			req("sello"),
			req("numeroCertificado"),
			aplicar(rel(hijo("http://www.sat.gob.mx/spei", "Ordenante"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/spei", "Beneficiario"))),
		},
		// cfd/spei/spei.xslt
		{"http://www.sat.gob.mx/spei", "Ordenante"}: {
			req("BancoEmisor"),
			req("Nombre"),
			req("TipoCuenta"),
			req("Cuenta"),
			req("RFC"),
		},
		// cfd/spei/spei.xslt
		{"http://www.sat.gob.mx/spei", "Beneficiario"}: {
			req("BancoReceptor"),
			req("Nombre"),
			req("TipoCuenta"),
			req("Cuenta"),
			req("RFC"),
			req("Concepto"),
			opc("IVA"),
			req("MontoPago"),
		},
		// cfd/nomina/nomina11.xslt
		{"http://www.sat.gob.mx/nomina", "Nomina"}: {
			siIgual("Version", "1.0",
				req("Version"),
				req("RegistroPatronal"),
				req("NumEmpleado"),
				req("CURP"),
				req("TipoRegimen"),
				opc("NumSeguridadSocial"),
				opc("CLABE"),
				opc("Banco"),
				conjunto(rel(hijo("http://www.sat.gob.mx/nomina", "Ingresos")),
					req("TotalGravado"),
				),
				conjunto(rel(hijo("http://www.sat.gob.mx/nomina", "Ingresos")),
					req("TotalExento"),
				),
				conjunto(rel(hijo("http://www.sat.gob.mx/nomina", "Descuentos")),
					req("Total"),
				),
			),
			siIgual("Version", "1.1",
				req("Version"),
				opc("RegistroPatronal"),
				req("NumEmpleado"),
				req("CURP"),
				req("TipoRegimen"),
				opc("NumSeguridadSocial"),
				req("FechaPago"),
				req("FechaInicialPago"),
				req("FechaFinalPago"),
				req("NumDiasPagados"),
				opc("Departamento"),
				opc("CLABE"),
				opc("Banco"),
				opc("FechaInicioRelLaboral"),
				opc("Antiguedad"),
				opc("Puesto"),
				opc("TipoContrato"),
				opc("TipoJornada"),
				req("PeriodicidadPago"),
				opc("SalarioBaseCotApor"),
				opc("RiesgoPuesto"),
				opc("SalarioDiarioIntegrado"),
				si(rel(hijo("http://www.sat.gob.mx/nomina", "Percepciones")),
					aplicar(rel(hijo("http://www.sat.gob.mx/nomina", "Percepciones"))),
				),
				si(rel(hijo("http://www.sat.gob.mx/nomina", "Deducciones")),
					aplicar(rel(hijo("http://www.sat.gob.mx/nomina", "Deducciones"))),
				),
				cada(rel(hijo("http://www.sat.gob.mx/nomina", "Incapacidades")),
					aplicar(rel()),
				),
				cada(rel(hijo("http://www.sat.gob.mx/nomina", "HorasExtras")),
					aplicar(rel()),
				),
			),
		},
		// cfd/nomina/nomina11.xslt
		{"http://www.sat.gob.mx/nomina", "Percepciones"}: {
			req("TotalGravado"),
			req("TotalExento"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina", "Percepcion")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina11.xslt
		{"http://www.sat.gob.mx/nomina", "Percepcion"}: {
			req("TipoPercepcion"),
			req("Clave"),
			req("Concepto"),
			req("ImporteGravado"),
			req("ImporteExento"),
		},
		// cfd/nomina/nomina11.xslt
		{"http://www.sat.gob.mx/nomina", "Deducciones"}: {
			req("TotalGravado"),
			req("TotalExento"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina", "Deduccion")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina11.xslt
		{"http://www.sat.gob.mx/nomina", "Deduccion"}: {
			req("TipoDeduccion"),
			req("Clave"),
			req("Concepto"),
			req("ImporteGravado"),
			req("ImporteExento"),
		},
		// cfd/nomina/nomina11.xslt
		{"http://www.sat.gob.mx/nomina", "Incapacidades"}: {
			cada(rel(hijo("http://www.sat.gob.mx/nomina", "Incapacidad")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina11.xslt
		{"http://www.sat.gob.mx/nomina", "Incapacidad"}: {
			req("DiasIncapacidad"),
			req("TipoIncapacidad"),
			req("Descuento"),
		},
		// cfd/nomina/nomina11.xslt
		{"http://www.sat.gob.mx/nomina", "HorasExtras"}: {
			cada(rel(hijo("http://www.sat.gob.mx/nomina", "HorasExtra")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina11.xslt
		{"http://www.sat.gob.mx/nomina", "HorasExtra"}: {
			req("Dias"),
			req("TipoHoras"),
			req("HorasExtra"),
			req("ImportePagado"),
		},
		// cfd/cfdiregistrofiscal/cfdiregistrofiscal.xslt
		{"http://www.sat.gob.mx/registrofiscal", "CFDIRegistroFiscal"}: {
			req("Version"),
			req("Folio"),
		},
		// cfd/pagoenespecie/pagoenespecie.xslt
		{"http://www.sat.gob.mx/pagoenespecie", "PagoEnEspecie"}: {
			req("Version"),
			req("CvePIC"),
			req("FolioSolDon"),
			req("PzaArtNombre"),
			req("PzaArtTecn"),
			req("PzaArtAProd"),
			req("PzaArtDim"),
		},
		// cfd/consumodecombustibles/consumodecombustibles.xslt
		{"http://www.sat.gob.mx/consumodecombustibles", "ConsumoDeCombustibles"}: {
			req("version"),
			req("tipoOperacion"),
			req("numeroDeCuenta"),
			opc("subTotal"),
			req("total"),
			aplicar(rel(hijo("http://www.sat.gob.mx/consumodecombustibles", "Conceptos"))),
		},
		// cfd/consumodecombustibles/consumodecombustibles.xslt
		{"http://www.sat.gob.mx/consumodecombustibles", "Conceptos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/consumodecombustibles", "ConceptoConsumoDeCombustibles")),
				aplicar(rel()),
			),
		},
		// cfd/consumodecombustibles/consumodecombustibles.xslt
		{"http://www.sat.gob.mx/consumodecombustibles", "ConceptoConsumoDeCombustibles"}: {
			req("identificador"),
			req("fecha"),
			req("rfc"),
			req("claveEstacion"),
			req("cantidad"),
			req("nombreCombustible"),
			req("folioOperacion"),
			req("valorUnitario"),
			req("importe"),
			cada(rel(hijo("http://www.sat.gob.mx/consumodecombustibles", "Determinados")),
				aplicar(rel()),
			),
		},
		// cfd/consumodecombustibles/consumodecombustibles.xslt
		{"http://www.sat.gob.mx/consumodecombustibles", "Determinados"}: {
			cada(rel(hijo("http://www.sat.gob.mx/consumodecombustibles", "Determinado")),
				aplicar(rel()),
			),
		},
		// cfd/consumodecombustibles/consumodecombustibles.xslt
		{"http://www.sat.gob.mx/consumodecombustibles", "Determinado"}: {
			req("impuesto"),
			req("tasa"),
			req("importe"),
		},
		// cfd/valesdedespensa/valesdedespensa.xslt
		{"http://www.sat.gob.mx/valesdedespensa", "ValesDeDespensa"}: {
			req("version"),
			req("tipoOperacion"),
			opc("registroPatronal"),
			req("numeroDeCuenta"),
			req("total"),
			aplicar(rel(hijo("http://www.sat.gob.mx/valesdedespensa", "Conceptos"))),
		},
		// cfd/valesdedespensa/valesdedespensa.xslt
		{"http://www.sat.gob.mx/valesdedespensa", "Conceptos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/valesdedespensa", "Concepto")),
				aplicar(rel()),
			),
		},
		// cfd/valesdedespensa/valesdedespensa.xslt
		{"http://www.sat.gob.mx/valesdedespensa", "Concepto"}: {
			req("identificador"),
			req("fecha"),
			req("rfc"),
			req("curp"),
			req("nombre"),
			opc("numSeguridadSocial"),
			req("importe"),
		},
		// cfd/aerolineas/aerolineas.xslt
		{"http://www.sat.gob.mx/aerolineas", "Aerolineas"}: {
			req("Version"),
			req("TUA"),
			aplicar(rel(hijo("http://www.sat.gob.mx/aerolineas", "OtrosCargos"))),
		},
		// cfd/aerolineas/aerolineas.xslt
		{"http://www.sat.gob.mx/aerolineas", "OtrosCargos"}: {
			req("TotalCargos"),
			cada(rel(hijo("http://www.sat.gob.mx/aerolineas", "Cargo")),
				aplicar(rel()),
			),
		},
		// cfd/aerolineas/aerolineas.xslt
		{"http://www.sat.gob.mx/aerolineas", "Cargo"}: {
			req("CodigoCargo"),
			req("Importe"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "NotariosPublicos"}: {
			req("Version"),
			aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DescInmuebles"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosOperacion"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosNotario"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosEnajenante"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosAdquiriente"))),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DescInmuebles"}: {
			cada(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DescInmueble")),
				aplicar(rel()),
			),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DescInmueble"}: {
			req("TipoInmueble"),
			req("Calle"),
			opc("NoExterior"),
			opc("NoInterior"),
			opc("Colonia"),
			opc("Localidad"),
			opc("Referencia"),
			req("Municipio"),
			req("Estado"),
			req("Pais"),
			req("CodigoPostal"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosOperacion"}: {
			req("NumInstrumentoNotarial"),
			req("FechaInstNotarial"),
			req("MontoOperacion"),
			req("Subtotal"),
			req("IVA"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosNotario"}: {
			req("CURP"),
			req("NumNotaria"),
			req("EntidadFederativa"),
			opc("Adscripcion"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosEnajenante"}: {
			req("CoproSocConyugalE"),
			si(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosUnEnajenante")),
				aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosUnEnajenante"))),
			),
			si(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosEnajenantesCopSC")),
				aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosEnajenantesCopSC"))),
			),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosUnEnajenante"}: {
			req("Nombre"),
			req("ApellidoPaterno"),
			opc("ApellidoMaterno"),
			req("RFC"),
			req("CURP"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosEnajenantesCopSC"}: {
			cada(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosEnajenanteCopSC")),
				aplicar(rel()),
			),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosEnajenanteCopSC"}: {
			req("Nombre"),
			opc("ApellidoPaterno"),
			opc("ApellidoMaterno"),
			req("RFC"),
			opc("CURP"),
			req("Porcentaje"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosAdquiriente"}: {
			req("CoproSocConyugalE"),
			si(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosUnAdquiriente")),
				aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosUnAdquiriente"))),
			),
			si(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosAdquirientesCopSC")),
				aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosAdquirientesCopSC"))),
			),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosUnAdquiriente"}: {
			req("Nombre"),
			opc("ApellidoPaterno"),
			opc("ApellidoMaterno"),
			req("RFC"),
			opc("CURP"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosAdquirientesCopSC"}: {
			cada(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosAdquirienteCopSC")),
				aplicar(rel()),
			),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosAdquirienteCopSC"}: {
			req("Nombre"),
			opc("ApellidoPaterno"),
			opc("ApellidoMaterno"),
			req("RFC"),
			opc("CURP"),
			req("Porcentaje"),
		},
		// cfd/vehiculousado/vehiculousado.xslt
		{"http://www.sat.gob.mx/vehiculousado", "VehiculoUsado"}: {
			req("Version"),
			req("montoAdquisicion"),
			req("montoEnajenacion"),
			req("claveVehicular"),
			req("marca"),
			req("tipo"),
			req("modelo"),
			opc("numeroMotor"),
			opc("numeroSerie"),
			opc("NIV"),
			req("valor"),
			aplicar(rel(hijo("http://www.sat.gob.mx/vehiculousado", "InformacionAduanera"))),
		},
		// cfd/vehiculousado/vehiculousado.xslt
		{"http://www.sat.gob.mx/vehiculousado", "InformacionAduanera"}: {
			req("numero"),
			req("fecha"),
			opc("aduana"),
		},
		// cfd/servicioparcialconstruccion/servicioparcialconstruccion.xslt
		{"http://www.sat.gob.mx/servicioparcialconstruccion", "parcialesconstruccion"}: {
			req("Version"),
			req("NumPerLicoAut"),
			aplicar(rel(hijo("http://www.sat.gob.mx/servicioparcialconstruccion", "Inmueble"))),
		},
		// cfd/servicioparcialconstruccion/servicioparcialconstruccion.xslt
		{"http://www.sat.gob.mx/servicioparcialconstruccion", "Inmueble"}: {
			req("Calle"),
			opc("NoExterior"),
			opc("NoInterior"),
			opc("Colonia"),
			opc("Localidad"),
			opc("Referencia"),
			req("Municipio"),
			req("Estado"),
			req("CodigoPostal"),
		},
		// cfd/certificadodestruccion/certificadodedestruccion.xslt
		{"http://www.sat.gob.mx/certificadodestruccion", "certificadodedestruccion"}: {
			req("Version"),
			req("Serie"),
			req("NumFolDesVeh"),
			aplicar(rel(hijo("http://www.sat.gob.mx/certificadodestruccion", "VehiculoDestruido"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/certificadodestruccion", "InformacionAduanera"))),
		},
		// cfd/certificadodestruccion/certificadodedestruccion.xslt
		{"http://www.sat.gob.mx/certificadodestruccion", "VehiculoDestruido"}: {
			req("Marca"),
			req("TipooClase"),
			req("Año"),
			opc("Modelo"),
			opc("NIV"),
			opc("NumSerie"),
			req("NumPlacas"),
			opc("NumMotor"),
			req("NumFolTarjCir"),
		},
		// cfd/certificadodestruccion/certificadodedestruccion.xslt
		{"http://www.sat.gob.mx/certificadodestruccion", "InformacionAduanera"}: {
			req("NumPedImp"),
			req("Fecha"),
			req("Aduana"),
		},
		// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt
		{"http://www.sat.gob.mx/renovacionysustitucionvehiculos", "renovacionysustitucionvehiculos"}: {
			req("Version"),
			req("TipoDeDecreto"),
			aplicar(rel(hijo("http://www.sat.gob.mx/renovacionysustitucionvehiculos", "DecretoRenovVehicular"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/renovacionysustitucionvehiculos", "DecretoSustitVehicular"))),
		},
		// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt
		{"http://www.sat.gob.mx/renovacionysustitucionvehiculos", "DecretoRenovVehicular"}: {
			req("VehEnaj"),
			cada(rel(hijo("http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculosUsadosEnajenadoPermAlFab")),
				aplicar(rel()),
			),
			aplicar(rel(hijo("http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculoNuvoSemEnajenadoFabAlPerm"))),
		},
		// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt
		{"http://www.sat.gob.mx/renovacionysustitucionvehiculos", "DecretoSustitVehicular"}: {
			req("VehEnaj"),
			aplicar(rel(hijo("http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculoUsadoEnajenadoPermAlFab"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculoNuvoSemEnajenadoFabAlPerm"))),
		},
		// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt
		{"http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculosUsadosEnajenadoPermAlFab"}: {
			req("PrecioVehUsado"),
			req("TipoVeh"),
			req("Marca"),
			req("TipooClase"),
			req("Año"),
			opc("Modelo"),
			opc("NIV"),
			opc("NumSerie"),
			req("NumPlacas"),
			opc("NumMotor"),
			req("NumFolTarjCir"),
			opc("NumPedIm"),
			opc("Aduana"),
			opc("FechaRegulVeh"),
			req("Foliofiscal"),
		},
		// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt
		{"http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculoNuvoSemEnajenadoFabAlPerm"}: {
			req("Año"),
			opc("Modelo"),
			req("NumPlacas"),
			opc("RFC"),
		},
		// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt
		{"http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculoUsadoEnajenadoPermAlFab"}: {
			req("PrecioVehUsado"),
			req("TipoVeh"),
			req("Marca"),
			req("TipooClase"),
			req("Año"),
			opc("Modelo"),
			opc("NIV"),
			opc("NumSerie"),
			req("NumPlacas"),
			opc("NumMotor"),
			req("NumFolTarjCir"),
			req("NumFolAvisoint"),
			req("NumPedIm"),
			req("Aduana"),
			req("FechaRegulVeh"),
			req("Foliofiscal"),
		},
		// cfd/arteantiguedades/obrasarteantiguedades.xslt
		{"http://www.sat.gob.mx/arteantiguedades", "obrasarteantiguedades"}: {
			req("Version"),
			req("TipoBien"),
			opc("OtrosTipoBien"),
			req("TituloAdquirido"),
			opc("OtrosTituloAdquirido"),
			opc("Subtotal"),
			opc("IVA"),
			req("FechaAdquisicion"),
			req("CaracterísticasDeObraoPieza"),
		},
		// cfd/acreditamiento/AcreditamientoIEPS10.xslt
		{"http://www.sat.gob.mx/acreditamiento", "acreditamientoIEPS"}: {
			req("Version"),
			req("TAR"),
		},
		// cfd/EstadoDeCuentaCombustible/ecc11.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible", "EstadoDeCuentaCombustible"}: {
			req("Version"),
			req("TipoOperacion"),
			req("NumeroDeCuenta"),
			req("SubTotal"),
			req("Total"),
			aplicar(rel(hijo("http://www.sat.gob.mx/EstadoDeCuentaCombustible", "Conceptos"))),
		},
		// cfd/EstadoDeCuentaCombustible/ecc11.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible", "Conceptos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/EstadoDeCuentaCombustible", "ConceptoEstadoDeCuentaCombustible")),
				aplicar(rel()),
			),
		},
		// cfd/EstadoDeCuentaCombustible/ecc11.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible", "Traslados"}: {
			cada(rel(hijo("http://www.sat.gob.mx/EstadoDeCuentaCombustible", "Traslado")),
				aplicar(rel()),
			),
		},
		// cfd/EstadoDeCuentaCombustible/ecc11.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible", "ConceptoEstadoDeCuentaCombustible"}: {
			req("Identificador"),
			req("Fecha"),
			req("Rfc"),
			req("ClaveEstacion"),
			opc("TAR"),
			req("Cantidad"),
			req("NoIdentificacion"),
			opc("Unidad"),
			req("NombreCombustible"),
			req("FolioOperacion"),
			req("ValorUnitario"),
			req("Importe"),
			aplicar(rel(hijo("http://www.sat.gob.mx/EstadoDeCuentaCombustible", "Traslados"))),
		},
		// cfd/EstadoDeCuentaCombustible/ecc11.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible", "Traslado"}: {
			req("Impuesto"),
			req("TasaoCuota"),
			req("Importe"),
		},
		// cfd/ComercioExterior/ComercioExterior10.xslt
		{"http://www.sat.gob.mx/ComercioExterior", "ComercioExterior"}: {
			req("Version"),
			req("TipoOperacion"),
			opc("ClaveDePedimento"),
			opc("CertificadoOrigen"),
			opc("NumCertificadoOrigen"),
			opc("NumeroExportadorConfiable"),
			opc("Incoterm"),
			opc("Subdivision"),
			opc("Observaciones"),
			opc("TipoCambioUSD"),
			opc("TotalUSD"),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior", "Emisor"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior", "Receptor"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior", "Destinatario"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior", "Mercancias"))),
		},
		// cfd/ComercioExterior/ComercioExterior10.xslt
		{"http://www.sat.gob.mx/ComercioExterior", "Emisor"}: {
			opc("Curp"),
		},
		// cfd/ComercioExterior/ComercioExterior10.xslt
		{"http://www.sat.gob.mx/ComercioExterior", "Receptor"}: {
			opc("Curp"),
			req("NumRegIdTrib"),
		},
		// cfd/ComercioExterior/ComercioExterior10.xslt
		{"http://www.sat.gob.mx/ComercioExterior", "Destinatario"}: {
			opc("NumRegIdTrib"),
			opc("Rfc"),
			opc("Curp"),
			opc("Nombre"),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior", "Domicilio"))),
		},
		// cfd/ComercioExterior/ComercioExterior10.xslt
		{"http://www.sat.gob.mx/ComercioExterior", "Mercancias"}: {
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior", "Mercancia")),
				aplicar(rel()),
			),
		},
		// cfd/ComercioExterior/ComercioExterior10.xslt
		{"http://www.sat.gob.mx/ComercioExterior", "Domicilio"}: {
			req("Calle"),
			opc("NumeroExterior"),
			opc("NumeroInterior"),
			opc("Colonia"),
			opc("Localidad"),
			opc("Referencia"),
			opc("Municipio"),
			req("Estado"),
			req("Pais"),
			req("CodigoPostal"),
		},
		// cfd/ComercioExterior/ComercioExterior10.xslt
		{"http://www.sat.gob.mx/ComercioExterior", "Mercancia"}: {
			req("NoIdentificacion"),
			opc("FraccionArancelaria"),
			opc("CantidadAduana"),
			opc("UnidadAduana"),
			opc("ValorUnitarioAduana"),
			req("ValorDolares"),
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior", "DescripcionesEspecificas")),
				aplicar(rel()),
			),
		},
		// cfd/ComercioExterior/ComercioExterior10.xslt
		{"http://www.sat.gob.mx/ComercioExterior", "DescripcionesEspecificas"}: {
			req("Marca"),
			opc("Modelo"),
			opc("SubModelo"),
			opc("NumeroSerie"),
		},
		// cfd/ine/ine10.xslt
		{"http://www.sat.gob.mx/ine", "INE"}: {
			req("Version"),
			req("TipoProceso"),
			opc("TipoComite"),
			opc("IdContabilidad"),
			cada(rel(hijo("http://www.sat.gob.mx/ine", "Entidad")),
				aplicar(rel()),
			),
		},
		// cfd/ine/ine10.xslt
		{"http://www.sat.gob.mx/ine", "Entidad"}: {
			req("ClaveEntidad"),
			opc("Ambito"),
			cada(rel(hijo("http://www.sat.gob.mx/ine", "Contabilidad")),
				aplicar(rel()),
			),
		},
		// cfd/ine/ine10.xslt
		{"http://www.sat.gob.mx/ine", "Contabilidad"}: {
			req("IdContabilidad"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Nomina"}: {
			req("Version"),
			req("TipoNomina"),
			req("FechaPago"),
			req("FechaInicialPago"),
			req("FechaFinalPago"),
			req("NumDiasPagados"),
			opc("TotalPercepciones"),
			opc("TotalDeducciones"),
			opc("TotalOtrosPagos"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Emisor")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Receptor")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Percepciones")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Deducciones")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "OtrosPagos")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Incapacidades")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Emisor"}: {
			opc("Curp"),
			opc("RegistroPatronal"),
			opc("RfcPatronOrigen"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "EntidadSNCF")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "EntidadSNCF"}: {
			req("OrigenRecurso"),
			opc("MontoRecursoPropio"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Receptor"}: {
			req("Curp"),
			opc("NumSeguridadSocial"),
			opc("FechaInicioRelLaboral"),
			opc("Antigüedad"),
			req("TipoContrato"),
			opc("Sindicalizado"),
			opc("TipoJornada"),
			req("TipoRegimen"),
			req("NumEmpleado"),
			opc("Departamento"),
			opc("Puesto"),
			opc("RiesgoPuesto"),
			req("PeriodicidadPago"),
			opc("Banco"),
			opc("CuentaBancaria"),
			opc("SalarioBaseCotApor"),
			opc("SalarioDiarioIntegrado"),
			req("ClaveEntFed"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "SubContratacion")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "SubContratacion"}: {
			req("RfcLabora"),
			req("PorcentajeTiempo"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Percepciones"}: {
			opc("TotalSueldos"),
			opc("TotalSeparacionIndemnizacion"),
			opc("TotalJubilacionPensionRetiro"),
			req("TotalGravado"),
			req("TotalExento"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Percepcion")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "JubilacionPensionRetiro")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "SeparacionIndemnizacion")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Percepcion"}: {
			req("TipoPercepcion"),
			req("Clave"),
			req("Concepto"),
			req("ImporteGravado"),
			req("ImporteExento"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "AccionesOTitulos")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "HorasExtra")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "AccionesOTitulos"}: {
			req("ValorMercado"),
			req("PrecioAlOtorgarse"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "HorasExtra"}: {
			req("Dias"),
			req("TipoHoras"),
			req("HorasExtra"),
			req("ImportePagado"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "JubilacionPensionRetiro"}: {
			opc("TotalUnaExhibicion"),
			opc("TotalParcialidad"),
			opc("MontoDiario"),
			req("IngresoAcumulable"),
			req("IngresoNoAcumulable"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "SeparacionIndemnizacion"}: {
			req("TotalPagado"),
			req("NumAñosServicio"),
			req("UltimoSueldoMensOrd"),
			req("IngresoAcumulable"),
			req("IngresoNoAcumulable"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Deducciones"}: {
			opc("TotalOtrasDeducciones"),
			opc("TotalImpuestosRetenidos"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Deduccion")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Deduccion"}: {
			req("TipoDeduccion"),
			req("Clave"),
			req("Concepto"),
			req("Importe"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "OtrosPagos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "OtroPago")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "OtroPago"}: {
			req("TipoOtroPago"),
			req("Clave"),
			req("Concepto"),
			req("Importe"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "SubsidioAlEmpleo")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "CompensacionSaldosAFavor")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "SubsidioAlEmpleo"}: {
			req("SubsidioCausado"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "CompensacionSaldosAFavor"}: {
			req("SaldoAFavor"),
			req("Año"),
			req("RemanenteSalFav"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Incapacidades"}: {
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Incapacidad")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Incapacidad"}: {
			req("DiasIncapacidad"),
			req("TipoIncapacidad"),
			opc("ImporteMonetario"),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "ComercioExterior"}: {
			req("Version"),
			opc("MotivoTraslado"),
			req("TipoOperacion"),
			opc("ClaveDePedimento"),
			opc("CertificadoOrigen"),
			opc("NumCertificadoOrigen"),
			opc("NumeroExportadorConfiable"),
			opc("Incoterm"),
			opc("Subdivision"),
			opc("Observaciones"),
			opc("TipoCambioUSD"),
			opc("TotalUSD"),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Emisor"))),
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Propietario")),
				aplicar(rel()),
			),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Receptor"))),
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Destinatario")),
				aplicar(rel()),
			),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Mercancias"))),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Emisor"}: {
			opc("Curp"),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Domicilio"))),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Propietario"}: {
			req("NumRegIdTrib"),
			req("ResidenciaFiscal"),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Receptor"}: {
			opc("NumRegIdTrib"),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Domicilio"))),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Destinatario"}: {
			opc("NumRegIdTrib"),
			opc("Nombre"),
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Domicilio")),
				aplicar(rel()),
			),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Mercancias"}: {
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Mercancia")),
				aplicar(rel()),
			),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Domicilio"}: {
			req("Calle"),
			opc("NumeroExterior"),
			opc("NumeroInterior"),
			opc("Colonia"),
			opc("Localidad"),
			opc("Referencia"),
			opc("Municipio"),
			req("Estado"),
			req("Pais"),
			req("CodigoPostal"),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Mercancia"}: {
			req("NoIdentificacion"),
			opc("FraccionArancelaria"),
			opc("CantidadAduana"),
			opc("UnidadAduana"),
			opc("ValorUnitarioAduana"),
			req("ValorDolares"),
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "DescripcionesEspecificas")),
				aplicar(rel()),
			),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "DescripcionesEspecificas"}: {
			req("Marca"),
			opc("Modelo"),
			opc("SubModelo"),
			opc("NumeroSerie"),
		},
		// cfd/3/cadenaoriginal_3_2/cadenaoriginal_3_2.xslt
		{"http://www.sat.gob.mx/cfd/3", "Comprobante"}: {
			req("version"),
			req("fecha"),
			req("tipoDeComprobante"),
			req("formaDePago"),
			opc("condicionesDePago"),
			req("subTotal"),
			opc("descuento"),
			opc("TipoCambio"),
			opc("Moneda"),
			req("total"),
			req("metodoDePago"),
			req("LugarExpedicion"),
			opc("NumCtaPago"),
			opc("FolioFiscalOrig"),
			opc("SerieFolioFiscalOrig"),
			opc("FechaFolioFiscalOrig"),
			opc("MontoFolioFiscalOrig"),
			aplicar(rel(hijo("http://www.sat.gob.mx/cfd/3", "Emisor"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/cfd/3", "Receptor"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/cfd/3", "Conceptos"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/cfd/3", "Impuestos"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/cfd/3", "Complemento"))),
		},
		// cfd/3/cadenaoriginal_3_2/cadenaoriginal_3_2.xslt
		{"http://www.sat.gob.mx/cfd/3", "Emisor"}: {
			req("rfc"),
			opc("nombre"),
			aplicar(rel(hijo("http://www.sat.gob.mx/cfd/3", "DomicilioFiscal"))),
			si(rel(hijo("http://www.sat.gob.mx/cfd/3", "ExpedidoEn")),
				conjunto(rel(hijo("http://www.sat.gob.mx/cfd/3", "ExpedidoEn")),
					opc("calle"),
					opc("noExterior"),
					opc("noInterior"),
					opc("colonia"),
					opc("localidad"),
					opc("referencia"),
					opc("municipio"),
					opc("estado"),
					req("pais"),
					opc("codigoPostal"),
				),
			),
			cada(rel(hijo("http://www.sat.gob.mx/cfd/3", "RegimenFiscal")),
				req("Regimen"),
			),
		},
		// cfd/3/cadenaoriginal_3_2/cadenaoriginal_3_2.xslt
		{"http://www.sat.gob.mx/cfd/3", "Receptor"}: {
			req("rfc"),
			opc("nombre"),
			si(rel(hijo("http://www.sat.gob.mx/cfd/3", "Domicilio")),
				conjunto(rel(hijo("http://www.sat.gob.mx/cfd/3", "Domicilio")),
					opc("calle"),
					opc("noExterior"),
					opc("noInterior"),
					opc("colonia"),
					opc("localidad"),
					opc("referencia"),
					opc("municipio"),
					opc("estado"),
					req("pais"),
					opc("codigoPostal"),
				),
			),
		},
		// cfd/3/cadenaoriginal_3_2/cadenaoriginal_3_2.xslt
		{"http://www.sat.gob.mx/cfd/3", "Conceptos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/cfd/3", "Concepto")),
				aplicar(rel()),
			),
		},
		// cfd/3/cadenaoriginal_3_2/cadenaoriginal_3_2.xslt
		{"http://www.sat.gob.mx/cfd/3", "Impuestos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/cfd/3", "Retenciones"), hijo("http://www.sat.gob.mx/cfd/3", "Retencion")),
				aplicar(rel()),
			),
			opc("totalImpuestosRetenidos"),
			cada(rel(hijo("http://www.sat.gob.mx/cfd/3", "Traslados"), hijo("http://www.sat.gob.mx/cfd/3", "Traslado")),
				aplicar(rel()),
			),
			opc("totalImpuestosTrasladados"),
		},
		// cfd/3/cadenaoriginal_3_2/cadenaoriginal_3_2.xslt
		{"http://www.sat.gob.mx/cfd/3", "Retencion"}: {
			req("impuesto"),
			req("importe"),
		},
		// cfd/3/cadenaoriginal_3_2/cadenaoriginal_3_2.xslt
		{"http://www.sat.gob.mx/cfd/3", "Traslado"}: {
			req("impuesto"),
			req("tasa"),
			req("importe"),
		},
		// cfd/3/cadenaoriginal_3_2/cadenaoriginal_3_2.xslt
		{"http://www.sat.gob.mx/cfd/3", "Complemento"}: {
			cada(rel(hijo("", "*")),
				aplicar(rel()),
			),
		},
		// cfd/3/cadenaoriginal_3_2/cadenaoriginal_3_2.xslt
		{"http://www.sat.gob.mx/cfd/3", "Concepto"}: {
			req("cantidad"),
			req("unidad"),
			opc("noIdentificacion"),
			req("descripcion"),
			req("valorUnitario"),
			req("importe"),
			cada(rel(desc("http://www.sat.gob.mx/cfd/3", "InformacionAduanera")),
				aplicar(rel()),
			),
			si(rel(hijo("http://www.sat.gob.mx/cfd/3", "CuentaPredial")),
				aplicar(rel(hijo("http://www.sat.gob.mx/cfd/3", "CuentaPredial"))),
			),
			si(rel(hijo("http://www.sat.gob.mx/cfd/3", "ComplementoConcepto")),
				aplicar(rel(hijo("http://www.sat.gob.mx/cfd/3", "ComplementoConcepto"))),
			),
		},
		// cfd/3/cadenaoriginal_3_2/cadenaoriginal_3_2.xslt
		{"http://www.sat.gob.mx/cfd/3", "InformacionAduanera"}: {
			req("numero"),
			req("fecha"),
			opc("aduana"),
		},
		// cfd/3/cadenaoriginal_3_2/cadenaoriginal_3_2.xslt
		{"http://www.sat.gob.mx/cfd/3", "CuentaPredial"}: {
			req("numero"),
		},
		// cfd/3/cadenaoriginal_3_2/cadenaoriginal_3_2.xslt
		{"http://www.sat.gob.mx/cfd/3", "ComplementoConcepto"}: {
			cada(rel(hijo("", "*")),
				aplicar(rel()),
			),
		},
		// cfd/3/cadenaoriginal_3_2/cadenaoriginal_3_2.xslt
		{"http://www.sat.gob.mx/cfd/3", "DomicilioFiscal"}: {
			req("calle"),
			opc("noExterior"),
			opc("noInterior"),
			opc("colonia"),
			opc("localidad"),
			opc("referencia"),
			req("municipio"),
			req("estado"),
			req("pais"),
			req("codigoPostal"),
		},
	},
}

// hojaCFDI33 sale de cfd/3/cadenaoriginal_3_3/cadenaoriginal_3_3.xslt y de lo que incluye.
//
// Archivos (sha256):
//   - cfd/3/cadenaoriginal_3_3/cadenaoriginal_3_3.xslt fe592d28936b393c5432908bd5245ac02043319daec0cb7df15c963ad936bd7d
//   - cfd/2/cadenaoriginal_2_0/utilerias.xslt d9ef68267747e0f6725295d4012684a4c051cb3f2265dcb2035d32bf40895016
//   - cfd/EstadoDeCuentaCombustible/ecc11.xslt d9bb8d26d511136e68ca8b986c7fedcb946aaffe0a29d986c30c30970ef219d6
//   - cfd/donat/donat11.xslt 1c72d07d9945fd7824471caaebaee47dafc3b56785a037549fcbee5e123c6d36
//   - cfd/divisas/divisas.xslt be7167e1d421b66ba17a6f1788c964c5bb1f207ed7ad9803166d60966a332b1d
//   - cfd/implocal/implocal.xslt 3c2ee6c3ecce9b2e74b2a225f3c17250259751a1307f58a96799ff119e2fafe6
//   - cfd/leyendasFiscales/leyendasFisc.xslt 68b0016e44c1829a5c7d5da340b69945100d8fda1243e4e279c2cdc8c745d1fb
//   - cfd/pfic/pfic.xslt 0e9e992e4e0fa6e13ee7b18ce5ddc51c55f3615236c7c51b80426811f9779bab
//   - cfd/TuristaPasajeroExtranjero/TuristaPasajeroExtranjero.xslt 06aee48f06ed53ab4fad9f2ddc950bca51840cfd4fb137ec8d14632e2c2caeb3
//   - cfd/nomina/nomina12.xslt a14cdb44995526aa5888036bedf7fac9f84ad0041f9393371b0bbbea0572649e
//   - cfd/cfdiregistrofiscal/cfdiregistrofiscal.xslt d824c4993e1f5940ff0605bdd3ea040669bef69355b551ab9da916f3a3746e7d
//   - cfd/pagoenespecie/pagoenespecie.xslt ed33ce278525a43db377d82fe641d42391d74b731f82015b1fedfa84654f1798
//   - cfd/aerolineas/aerolineas.xslt d000d0ed64a45a05a7f59b1eb2399d55d1588f18056ed37e8bc642ca2fe655c0
//   - cfd/valesdedespensa/valesdedespensa.xslt e22cf591f5669b78b3306c33e2434acbf7f40c3671e318c3fb2b736a2bf6ef8c
//   - cfd/consumodecombustibles/consumodecombustibles.xslt e1141cfa02dc67de719f9e31c6abc997c19831715544a2d525ad14dadde4b623
//   - cfd/notariospublicos/notariospublicos.xslt b3aeebdbfa1858b60d9f8b437f9f07a95f600e712b52aa0a8c08a7f924fe9ce7
//   - cfd/vehiculousado/vehiculousado.xslt 226664fcdcbc51e57abb260104b59a06d3bba62a62ae92fbbfeee6f16791ae2c
//   - cfd/servicioparcialconstruccion/servicioparcialconstruccion.xslt b3b9a8983a673a352ac291c1e0a9605e983dc0af8a4ecb4f6ec0cb13e258fe85
//   - cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt 347bf6dfdae5c0b18d6a107c8c339f212763496fa4ae1fa748d0354a7c3338d2
//   - cfd/certificadodestruccion/certificadodedestruccion.xslt 0b306d8142eb118f47d3ab6c967a6be6c9c31436678e776e571ce496018b68d5
//   - cfd/arteantiguedades/obrasarteantiguedades.xslt 2d11a854ff914b229ef7d2d43cdfd6296315d040c36e0ad6c721ace08ba89d1d
//   - cfd/ComercioExterior11/ComercioExterior11.xslt bb9e7c3d2a2a3e8af9c65498fd1422b0f400e0787d2418b460d51ece6c988f9d
//   - cfd/ine/ine11.xslt 02781b96a0da332693f719f893f7a72117196b073c5d0df48e7fa6a4ed444882
//   - cfd/iedu/iedu.xslt 2e9ddb05e5cd222cded57a5306b43f954f15a2ac9f8ddf3e4818156cd703a49d
//   - cfd/ventavehiculos/ventavehiculos11.xslt 95fd6522700a5125dacc363e07fac8ff2f470b2e23aec5bbe3d3f684c8a4152f
//   - cfd/terceros/terceros11.xslt 62b8986466db9b1032462bfbe0cf3514caa8d5e1b8b5906caeab44179853996a
//   - cfd/Pagos/Pagos10.xslt 65337d977e4697733b7cf9ddcc1dcececb485a7e43ffcca149b706e5fa60f5ec
//   - cfd/EstadoDeCuentaCombustible/ecc12.xslt 9ce3097989da2560e3fcb6e45928cd7d90692832b2d29c0c4127732349f3dba5
//   - cfd/consumodecombustibles/consumodeCombustibles11.xslt 610de8017f569d6e66669fc1f8b2303c5b4ac970e9b17e44bcd5d1b0e4ef7bf1
//   - cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt 1d735ccdd03040a19469795e6703fbd2c56a2c08c708111278af6d885e98621c
//   - cfd/IngresosHidrocarburos10/IngresosHidrocarburos.xslt 2d17335608f036214129210b06da1a132ec0fc06b315f97dd3e07d09641303d9
//   - cfd/CartaPorte/CartaPorte.xslt 32fec3089d94c1025a2e7112eda480d2b49167bf9b71335d0434c33c012b4829
//   - cfd/CartaPorte/CartaPorte20.xslt 0e8548e036920208c8e909acd64f6e7d4de6742e75c28617c7ac8edb97c55766
//
// Includes omitidos (complementos que la librería no lee):
//   - cfd/detallista/detallista.xslt
//
// Plantillas match="/" de hojas incluidas, que la de la raíz sustituye:
//   - cfd/vehiculousado/vehiculousado.xslt
var hojaCFDI33 = &hojaCadena{
	nombre: "cadenaoriginal_3_3",
	raiz: []pasoCadena{
		txt("|"),
		aplicar(abs(hijo("http://www.sat.gob.mx/cfd/3", "Comprobante"))),
		txt("||"),
	},
	plantillas: map[nombreXML][]pasoCadena{
		// cfd/EstadoDeCuentaCombustible/ecc11.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible", "EstadoDeCuentaCombustible"}: {
			req("Version"),
			req("TipoOperacion"),
			req("NumeroDeCuenta"),
			req("SubTotal"),
			req("Total"),
			aplicar(rel(hijo("http://www.sat.gob.mx/EstadoDeCuentaCombustible", "Conceptos"))),
		},
		// cfd/EstadoDeCuentaCombustible/ecc11.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible", "Conceptos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/EstadoDeCuentaCombustible", "ConceptoEstadoDeCuentaCombustible")),
				aplicar(rel()),
			),
		},
		// cfd/EstadoDeCuentaCombustible/ecc11.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible", "Traslados"}: {
			cada(rel(hijo("http://www.sat.gob.mx/EstadoDeCuentaCombustible", "Traslado")),
				aplicar(rel()),
			),
		},
		// cfd/EstadoDeCuentaCombustible/ecc11.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible", "ConceptoEstadoDeCuentaCombustible"}: {
			req("Identificador"),
			req("Fecha"),
			req("Rfc"),
			req("ClaveEstacion"),
			opc("TAR"),
			req("Cantidad"),
			req("NoIdentificacion"),
			opc("Unidad"),
			req("NombreCombustible"),
			req("FolioOperacion"),
			req("ValorUnitario"),
			req("Importe"),
			aplicar(rel(hijo("http://www.sat.gob.mx/EstadoDeCuentaCombustible", "Traslados"))),
		},
		// cfd/EstadoDeCuentaCombustible/ecc11.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible", "Traslado"}: {
			req("Impuesto"),
			req("TasaoCuota"),
			req("Importe"),
		},
		// cfd/donat/donat11.xslt
		{"http://www.sat.gob.mx/donat", "Donatarias"}: {
			req("version"),
			req("noAutorizacion"),
			req("fechaAutorizacion"),
			req("leyenda"),
		},
		// cfd/divisas/divisas.xslt
		{"http://www.sat.gob.mx/divisas", "Divisas"}: {
			req("version"),
			req("tipoOperacion"),
		},
		// cfd/implocal/implocal.xslt
		{"http://www.sat.gob.mx/implocal", "ImpuestosLocales"}: {
			req("version"),
			req("TotaldeRetenciones"),
			req("TotaldeTraslados"),
			cada(rel(hijo("http://www.sat.gob.mx/implocal", "RetencionesLocales")),
				req("ImpLocRetenido"),
				req("TasadeRetencion"),
				req("Importe"),
			),
			cada(rel(hijo("http://www.sat.gob.mx/implocal", "TrasladosLocales")),
				req("ImpLocTrasladado"),
				req("TasadeTraslado"),
				req("Importe"),
			),
		},
		// cfd/leyendasFiscales/leyendasFisc.xslt
		{"http://www.sat.gob.mx/leyendasFiscales", "LeyendasFiscales"}: {
			req("version"),
			cada(rel(hijo("http://www.sat.gob.mx/leyendasFiscales", "Leyenda")),
				aplicar(rel()),
			),
		},
		// cfd/leyendasFiscales/leyendasFisc.xslt
		{"http://www.sat.gob.mx/leyendasFiscales", "Leyenda"}: {
			opc("disposicionFiscal"),
			opc("norma"),
			req("textoLeyenda"),
		},
		// cfd/pfic/pfic.xslt
		{"http://www.sat.gob.mx/pfic", "PFintegranteCoordinado"}: {
			req("version"),
			req("ClaveVehicular"),
			req("Placa"),
			opc("RFCPF"),
		},
		// cfd/TuristaPasajeroExtranjero/TuristaPasajeroExtranjero.xslt
		{"http://www.sat.gob.mx/TuristaPasajeroExtranjero", "TuristaPasajeroExtranjero"}: {
			req("version"),
			req("fechadeTransito"),
			req("tipoTransito"),
			aplicar(rel(hijo("http://www.sat.gob.mx/TuristaPasajeroExtranjero", "datosTransito"))),
		},
		// cfd/TuristaPasajeroExtranjero/TuristaPasajeroExtranjero.xslt
		{"http://www.sat.gob.mx/TuristaPasajeroExtranjero", "datosTransito"}: {
			req("Via"),
			req("TipoId"),
			req("NumeroId"),
			req("Nacionalidad"),
			req("EmpresaTransporte"),
			opc("IdTransporte"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Nomina"}: {
			req("Version"),
			req("TipoNomina"),
			req("FechaPago"),
			req("FechaInicialPago"),
			req("FechaFinalPago"),
			req("NumDiasPagados"),
			opc("TotalPercepciones"),
			opc("TotalDeducciones"),
			opc("TotalOtrosPagos"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Emisor")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Receptor")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Percepciones")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Deducciones")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "OtrosPagos")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Incapacidades")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Emisor"}: {
			opc("Curp"),
			opc("RegistroPatronal"),
			opc("RfcPatronOrigen"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "EntidadSNCF")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "EntidadSNCF"}: {
			req("OrigenRecurso"),
			opc("MontoRecursoPropio"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Receptor"}: {
			req("Curp"),
			opc("NumSeguridadSocial"),
			opc("FechaInicioRelLaboral"),
			opc("Antigüedad"),
			req("TipoContrato"),
			opc("Sindicalizado"),
			opc("TipoJornada"),
			req("TipoRegimen"),
			req("NumEmpleado"),
			opc("Departamento"),
			opc("Puesto"),
			opc("RiesgoPuesto"),
			req("PeriodicidadPago"),
			opc("Banco"),
			opc("CuentaBancaria"),
			opc("SalarioBaseCotApor"),
			opc("SalarioDiarioIntegrado"),
			req("ClaveEntFed"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "SubContratacion")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "SubContratacion"}: {
			req("RfcLabora"),
			req("PorcentajeTiempo"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Percepciones"}: {
			opc("TotalSueldos"),
			opc("TotalSeparacionIndemnizacion"),
			opc("TotalJubilacionPensionRetiro"),
			req("TotalGravado"),
			req("TotalExento"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Percepcion")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "JubilacionPensionRetiro")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "SeparacionIndemnizacion")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Percepcion"}: {
			req("TipoPercepcion"),
			req("Clave"),
			req("Concepto"),
			req("ImporteGravado"),
			req("ImporteExento"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "AccionesOTitulos")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "HorasExtra")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "AccionesOTitulos"}: {
			req("ValorMercado"),
			req("PrecioAlOtorgarse"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "HorasExtra"}: {
			req("Dias"),
			req("TipoHoras"),
			req("HorasExtra"),
			req("ImportePagado"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "JubilacionPensionRetiro"}: {
			opc("TotalUnaExhibicion"),
			opc("TotalParcialidad"),
			opc("MontoDiario"),
			req("IngresoAcumulable"),
			req("IngresoNoAcumulable"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "SeparacionIndemnizacion"}: {
			req("TotalPagado"),
			req("NumAñosServicio"),
			req("UltimoSueldoMensOrd"),
			req("IngresoAcumulable"),
			req("IngresoNoAcumulable"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Deducciones"}: {
			opc("TotalOtrasDeducciones"),
			opc("TotalImpuestosRetenidos"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Deduccion")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Deduccion"}: {
			req("TipoDeduccion"),
			req("Clave"),
			req("Concepto"),
			req("Importe"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "OtrosPagos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "OtroPago")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "OtroPago"}: {
			req("TipoOtroPago"),
			req("Clave"),
			req("Concepto"),
			req("Importe"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "SubsidioAlEmpleo")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "CompensacionSaldosAFavor")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "SubsidioAlEmpleo"}: {
			req("SubsidioCausado"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "CompensacionSaldosAFavor"}: {
			req("SaldoAFavor"),
			req("Año"),
			req("RemanenteSalFav"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Incapacidades"}: {
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Incapacidad")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Incapacidad"}: {
			req("DiasIncapacidad"),
			req("TipoIncapacidad"),
			opc("ImporteMonetario"),
		},
		// cfd/cfdiregistrofiscal/cfdiregistrofiscal.xslt
		{"http://www.sat.gob.mx/registrofiscal", "CFDIRegistroFiscal"}: {
			req("Version"),
			req("Folio"),
		},
		// cfd/pagoenespecie/pagoenespecie.xslt
		{"http://www.sat.gob.mx/pagoenespecie", "PagoEnEspecie"}: {
			req("Version"),
			req("CvePIC"),
			req("FolioSolDon"),
			req("PzaArtNombre"),
			req("PzaArtTecn"),
			req("PzaArtAProd"),
			req("PzaArtDim"),
		},
		// cfd/aerolineas/aerolineas.xslt
		{"http://www.sat.gob.mx/aerolineas", "Aerolineas"}: {
			req("Version"),
			req("TUA"),
			aplicar(rel(hijo("http://www.sat.gob.mx/aerolineas", "OtrosCargos"))),
		},
		// cfd/aerolineas/aerolineas.xslt
		{"http://www.sat.gob.mx/aerolineas", "OtrosCargos"}: {
			req("TotalCargos"),
			cada(rel(hijo("http://www.sat.gob.mx/aerolineas", "Cargo")),
				aplicar(rel()),
			),
		},
		// cfd/aerolineas/aerolineas.xslt
		{"http://www.sat.gob.mx/aerolineas", "Cargo"}: {
			req("CodigoCargo"),
			req("Importe"),
		},
		// cfd/valesdedespensa/valesdedespensa.xslt
		{"http://www.sat.gob.mx/valesdedespensa", "ValesDeDespensa"}: {
			req("version"),
			req("tipoOperacion"),
			opc("registroPatronal"),
			req("numeroDeCuenta"),
			req("total"),
			aplicar(rel(hijo("http://www.sat.gob.mx/valesdedespensa", "Conceptos"))),
		},
		// cfd/valesdedespensa/valesdedespensa.xslt
		{"http://www.sat.gob.mx/valesdedespensa", "Conceptos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/valesdedespensa", "Concepto")),
				aplicar(rel()),
			),
		},
		// cfd/valesdedespensa/valesdedespensa.xslt
		{"http://www.sat.gob.mx/valesdedespensa", "Concepto"}: {
			req("identificador"),
			req("fecha"),
			req("rfc"),
			req("curp"),
			req("nombre"),
			opc("numSeguridadSocial"),
			req("importe"),
		},
		// cfd/consumodecombustibles/consumodecombustibles.xslt
		{"http://www.sat.gob.mx/consumodecombustibles", "ConsumoDeCombustibles"}: {
			req("version"),
			req("tipoOperacion"),
			req("numeroDeCuenta"),
			opc("subTotal"),
			req("total"),
			aplicar(rel(hijo("http://www.sat.gob.mx/consumodecombustibles", "Conceptos"))),
		},
		// cfd/consumodecombustibles/consumodecombustibles.xslt
		{"http://www.sat.gob.mx/consumodecombustibles", "Conceptos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/consumodecombustibles", "ConceptoConsumoDeCombustibles")),
				aplicar(rel()),
			),
		},
		// cfd/consumodecombustibles/consumodecombustibles.xslt
		{"http://www.sat.gob.mx/consumodecombustibles", "ConceptoConsumoDeCombustibles"}: {
			req("identificador"),
			req("fecha"),
			req("rfc"),
			req("claveEstacion"),
			req("cantidad"),
			req("nombreCombustible"),
			req("folioOperacion"),
			req("valorUnitario"),
			req("importe"),
			cada(rel(hijo("http://www.sat.gob.mx/consumodecombustibles", "Determinados")),
				aplicar(rel()),
			),
		},
		// cfd/consumodecombustibles/consumodecombustibles.xslt
		{"http://www.sat.gob.mx/consumodecombustibles", "Determinados"}: {
			cada(rel(hijo("http://www.sat.gob.mx/consumodecombustibles", "Determinado")),
				aplicar(rel()),
			),
		},
		// cfd/consumodecombustibles/consumodecombustibles.xslt
		{"http://www.sat.gob.mx/consumodecombustibles", "Determinado"}: {
			req("impuesto"),
			req("tasa"),
			req("importe"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "NotariosPublicos"}: {
			req("Version"),
			aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DescInmuebles"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosOperacion"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosNotario"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosEnajenante"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosAdquiriente"))),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DescInmuebles"}: {
			cada(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DescInmueble")),
				aplicar(rel()),
			),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DescInmueble"}: {
			req("TipoInmueble"),
			req("Calle"),
			opc("NoExterior"),
			opc("NoInterior"),
			opc("Colonia"),
			opc("Localidad"),
			opc("Referencia"),
			req("Municipio"),
			req("Estado"),
			req("Pais"),
			req("CodigoPostal"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosOperacion"}: {
			req("NumInstrumentoNotarial"),
			req("FechaInstNotarial"),
			req("MontoOperacion"),
			req("Subtotal"),
			req("IVA"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosNotario"}: {
			req("CURP"),
			req("NumNotaria"),
			req("EntidadFederativa"),
			opc("Adscripcion"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosEnajenante"}: {
			req("CoproSocConyugalE"),
			si(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosUnEnajenante")),
				aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosUnEnajenante"))),
			),
			si(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosEnajenantesCopSC")),
				aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosEnajenantesCopSC"))),
			),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosUnEnajenante"}: {
			req("Nombre"),
			req("ApellidoPaterno"),
			opc("ApellidoMaterno"),
			req("RFC"),
			req("CURP"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosEnajenantesCopSC"}: {
			cada(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosEnajenanteCopSC")),
				aplicar(rel()),
			),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosEnajenanteCopSC"}: {
			req("Nombre"),
			opc("ApellidoPaterno"),
			opc("ApellidoMaterno"),
			req("RFC"),
			opc("CURP"),
			req("Porcentaje"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosAdquiriente"}: {
			req("CoproSocConyugalE"),
			si(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosUnAdquiriente")),
				aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosUnAdquiriente"))),
			),
			si(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosAdquirientesCopSC")),
				aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosAdquirientesCopSC"))),
			),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosUnAdquiriente"}: {
			req("Nombre"),
			opc("ApellidoPaterno"),
			opc("ApellidoMaterno"),
			req("RFC"),
			opc("CURP"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosAdquirientesCopSC"}: {
			cada(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosAdquirienteCopSC")),
				aplicar(rel()),
			),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosAdquirienteCopSC"}: {
			req("Nombre"),
			opc("ApellidoPaterno"),
			opc("ApellidoMaterno"),
			req("RFC"),
			opc("CURP"),
			req("Porcentaje"),
		},
		// cfd/vehiculousado/vehiculousado.xslt
		{"http://www.sat.gob.mx/vehiculousado", "VehiculoUsado"}: {
			req("Version"),
			req("montoAdquisicion"),
			req("montoEnajenacion"),
			req("claveVehicular"),
			req("marca"),
			req("tipo"),
			req("modelo"),
			opc("numeroMotor"),
			opc("numeroSerie"),
			opc("NIV"),
			req("valor"),
			aplicar(rel(hijo("http://www.sat.gob.mx/vehiculousado", "InformacionAduanera"))),
		},
		// cfd/vehiculousado/vehiculousado.xslt
		{"http://www.sat.gob.mx/vehiculousado", "InformacionAduanera"}: {
			req("numero"),
			req("fecha"),
			opc("aduana"),
		},
		// cfd/servicioparcialconstruccion/servicioparcialconstruccion.xslt
		{"http://www.sat.gob.mx/servicioparcialconstruccion", "parcialesconstruccion"}: {
			req("Version"),
			req("NumPerLicoAut"),
			aplicar(rel(hijo("http://www.sat.gob.mx/servicioparcialconstruccion", "Inmueble"))),
		},
		// cfd/servicioparcialconstruccion/servicioparcialconstruccion.xslt
		{"http://www.sat.gob.mx/servicioparcialconstruccion", "Inmueble"}: {
			req("Calle"),
			opc("NoExterior"),
			opc("NoInterior"),
			opc("Colonia"),
			opc("Localidad"),
			opc("Referencia"),
			req("Municipio"),
			req("Estado"),
			req("CodigoPostal"),
		},
		// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt
		{"http://www.sat.gob.mx/renovacionysustitucionvehiculos", "renovacionysustitucionvehiculos"}: {
			req("Version"),
			req("TipoDeDecreto"),
			aplicar(rel(hijo("http://www.sat.gob.mx/renovacionysustitucionvehiculos", "DecretoRenovVehicular"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/renovacionysustitucionvehiculos", "DecretoSustitVehicular"))),
		},
		// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt
		{"http://www.sat.gob.mx/renovacionysustitucionvehiculos", "DecretoRenovVehicular"}: {
			req("VehEnaj"),
			cada(rel(hijo("http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculosUsadosEnajenadoPermAlFab")),
				aplicar(rel()),
			),
			aplicar(rel(hijo("http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculoNuvoSemEnajenadoFabAlPerm"))),
		},
		// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt
		{"http://www.sat.gob.mx/renovacionysustitucionvehiculos", "DecretoSustitVehicular"}: {
			req("VehEnaj"),
			aplicar(rel(hijo("http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculoUsadoEnajenadoPermAlFab"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculoNuvoSemEnajenadoFabAlPerm"))),
		},
		// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt
		{"http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculosUsadosEnajenadoPermAlFab"}: {
			req("PrecioVehUsado"),
			req("TipoVeh"),
			req("Marca"),
			req("TipooClase"),
			req("Año"),
			opc("Modelo"),
			opc("NIV"),
			opc("NumSerie"),
			req("NumPlacas"),
			opc("NumMotor"),
			req("NumFolTarjCir"),
			opc("NumPedIm"),
			opc("Aduana"),
			opc("FechaRegulVeh"),
			req("Foliofiscal"),
		},
		// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt
		{"http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculoNuvoSemEnajenadoFabAlPerm"}: {
			req("Año"),
			opc("Modelo"),
			req("NumPlacas"),
			opc("RFC"),
		},
		// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt
		{"http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculoUsadoEnajenadoPermAlFab"}: {
			req("PrecioVehUsado"),
			req("TipoVeh"),
			req("Marca"),
			req("TipooClase"),
			req("Año"),
			opc("Modelo"),
			opc("NIV"),
			opc("NumSerie"),
			req("NumPlacas"),
			opc("NumMotor"),
			req("NumFolTarjCir"),
			req("NumFolAvisoint"),
			req("NumPedIm"),
			req("Aduana"),
			req("FechaRegulVeh"),
			req("Foliofiscal"),
		},
		// cfd/certificadodestruccion/certificadodedestruccion.xslt
		{"http://www.sat.gob.mx/certificadodestruccion", "certificadodedestruccion"}: {
			req("Version"),
			req("Serie"),
			req("NumFolDesVeh"),
			aplicar(rel(hijo("http://www.sat.gob.mx/certificadodestruccion", "VehiculoDestruido"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/certificadodestruccion", "InformacionAduanera"))),
		},
		// cfd/certificadodestruccion/certificadodedestruccion.xslt
		{"http://www.sat.gob.mx/certificadodestruccion", "VehiculoDestruido"}: {
			req("Marca"),
			req("TipooClase"),
			req("Año"),
			opc("Modelo"),
			opc("NIV"),
			opc("NumSerie"),
			req("NumPlacas"),
			opc("NumMotor"),
			req("NumFolTarjCir"),
		},
		// cfd/certificadodestruccion/certificadodedestruccion.xslt
		{"http://www.sat.gob.mx/certificadodestruccion", "InformacionAduanera"}: {
			req("NumPedImp"),
			req("Fecha"),
			req("Aduana"),
		},
		// cfd/arteantiguedades/obrasarteantiguedades.xslt
		{"http://www.sat.gob.mx/arteantiguedades", "obrasarteantiguedades"}: {
			req("Version"),
			req("TipoBien"),
			opc("OtrosTipoBien"),
			req("TituloAdquirido"),
			opc("OtrosTituloAdquirido"),
			opc("Subtotal"),
			opc("IVA"),
			req("FechaAdquisicion"),
			req("CaracterísticasDeObraoPieza"),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "ComercioExterior"}: {
			req("Version"),
			opc("MotivoTraslado"),
			req("TipoOperacion"),
			opc("ClaveDePedimento"),
			opc("CertificadoOrigen"),
			opc("NumCertificadoOrigen"),
			opc("NumeroExportadorConfiable"),
			opc("Incoterm"),
			opc("Subdivision"),
			opc("Observaciones"),
			opc("TipoCambioUSD"),
			opc("TotalUSD"),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Emisor"))),
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Propietario")),
				aplicar(rel()),
			),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Receptor"))),
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Destinatario")),
				aplicar(rel()),
			),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Mercancias"))),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Emisor"}: {
			opc("Curp"),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Domicilio"))),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Propietario"}: {
			req("NumRegIdTrib"),
			req("ResidenciaFiscal"),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Receptor"}: {
			opc("NumRegIdTrib"),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Domicilio"))),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Destinatario"}: {
			opc("NumRegIdTrib"),
			opc("Nombre"),
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Domicilio")),
				aplicar(rel()),
			),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Mercancias"}: {
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Mercancia")),
				aplicar(rel()),
			),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Domicilio"}: {
			req("Calle"),
			opc("NumeroExterior"),
			opc("NumeroInterior"),
			opc("Colonia"),
			opc("Localidad"),
			opc("Referencia"),
			opc("Municipio"),
			req("Estado"),
			req("Pais"),
			req("CodigoPostal"),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Mercancia"}: {
			req("NoIdentificacion"),
			opc("FraccionArancelaria"),
			opc("CantidadAduana"),
			opc("UnidadAduana"),
			opc("ValorUnitarioAduana"),
			req("ValorDolares"),
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "DescripcionesEspecificas")),
				aplicar(rel()),
			),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "DescripcionesEspecificas"}: {
			req("Marca"),
			opc("Modelo"),
			opc("SubModelo"),
			opc("NumeroSerie"),
		},
		// cfd/ine/ine11.xslt
		{"http://www.sat.gob.mx/ine", "INE"}: {
			req("Version"),
			req("TipoProceso"),
			opc("TipoComite"),
			opc("IdContabilidad"),
			cada(rel(hijo("http://www.sat.gob.mx/ine", "Entidad")),
				aplicar(rel()),
			),
		},
		// cfd/ine/ine11.xslt
		{"http://www.sat.gob.mx/ine", "Entidad"}: {
			req("ClaveEntidad"),
			opc("Ambito"),
			cada(rel(hijo("http://www.sat.gob.mx/ine", "Contabilidad")),
				aplicar(rel()),
			),
		},
		// cfd/ine/ine11.xslt
		{"http://www.sat.gob.mx/ine", "Contabilidad"}: {
			req("IdContabilidad"),
		},
		// cfd/iedu/iedu.xslt
		{"http://www.sat.gob.mx/iedu", "instEducativas"}: {
			req("version"),
			req("nombreAlumno"),
			req("CURP"),
			req("nivelEducativo"),
			req("autRVOE"),
			opc("rfcPago"),
		},
		// cfd/ventavehiculos/ventavehiculos11.xslt
		{"http://www.sat.gob.mx/ventavehiculos", "VentaVehiculos"}: {
			req("version"),
			req("ClaveVehicular"),
			siIgual("version", "1.1",
				req("Niv"),
			),
			cada(rel(desc("http://www.sat.gob.mx/ventavehiculos", "InformacionAduanera")),
				aplicar(rel()),
			),
		},
		// cfd/ventavehiculos/ventavehiculos11.xslt
		{"http://www.sat.gob.mx/ventavehiculos", "InformacionAduanera"}: {
			req("numero"),
			req("fecha"),
			opc("aduana"),
		},
		// cfd/terceros/terceros11.xslt
		{"http://www.sat.gob.mx/terceros", "PorCuentadeTerceros"}: {
			req("version"),
			req("rfc"),
			opc("nombre"),
			aplicar(rel(desc("http://www.sat.gob.mx/terceros", "InformacionFiscalTercero"))),
			cada(rel(desc("http://www.sat.gob.mx/terceros", "InformacionAduanera")),
				aplicar(rel()),
			),
			si(rel(hijo("http://www.sat.gob.mx/terceros", "CuentaPredial")),
				aplicar(rel(hijo("http://www.sat.gob.mx/terceros", "CuentaPredial"))),
			),
			cada(rel(desc("http://www.sat.gob.mx/terceros", "Retenciones"), hijo("http://www.sat.gob.mx/terceros", "Retencion")),
				aplicar(rel()),
			),
			cada(rel(desc("http://www.sat.gob.mx/terceros", "Traslados"), hijo("http://www.sat.gob.mx/terceros", "Traslado")),
				aplicar(rel()),
			),
		},
		// cfd/terceros/terceros11.xslt
		{"http://www.sat.gob.mx/terceros", "Retencion"}: {
			req("impuesto"),
			req("importe"),
		},
		// cfd/terceros/terceros11.xslt
		{"http://www.sat.gob.mx/terceros", "Traslado"}: {
			req("impuesto"),
			req("tasa"),
			req("importe"),
		},
		// cfd/terceros/terceros11.xslt
		{"http://www.sat.gob.mx/terceros", "InformacionAduanera"}: {
			req("numero"),
			req("fecha"),
			opc("aduana"),
		},
		// cfd/terceros/terceros11.xslt
		{"http://www.sat.gob.mx/terceros", "CuentaPredial"}: {
			req("numero"),
		},
		// cfd/terceros/terceros11.xslt
		{"http://www.sat.gob.mx/terceros", "InformacionFiscalTercero"}: {
			req("calle"),
			opc("noExterior"),
			opc("noInterior"),
			opc("colonia"),
			opc("localidad"),
			opc("referencia"),
			req("municipio"),
			req("estado"),
			req("pais"),
			req("codigoPostal"),
		},
		// cfd/Pagos/Pagos10.xslt
		{"http://www.sat.gob.mx/Pagos", "Pagos"}: {
			req("Version"),
			cada(rel(hijo("http://www.sat.gob.mx/Pagos", "Pago")),
				aplicar(rel()),
			),
		},
		// cfd/Pagos/Pagos10.xslt
		{"http://www.sat.gob.mx/Pagos", "Pago"}: {
			req("FechaPago"),
			req("FormaDePagoP"),
			req("MonedaP"),
			opc("TipoCambioP"),
			req("Monto"),
			opc("NumOperacion"),
			opc("RfcEmisorCtaOrd"),
			opc("NomBancoOrdExt"),
			opc("CtaOrdenante"),
			opc("RfcEmisorCtaBen"),
			opc("CtaBeneficiario"),
			opc("TipoCadPago"),
			opc("CertPago"),
			opc("CadPago"),
			opc("SelloPago"),
			cada(rel(hijo("http://www.sat.gob.mx/Pagos", "DoctoRelacionado")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/Pagos", "Impuestos")),
				aplicar(rel()),
			),
		},
		// cfd/Pagos/Pagos10.xslt
		{"http://www.sat.gob.mx/Pagos", "DoctoRelacionado"}: {
			req("IdDocumento"),
			opc("Serie"),
			opc("Folio"),
			req("MonedaDR"),
			opc("TipoCambioDR"),
			req("MetodoDePagoDR"),
			opc("NumParcialidad"),
			opc("ImpSaldoAnt"),
			opc("ImpPagado"),
			opc("ImpSaldoInsoluto"),
		},
		// cfd/Pagos/Pagos10.xslt
		{"http://www.sat.gob.mx/Pagos", "Impuestos"}: {
			opc("TotalImpuestosRetenidos"),
			opc("TotalImpuestosTrasladados"),
			aplicar(rel(hijo("http://www.sat.gob.mx/Pagos", "Retenciones"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/Pagos", "Traslados"))),
		},
		// cfd/Pagos/Pagos10.xslt
		{"http://www.sat.gob.mx/Pagos", "Retenciones"}: {
			cada(rel(hijo("http://www.sat.gob.mx/Pagos", "Retencion")),
				aplicar(rel()),
			),
		},
		// cfd/Pagos/Pagos10.xslt
		{"http://www.sat.gob.mx/Pagos", "Traslados"}: {
			cada(rel(hijo("http://www.sat.gob.mx/Pagos", "Traslado")),
				aplicar(rel()),
			),
		},
		// cfd/Pagos/Pagos10.xslt
		{"http://www.sat.gob.mx/Pagos", "Retencion"}: {
			req("Impuesto"),
			req("Importe"),
		},
		// cfd/Pagos/Pagos10.xslt
		{"http://www.sat.gob.mx/Pagos", "Traslado"}: {
			req("Impuesto"),
			req("TipoFactor"),
			req("TasaOCuota"),
			req("Importe"),
		},
		// cfd/EstadoDeCuentaCombustible/ecc12.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible12", "EstadoDeCuentaCombustible"}: {
			req("Version"),
			req("TipoOperacion"),
			req("NumeroDeCuenta"),
			req("SubTotal"),
			req("Total"),
			aplicar(rel(hijo("http://www.sat.gob.mx/EstadoDeCuentaCombustible12", "Conceptos"))),
		},
		// cfd/EstadoDeCuentaCombustible/ecc12.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible12", "Conceptos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/EstadoDeCuentaCombustible12", "ConceptoEstadoDeCuentaCombustible")),
				aplicar(rel()),
			),
		},
		// cfd/EstadoDeCuentaCombustible/ecc12.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible12", "Traslados"}: {
			cada(rel(hijo("http://www.sat.gob.mx/EstadoDeCuentaCombustible12", "Traslado")),
				aplicar(rel()),
			),
		},
		// cfd/EstadoDeCuentaCombustible/ecc12.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible12", "ConceptoEstadoDeCuentaCombustible"}: {
			req("Identificador"),
			req("Fecha"),
			req("Rfc"),
			req("ClaveEstacion"),
			req("Cantidad"),
			req("TipoCombustible"),
			opc("Unidad"),
			req("NombreCombustible"),
			req("FolioOperacion"),
			req("ValorUnitario"),
			req("Importe"),
			aplicar(rel(hijo("http://www.sat.gob.mx/EstadoDeCuentaCombustible12", "Traslados"))),
		},
		// cfd/EstadoDeCuentaCombustible/ecc12.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible12", "Traslado"}: {
			req("Impuesto"),
			req("TasaOCuota"),
			req("Importe"),
		},
		// cfd/consumodecombustibles/consumodeCombustibles11.xslt
		{"http://www.sat.gob.mx/ConsumoDeCombustibles11", "ConsumoDeCombustibles"}: {
			req("version"),
			req("tipoOperacion"),
			req("numeroDeCuenta"),
			opc("subTotal"),
			req("total"),
			aplicar(rel(hijo("http://www.sat.gob.mx/ConsumoDeCombustibles11", "Conceptos"))),
		},
		// cfd/consumodecombustibles/consumodeCombustibles11.xslt
		{"http://www.sat.gob.mx/ConsumoDeCombustibles11", "Conceptos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/ConsumoDeCombustibles11", "ConceptoConsumoDeCombustibles")),
				aplicar(rel()),
			),
		},
		// cfd/consumodecombustibles/consumodeCombustibles11.xslt
		{"http://www.sat.gob.mx/ConsumoDeCombustibles11", "ConceptoConsumoDeCombustibles"}: {
			req("identificador"),
			req("fecha"),
			req("rfc"),
			req("claveEstacion"),
			req("tipoCombustible"),
			req("cantidad"),
			req("nombreCombustible"),
			req("folioOperacion"),
			req("valorUnitario"),
			req("importe"),
			aplicar(rel(hijo("http://www.sat.gob.mx/ConsumoDeCombustibles11", "Determinados"))),
		},
		// cfd/consumodecombustibles/consumodeCombustibles11.xslt
		{"http://www.sat.gob.mx/ConsumoDeCombustibles11", "Determinados"}: {
			cada(rel(hijo("http://www.sat.gob.mx/ConsumoDeCombustibles11", "Determinado")),
				aplicar(rel()),
			),
		},
		// cfd/consumodecombustibles/consumodeCombustibles11.xslt
		{"http://www.sat.gob.mx/ConsumoDeCombustibles11", "Determinado"}: {
			req("impuesto"),
			req("tasaOCuota"),
			req("importe"),
		},
		// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt
		{"http://www.sat.gob.mx/GastosHidrocarburos10", "GastosHidrocarburos"}: {
			req("Version"),
			req("NumeroContrato"),
			opc("AreaContractual"),
			cada(rel(hijo("http://www.sat.gob.mx/GastosHidrocarburos10", "Erogacion")),
				aplicar(rel()),
			),
		},
		// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt
		{"http://www.sat.gob.mx/GastosHidrocarburos10", "Erogacion"}: {
			req("TipoErogacion"),
			req("MontocuErogacion"),
			req("Porcentaje"),
			cada(rel(hijo("http://www.sat.gob.mx/GastosHidrocarburos10", "DocumentoRelacionado")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/GastosHidrocarburos10", "Actividades")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/GastosHidrocarburos10", "CentroCostos")),
				aplicar(rel()),
			),
		},
		// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt
		{"http://www.sat.gob.mx/GastosHidrocarburos10", "DocumentoRelacionado"}: {
			req("OrigenErogacion"),
			opc("FolioFiscalVinculado"),
			opc("RFCProveedor"),
			req("MontoTotalIVA"),
			opc("MontoRetencionISR"),
			opc("MontoRetencionIVA"),
			opc("MontoRetencionOtrosImpuestos"),
			opc("NumeroPedimentoVinculado"),
			opc("ClavePedimentoVinculado"),
			opc("ClavePagoPedimentoVinculado"),
			opc("MontoIVAPedimento"),
			opc("OtrosImpuestosPagadosPedimento"),
			req("FechaFolioFiscalVinculado"),
			req("Mes"),
			req("MontoTotalErogaciones"),
		},
		// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt
		{"http://www.sat.gob.mx/GastosHidrocarburos10", "Actividades"}: {
			opc("ActividadRelacionada"),
			cada(rel(hijo("http://www.sat.gob.mx/GastosHidrocarburos10", "SubActividades")),
				aplicar(rel()),
			),
		},
		// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt
		{"http://www.sat.gob.mx/GastosHidrocarburos10", "SubActividades"}: {
			opc("SubActividadRelacionada"),
			cada(rel(hijo("http://www.sat.gob.mx/GastosHidrocarburos10", "Tareas")),
				aplicar(rel()),
			),
		},
		// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt
		{"http://www.sat.gob.mx/GastosHidrocarburos10", "Tareas"}: {
			opc("TareaRelacionada"),
		},
		// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt
		{"http://www.sat.gob.mx/GastosHidrocarburos10", "CentroCostos"}: {
			opc("Campo"),
			cada(rel(hijo("http://www.sat.gob.mx/GastosHidrocarburos10", "Yacimientos")),
				aplicar(rel()),
			),
		},
		// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt
		{"http://www.sat.gob.mx/GastosHidrocarburos10", "Yacimientos"}: {
			opc("Yacimiento"),
			cada(rel(hijo("http://www.sat.gob.mx/GastosHidrocarburos10", "Pozos")),
				aplicar(rel()),
			),
		},
		// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt
		{"http://www.sat.gob.mx/GastosHidrocarburos10", "Pozos"}: {
			opc("Pozo"),
		},
		// cfd/IngresosHidrocarburos10/IngresosHidrocarburos.xslt
		{"http://www.sat.gob.mx/IngresosHidrocarburos10", "IngresosHidrocarburos"}: {
			req("Version"),
			req("NumeroContrato"),
			req("ContraprestacionPagadaOperador"),
			req("Porcentaje"),
			cada(rel(hijo("http://www.sat.gob.mx/IngresosHidrocarburos10", "DocumentoRelacionado")),
				aplicar(rel()),
			),
		},
		// cfd/IngresosHidrocarburos10/IngresosHidrocarburos.xslt
		{"http://www.sat.gob.mx/IngresosHidrocarburos10", "DocumentoRelacionado"}: {
			req("FolioFiscalVinculado"),
			req("FechaFolioFiscalVinculado"),
			req("Mes"),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "CartaPorte"}: {
			req("Version"),
			req("TranspInternac"),
			opc("EntradaSalidaMerc"),
			opc("ViaEntradaSalida"),
			opc("TotalDistRec"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Ubicaciones")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Mercancias")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "FiguraTransporte")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "Ubicaciones"}: {
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Ubicacion")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "Ubicacion"}: {
			opc("TipoEstacion"),
			opc("DistanciaRecorrida"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Origen")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Destino")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Domicilio")),
				req("Calle"),
				opc("NumeroExterior"),
				opc("NumeroInterior"),
				opc("Colonia"),
				opc("Localidad"),
				opc("Referencia"),
				opc("Municipio"),
				req("Estado"),
				req("Pais"),
				req("CodigoPostal"),
			),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "Origen"}: {
			opc("IDOrigen"),
			opc("RFCRemitente"),
			opc("NombreRemitente"),
			opc("NumRegIdTrib"),
			opc("ResidenciaFiscal"),
			opc("NumEstacion"),
			opc("NombreEstacion"),
			opc("NavegacionTrafico"),
			req("FechaHoraSalida"),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "Destino"}: {
			opc("IDDestino"),
			opc("RFCDestinatario"),
			opc("NombreDestinatario"),
			opc("NumRegIdTrib"),
			opc("ResidenciaFiscal"),
			opc("NumEstacion"),
			opc("NombreEstacion"),
			opc("NavegacionTrafico"),
			req("FechaHoraProgLlegada"),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "Mercancias"}: {
			opc("PesoBrutoTotal"),
			opc("UnidadPeso"),
			opc("PesoNetoTotal"),
			req("NumTotalMercancias"),
			opc("CargoPorTasacion"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Mercancia")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "AutotransporteFederal")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "TransporteMaritimo")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "TransporteAereo")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "TransporteFerroviario")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "Mercancia"}: {
			opc("BienesTransp"),
			opc("ClaveSTCC"),
			opc("Descripcion"),
			opc("Cantidad"),
			opc("ClaveUnidad"),
			opc("Unidad"),
			opc("Dimensiones"),
			opc("MaterialPeligroso"),
			opc("CveMaterialPeligroso"),
			opc("Embalaje"),
			opc("DescripEmbalaje"),
			req("PesoEnKg"),
			opc("ValorMercancia"),
			opc("Moneda"),
			opc("FraccionArancelaria"),
			opc("UUIDComercioExt"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "CantidadTransporta")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "DetalleMercancia")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "CantidadTransporta"}: {
			req("Cantidad"),
			req("IDOrigen"),
			req("IDDestino"),
			opc("CvesTransporte"),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "DetalleMercancia"}: {
			req("UnidadPeso"),
			req("PesoBruto"),
			req("PesoNeto"),
			req("PesoTara"),
			opc("NumPiezas"),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "AutotransporteFederal"}: {
			req("PermSCT"),
			req("NumPermisoSCT"),
			req("NombreAseg"),
			req("NumPolizaSeguro"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "IdentificacionVehicular")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Remolques")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "IdentificacionVehicular"}: {
			req("ConfigVehicular"),
			req("PlacaVM"),
			req("AnioModeloVM"),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "Remolques"}: {
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Remolque")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "Remolque"}: {
			req("SubTipoRem"),
			req("Placa"),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "TransporteMaritimo"}: {
			opc("PermSCT"),
			opc("NumPermisoSCT"),
			opc("NombreAseg"),
			opc("NumPolizaSeguro"),
			req("TipoEmbarcacion"),
			req("Matricula"),
			req("NumeroOMI"),
			opc("AnioEmbarcacion"),
			opc("NombreEmbarc"),
			req("NacionalidadEmbarc"),
			req("UnidadesDeArqBruto"),
			req("TipoCarga"),
			req("NumCertITC"),
			opc("Eslora"),
			opc("Manga"),
			opc("Calado"),
			opc("LineaNaviera"),
			req("NombreAgenteNaviero"),
			req("NumAutorizacionNaviero"),
			opc("NumViaje"),
			opc("NumConocEmbarc"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Contenedor")),
				req("MatriculaContenedor"),
				req("TipoContenedor"),
				opc("NumPrecinto"),
			),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "TransporteAereo"}: {
			req("PermSCT"),
			req("NumPermisoSCT"),
			req("MatriculaAeronave"),
			opc("NombreAseg"),
			opc("NumPolizaSeguro"),
			req("NumeroGuia"),
			opc("LugarContrato"),
			opc("RFCTransportista"),
			req("CodigoTransportista"),
			opc("NumRegIdTribTranspor"),
			opc("ResidenciaFiscalTranspor"),
			opc("NombreTransportista"),
			opc("RFCEmbarcador"),
			opc("NumRegIdTribEmbarc"),
			opc("ResidenciaFiscalEmbarc"),
			opc("NombreEmbarcador"),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "TransporteFerroviario"}: {
			req("TipoDeServicio"),
			opc("NombreAseg"),
			opc("NumPolizaSeguro"),
			opc("Concesionario"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "DerechosDePaso")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Carro")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "DerechosDePaso"}: {
			req("TipoDerechoDePaso"),
			req("KilometrajePagado"),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "Carro"}: {
			req("TipoCarro"),
			req("MatriculaCarro"),
			req("GuiaCarro"),
			req("ToneladasNetasCarro"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Contenedor")),
				req("TipoContenedor"),
				req("PesoContenedorVacio"),
				req("PesoNetoMercancia"),
			),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "FiguraTransporte"}: {
			req("CveTransporte"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Operadores")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Propietario")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Arrendatario")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Notificado")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "Operadores"}: {
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Operador")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "Operador"}: {
			opc("RFCOperador"),
			opc("NumLicencia"),
			opc("NombreOperador"),
			opc("NumRegIdTribOperador"),
			opc("ResidenciaFiscalOperador"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Domicilio")),
				req("Calle"),
				opc("NumeroExterior"),
				opc("NumeroInterior"),
				opc("Colonia"),
				opc("Localidad"),
				opc("Referencia"),
				opc("Municipio"),
				req("Estado"),
				req("Pais"),
				req("CodigoPostal"),
			),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "Propietario"}: {
			opc("RFCPropietario"),
			opc("NombrePropietario"),
			opc("NumRegIdTribPropietario"),
			opc("ResidenciaFiscalPropietario"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Domicilio")),
				req("Calle"),
				opc("NumeroExterior"),
				opc("NumeroInterior"),
				opc("Colonia"),
				opc("Localidad"),
				opc("Referencia"),
				opc("Municipio"),
				req("Estado"),
				req("Pais"),
				req("CodigoPostal"),
			),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "Arrendatario"}: {
			opc("RFCArrendatario"),
			opc("NombreArrendatario"),
			opc("NumRegIdTribArrendatario"),
			opc("ResidenciaFiscalArrendatario"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Domicilio")),
				req("Calle"),
				opc("NumeroExterior"),
				opc("NumeroInterior"),
				opc("Colonia"),
				opc("Localidad"),
				opc("Referencia"),
				opc("Municipio"),
				req("Estado"),
				req("Pais"),
				req("CodigoPostal"),
			),
		},
		// cfd/CartaPorte/CartaPorte.xslt
		{"http://www.sat.gob.mx/CartaPorte", "Notificado"}: {
			opc("RFCNotificado"),
			opc("NombreNotificado"),
			opc("NumRegIdTribNotificado"),
			opc("ResidenciaFiscalNotificado"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte", "Domicilio")),
				req("Calle"),
				opc("NumeroExterior"),
				opc("NumeroInterior"),
				opc("Colonia"),
				opc("Localidad"),
				opc("Referencia"),
				opc("Municipio"),
				req("Estado"),
				req("Pais"),
				req("CodigoPostal"),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "CartaPorte"}: {
			req("Version"),
			req("TranspInternac"),
			opc("EntradaSalidaMerc"),
			opc("PaisOrigenDestino"),
			opc("ViaEntradaSalida"),
			opc("TotalDistRec"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Ubicaciones")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Mercancias")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "FiguraTransporte")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Ubicaciones"}: {
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Ubicacion")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Ubicacion"}: {
			req("TipoUbicacion"),
			opc("IDUbicacion"),
			req("RFCRemitenteDestinatario"),
			opc("NombreRemitenteDestinatario"),
			opc("NumRegIdTrib"),
			opc("ResidenciaFiscal"),
			opc("NumEstacion"),
			opc("NombreEstacion"),
			opc("NavegacionTrafico"),
			req("FechaHoraSalidaLlegada"),
			opc("TipoEstacion"),
			opc("DistanciaRecorrida"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Domicilio")),
				opc("Calle"),
				opc("NumeroExterior"),
				opc("NumeroInterior"),
				opc("Colonia"),
				opc("Localidad"),
				opc("Referencia"),
				opc("Municipio"),
				req("Estado"),
				req("Pais"),
				req("CodigoPostal"),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Mercancias"}: {
			req("PesoBrutoTotal"),
			req("UnidadPeso"),
			opc("PesoNetoTotal"),
			req("NumTotalMercancias"),
			opc("CargoPorTasacion"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Mercancia")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Autotransporte")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "TransporteMaritimo")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "TransporteAereo")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "TransporteFerroviario")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Mercancia"}: {
			req("BienesTransp"),
			opc("ClaveSTCC"),
			req("Descripcion"),
			req("Cantidad"),
			req("ClaveUnidad"),
			opc("Unidad"),
			opc("Dimensiones"),
			opc("MaterialPeligroso"),
			opc("CveMaterialPeligroso"),
			opc("Embalaje"),
			opc("DescripEmbalaje"),
			req("PesoEnKg"),
			opc("ValorMercancia"),
			opc("Moneda"),
			opc("FraccionArancelaria"),
			opc("UUIDComercioExt"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Pedimentos")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "GuiasIdentificacion")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "CantidadTransporta")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "DetalleMercancia")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Pedimentos"}: {
			req("Pedimento"),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "GuiasIdentificacion"}: {
			req("NumeroGuiaIdentificacion"),
			req("DescripGuiaIdentificacion"),
			req("PesoGuiaIdentificacion"),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "CantidadTransporta"}: {
			req("Cantidad"),
			req("IDOrigen"),
			req("IDDestino"),
			opc("CvesTransporte"),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "DetalleMercancia"}: {
			req("UnidadPesoMerc"),
			req("PesoBruto"),
			req("PesoNeto"),
			req("PesoTara"),
			opc("NumPiezas"),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Autotransporte"}: {
			req("PermSCT"),
			req("NumPermisoSCT"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "IdentificacionVehicular")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Seguros")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Remolques")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "IdentificacionVehicular"}: {
			req("ConfigVehicular"),
			req("PlacaVM"),
			req("AnioModeloVM"),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Seguros"}: {
			req("AseguraRespCivil"),
			req("PolizaRespCivil"),
			opc("AseguraMedAmbiente"),
			opc("PolizaMedAmbiente"),
			opc("AseguraCarga"),
			opc("PolizaCarga"),
			opc("PrimaSeguro"),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Remolques"}: {
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Remolque")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Remolque"}: {
			req("SubTipoRem"),
			req("Placa"),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "TransporteMaritimo"}: {
			opc("PermSCT"),
			opc("NumPermisoSCT"),
			opc("NombreAseg"),
			opc("NumPolizaSeguro"),
			req("TipoEmbarcacion"),
			req("Matricula"),
			req("NumeroOMI"),
			opc("AnioEmbarcacion"),
			opc("NombreEmbarc"),
			req("NacionalidadEmbarc"),
			req("UnidadesDeArqBruto"),
			req("TipoCarga"),
			req("NumCertITC"),
			opc("Eslora"),
			opc("Manga"),
			opc("Calado"),
			opc("LineaNaviera"),
			req("NombreAgenteNaviero"),
			req("NumAutorizacionNaviero"),
			opc("NumViaje"),
			opc("NumConocEmbarc"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Contenedor")),
				req("MatriculaContenedor"),
				req("TipoContenedor"),
				opc("NumPrecinto"),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "TransporteAereo"}: {
			req("PermSCT"),
			req("NumPermisoSCT"),
			opc("MatriculaAeronave"),
			opc("NombreAseg"),
			opc("NumPolizaSeguro"),
			req("NumeroGuia"),
			opc("LugarContrato"),
			req("CodigoTransportista"),
			opc("RFCEmbarcador"),
			opc("NumRegIdTribEmbarc"),
			opc("ResidenciaFiscalEmbarc"),
			opc("NombreEmbarcador"),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "TransporteFerroviario"}: {
			req("TipoDeServicio"),
			req("TipoDeTrafico"),
			opc("NombreAseg"),
			opc("NumPolizaSeguro"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "DerechosDePaso")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Carro")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "DerechosDePaso"}: {
			req("TipoDerechoDePaso"),
			req("KilometrajePagado"),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Carro"}: {
			req("TipoCarro"),
			req("MatriculaCarro"),
			req("GuiaCarro"),
			req("ToneladasNetasCarro"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Contenedor")),
				req("TipoContenedor"),
				req("PesoContenedorVacio"),
				req("PesoNetoMercancia"),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "FiguraTransporte"}: {
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "TiposFigura")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "TiposFigura"}: {
			req("TipoFigura"),
			opc("RFCFigura"),
			opc("NumLicencia"),
			opc("NombreFigura"),
			opc("NumRegIdTribFigura"),
			opc("ResidenciaFiscalFigura"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "PartesTransporte")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Domicilio")),
				opc("Calle"),
				opc("NumeroExterior"),
				opc("NumeroInterior"),
				opc("Colonia"),
				opc("Localidad"),
				opc("Referencia"),
				opc("Municipio"),
				req("Estado"),
				req("Pais"),
				req("CodigoPostal"),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "PartesTransporte"}: {
			req("ParteTransporte"),
		},
		// cfd/3/cadenaoriginal_3_3/cadenaoriginal_3_3.xslt
		{"http://www.sat.gob.mx/cfd/3", "Comprobante"}: {
			req("Version"),
			opc("Serie"),
			opc("Folio"),
			req("Fecha"),
			opc("FormaPago"),
			req("NoCertificado"),
			opc("CondicionesDePago"),
			req("SubTotal"),
			opc("Descuento"),
			req("Moneda"),
			opc("TipoCambio"),
			req("Total"),
			req("TipoDeComprobante"),
			opc("MetodoPago"),
			req("LugarExpedicion"),
			opc("Confirmacion"),
			aplicar(rel(hijo("http://www.sat.gob.mx/cfd/3", "CfdiRelacionados"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/cfd/3", "Emisor"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/cfd/3", "Receptor"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/cfd/3", "Conceptos"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/cfd/3", "Impuestos"))),
			cada(rel(hijo("http://www.sat.gob.mx/cfd/3", "Complemento")),
				aplicar(rel()),
			),
		},
		// cfd/3/cadenaoriginal_3_3/cadenaoriginal_3_3.xslt
		{"http://www.sat.gob.mx/cfd/3", "CfdiRelacionados"}: {
			req("TipoRelacion"),
			cada(rel(hijo("http://www.sat.gob.mx/cfd/3", "CfdiRelacionado")),
				req("UUID"),
			),
		},
		// cfd/3/cadenaoriginal_3_3/cadenaoriginal_3_3.xslt
		{"http://www.sat.gob.mx/cfd/3", "Emisor"}: {
			req("Rfc"),
			opc("Nombre"),
			req("RegimenFiscal"),
		},
		// cfd/3/cadenaoriginal_3_3/cadenaoriginal_3_3.xslt
		{"http://www.sat.gob.mx/cfd/3", "Receptor"}: {
			req("Rfc"),
			opc("Nombre"),
			opc("ResidenciaFiscal"),
			opc("NumRegIdTrib"),
			req("UsoCFDI"),
		},
		// cfd/3/cadenaoriginal_3_3/cadenaoriginal_3_3.xslt
		{"http://www.sat.gob.mx/cfd/3", "Conceptos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/cfd/3", "Concepto")),
				aplicar(rel()),
			),
		},
		// cfd/3/cadenaoriginal_3_3/cadenaoriginal_3_3.xslt
		{"http://www.sat.gob.mx/cfd/3", "Concepto"}: {
			req("ClaveProdServ"),
			opc("NoIdentificacion"),
			req("Cantidad"),
			req("ClaveUnidad"),
			opc("Unidad"),
			req("Descripcion"),
			req("ValorUnitario"),
			req("Importe"),
			opc("Descuento"),
			cada(rel(hijo("http://www.sat.gob.mx/cfd/3", "Impuestos"), hijo("http://www.sat.gob.mx/cfd/3", "Traslados"), hijo("http://www.sat.gob.mx/cfd/3", "Traslado")),
				req("Base"),
				req("Impuesto"),
				req("TipoFactor"),
				opc("TasaOCuota"),
				opc("Importe"),
			),
			cada(rel(hijo("http://www.sat.gob.mx/cfd/3", "Impuestos"), hijo("http://www.sat.gob.mx/cfd/3", "Retenciones"), hijo("http://www.sat.gob.mx/cfd/3", "Retencion")),
				req("Base"),
				req("Impuesto"),
				req("TipoFactor"),
				req("TasaOCuota"),
				req("Importe"),
			),
			cada(rel(hijo("http://www.sat.gob.mx/cfd/3", "InformacionAduanera")),
				aplicar(rel()),
			),
			si(rel(hijo("http://www.sat.gob.mx/cfd/3", "CuentaPredial")),
				aplicar(rel(hijo("http://www.sat.gob.mx/cfd/3", "CuentaPredial"))),
			),
			si(rel(hijo("http://www.sat.gob.mx/cfd/3", "ComplementoConcepto")),
				aplicar(rel(hijo("http://www.sat.gob.mx/cfd/3", "ComplementoConcepto"))),
			),
			cada(rel(desc("http://www.sat.gob.mx/cfd/3", "Parte")),
				aplicar(rel()),
			),
		},
		// cfd/3/cadenaoriginal_3_3/cadenaoriginal_3_3.xslt
		{"http://www.sat.gob.mx/cfd/3", "InformacionAduanera"}: {
			req("NumeroPedimento"),
		},
		// cfd/3/cadenaoriginal_3_3/cadenaoriginal_3_3.xslt
		{"http://www.sat.gob.mx/cfd/3", "CuentaPredial"}: {
			req("Numero"),
		},
		// cfd/3/cadenaoriginal_3_3/cadenaoriginal_3_3.xslt
		{"http://www.sat.gob.mx/cfd/3", "ComplementoConcepto"}: {
			cada(rel(hijo("", "*")),
				aplicar(rel()),
			),
		},
		// cfd/3/cadenaoriginal_3_3/cadenaoriginal_3_3.xslt
		{"http://www.sat.gob.mx/cfd/3", "Parte"}: {
			req("ClaveProdServ"),
			opc("NoIdentificacion"),
			req("Cantidad"),
			opc("Unidad"),
			req("Descripcion"),
			opc("ValorUnitario"),
			opc("Importe"),
			cada(rel(desc("http://www.sat.gob.mx/cfd/3", "InformacionAduanera")),
				aplicar(rel()),
			),
		},
		// cfd/3/cadenaoriginal_3_3/cadenaoriginal_3_3.xslt
		{"http://www.sat.gob.mx/cfd/3", "Complemento"}: {
			cada(rel(hijo("", "*")),
				aplicar(rel()),
			),
		},
		// cfd/3/cadenaoriginal_3_3/cadenaoriginal_3_3.xslt
		{"http://www.sat.gob.mx/cfd/3", "Impuestos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/cfd/3", "Retenciones"), hijo("http://www.sat.gob.mx/cfd/3", "Retencion")),
				req("Impuesto"),
				req("Importe"),
			),
			opc("TotalImpuestosRetenidos"),
			cada(rel(hijo("http://www.sat.gob.mx/cfd/3", "Traslados"), hijo("http://www.sat.gob.mx/cfd/3", "Traslado")),
				req("Impuesto"),
				req("TipoFactor"),
				req("TasaOCuota"),
				req("Importe"),
			),
			opc("TotalImpuestosTrasladados"),
		},
	},
}

// hojaCFDI40 sale de cfd/4/cadenaoriginal_4_0/cadenaoriginal_4_0.xslt y de lo que incluye.
//
// Archivos (sha256):
//   - cfd/4/cadenaoriginal_4_0/cadenaoriginal_4_0.xslt 11dafd930e26127fbbd6c5714067286e04d6468ed1f37e4ec5647e5f21894729
//   - cfd/2/cadenaoriginal_2_0/utilerias.xslt d9ef68267747e0f6725295d4012684a4c051cb3f2265dcb2035d32bf40895016
//   - cfd/donat/donat11.xslt 1c72d07d9945fd7824471caaebaee47dafc3b56785a037549fcbee5e123c6d36
//   - cfd/divisas/divisas.xslt be7167e1d421b66ba17a6f1788c964c5bb1f207ed7ad9803166d60966a332b1d
//   - cfd/implocal/implocal.xslt 3c2ee6c3ecce9b2e74b2a225f3c17250259751a1307f58a96799ff119e2fafe6
//   - cfd/leyendasFiscales/leyendasFisc.xslt 68b0016e44c1829a5c7d5da340b69945100d8fda1243e4e279c2cdc8c745d1fb
//   - cfd/pfic/pfic.xslt 0e9e992e4e0fa6e13ee7b18ce5ddc51c55f3615236c7c51b80426811f9779bab
//   - cfd/TuristaPasajeroExtranjero/TuristaPasajeroExtranjero.xslt 06aee48f06ed53ab4fad9f2ddc950bca51840cfd4fb137ec8d14632e2c2caeb3
//   - cfd/nomina/nomina12.xslt a14cdb44995526aa5888036bedf7fac9f84ad0041f9393371b0bbbea0572649e
//   - cfd/cfdiregistrofiscal/cfdiregistrofiscal.xslt d824c4993e1f5940ff0605bdd3ea040669bef69355b551ab9da916f3a3746e7d
//   - cfd/pagoenespecie/pagoenespecie.xslt ed33ce278525a43db377d82fe641d42391d74b731f82015b1fedfa84654f1798
//   - cfd/aerolineas/aerolineas.xslt d000d0ed64a45a05a7f59b1eb2399d55d1588f18056ed37e8bc642ca2fe655c0
//   - cfd/valesdedespensa/valesdedespensa.xslt e22cf591f5669b78b3306c33e2434acbf7f40c3671e318c3fb2b736a2bf6ef8c
//   - cfd/notariospublicos/notariospublicos.xslt b3aeebdbfa1858b60d9f8b437f9f07a95f600e712b52aa0a8c08a7f924fe9ce7
//   - cfd/vehiculousado/vehiculousado.xslt 226664fcdcbc51e57abb260104b59a06d3bba62a62ae92fbbfeee6f16791ae2c
//   - cfd/servicioparcialconstruccion/servicioparcialconstruccion.xslt b3b9a8983a673a352ac291c1e0a9605e983dc0af8a4ecb4f6ec0cb13e258fe85
//   - cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt 347bf6dfdae5c0b18d6a107c8c339f212763496fa4ae1fa748d0354a7c3338d2
//   - cfd/certificadodestruccion/certificadodedestruccion.xslt 0b306d8142eb118f47d3ab6c967a6be6c9c31436678e776e571ce496018b68d5
//   - cfd/arteantiguedades/obrasarteantiguedades.xslt 2d11a854ff914b229ef7d2d43cdfd6296315d040c36e0ad6c721ace08ba89d1d
//   - cfd/ComercioExterior11/ComercioExterior11.xslt bb9e7c3d2a2a3e8af9c65498fd1422b0f400e0787d2418b460d51ece6c988f9d
//   - cfd/ComercioExterior20/ComercioExterior20.xslt d1ac00ab5fd8a424e8c2e714bb998d3933cc78398260e16a451b12e6378c1802
//   - cfd/ine/ine11.xslt 02781b96a0da332693f719f893f7a72117196b073c5d0df48e7fa6a4ed444882
//   - cfd/iedu/iedu.xslt 2e9ddb05e5cd222cded57a5306b43f954f15a2ac9f8ddf3e4818156cd703a49d
//   - cfd/ventavehiculos/ventavehiculos11.xslt 95fd6522700a5125dacc363e07fac8ff2f470b2e23aec5bbe3d3f684c8a4152f
//   - cfd/EstadoDeCuentaCombustible/ecc12.xslt 9ce3097989da2560e3fcb6e45928cd7d90692832b2d29c0c4127732349f3dba5
//   - cfd/consumodecombustibles/consumodeCombustibles11.xslt 610de8017f569d6e66669fc1f8b2303c5b4ac970e9b17e44bcd5d1b0e4ef7bf1
//   - cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt 1d735ccdd03040a19469795e6703fbd2c56a2c08c708111278af6d885e98621c
//   - cfd/IngresosHidrocarburos10/IngresosHidrocarburos.xslt 2d17335608f036214129210b06da1a132ec0fc06b315f97dd3e07d09641303d9
//   - cfd/CartaPorte/CartaPorte20.xslt 0e8548e036920208c8e909acd64f6e7d4de6742e75c28617c7ac8edb97c55766
//   - cfd/Pagos/Pagos20.xslt c731c30e75ce92f25a30b9b29c91f8b059fb1404f3af0fb1a413d204f9b59a84
//   - cfd/CartaPorte/CartaPorte30.xslt 4817b2640e11ec6f01ed8eaeaaa196d459702e4f71d136d86e1c59e7e76da4ea
//   - cfd/CartaPorte/CartaPorte31.xslt c5e349ca4a3f443489aa914516998b8f63d2c104687533d6181114efb69f2b02
//   - cfd/HidrocarburosPetro/hidrocarburospetroliferos.xslt c207c8dcf8ef6336e69ac09dc2484fbe1d81ef86cabcc9d2903a9fa1ddc9ae27
//
// Includes omitidos (complementos que la librería no lee):
//   - cfd/detallista/detallista.xslt
//
// Plantillas match="/" de hojas incluidas, que la de la raíz sustituye:
//   - cfd/vehiculousado/vehiculousado.xslt
var hojaCFDI40 = &hojaCadena{
	nombre: "cadenaoriginal_4_0",
	raiz: []pasoCadena{
		txt("|"),
		aplicar(abs(hijo("http://www.sat.gob.mx/cfd/4", "Comprobante"))),
		txt("||"),
	},
	plantillas: map[nombreXML][]pasoCadena{
		// cfd/donat/donat11.xslt
		{"http://www.sat.gob.mx/donat", "Donatarias"}: {
			req("version"),
			req("noAutorizacion"),
			req("fechaAutorizacion"),
			req("leyenda"),
		},
		// cfd/divisas/divisas.xslt
		{"http://www.sat.gob.mx/divisas", "Divisas"}: {
			req("version"),
			req("tipoOperacion"),
		},
		// cfd/implocal/implocal.xslt
		{"http://www.sat.gob.mx/implocal", "ImpuestosLocales"}: {
			req("version"),
			req("TotaldeRetenciones"),
			req("TotaldeTraslados"),
			cada(rel(hijo("http://www.sat.gob.mx/implocal", "RetencionesLocales")),
				req("ImpLocRetenido"),
				req("TasadeRetencion"),
				req("Importe"),
			),
			cada(rel(hijo("http://www.sat.gob.mx/implocal", "TrasladosLocales")),
				req("ImpLocTrasladado"),
				req("TasadeTraslado"),
				req("Importe"),
			),
		},
		// cfd/leyendasFiscales/leyendasFisc.xslt
		{"http://www.sat.gob.mx/leyendasFiscales", "LeyendasFiscales"}: {
			req("version"),
			cada(rel(hijo("http://www.sat.gob.mx/leyendasFiscales", "Leyenda")),
				aplicar(rel()),
			),
		},
		// cfd/leyendasFiscales/leyendasFisc.xslt
		{"http://www.sat.gob.mx/leyendasFiscales", "Leyenda"}: {
			opc("disposicionFiscal"),
			opc("norma"),
			req("textoLeyenda"),
		},
		// cfd/pfic/pfic.xslt
		{"http://www.sat.gob.mx/pfic", "PFintegranteCoordinado"}: {
			req("version"),
			req("ClaveVehicular"),
			req("Placa"),
			opc("RFCPF"),
		},
		// cfd/TuristaPasajeroExtranjero/TuristaPasajeroExtranjero.xslt
		{"http://www.sat.gob.mx/TuristaPasajeroExtranjero", "TuristaPasajeroExtranjero"}: {
			req("version"),
			req("fechadeTransito"),
			req("tipoTransito"),
			aplicar(rel(hijo("http://www.sat.gob.mx/TuristaPasajeroExtranjero", "datosTransito"))),
		},
		// cfd/TuristaPasajeroExtranjero/TuristaPasajeroExtranjero.xslt
		{"http://www.sat.gob.mx/TuristaPasajeroExtranjero", "datosTransito"}: {
			req("Via"),
			req("TipoId"),
			req("NumeroId"),
			req("Nacionalidad"),
			req("EmpresaTransporte"),
			opc("IdTransporte"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Nomina"}: {
			req("Version"),
			req("TipoNomina"),
			req("FechaPago"),
			req("FechaInicialPago"),
			req("FechaFinalPago"),
			req("NumDiasPagados"),
			opc("TotalPercepciones"),
			opc("TotalDeducciones"),
			opc("TotalOtrosPagos"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Emisor")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Receptor")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Percepciones")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Deducciones")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "OtrosPagos")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Incapacidades")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Emisor"}: {
			opc("Curp"),
			opc("RegistroPatronal"),
			opc("RfcPatronOrigen"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "EntidadSNCF")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "EntidadSNCF"}: {
			req("OrigenRecurso"),
			opc("MontoRecursoPropio"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Receptor"}: {
			req("Curp"),
			opc("NumSeguridadSocial"),
			opc("FechaInicioRelLaboral"),
			opc("Antigüedad"),
			req("TipoContrato"),
			opc("Sindicalizado"),
			opc("TipoJornada"),
			req("TipoRegimen"),
			req("NumEmpleado"),
			opc("Departamento"),
			opc("Puesto"),
			opc("RiesgoPuesto"),
			req("PeriodicidadPago"),
			opc("Banco"),
			opc("CuentaBancaria"),
			opc("SalarioBaseCotApor"),
			opc("SalarioDiarioIntegrado"),
			req("ClaveEntFed"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "SubContratacion")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "SubContratacion"}: {
			req("RfcLabora"),
			req("PorcentajeTiempo"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Percepciones"}: {
			opc("TotalSueldos"),
			opc("TotalSeparacionIndemnizacion"),
			opc("TotalJubilacionPensionRetiro"),
			req("TotalGravado"),
			req("TotalExento"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Percepcion")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "JubilacionPensionRetiro")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "SeparacionIndemnizacion")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Percepcion"}: {
			req("TipoPercepcion"),
			req("Clave"),
			req("Concepto"),
			req("ImporteGravado"),
			req("ImporteExento"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "AccionesOTitulos")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "HorasExtra")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "AccionesOTitulos"}: {
			req("ValorMercado"),
			req("PrecioAlOtorgarse"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "HorasExtra"}: {
			req("Dias"),
			req("TipoHoras"),
			req("HorasExtra"),
			req("ImportePagado"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "JubilacionPensionRetiro"}: {
			opc("TotalUnaExhibicion"),
			opc("TotalParcialidad"),
			opc("MontoDiario"),
			req("IngresoAcumulable"),
			req("IngresoNoAcumulable"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "SeparacionIndemnizacion"}: {
			req("TotalPagado"),
			req("NumAñosServicio"),
			req("UltimoSueldoMensOrd"),
			req("IngresoAcumulable"),
			req("IngresoNoAcumulable"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Deducciones"}: {
			opc("TotalOtrasDeducciones"),
			opc("TotalImpuestosRetenidos"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Deduccion")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Deduccion"}: {
			req("TipoDeduccion"),
			req("Clave"),
			req("Concepto"),
			req("Importe"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "OtrosPagos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "OtroPago")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "OtroPago"}: {
			req("TipoOtroPago"),
			req("Clave"),
			req("Concepto"),
			req("Importe"),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "SubsidioAlEmpleo")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "CompensacionSaldosAFavor")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "SubsidioAlEmpleo"}: {
			req("SubsidioCausado"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "CompensacionSaldosAFavor"}: {
			req("SaldoAFavor"),
			req("Año"),
			req("RemanenteSalFav"),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Incapacidades"}: {
			cada(rel(hijo("http://www.sat.gob.mx/nomina12", "Incapacidad")),
				aplicar(rel()),
			),
		},
		// cfd/nomina/nomina12.xslt
		{"http://www.sat.gob.mx/nomina12", "Incapacidad"}: {
			req("DiasIncapacidad"),
			req("TipoIncapacidad"),
			opc("ImporteMonetario"),
		},
		// cfd/cfdiregistrofiscal/cfdiregistrofiscal.xslt
		{"http://www.sat.gob.mx/registrofiscal", "CFDIRegistroFiscal"}: {
			req("Version"),
			req("Folio"),
		},
		// cfd/pagoenespecie/pagoenespecie.xslt
		{"http://www.sat.gob.mx/pagoenespecie", "PagoEnEspecie"}: {
			req("Version"),
			req("CvePIC"),
			req("FolioSolDon"),
			req("PzaArtNombre"),
			req("PzaArtTecn"),
			req("PzaArtAProd"),
			req("PzaArtDim"),
		},
		// cfd/aerolineas/aerolineas.xslt
		{"http://www.sat.gob.mx/aerolineas", "Aerolineas"}: {
			req("Version"),
			req("TUA"),
			aplicar(rel(hijo("http://www.sat.gob.mx/aerolineas", "OtrosCargos"))),
		},
		// cfd/aerolineas/aerolineas.xslt
		{"http://www.sat.gob.mx/aerolineas", "OtrosCargos"}: {
			req("TotalCargos"),
			cada(rel(hijo("http://www.sat.gob.mx/aerolineas", "Cargo")),
				aplicar(rel()),
			),
		},
		// cfd/aerolineas/aerolineas.xslt
		{"http://www.sat.gob.mx/aerolineas", "Cargo"}: {
			req("CodigoCargo"),
			req("Importe"),
		},
		// cfd/valesdedespensa/valesdedespensa.xslt
		{"http://www.sat.gob.mx/valesdedespensa", "ValesDeDespensa"}: {
			req("version"),
			req("tipoOperacion"),
			opc("registroPatronal"),
			req("numeroDeCuenta"),
			req("total"),
			aplicar(rel(hijo("http://www.sat.gob.mx/valesdedespensa", "Conceptos"))),
		},
		// cfd/valesdedespensa/valesdedespensa.xslt
		{"http://www.sat.gob.mx/valesdedespensa", "Conceptos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/valesdedespensa", "Concepto")),
				aplicar(rel()),
			),
		},
		// cfd/valesdedespensa/valesdedespensa.xslt
		{"http://www.sat.gob.mx/valesdedespensa", "Concepto"}: {
			req("identificador"),
			req("fecha"),
			req("rfc"),
			req("curp"),
			req("nombre"),
			opc("numSeguridadSocial"),
			req("importe"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "NotariosPublicos"}: {
			req("Version"),
			aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DescInmuebles"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosOperacion"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosNotario"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosEnajenante"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosAdquiriente"))),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DescInmuebles"}: {
			cada(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DescInmueble")),
				aplicar(rel()),
			),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DescInmueble"}: {
			req("TipoInmueble"),
			req("Calle"),
			opc("NoExterior"),
			opc("NoInterior"),
			opc("Colonia"),
			opc("Localidad"),
			opc("Referencia"),
			req("Municipio"),
			req("Estado"),
			req("Pais"),
			req("CodigoPostal"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosOperacion"}: {
			req("NumInstrumentoNotarial"),
			req("FechaInstNotarial"),
			req("MontoOperacion"),
			req("Subtotal"),
			req("IVA"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosNotario"}: {
			req("CURP"),
			req("NumNotaria"),
			req("EntidadFederativa"),
			opc("Adscripcion"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosEnajenante"}: {
			req("CoproSocConyugalE"),
			si(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosUnEnajenante")),
				aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosUnEnajenante"))),
			),
			si(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosEnajenantesCopSC")),
				aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosEnajenantesCopSC"))),
			),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosUnEnajenante"}: {
			req("Nombre"),
			req("ApellidoPaterno"),
			opc("ApellidoMaterno"),
			req("RFC"),
			req("CURP"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosEnajenantesCopSC"}: {
			cada(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosEnajenanteCopSC")),
				aplicar(rel()),
			),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosEnajenanteCopSC"}: {
			req("Nombre"),
			opc("ApellidoPaterno"),
			opc("ApellidoMaterno"),
			req("RFC"),
			opc("CURP"),
			req("Porcentaje"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosAdquiriente"}: {
			req("CoproSocConyugalE"),
			si(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosUnAdquiriente")),
				aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosUnAdquiriente"))),
			),
			si(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosAdquirientesCopSC")),
				aplicar(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosAdquirientesCopSC"))),
			),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosUnAdquiriente"}: {
			req("Nombre"),
			opc("ApellidoPaterno"),
			opc("ApellidoMaterno"),
			req("RFC"),
			opc("CURP"),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosAdquirientesCopSC"}: {
			cada(rel(hijo("http://www.sat.gob.mx/notariospublicos", "DatosAdquirienteCopSC")),
				aplicar(rel()),
			),
		},
		// cfd/notariospublicos/notariospublicos.xslt
		{"http://www.sat.gob.mx/notariospublicos", "DatosAdquirienteCopSC"}: {
			req("Nombre"),
			opc("ApellidoPaterno"),
			opc("ApellidoMaterno"),
			req("RFC"),
			opc("CURP"),
			req("Porcentaje"),
		},
		// cfd/vehiculousado/vehiculousado.xslt
		{"http://www.sat.gob.mx/vehiculousado", "VehiculoUsado"}: {
			req("Version"),
			req("montoAdquisicion"),
			req("montoEnajenacion"),
			req("claveVehicular"),
			req("marca"),
			req("tipo"),
			req("modelo"),
			opc("numeroMotor"),
			opc("numeroSerie"),
			opc("NIV"),
			req("valor"),
			aplicar(rel(hijo("http://www.sat.gob.mx/vehiculousado", "InformacionAduanera"))),
		},
		// cfd/vehiculousado/vehiculousado.xslt
		{"http://www.sat.gob.mx/vehiculousado", "InformacionAduanera"}: {
			req("numero"),
			req("fecha"),
			opc("aduana"),
		},
		// cfd/servicioparcialconstruccion/servicioparcialconstruccion.xslt
		{"http://www.sat.gob.mx/servicioparcialconstruccion", "parcialesconstruccion"}: {
			req("Version"),
			req("NumPerLicoAut"),
			aplicar(rel(hijo("http://www.sat.gob.mx/servicioparcialconstruccion", "Inmueble"))),
		},
		// cfd/servicioparcialconstruccion/servicioparcialconstruccion.xslt
		{"http://www.sat.gob.mx/servicioparcialconstruccion", "Inmueble"}: {
			req("Calle"),
			opc("NoExterior"),
			opc("NoInterior"),
			opc("Colonia"),
			opc("Localidad"),
			opc("Referencia"),
			req("Municipio"),
			req("Estado"),
			req("CodigoPostal"),
		},
		// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt
		{"http://www.sat.gob.mx/renovacionysustitucionvehiculos", "renovacionysustitucionvehiculos"}: {
			req("Version"),
			req("TipoDeDecreto"),
			aplicar(rel(hijo("http://www.sat.gob.mx/renovacionysustitucionvehiculos", "DecretoRenovVehicular"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/renovacionysustitucionvehiculos", "DecretoSustitVehicular"))),
		},
		// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt
		{"http://www.sat.gob.mx/renovacionysustitucionvehiculos", "DecretoRenovVehicular"}: {
			req("VehEnaj"),
			cada(rel(hijo("http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculosUsadosEnajenadoPermAlFab")),
				aplicar(rel()),
			),
			aplicar(rel(hijo("http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculoNuvoSemEnajenadoFabAlPerm"))),
		},
		// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt
		{"http://www.sat.gob.mx/renovacionysustitucionvehiculos", "DecretoSustitVehicular"}: {
			req("VehEnaj"),
			aplicar(rel(hijo("http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculoUsadoEnajenadoPermAlFab"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculoNuvoSemEnajenadoFabAlPerm"))),
		},
		// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt
		{"http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculosUsadosEnajenadoPermAlFab"}: {
			req("PrecioVehUsado"),
			req("TipoVeh"),
			req("Marca"),
			req("TipooClase"),
			req("Año"),
			opc("Modelo"),
			opc("NIV"),
			opc("NumSerie"),
			req("NumPlacas"),
			opc("NumMotor"),
			req("NumFolTarjCir"),
			opc("NumPedIm"),
			opc("Aduana"),
			opc("FechaRegulVeh"),
			req("Foliofiscal"),
		},
		// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt
		{"http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculoNuvoSemEnajenadoFabAlPerm"}: {
			req("Año"),
			opc("Modelo"),
			req("NumPlacas"),
			opc("RFC"),
		},
		// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt
		{"http://www.sat.gob.mx/renovacionysustitucionvehiculos", "VehiculoUsadoEnajenadoPermAlFab"}: {
			req("PrecioVehUsado"),
			req("TipoVeh"),
			req("Marca"),
			req("TipooClase"),
			req("Año"),
			opc("Modelo"),
			opc("NIV"),
			opc("NumSerie"),
			req("NumPlacas"),
			opc("NumMotor"),
			req("NumFolTarjCir"),
			req("NumFolAvisoint"),
			req("NumPedIm"),
			req("Aduana"),
			req("FechaRegulVeh"),
			req("Foliofiscal"),
		},
		// cfd/certificadodestruccion/certificadodedestruccion.xslt
		{"http://www.sat.gob.mx/certificadodestruccion", "certificadodedestruccion"}: {
			req("Version"),
			req("Serie"),
			req("NumFolDesVeh"),
			aplicar(rel(hijo("http://www.sat.gob.mx/certificadodestruccion", "VehiculoDestruido"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/certificadodestruccion", "InformacionAduanera"))),
		},
		// cfd/certificadodestruccion/certificadodedestruccion.xslt
		{"http://www.sat.gob.mx/certificadodestruccion", "VehiculoDestruido"}: {
			req("Marca"),
			req("TipooClase"),
			req("Año"),
			opc("Modelo"),
			opc("NIV"),
			opc("NumSerie"),
			req("NumPlacas"),
			opc("NumMotor"),
			req("NumFolTarjCir"),
		},
		// cfd/certificadodestruccion/certificadodedestruccion.xslt
		{"http://www.sat.gob.mx/certificadodestruccion", "InformacionAduanera"}: {
			req("NumPedImp"),
			req("Fecha"),
			req("Aduana"),
		},
		// cfd/arteantiguedades/obrasarteantiguedades.xslt
		{"http://www.sat.gob.mx/arteantiguedades", "obrasarteantiguedades"}: {
			req("Version"),
			req("TipoBien"),
			opc("OtrosTipoBien"),
			req("TituloAdquirido"),
			opc("OtrosTituloAdquirido"),
			opc("Subtotal"),
			opc("IVA"),
			req("FechaAdquisicion"),
			req("CaracterísticasDeObraoPieza"),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "ComercioExterior"}: {
			req("Version"),
			opc("MotivoTraslado"),
			req("TipoOperacion"),
			opc("ClaveDePedimento"),
			opc("CertificadoOrigen"),
			opc("NumCertificadoOrigen"),
			opc("NumeroExportadorConfiable"),
			opc("Incoterm"),
			opc("Subdivision"),
			opc("Observaciones"),
			opc("TipoCambioUSD"),
			opc("TotalUSD"),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Emisor"))),
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Propietario")),
				aplicar(rel()),
			),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Receptor"))),
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Destinatario")),
				aplicar(rel()),
			),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Mercancias"))),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Emisor"}: {
			opc("Curp"),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Domicilio"))),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Propietario"}: {
			req("NumRegIdTrib"),
			req("ResidenciaFiscal"),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Receptor"}: {
			opc("NumRegIdTrib"),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Domicilio"))),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Destinatario"}: {
			opc("NumRegIdTrib"),
			opc("Nombre"),
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Domicilio")),
				aplicar(rel()),
			),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Mercancias"}: {
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "Mercancia")),
				aplicar(rel()),
			),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Domicilio"}: {
			req("Calle"),
			opc("NumeroExterior"),
			opc("NumeroInterior"),
			opc("Colonia"),
			opc("Localidad"),
			opc("Referencia"),
			opc("Municipio"),
			req("Estado"),
			req("Pais"),
			req("CodigoPostal"),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "Mercancia"}: {
			req("NoIdentificacion"),
			opc("FraccionArancelaria"),
			opc("CantidadAduana"),
			opc("UnidadAduana"),
			opc("ValorUnitarioAduana"),
			req("ValorDolares"),
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior11", "DescripcionesEspecificas")),
				aplicar(rel()),
			),
		},
		// cfd/ComercioExterior11/ComercioExterior11.xslt
		{"http://www.sat.gob.mx/ComercioExterior11", "DescripcionesEspecificas"}: {
			req("Marca"),
			opc("Modelo"),
			opc("SubModelo"),
			opc("NumeroSerie"),
		},
		// cfd/ComercioExterior20/ComercioExterior20.xslt
		{"http://www.sat.gob.mx/ComercioExterior20", "ComercioExterior"}: {
			req("Version"),
			opc("MotivoTraslado"),
			req("ClaveDePedimento"),
			req("CertificadoOrigen"),
			opc("NumCertificadoOrigen"),
			opc("NumeroExportadorConfiable"),
			opc("Incoterm"),
			opc("Observaciones"),
			req("TipoCambioUSD"),
			req("TotalUSD"),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior20", "Emisor"))),
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior20", "Propietario")),
				aplicar(rel()),
			),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior20", "Receptor"))),
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior20", "Destinatario")),
				aplicar(rel()),
			),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior20", "Mercancias"))),
		},
		// cfd/ComercioExterior20/ComercioExterior20.xslt
		{"http://www.sat.gob.mx/ComercioExterior20", "Emisor"}: {
			opc("Curp"),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior20", "Domicilio"))),
		},
		// cfd/ComercioExterior20/ComercioExterior20.xslt
		{"http://www.sat.gob.mx/ComercioExterior20", "Propietario"}: {
			req("NumRegIdTrib"),
			req("ResidenciaFiscal"),
		},
		// cfd/ComercioExterior20/ComercioExterior20.xslt
		{"http://www.sat.gob.mx/ComercioExterior20", "Receptor"}: {
			opc("NumRegIdTrib"),
			aplicar(rel(hijo("http://www.sat.gob.mx/ComercioExterior20", "Domicilio"))),
		},
		// cfd/ComercioExterior20/ComercioExterior20.xslt
		{"http://www.sat.gob.mx/ComercioExterior20", "Destinatario"}: {
			opc("NumRegIdTrib"),
			opc("Nombre"),
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior20", "Domicilio")),
				aplicar(rel()),
			),
		},
		// cfd/ComercioExterior20/ComercioExterior20.xslt
		{"http://www.sat.gob.mx/ComercioExterior20", "Mercancias"}: {
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior20", "Mercancia")),
				aplicar(rel()),
			),
		},
		// cfd/ComercioExterior20/ComercioExterior20.xslt
		{"http://www.sat.gob.mx/ComercioExterior20", "Domicilio"}: {
			req("Calle"),
			opc("NumeroExterior"),
			opc("NumeroInterior"),
			opc("Colonia"),
			opc("Localidad"),
			opc("Referencia"),
			opc("Municipio"),
			req("Estado"),
			req("Pais"),
			req("CodigoPostal"),
		},
		// cfd/ComercioExterior20/ComercioExterior20.xslt
		{"http://www.sat.gob.mx/ComercioExterior20", "Mercancia"}: {
			req("NoIdentificacion"),
			opc("FraccionArancelaria"),
			opc("CantidadAduana"),
			opc("UnidadAduana"),
			opc("ValorUnitarioAduana"),
			req("ValorDolares"),
			cada(rel(hijo("http://www.sat.gob.mx/ComercioExterior20", "DescripcionesEspecificas")),
				aplicar(rel()),
			),
		},
		// cfd/ComercioExterior20/ComercioExterior20.xslt
		{"http://www.sat.gob.mx/ComercioExterior20", "DescripcionesEspecificas"}: {
			req("Marca"),
			opc("Modelo"),
			opc("SubModelo"),
			opc("NumeroSerie"),
		},
		// cfd/ine/ine11.xslt
		{"http://www.sat.gob.mx/ine", "INE"}: {
			req("Version"),
			req("TipoProceso"),
			opc("TipoComite"),
			opc("IdContabilidad"),
			cada(rel(hijo("http://www.sat.gob.mx/ine", "Entidad")),
				aplicar(rel()),
			),
		},
		// cfd/ine/ine11.xslt
		{"http://www.sat.gob.mx/ine", "Entidad"}: {
			req("ClaveEntidad"),
			opc("Ambito"),
			cada(rel(hijo("http://www.sat.gob.mx/ine", "Contabilidad")),
				aplicar(rel()),
			),
		},
		// cfd/ine/ine11.xslt
		{"http://www.sat.gob.mx/ine", "Contabilidad"}: {
			req("IdContabilidad"),
		},
		// cfd/iedu/iedu.xslt
		{"http://www.sat.gob.mx/iedu", "instEducativas"}: {
			req("version"),
			req("nombreAlumno"),
			req("CURP"),
			req("nivelEducativo"),
			req("autRVOE"),
			opc("rfcPago"),
		},
		// cfd/ventavehiculos/ventavehiculos11.xslt
		{"http://www.sat.gob.mx/ventavehiculos", "VentaVehiculos"}: {
			req("version"),
			req("ClaveVehicular"),
			siIgual("version", "1.1",
				req("Niv"),
			),
			cada(rel(desc("http://www.sat.gob.mx/ventavehiculos", "InformacionAduanera")),
				aplicar(rel()),
			),
		},
		// cfd/ventavehiculos/ventavehiculos11.xslt
		{"http://www.sat.gob.mx/ventavehiculos", "InformacionAduanera"}: {
			req("numero"),
			req("fecha"),
			opc("aduana"),
		},
		// cfd/EstadoDeCuentaCombustible/ecc12.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible12", "EstadoDeCuentaCombustible"}: {
			req("Version"),
			req("TipoOperacion"),
			req("NumeroDeCuenta"),
			req("SubTotal"),
			req("Total"),
			aplicar(rel(hijo("http://www.sat.gob.mx/EstadoDeCuentaCombustible12", "Conceptos"))),
		},
		// cfd/EstadoDeCuentaCombustible/ecc12.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible12", "Conceptos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/EstadoDeCuentaCombustible12", "ConceptoEstadoDeCuentaCombustible")),
				aplicar(rel()),
			),
		},
		// cfd/EstadoDeCuentaCombustible/ecc12.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible12", "Traslados"}: {
			cada(rel(hijo("http://www.sat.gob.mx/EstadoDeCuentaCombustible12", "Traslado")),
				aplicar(rel()),
			),
		},
		// cfd/EstadoDeCuentaCombustible/ecc12.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible12", "ConceptoEstadoDeCuentaCombustible"}: {
			req("Identificador"),
			req("Fecha"),
			req("Rfc"),
			req("ClaveEstacion"),
			req("Cantidad"),
			req("TipoCombustible"),
			opc("Unidad"),
			req("NombreCombustible"),
			req("FolioOperacion"),
			req("ValorUnitario"),
			req("Importe"),
			aplicar(rel(hijo("http://www.sat.gob.mx/EstadoDeCuentaCombustible12", "Traslados"))),
		},
		// cfd/EstadoDeCuentaCombustible/ecc12.xslt
		{"http://www.sat.gob.mx/EstadoDeCuentaCombustible12", "Traslado"}: {
			req("Impuesto"),
			req("TasaOCuota"),
			req("Importe"),
		},
		// cfd/consumodecombustibles/consumodeCombustibles11.xslt
		{"http://www.sat.gob.mx/ConsumoDeCombustibles11", "ConsumoDeCombustibles"}: {
			req("version"),
			req("tipoOperacion"),
			req("numeroDeCuenta"),
			opc("subTotal"),
			req("total"),
			aplicar(rel(hijo("http://www.sat.gob.mx/ConsumoDeCombustibles11", "Conceptos"))),
		},
		// cfd/consumodecombustibles/consumodeCombustibles11.xslt
		{"http://www.sat.gob.mx/ConsumoDeCombustibles11", "Conceptos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/ConsumoDeCombustibles11", "ConceptoConsumoDeCombustibles")),
				aplicar(rel()),
			),
		},
		// cfd/consumodecombustibles/consumodeCombustibles11.xslt
		{"http://www.sat.gob.mx/ConsumoDeCombustibles11", "ConceptoConsumoDeCombustibles"}: {
			req("identificador"),
			req("fecha"),
			req("rfc"),
			req("claveEstacion"),
			req("tipoCombustible"),
			req("cantidad"),
			req("nombreCombustible"),
			req("folioOperacion"),
			req("valorUnitario"),
			req("importe"),
			aplicar(rel(hijo("http://www.sat.gob.mx/ConsumoDeCombustibles11", "Determinados"))),
		},
		// cfd/consumodecombustibles/consumodeCombustibles11.xslt
		{"http://www.sat.gob.mx/ConsumoDeCombustibles11", "Determinados"}: {
			cada(rel(hijo("http://www.sat.gob.mx/ConsumoDeCombustibles11", "Determinado")),
				aplicar(rel()),
			),
		},
		// cfd/consumodecombustibles/consumodeCombustibles11.xslt
		{"http://www.sat.gob.mx/ConsumoDeCombustibles11", "Determinado"}: {
			req("impuesto"),
			req("tasaOCuota"),
			req("importe"),
		},
		// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt
		{"http://www.sat.gob.mx/GastosHidrocarburos10", "GastosHidrocarburos"}: {
			req("Version"),
			req("NumeroContrato"),
			opc("AreaContractual"),
			cada(rel(hijo("http://www.sat.gob.mx/GastosHidrocarburos10", "Erogacion")),
				aplicar(rel()),
			),
		},
		// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt
		{"http://www.sat.gob.mx/GastosHidrocarburos10", "Erogacion"}: {
			req("TipoErogacion"),
			req("MontocuErogacion"),
			req("Porcentaje"),
			cada(rel(hijo("http://www.sat.gob.mx/GastosHidrocarburos10", "DocumentoRelacionado")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/GastosHidrocarburos10", "Actividades")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/GastosHidrocarburos10", "CentroCostos")),
				aplicar(rel()),
			),
		},
		// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt
		{"http://www.sat.gob.mx/GastosHidrocarburos10", "DocumentoRelacionado"}: {
			req("OrigenErogacion"),
			opc("FolioFiscalVinculado"),
			opc("RFCProveedor"),
			req("MontoTotalIVA"),
			opc("MontoRetencionISR"),
			opc("MontoRetencionIVA"),
			opc("MontoRetencionOtrosImpuestos"),
			opc("NumeroPedimentoVinculado"),
			opc("ClavePedimentoVinculado"),
			opc("ClavePagoPedimentoVinculado"),
			opc("MontoIVAPedimento"),
			opc("OtrosImpuestosPagadosPedimento"),
			req("FechaFolioFiscalVinculado"),
			req("Mes"),
			req("MontoTotalErogaciones"),
		},
		// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt
		{"http://www.sat.gob.mx/GastosHidrocarburos10", "Actividades"}: {
			opc("ActividadRelacionada"),
			cada(rel(hijo("http://www.sat.gob.mx/GastosHidrocarburos10", "SubActividades")),
				aplicar(rel()),
			),
		},
		// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt
		{"http://www.sat.gob.mx/GastosHidrocarburos10", "SubActividades"}: {
			opc("SubActividadRelacionada"),
			cada(rel(hijo("http://www.sat.gob.mx/GastosHidrocarburos10", "Tareas")),
				aplicar(rel()),
			),
		},
		// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt
		{"http://www.sat.gob.mx/GastosHidrocarburos10", "Tareas"}: {
			opc("TareaRelacionada"),
		},
		// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt
		{"http://www.sat.gob.mx/GastosHidrocarburos10", "CentroCostos"}: {
			opc("Campo"),
			cada(rel(hijo("http://www.sat.gob.mx/GastosHidrocarburos10", "Yacimientos")),
				aplicar(rel()),
			),
		},
		// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt
		{"http://www.sat.gob.mx/GastosHidrocarburos10", "Yacimientos"}: {
			opc("Yacimiento"),
			cada(rel(hijo("http://www.sat.gob.mx/GastosHidrocarburos10", "Pozos")),
				aplicar(rel()),
			),
		},
		// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt
		{"http://www.sat.gob.mx/GastosHidrocarburos10", "Pozos"}: {
			opc("Pozo"),
		},
		// cfd/IngresosHidrocarburos10/IngresosHidrocarburos.xslt
		{"http://www.sat.gob.mx/IngresosHidrocarburos10", "IngresosHidrocarburos"}: {
			req("Version"),
			req("NumeroContrato"),
			req("ContraprestacionPagadaOperador"),
			req("Porcentaje"),
			cada(rel(hijo("http://www.sat.gob.mx/IngresosHidrocarburos10", "DocumentoRelacionado")),
				aplicar(rel()),
			),
		},
		// cfd/IngresosHidrocarburos10/IngresosHidrocarburos.xslt
		{"http://www.sat.gob.mx/IngresosHidrocarburos10", "DocumentoRelacionado"}: {
			req("FolioFiscalVinculado"),
			req("FechaFolioFiscalVinculado"),
			req("Mes"),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "CartaPorte"}: {
			req("Version"),
			req("TranspInternac"),
			opc("EntradaSalidaMerc"),
			opc("PaisOrigenDestino"),
			opc("ViaEntradaSalida"),
			opc("TotalDistRec"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Ubicaciones")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Mercancias")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "FiguraTransporte")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Ubicaciones"}: {
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Ubicacion")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Ubicacion"}: {
			req("TipoUbicacion"),
			opc("IDUbicacion"),
			req("RFCRemitenteDestinatario"),
			opc("NombreRemitenteDestinatario"),
			opc("NumRegIdTrib"),
			opc("ResidenciaFiscal"),
			opc("NumEstacion"),
			opc("NombreEstacion"),
			opc("NavegacionTrafico"),
			req("FechaHoraSalidaLlegada"),
			opc("TipoEstacion"),
			opc("DistanciaRecorrida"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Domicilio")),
				opc("Calle"),
				opc("NumeroExterior"),
				opc("NumeroInterior"),
				opc("Colonia"),
				opc("Localidad"),
				opc("Referencia"),
				opc("Municipio"),
				req("Estado"),
				req("Pais"),
				req("CodigoPostal"),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Mercancias"}: {
			req("PesoBrutoTotal"),
			req("UnidadPeso"),
			opc("PesoNetoTotal"),
			req("NumTotalMercancias"),
			opc("CargoPorTasacion"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Mercancia")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Autotransporte")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "TransporteMaritimo")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "TransporteAereo")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "TransporteFerroviario")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Mercancia"}: {
			req("BienesTransp"),
			opc("ClaveSTCC"),
			req("Descripcion"),
			req("Cantidad"),
			req("ClaveUnidad"),
			opc("Unidad"),
			opc("Dimensiones"),
			opc("MaterialPeligroso"),
			opc("CveMaterialPeligroso"),
			opc("Embalaje"),
			opc("DescripEmbalaje"),
			req("PesoEnKg"),
			opc("ValorMercancia"),
			opc("Moneda"),
			opc("FraccionArancelaria"),
			opc("UUIDComercioExt"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Pedimentos")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "GuiasIdentificacion")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "CantidadTransporta")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "DetalleMercancia")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Pedimentos"}: {
			req("Pedimento"),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "GuiasIdentificacion"}: {
			req("NumeroGuiaIdentificacion"),
			req("DescripGuiaIdentificacion"),
			req("PesoGuiaIdentificacion"),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "CantidadTransporta"}: {
			req("Cantidad"),
			req("IDOrigen"),
			req("IDDestino"),
			opc("CvesTransporte"),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "DetalleMercancia"}: {
			req("UnidadPesoMerc"),
			req("PesoBruto"),
			req("PesoNeto"),
			req("PesoTara"),
			opc("NumPiezas"),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Autotransporte"}: {
			req("PermSCT"),
			req("NumPermisoSCT"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "IdentificacionVehicular")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Seguros")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Remolques")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "IdentificacionVehicular"}: {
			req("ConfigVehicular"),
			req("PlacaVM"),
			req("AnioModeloVM"),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Seguros"}: {
			req("AseguraRespCivil"),
			req("PolizaRespCivil"),
			opc("AseguraMedAmbiente"),
			opc("PolizaMedAmbiente"),
			opc("AseguraCarga"),
			opc("PolizaCarga"),
			opc("PrimaSeguro"),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Remolques"}: {
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Remolque")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Remolque"}: {
			req("SubTipoRem"),
			req("Placa"),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "TransporteMaritimo"}: {
			opc("PermSCT"),
			opc("NumPermisoSCT"),
			opc("NombreAseg"),
			opc("NumPolizaSeguro"),
			req("TipoEmbarcacion"),
			req("Matricula"),
			req("NumeroOMI"),
			opc("AnioEmbarcacion"),
			opc("NombreEmbarc"),
			req("NacionalidadEmbarc"),
			req("UnidadesDeArqBruto"),
			req("TipoCarga"),
			req("NumCertITC"),
			opc("Eslora"),
			opc("Manga"),
			opc("Calado"),
			opc("LineaNaviera"),
			req("NombreAgenteNaviero"),
			req("NumAutorizacionNaviero"),
			opc("NumViaje"),
			opc("NumConocEmbarc"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Contenedor")),
				req("MatriculaContenedor"),
				req("TipoContenedor"),
				opc("NumPrecinto"),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "TransporteAereo"}: {
			req("PermSCT"),
			req("NumPermisoSCT"),
			opc("MatriculaAeronave"),
			opc("NombreAseg"),
			opc("NumPolizaSeguro"),
			req("NumeroGuia"),
			opc("LugarContrato"),
			req("CodigoTransportista"),
			opc("RFCEmbarcador"),
			opc("NumRegIdTribEmbarc"),
			opc("ResidenciaFiscalEmbarc"),
			opc("NombreEmbarcador"),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "TransporteFerroviario"}: {
			req("TipoDeServicio"),
			req("TipoDeTrafico"),
			opc("NombreAseg"),
			opc("NumPolizaSeguro"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "DerechosDePaso")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Carro")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "DerechosDePaso"}: {
			req("TipoDerechoDePaso"),
			req("KilometrajePagado"),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "Carro"}: {
			req("TipoCarro"),
			req("MatriculaCarro"),
			req("GuiaCarro"),
			req("ToneladasNetasCarro"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Contenedor")),
				req("TipoContenedor"),
				req("PesoContenedorVacio"),
				req("PesoNetoMercancia"),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "FiguraTransporte"}: {
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "TiposFigura")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "TiposFigura"}: {
			req("TipoFigura"),
			opc("RFCFigura"),
			opc("NumLicencia"),
			opc("NombreFigura"),
			opc("NumRegIdTribFigura"),
			opc("ResidenciaFiscalFigura"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "PartesTransporte")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte20", "Domicilio")),
				opc("Calle"),
				opc("NumeroExterior"),
				opc("NumeroInterior"),
				opc("Colonia"),
				opc("Localidad"),
				opc("Referencia"),
				opc("Municipio"),
				req("Estado"),
				req("Pais"),
				req("CodigoPostal"),
			),
		},
		// cfd/CartaPorte/CartaPorte20.xslt
		{"http://www.sat.gob.mx/CartaPorte20", "PartesTransporte"}: {
			req("ParteTransporte"),
		},
		// cfd/Pagos/Pagos20.xslt
		{"http://www.sat.gob.mx/Pagos20", "Pagos"}: {
			req("Version"),
			cada(rel(hijo("http://www.sat.gob.mx/Pagos20", "Totales")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/Pagos20", "Pago")),
				aplicar(rel()),
			),
		},
		// cfd/Pagos/Pagos20.xslt
		{"http://www.sat.gob.mx/Pagos20", "Totales"}: {
			opc("TotalRetencionesIVA"),
			opc("TotalRetencionesISR"),
			opc("TotalRetencionesIEPS"),
			opc("TotalTrasladosBaseIVA16"),
			opc("TotalTrasladosImpuestoIVA16"),
			opc("TotalTrasladosBaseIVA8"),
			opc("TotalTrasladosImpuestoIVA8"),
			opc("TotalTrasladosBaseIVA0"),
			opc("TotalTrasladosImpuestoIVA0"),
			opc("TotalTrasladosBaseIVAExento"),
			req("MontoTotalPagos"),
		},
		// cfd/Pagos/Pagos20.xslt
		{"http://www.sat.gob.mx/Pagos20", "Pago"}: {
			req("FechaPago"),
			req("FormaDePagoP"),
			req("MonedaP"),
			opc("TipoCambioP"),
			req("Monto"),
			opc("NumOperacion"),
			opc("RfcEmisorCtaOrd"),
			opc("NomBancoOrdExt"),
			opc("CtaOrdenante"),
			opc("RfcEmisorCtaBen"),
			opc("CtaBeneficiario"),
			opc("TipoCadPago"),
			opc("CertPago"),
			opc("CadPago"),
			opc("SelloPago"),
			cada(rel(hijo("http://www.sat.gob.mx/Pagos20", "DoctoRelacionado")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/Pagos20", "ImpuestosP")),
				aplicar(rel()),
			),
		},
		// cfd/Pagos/Pagos20.xslt
		{"http://www.sat.gob.mx/Pagos20", "DoctoRelacionado"}: {
			req("IdDocumento"),
			opc("Serie"),
			opc("Folio"),
			req("MonedaDR"),
			opc("EquivalenciaDR"),
			req("NumParcialidad"),
			req("ImpSaldoAnt"),
			req("ImpPagado"),
			req("ImpSaldoInsoluto"),
			req("ObjetoImpDR"),
			cada(rel(hijo("http://www.sat.gob.mx/Pagos20", "ImpuestosDR"), hijo("http://www.sat.gob.mx/Pagos20", "RetencionesDR"), hijo("http://www.sat.gob.mx/Pagos20", "RetencionDR")),
				req("BaseDR"),
				req("ImpuestoDR"),
				req("TipoFactorDR"),
				req("TasaOCuotaDR"),
				req("ImporteDR"),
			),
			cada(rel(hijo("http://www.sat.gob.mx/Pagos20", "ImpuestosDR"), hijo("http://www.sat.gob.mx/Pagos20", "TrasladosDR"), hijo("http://www.sat.gob.mx/Pagos20", "TrasladoDR")),
				req("BaseDR"),
				req("ImpuestoDR"),
				req("TipoFactorDR"),
				opc("TasaOCuotaDR"),
				opc("ImporteDR"),
			),
		},
		// cfd/Pagos/Pagos20.xslt
		{"http://www.sat.gob.mx/Pagos20", "ImpuestosP"}: {
			aplicar(rel(hijo("http://www.sat.gob.mx/Pagos20", "RetencionesP"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/Pagos20", "TrasladosP"))),
		},
		// cfd/Pagos/Pagos20.xslt
		{"http://www.sat.gob.mx/Pagos20", "RetencionesP"}: {
			cada(rel(hijo("http://www.sat.gob.mx/Pagos20", "RetencionP")),
				aplicar(rel()),
			),
		},
		// cfd/Pagos/Pagos20.xslt
		{"http://www.sat.gob.mx/Pagos20", "TrasladosP"}: {
			cada(rel(hijo("http://www.sat.gob.mx/Pagos20", "TrasladoP")),
				aplicar(rel()),
			),
		},
		// cfd/Pagos/Pagos20.xslt
		{"http://www.sat.gob.mx/Pagos20", "RetencionP"}: {
			req("ImpuestoP"),
			req("ImporteP"),
		},
		// cfd/Pagos/Pagos20.xslt
		{"http://www.sat.gob.mx/Pagos20", "TrasladoP"}: {
			req("BaseP"),
			req("ImpuestoP"),
			req("TipoFactorP"),
			opc("TasaOCuotaP"),
			opc("ImporteP"),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "CartaPorte"}: {
			req("Version"),
			req("IdCCP"),
			req("TranspInternac"),
			opc("RegimenAduanero"),
			opc("EntradaSalidaMerc"),
			opc("PaisOrigenDestino"),
			opc("ViaEntradaSalida"),
			opc("TotalDistRec"),
			opc("RegistroISTMO"),
			opc("UbicacionPoloOrigen"),
			opc("UbicacionPoloDestino"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "Ubicaciones")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "Mercancias")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "FiguraTransporte")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "Ubicaciones"}: {
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "Ubicacion")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "Ubicacion"}: {
			req("TipoUbicacion"),
			opc("IDUbicacion"),
			req("RFCRemitenteDestinatario"),
			opc("NombreRemitenteDestinatario"),
			opc("NumRegIdTrib"),
			opc("ResidenciaFiscal"),
			opc("NumEstacion"),
			opc("NombreEstacion"),
			opc("NavegacionTrafico"),
			req("FechaHoraSalidaLlegada"),
			opc("TipoEstacion"),
			opc("DistanciaRecorrida"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "Domicilio")),
				opc("Calle"),
				opc("NumeroExterior"),
				opc("NumeroInterior"),
				opc("Colonia"),
				opc("Localidad"),
				opc("Referencia"),
				opc("Municipio"),
				req("Estado"),
				req("Pais"),
				req("CodigoPostal"),
			),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "Mercancias"}: {
			req("PesoBrutoTotal"),
			req("UnidadPeso"),
			opc("PesoNetoTotal"),
			req("NumTotalMercancias"),
			opc("CargoPorTasacion"),
			opc("LogisticaInversaRecoleccionDevolucion"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "Mercancia")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "Autotransporte")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "TransporteMaritimo")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "TransporteAereo")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "TransporteFerroviario")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "Mercancia"}: {
			req("BienesTransp"),
			opc("ClaveSTCC"),
			req("Descripcion"),
			req("Cantidad"),
			req("ClaveUnidad"),
			opc("Unidad"),
			opc("Dimensiones"),
			opc("MaterialPeligroso"),
			opc("CveMaterialPeligroso"),
			opc("Embalaje"),
			opc("DescripEmbalaje"),
			opc("SectorCOFEPRIS"),
			opc("NombreIngredienteActivo"),
			opc("NomQuimico"),
			opc("DenominacionGenericaProd"),
			opc("DenominacionDistintivaProd"),
			opc("Fabricante"),
			opc("FechaCaducidad"),
			opc("LoteMedicamento"),
			opc("FormaFarmaceutica"),
			opc("CondicionesEspTransp"),
			opc("RegistroSanitarioFolioAutorizacion"),
			opc("PermisoImportacion"),
			opc("FolioImpoVUCEM"),
			opc("NumCAS"),
			opc("RazonSocialEmpImp"),
			opc("NumRegSanPlagCOFEPRIS"),
			opc("DatosFabricante"),
			opc("DatosFormulador"),
			opc("DatosMaquilador"),
			opc("UsoAutorizado"),
			req("PesoEnKg"),
			opc("ValorMercancia"),
			opc("Moneda"),
			opc("FraccionArancelaria"),
			opc("UUIDComercioExt"),
			opc("TipoMateria"),
			opc("DescripcionMateria"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "DocumentacionAduanera")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "GuiasIdentificacion")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "CantidadTransporta")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "DetalleMercancia")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "DocumentacionAduanera"}: {
			req("TipoDocumento"),
			opc("NumPedimento"),
			opc("IdentDocAduanero"),
			opc("RFCImpo"),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "GuiasIdentificacion"}: {
			req("NumeroGuiaIdentificacion"),
			req("DescripGuiaIdentificacion"),
			req("PesoGuiaIdentificacion"),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "CantidadTransporta"}: {
			req("Cantidad"),
			req("IDOrigen"),
			req("IDDestino"),
			opc("CvesTransporte"),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "DetalleMercancia"}: {
			req("UnidadPesoMerc"),
			req("PesoBruto"),
			req("PesoNeto"),
			req("PesoTara"),
			opc("NumPiezas"),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "Autotransporte"}: {
			req("PermSCT"),
			req("NumPermisoSCT"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "IdentificacionVehicular")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "Seguros")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "Remolques")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "IdentificacionVehicular"}: {
			req("ConfigVehicular"),
			req("PesoBrutoVehicular"),
			req("PlacaVM"),
			req("AnioModeloVM"),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "Seguros"}: {
			req("AseguraRespCivil"),
			req("PolizaRespCivil"),
			opc("AseguraMedAmbiente"),
			opc("PolizaMedAmbiente"),
			opc("AseguraCarga"),
			opc("PolizaCarga"),
			opc("PrimaSeguro"),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "Remolques"}: {
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "Remolque")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "Remolque"}: {
			req("SubTipoRem"),
			req("Placa"),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "TransporteMaritimo"}: {
			opc("PermSCT"),
			opc("NumPermisoSCT"),
			opc("NombreAseg"),
			opc("NumPolizaSeguro"),
			req("TipoEmbarcacion"),
			req("Matricula"),
			req("NumeroOMI"),
			opc("AnioEmbarcacion"),
			opc("NombreEmbarc"),
			req("NacionalidadEmbarc"),
			req("UnidadesDeArqBruto"),
			req("TipoCarga"),
			opc("Eslora"),
			opc("Manga"),
			opc("Calado"),
			opc("Puntal"),
			opc("LineaNaviera"),
			req("NombreAgenteNaviero"),
			req("NumAutorizacionNaviero"),
			opc("NumViaje"),
			opc("NumConocEmbarc"),
			opc("PermisoTempNavegacion"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "Contenedor")),
				req("TipoContenedor"),
				opc("MatriculaContenedor"),
				opc("NumPrecinto"),
				opc("IdCCPRelacionado"),
				opc("PlacaVMCCP"),
				opc("FechaCertificacionCCP"),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "RemolquesCCP")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "RemolquesCCP"}: {
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "RemolqueCCP")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "RemolqueCCP"}: {
			req("SubTipoRemCCP"),
			req("PlacaCCP"),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "TransporteAereo"}: {
			req("PermSCT"),
			req("NumPermisoSCT"),
			opc("MatriculaAeronave"),
			opc("NombreAseg"),
			opc("NumPolizaSeguro"),
			req("NumeroGuia"),
			opc("LugarContrato"),
			req("CodigoTransportista"),
			opc("RFCEmbarcador"),
			opc("NumRegIdTribEmbarc"),
			opc("ResidenciaFiscalEmbarc"),
			opc("NombreEmbarcador"),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "TransporteFerroviario"}: {
			req("TipoDeServicio"),
			req("TipoDeTrafico"),
			opc("NombreAseg"),
			opc("NumPolizaSeguro"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "DerechosDePaso")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "Carro")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "DerechosDePaso"}: {
			req("TipoDerechoDePaso"),
			req("KilometrajePagado"),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "Carro"}: {
			req("TipoCarro"),
			req("MatriculaCarro"),
			req("GuiaCarro"),
			req("ToneladasNetasCarro"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "Contenedor")),
				req("TipoContenedor"),
				req("PesoContenedorVacio"),
				req("PesoNetoMercancia"),
			),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "FiguraTransporte"}: {
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "TiposFigura")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "TiposFigura"}: {
			req("TipoFigura"),
			opc("RFCFigura"),
			opc("NumLicencia"),
			req("NombreFigura"),
			opc("NumRegIdTribFigura"),
			opc("ResidenciaFiscalFigura"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "PartesTransporte")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte30", "Domicilio")),
				opc("Calle"),
				opc("NumeroExterior"),
				opc("NumeroInterior"),
				opc("Colonia"),
				opc("Localidad"),
				opc("Referencia"),
				opc("Municipio"),
				req("Estado"),
				req("Pais"),
				req("CodigoPostal"),
			),
		},
		// cfd/CartaPorte/CartaPorte30.xslt
		{"http://www.sat.gob.mx/CartaPorte30", "PartesTransporte"}: {
			req("ParteTransporte"),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "CartaPorte"}: {
			req("Version"),
			req("IdCCP"),
			req("TranspInternac"),
			opc("EntradaSalidaMerc"),
			opc("PaisOrigenDestino"),
			opc("ViaEntradaSalida"),
			opc("TotalDistRec"),
			opc("RegistroISTMO"),
			opc("UbicacionPoloOrigen"),
			opc("UbicacionPoloDestino"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "RegimenesAduaneros")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "Ubicaciones")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "Mercancias")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "FiguraTransporte")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "RegimenesAduaneros"}: {
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "RegimenAduaneroCCP")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "RegimenAduaneroCCP"}: {
			req("RegimenAduanero"),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "Ubicaciones"}: {
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "Ubicacion")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "Ubicacion"}: {
			req("TipoUbicacion"),
			opc("IDUbicacion"),
			req("RFCRemitenteDestinatario"),
			opc("NombreRemitenteDestinatario"),
			opc("NumRegIdTrib"),
			opc("ResidenciaFiscal"),
			opc("NumEstacion"),
			opc("NombreEstacion"),
			opc("NavegacionTrafico"),
			req("FechaHoraSalidaLlegada"),
			opc("TipoEstacion"),
			opc("DistanciaRecorrida"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "Domicilio")),
				opc("Calle"),
				opc("NumeroExterior"),
				opc("NumeroInterior"),
				opc("Colonia"),
				opc("Localidad"),
				opc("Referencia"),
				opc("Municipio"),
				req("Estado"),
				req("Pais"),
				req("CodigoPostal"),
			),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "Mercancias"}: {
			req("PesoBrutoTotal"),
			req("UnidadPeso"),
			opc("PesoNetoTotal"),
			req("NumTotalMercancias"),
			opc("CargoPorTasacion"),
			opc("LogisticaInversaRecoleccionDevolucion"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "Mercancia")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "Autotransporte")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "TransporteMaritimo")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "TransporteAereo")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "TransporteFerroviario")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "Mercancia"}: {
			req("BienesTransp"),
			opc("ClaveSTCC"),
			req("Descripcion"),
			req("Cantidad"),
			req("ClaveUnidad"),
			opc("Unidad"),
			opc("Dimensiones"),
			opc("MaterialPeligroso"),
			opc("CveMaterialPeligroso"),
			opc("Embalaje"),
			opc("DescripEmbalaje"),
			opc("SectorCOFEPRIS"),
			opc("NombreIngredienteActivo"),
			opc("NomQuimico"),
			opc("DenominacionGenericaProd"),
			opc("DenominacionDistintivaProd"),
			opc("Fabricante"),
			opc("FechaCaducidad"),
			opc("LoteMedicamento"),
			opc("FormaFarmaceutica"),
			opc("CondicionesEspTransp"),
			opc("RegistroSanitarioFolioAutorizacion"),
			opc("PermisoImportacion"),
			opc("FolioImpoVUCEM"),
			opc("NumCAS"),
			opc("RazonSocialEmpImp"),
			opc("NumRegSanPlagCOFEPRIS"),
			opc("DatosFabricante"),
			opc("DatosFormulador"),
			opc("DatosMaquilador"),
			opc("UsoAutorizado"),
			req("PesoEnKg"),
			opc("ValorMercancia"),
			opc("Moneda"),
			opc("FraccionArancelaria"),
			opc("UUIDComercioExt"),
			opc("TipoMateria"),
			opc("DescripcionMateria"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "DocumentacionAduanera")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "GuiasIdentificacion")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "CantidadTransporta")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "DetalleMercancia")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "DocumentacionAduanera"}: {
			req("TipoDocumento"),
			opc("NumPedimento"),
			opc("IdentDocAduanero"),
			opc("RFCImpo"),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "GuiasIdentificacion"}: {
			req("NumeroGuiaIdentificacion"),
			req("DescripGuiaIdentificacion"),
			req("PesoGuiaIdentificacion"),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "CantidadTransporta"}: {
			req("Cantidad"),
			req("IDOrigen"),
			req("IDDestino"),
			opc("CvesTransporte"),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "DetalleMercancia"}: {
			req("UnidadPesoMerc"),
			req("PesoBruto"),
			req("PesoNeto"),
			req("PesoTara"),
			opc("NumPiezas"),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "Autotransporte"}: {
			req("PermSCT"),
			req("NumPermisoSCT"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "IdentificacionVehicular")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "Seguros")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "Remolques")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "IdentificacionVehicular"}: {
			req("ConfigVehicular"),
			req("PesoBrutoVehicular"),
			req("PlacaVM"),
			req("AnioModeloVM"),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "Seguros"}: {
			req("AseguraRespCivil"),
			req("PolizaRespCivil"),
			opc("AseguraMedAmbiente"),
			opc("PolizaMedAmbiente"),
			opc("AseguraCarga"),
			opc("PolizaCarga"),
			opc("PrimaSeguro"),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "Remolques"}: {
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "Remolque")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "Remolque"}: {
			req("SubTipoRem"),
			req("Placa"),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "TransporteMaritimo"}: {
			opc("PermSCT"),
			opc("NumPermisoSCT"),
			opc("NombreAseg"),
			opc("NumPolizaSeguro"),
			req("TipoEmbarcacion"),
			req("Matricula"),
			req("NumeroOMI"),
			opc("AnioEmbarcacion"),
			opc("NombreEmbarc"),
			req("NacionalidadEmbarc"),
			req("UnidadesDeArqBruto"),
			req("TipoCarga"),
			opc("Eslora"),
			opc("Manga"),
			opc("Calado"),
			opc("Puntal"),
			opc("LineaNaviera"),
			req("NombreAgenteNaviero"),
			req("NumAutorizacionNaviero"),
			opc("NumViaje"),
			opc("NumConocEmbarc"),
			opc("PermisoTempNavegacion"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "Contenedor")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "Contenedor"}: {
			req("TipoContenedor"),
			opc("MatriculaContenedor"),
			opc("NumPrecinto"),
			opc("IdCCPRelacionado"),
			opc("PlacaVMCCP"),
			opc("FechaCertificacionCCP"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "RemolquesCCP")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "RemolquesCCP"}: {
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "RemolqueCCP")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "RemolqueCCP"}: {
			req("SubTipoRemCCP"),
			req("PlacaCCP"),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "TransporteAereo"}: {
			req("PermSCT"),
			req("NumPermisoSCT"),
			opc("MatriculaAeronave"),
			opc("NombreAseg"),
			opc("NumPolizaSeguro"),
			req("NumeroGuia"),
			opc("LugarContrato"),
			req("CodigoTransportista"),
			opc("RFCEmbarcador"),
			opc("NumRegIdTribEmbarc"),
			opc("ResidenciaFiscalEmbarc"),
			opc("NombreEmbarcador"),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "TransporteFerroviario"}: {
			req("TipoDeServicio"),
			req("TipoDeTrafico"),
			opc("NombreAseg"),
			opc("NumPolizaSeguro"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "DerechosDePaso")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "Carro")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "DerechosDePaso"}: {
			req("TipoDerechoDePaso"),
			req("KilometrajePagado"),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "Carro"}: {
			req("TipoCarro"),
			req("MatriculaCarro"),
			req("GuiaCarro"),
			req("ToneladasNetasCarro"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "Contenedor")),
				req("TipoContenedor"),
				req("PesoContenedorVacio"),
				req("PesoNetoMercancia"),
			),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "FiguraTransporte"}: {
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "TiposFigura")),
				aplicar(rel()),
			),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "TiposFigura"}: {
			req("TipoFigura"),
			opc("RFCFigura"),
			opc("NumLicencia"),
			req("NombreFigura"),
			opc("NumRegIdTribFigura"),
			opc("ResidenciaFiscalFigura"),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "PartesTransporte")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/CartaPorte31", "Domicilio")),
				opc("Calle"),
				opc("NumeroExterior"),
				opc("NumeroInterior"),
				opc("Colonia"),
				opc("Localidad"),
				opc("Referencia"),
				opc("Municipio"),
				req("Estado"),
				req("Pais"),
				req("CodigoPostal"),
			),
		},
		// cfd/CartaPorte/CartaPorte31.xslt
		{"http://www.sat.gob.mx/CartaPorte31", "PartesTransporte"}: {
			req("ParteTransporte"),
		},
		// cfd/HidrocarburosPetro/hidrocarburospetroliferos.xslt
		{"http://www.sat.gob.mx/hidrocarburospetroliferos", "HidroYPetro"}: {
			req("Version"),
			req("TipoPermiso"),
			req("NumeroPermiso"),
			req("ClaveHYP"),
			req("SubProductoHYP"),
		},
		// cfd/4/cadenaoriginal_4_0/cadenaoriginal_4_0.xslt
		{"http://www.sat.gob.mx/cfd/4", "Comprobante"}: {
			req("Version"),
			opc("Serie"),
			opc("Folio"),
			req("Fecha"),
			opc("FormaPago"),
			req("NoCertificado"),
			opc("CondicionesDePago"),
			req("SubTotal"),
			opc("Descuento"),
			req("Moneda"),
			opc("TipoCambio"),
			req("Total"),
			req("TipoDeComprobante"),
			req("Exportacion"),
			opc("MetodoPago"),
			req("LugarExpedicion"),
			opc("Confirmacion"),
			aplicar(rel(hijo("http://www.sat.gob.mx/cfd/4", "InformacionGlobal"))),
			cada(rel(hijo("http://www.sat.gob.mx/cfd/4", "CfdiRelacionados")),
				aplicar(rel()),
			),
			aplicar(rel(hijo("http://www.sat.gob.mx/cfd/4", "Emisor"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/cfd/4", "Receptor"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/cfd/4", "Conceptos"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/cfd/4", "Impuestos"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/cfd/4", "Complemento"))),
		},
		// cfd/4/cadenaoriginal_4_0/cadenaoriginal_4_0.xslt
		{"http://www.sat.gob.mx/cfd/4", "InformacionGlobal"}: {
			req("Periodicidad"),
			req("Meses"),
			req("Año"),
		},
		// cfd/4/cadenaoriginal_4_0/cadenaoriginal_4_0.xslt
		{"http://www.sat.gob.mx/cfd/4", "CfdiRelacionados"}: {
			req("TipoRelacion"),
			cada(rel(hijo("http://www.sat.gob.mx/cfd/4", "CfdiRelacionado")),
				req("UUID"),
			),
		},
		// cfd/4/cadenaoriginal_4_0/cadenaoriginal_4_0.xslt
		{"http://www.sat.gob.mx/cfd/4", "Emisor"}: {
			req("Rfc"),
			req("Nombre"),
			req("RegimenFiscal"),
			opc("FacAtrAdquirente"),
		},
		// cfd/4/cadenaoriginal_4_0/cadenaoriginal_4_0.xslt
		{"http://www.sat.gob.mx/cfd/4", "Receptor"}: {
			req("Rfc"),
			req("Nombre"),
			req("DomicilioFiscalReceptor"),
			opc("ResidenciaFiscal"),
			opc("NumRegIdTrib"),
			req("RegimenFiscalReceptor"),
			req("UsoCFDI"),
		},
		// cfd/4/cadenaoriginal_4_0/cadenaoriginal_4_0.xslt
		{"http://www.sat.gob.mx/cfd/4", "Conceptos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/cfd/4", "Concepto")),
				aplicar(rel()),
			),
		},
		// cfd/4/cadenaoriginal_4_0/cadenaoriginal_4_0.xslt
		{"http://www.sat.gob.mx/cfd/4", "Concepto"}: {
			req("ClaveProdServ"),
			opc("NoIdentificacion"),
			req("Cantidad"),
			req("ClaveUnidad"),
			opc("Unidad"),
			req("Descripcion"),
			req("ValorUnitario"),
			req("Importe"),
			opc("Descuento"),
			req("ObjetoImp"),
			cada(rel(hijo("http://www.sat.gob.mx/cfd/4", "Impuestos"), hijo("http://www.sat.gob.mx/cfd/4", "Traslados"), hijo("http://www.sat.gob.mx/cfd/4", "Traslado")),
				req("Base"),
				req("Impuesto"),
				req("TipoFactor"),
				opc("TasaOCuota"),
				opc("Importe"),
			),
			cada(rel(hijo("http://www.sat.gob.mx/cfd/4", "Impuestos"), hijo("http://www.sat.gob.mx/cfd/4", "Retenciones"), hijo("http://www.sat.gob.mx/cfd/4", "Retencion")),
				req("Base"),
				req("Impuesto"),
				req("TipoFactor"),
				req("TasaOCuota"),
				req("Importe"),
			),
			cada(rel(hijo("http://www.sat.gob.mx/cfd/4", "ACuentaTerceros")),
				aplicar(rel()),
			),
			cada(rel(hijo("http://www.sat.gob.mx/cfd/4", "InformacionAduanera")),
				aplicar(rel()),
			),
			si(rel(hijo("http://www.sat.gob.mx/cfd/4", "CuentaPredial")),
				aplicar(rel(hijo("http://www.sat.gob.mx/cfd/4", "CuentaPredial"))),
			),
			si(rel(hijo("http://www.sat.gob.mx/cfd/4", "ComplementoConcepto")),
				aplicar(rel(hijo("http://www.sat.gob.mx/cfd/4", "ComplementoConcepto"))),
			),
			cada(rel(desc("http://www.sat.gob.mx/cfd/4", "Parte")),
				aplicar(rel()),
			),
		},
		// cfd/4/cadenaoriginal_4_0/cadenaoriginal_4_0.xslt
		{"http://www.sat.gob.mx/cfd/4", "ACuentaTerceros"}: {
			req("RfcACuentaTerceros"),
			req("NombreACuentaTerceros"),
			req("RegimenFiscalACuentaTerceros"),
			req("DomicilioFiscalACuentaTerceros"),
		},
		// cfd/4/cadenaoriginal_4_0/cadenaoriginal_4_0.xslt
		{"http://www.sat.gob.mx/cfd/4", "InformacionAduanera"}: {
			req("NumeroPedimento"),
		},
		// cfd/4/cadenaoriginal_4_0/cadenaoriginal_4_0.xslt
		{"http://www.sat.gob.mx/cfd/4", "CuentaPredial"}: {
			req("Numero"),
		},
		// cfd/4/cadenaoriginal_4_0/cadenaoriginal_4_0.xslt
		{"http://www.sat.gob.mx/cfd/4", "ComplementoConcepto"}: {
			cada(rel(hijo("", "*")),
				aplicar(rel()),
			),
		},
		// cfd/4/cadenaoriginal_4_0/cadenaoriginal_4_0.xslt
		{"http://www.sat.gob.mx/cfd/4", "Parte"}: {
			req("ClaveProdServ"),
			opc("NoIdentificacion"),
			req("Cantidad"),
			opc("Unidad"),
			req("Descripcion"),
			opc("ValorUnitario"),
			opc("Importe"),
			cada(rel(desc("http://www.sat.gob.mx/cfd/4", "InformacionAduanera")),
				aplicar(rel()),
			),
		},
		// cfd/4/cadenaoriginal_4_0/cadenaoriginal_4_0.xslt
		{"http://www.sat.gob.mx/cfd/4", "Complemento"}: {
			cada(rel(hijo("", "*")),
				aplicar(rel()),
			),
		},
		// cfd/4/cadenaoriginal_4_0/cadenaoriginal_4_0.xslt
		{"http://www.sat.gob.mx/cfd/4", "Impuestos"}: {
			cada(rel(hijo("http://www.sat.gob.mx/cfd/4", "Retenciones"), hijo("http://www.sat.gob.mx/cfd/4", "Retencion")),
				req("Impuesto"),
				req("Importe"),
			),
			opc("TotalImpuestosRetenidos"),
			cada(rel(hijo("http://www.sat.gob.mx/cfd/4", "Traslados"), hijo("http://www.sat.gob.mx/cfd/4", "Traslado")),
				req("Base"),
				req("Impuesto"),
				req("TipoFactor"),
				opc("TasaOCuota"),
				opc("Importe"),
			),
			opc("TotalImpuestosTrasladados"),
		},
	},
}

// hojaTFD10 sale de timbrefiscaldigital/cadenaoriginal_TFD_1_0.xslt y de lo que incluye.
//
// Archivos (sha256):
//   - timbrefiscaldigital/cadenaoriginal_TFD_1_0.xslt ba786b756baaceac4675282b60b793f94da51a4a703bf565bffbeed73289ce55
var hojaTFD10 = &hojaCadena{
	nombre: "cadenaoriginal_TFD_1_0",
	raiz: []pasoCadena{
		txt("|"),
		aplicar(abs(hijo("http://www.sat.gob.mx/TimbreFiscalDigital", "TimbreFiscalDigital"))),
		txt("||"),
	},
	plantillas: map[nombreXML][]pasoCadena{
		// timbrefiscaldigital/cadenaoriginal_TFD_1_0.xslt
		{"http://www.sat.gob.mx/TimbreFiscalDigital", "TimbreFiscalDigital"}: {
			req("version"),
			req("UUID"),
			req("FechaTimbrado"),
			req("selloCFD"),
			req("noCertificadoSAT"),
		},
	},
}

// hojaTFD11 sale de cfd/TimbreFiscalDigital/cadenaoriginal_TFD_1_1.xslt y de lo que incluye.
//
// Archivos (sha256):
//   - cfd/TimbreFiscalDigital/cadenaoriginal_TFD_1_1.xslt a1be6a85c4ac9e739ec9982987b36f2aac5dcf43eff41a332c59fa6b9a5a7126
var hojaTFD11 = &hojaCadena{
	nombre: "cadenaoriginal_TFD_1_1",
	raiz: []pasoCadena{
		txt("|"),
		aplicar(abs(hijo("http://www.sat.gob.mx/TimbreFiscalDigital", "TimbreFiscalDigital"))),
		txt("||"),
	},
	plantillas: map[nombreXML][]pasoCadena{
		// cfd/TimbreFiscalDigital/cadenaoriginal_TFD_1_1.xslt
		{"http://www.sat.gob.mx/TimbreFiscalDigital", "TimbreFiscalDigital"}: {
			req("Version"),
			req("UUID"),
			req("FechaTimbrado"),
			req("RfcProvCertif"),
			opc("Leyenda"),
			req("SelloCFD"),
			req("NoCertificadoSAT"),
		},
	},
}

// hojaRetenciones10 sale de ../esquemas/retencionpago/1/retenciones.xslt y de lo que incluye.
//
// Archivos (sha256):
//   - ../esquemas/retencionpago/1/retenciones.xslt 30fefa9427bed22d09d8f61926c2fe0153d5df572161b88e161da41a39015336
//   - cfd/2/cadenaoriginal_2_0/utilerias.xslt d9ef68267747e0f6725295d4012684a4c051cb3f2265dcb2035d32bf40895016
//   - ../esquemas/retencionpago/1/arrendamientoenfideicomiso/arrendamientoenfideicomiso.xslt f99691b8b6399bf09202b47e2d2c428181c54ff87788d925f25563d73c1c12a2
//   - ../esquemas/retencionpago/1/dividendos/dividendos.xslt 26925738c0872fde8f1168a9f9cd853bb3f71c781351c5be365e301e2e8388d0
//   - ../esquemas/retencionpago/1/enajenaciondeacciones/enajenaciondeacciones.xslt 999dbdbe864d6f73b2f667ba51d05b9a618c10bbec6b17d3763996f9cbb67a1f
//   - ../esquemas/retencionpago/1/fideicomisonoempresarial/fideicomisonoempresarial.xslt 2137b4185663e614f048c8831a267d82fc8217fc88470016ba2bfae96ebf1a2b
//   - ../esquemas/retencionpago/1/intereses/intereses.xslt a78097f3a6c652456dc84f56548b88bbbfbfc471a939686a661d902ca91cec37
//   - ../esquemas/retencionpago/1/intereseshipotecarios/intereseshipotecarios.xslt d41cc75a02163d830798b1f232619ea314db1906f4274f088dd46f1b1b911c3b
//   - ../esquemas/retencionpago/1/pagosaextranjeros/pagosaextranjeros.xslt 6328ceece28199899f8f0a3bc5a67c407ba1f242a77bc996b43fb8c9855ec114
//   - ../esquemas/retencionpago/1/planesderetiro/planesderetiro.xslt be0614e3f443140d5ce30f3b09cc4967f4f4b5f0f926227f7810c5d0e65e80a9
//   - ../esquemas/retencionpago/1/premios/premios.xslt 48cb784e1e1116f447b2fb5f53bfb33458bcf48067dacce65474a0c209d73a25
//   - ../esquemas/retencionpago/1/operacionesconderivados/operacionesconderivados.xslt e72ed73fb0b9ae3b4143b07347b73d2de6a2f34657d8157e874b4aad85e77c5d
//   - ../esquemas/retencionpago/1/sectorfinanciero/sectorfinanciero.xslt 2d5575b989462a31a4415b54fca08e2b658daa5f45d2de6fc2347e4481b18468
//   - ../esquemas/retencionpago/1/planesderetiro11/planesderetiro11.xslt 7b7cca0c22ec8bcdc4a759c1f61c5822bd9169829e3ab718495037f90c22379f
//   - ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xslt a92c363dedf55bd4f066cd08d964a57b7ba4a1962fd1d283b3fe567634aeda83
var hojaRetenciones10 = &hojaCadena{
	nombre: "retenciones",
	raiz: []pasoCadena{
		txt("|"),
		aplicar(abs(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1", "Retenciones"))),
		txt("||"),
	},
	plantillas: map[nombreXML][]pasoCadena{
		// ../esquemas/retencionpago/1/arrendamientoenfideicomiso/arrendamientoenfideicomiso.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/arrendamientoenfideicomiso", "Arrendamientoenfideicomiso"}: {
			req("Version"),
			req("PagProvEfecPorFiduc"),
			req("RendimFideicom"),
			req("DeduccCorresp"),
			opc("MontTotRet"),
			opc("MontResFiscDistFibras"),
			opc("MontOtrosConceptDistr"),
			opc("DescrMontOtrosConceptDistr"),
		},
		// ../esquemas/retencionpago/1/dividendos/dividendos.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/dividendos", "Dividendos"}: {
			req("Version"),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/dividendos", "DividOUtil"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/dividendos", "Remanente"))),
		},
		// ../esquemas/retencionpago/1/dividendos/dividendos.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/dividendos", "DividOUtil"}: {
			req("CveTipDivOUtil"),
			req("MontISRAcredRetMexico"),
			req("MontISRAcredRetExtranjero"),
			opc("MontRetExtDivExt"),
			req("TipoSocDistrDiv"),
			opc("MontISRAcredNal"),
			opc("MontDivAcumNal"),
			opc("MontDivAcumExt"),
		},
		// ../esquemas/retencionpago/1/dividendos/dividendos.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/dividendos", "Remanente"}: {
			opc("ProporcionRem"),
		},
		// ../esquemas/retencionpago/1/enajenaciondeacciones/enajenaciondeacciones.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/enajenaciondeacciones", "EnajenaciondeAcciones"}: {
			req("Version"),
			req("ContratoIntermediacion"),
			req("Ganancia"),
			req("Perdida"),
		},
		// ../esquemas/retencionpago/1/fideicomisonoempresarial/fideicomisonoempresarial.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "Fideicomisonoempresarial"}: {
			req("Version"),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "IngresosOEntradas"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "DeduccOSalidas"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "RetEfectFideicomiso"))),
		},
		// ../esquemas/retencionpago/1/fideicomisonoempresarial/fideicomisonoempresarial.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "IngresosOEntradas"}: {
			req("MontTotEntradasPeriodo"),
			req("PartPropAcumDelFideicom"),
			req("PropDelMontTot"),
			si(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "IntegracIngresos")),
				conjunto(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "IntegracIngresos")),
					req("Concepto"),
				),
			),
		},
		// ../esquemas/retencionpago/1/fideicomisonoempresarial/fideicomisonoempresarial.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "DeduccOSalidas"}: {
			req("MontTotEgresPeriodo"),
			req("PartPropDelFideicom"),
			req("PropDelMontTot"),
			si(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "IntegracEgresos")),
				conjunto(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "IntegracEgresos")),
					req("ConceptoS"),
				),
			),
		},
		// ../esquemas/retencionpago/1/fideicomisonoempresarial/fideicomisonoempresarial.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "RetEfectFideicomiso"}: {
			req("MontRetRelPagFideic"),
			req("DescRetRelPagFideic"),
		},
		// ../esquemas/retencionpago/1/intereses/intereses.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/intereses", "Intereses"}: {
			req("Version"),
			req("SistFinanciero"),
			req("RetiroAORESRetInt"),
			req("OperFinancDerivad"),
			req("MontIntNominal"),
			req("MontIntReal"),
			req("Perdida"),
		},
		// ../esquemas/retencionpago/1/intereseshipotecarios/intereseshipotecarios.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/intereseshipotecarios", "Intereseshipotecarios"}: {
			req("Version"),
			req("CreditoDeInstFinanc"),
			req("SaldoInsoluto"),
			opc("PropDeducDelCredit"),
			opc("MontTotIntNominalesDev"),
			opc("MontTotIntNominalesDevYPag"),
			opc("MontTotIntRealPagDeduc"),
			opc("NumContrato"),
		},
		// ../esquemas/retencionpago/1/pagosaextranjeros/pagosaextranjeros.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/pagosaextranjeros", "Pagosaextranjeros"}: {
			req("Version"),
			req("EsBenefEfectDelCobro"),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/pagosaextranjeros", "NoBeneficiario"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/pagosaextranjeros", "Beneficiario"))),
		},
		// ../esquemas/retencionpago/1/pagosaextranjeros/pagosaextranjeros.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/pagosaextranjeros", "NoBeneficiario"}: {
			req("PaisDeResidParaEfecFisc"),
			req("ConceptoPago"),
			req("DescripcionConcepto"),
		},
		// ../esquemas/retencionpago/1/pagosaextranjeros/pagosaextranjeros.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/pagosaextranjeros", "Beneficiario"}: {
			req("RFC"),
			req("CURP"),
			req("NomDenRazSocB"),
			req("ConceptoPago"),
			req("DescripcionConcepto"),
		},
		// ../esquemas/retencionpago/1/planesderetiro/planesderetiro.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/planesderetiro", "Planesderetiro"}: {
			req("Version"),
			req("SistemaFinanc"),
			opc("MontTotAportAnioInmAnterior"),
			req("MontIntRealesDevengAniooInmAnt"),
			req("HuboRetirosAnioInmAntPer"),
			opc("MontTotRetiradoAnioInmAntPer"),
			opc("MontTotExentRetiradoAnioInmAnt"),
			opc("MontTotExedenteAnioInmAnt"),
			req("HuboRetirosAnioInmAnt"),
			opc("MontTotRetiradoAnioInmAnt"),
		},
		// ../esquemas/retencionpago/1/premios/premios.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/premios", "Premios"}: {
			req("Version"),
			req("EntidadFederativa"),
			req("MontTotPago"),
			req("MontTotPagoGrav"),
			req("MontTotPagoExent"),
		},
		// ../esquemas/retencionpago/1/operacionesconderivados/operacionesconderivados.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/operacionesconderivados", "Operacionesconderivados"}: {
			req("Version"),
			req("MontGanAcum"),
			req("MontPerdDed"),
		},
		// ../esquemas/retencionpago/1/sectorfinanciero/sectorfinanciero.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/sectorfinanciero", "SectorFinanciero"}: {
			req("Version"),
			req("IdFideicom"),
			opc("NomFideicom"),
			req("DescripFideicom"),
		},
		// ../esquemas/retencionpago/1/planesderetiro11/planesderetiro11.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/planesderetiro11", "Planesderetiro"}: {
			req("Version"),
			req("SistemaFinanc"),
			opc("MontTotAportAnioInmAnterior"),
			req("MontIntRealesDevengAniooInmAnt"),
			req("HuboRetirosAnioInmAntPer"),
			opc("MontTotRetiradoAnioInmAntPer"),
			opc("MontTotExentRetiradoAnioInmAnt"),
			opc("MontTotExedenteAnioInmAnt"),
			req("HuboRetirosAnioInmAnt"),
			opc("MontTotRetiradoAnioInmAnt"),
			opc("NumReferencia"),
			cada(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/planesderetiro11", "AportacionesODepositos")),
				aplicar(rel()),
			),
		},
		// ../esquemas/retencionpago/1/planesderetiro11/planesderetiro11.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/planesderetiro11", "AportacionesODepositos"}: {
			req("TipoAportacionODeposito"),
			req("MontAportODep"),
			opc("RFCFiduciaria"),
		},
		// ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ServiciosPlataformasTecnologicas"}: {
			req("Version"),
			req("Periodicidad"),
			req("NumServ"),
			req("MonTotServSIVA"),
			req("TotalIVATrasladado"),
			req("TotalIVARetenido"),
			req("TotalISRRetenido"),
			req("DifIVAEntregadoPrestServ"),
			req("MonTotalporUsoPlataforma"),
			opc("MonTotalContribucionGubernamental"),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "Servicios"))),
		},
		// ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "Servicios"}: {
			cada(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "DetallesDelServicio")),
				aplicar(rel()),
			),
		},
		// ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "DetallesDelServicio"}: {
			req("FormaPagoServ"),
			req("TipoDeServ"),
			opc("SubTipServ"),
			opc("RFCTerceroAutorizado"),
			req("FechaServ"),
			req("PrecioServSinIVA"),
			si(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ImpuestosTrasladadosdelServicio")),
				aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ImpuestosTrasladadosdelServicio"))),
			),
			si(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ContribucionGubernamental")),
				aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ContribucionGubernamental"))),
			),
			si(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ComisionDelServicio")),
				aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ComisionDelServicio"))),
			),
		},
		// ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ImpuestosTrasladadosdelServicio"}: {
			req("Base"),
			req("Impuesto"),
			req("TipoFactor"),
			req("TasaCuota"),
			req("Importe"),
		},
		// ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ContribucionGubernamental"}: {
			req("ImpContrib"),
			req("EntidadDondePagaLaContribucion"),
		},
		// ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ComisionDelServicio"}: {
			opc("Base"),
			opc("Porcentaje"),
			req("Importe"),
		},
		// ../esquemas/retencionpago/1/retenciones.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1", "Retenciones"}: {
			req("Version"),
			req("NumCert"),
			opc("FolioInt"),
			req("FechaExp"),
			req("CveRetenc"),
			opc("DescRetenc"),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1", "Emisor"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1", "Receptor"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1", "Periodo"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1", "Totales"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1", "Complemento"))),
		},
		// ../esquemas/retencionpago/1/retenciones.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1", "Emisor"}: {
			req("RFCEmisor"),
			opc("NomDenRazSocE"),
			opc("CURPE"),
		},
		// ../esquemas/retencionpago/1/retenciones.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1", "Receptor"}: {
			req("Nacionalidad"),
			si(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1", "Nacional")),
				conjunto(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1", "Nacional")),
					req("RFCRecep"),
				),
			),
			si(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1", "Nacional")),
				conjunto(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1", "Nacional")),
					opc("NomDenRazSocR"),
				),
			),
			si(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1", "Nacional")),
				conjunto(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1", "Nacional")),
					opc("CURPR"),
				),
			),
			si(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1", "Extranjero")),
				conjunto(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1", "Extranjero")),
					opc("NumRegIdTrib"),
				),
			),
			si(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1", "Extranjero")),
				conjunto(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1", "Extranjero")),
					req("NomDenRazSocR"),
				),
			),
		},
		// ../esquemas/retencionpago/1/retenciones.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1", "Periodo"}: {
			req("MesIni"),
			req("MesFin"),
			req("Ejerc"),
		},
		// ../esquemas/retencionpago/1/retenciones.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1", "Totales"}: {
			req("montoTotOperacion"),
			req("montoTotGrav"),
			req("montoTotExent"),
			req("montoTotRet"),
			cada(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1", "ImpRetenidos")),
				aplicar(rel()),
				opc("BaseRet"),
				opc("Impuesto"),
				req("montoRet"),
				req("TipoPagoRet"),
			),
		},
		// ../esquemas/retencionpago/1/retenciones.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1", "Complemento"}: {
			cada(rel(hijo("", "*")),
				aplicar(rel()),
			),
		},
	},
	integradas: map[nombreXML]bool{
		{"http://www.sat.gob.mx/esquemas/retencionpago/1", "ImpRetenidos"}: true,
	},
}

// hojaRetenciones20 sale de ../esquemas/retencionpago/2/retenciones.xslt y de lo que incluye.
//
// Archivos (sha256):
//   - ../esquemas/retencionpago/2/retenciones.xslt ffe2298fb076a798e13e7e7f1daffaaff24110ebdfbd0df87083ba3c6529ac76
//   - cfd/2/cadenaoriginal_2_0/utilerias.xslt d9ef68267747e0f6725295d4012684a4c051cb3f2265dcb2035d32bf40895016
//   - ../esquemas/retencionpago/1/arrendamientoenfideicomiso/arrendamientoenfideicomiso.xslt f99691b8b6399bf09202b47e2d2c428181c54ff87788d925f25563d73c1c12a2
//   - ../esquemas/retencionpago/1/dividendos/dividendos.xslt 26925738c0872fde8f1168a9f9cd853bb3f71c781351c5be365e301e2e8388d0
//   - ../esquemas/retencionpago/1/enajenaciondeacciones/enajenaciondeacciones.xslt 999dbdbe864d6f73b2f667ba51d05b9a618c10bbec6b17d3763996f9cbb67a1f
//   - ../esquemas/retencionpago/1/fideicomisonoempresarial/fideicomisonoempresarial.xslt 2137b4185663e614f048c8831a267d82fc8217fc88470016ba2bfae96ebf1a2b
//   - ../esquemas/retencionpago/1/intereses/intereses.xslt a78097f3a6c652456dc84f56548b88bbbfbfc471a939686a661d902ca91cec37
//   - ../esquemas/retencionpago/1/intereseshipotecarios/intereseshipotecarios.xslt d41cc75a02163d830798b1f232619ea314db1906f4274f088dd46f1b1b911c3b
//   - ../esquemas/retencionpago/1/pagosaextranjeros/pagosaextranjeros.xslt 6328ceece28199899f8f0a3bc5a67c407ba1f242a77bc996b43fb8c9855ec114
//   - ../esquemas/retencionpago/1/premios/premios.xslt 48cb784e1e1116f447b2fb5f53bfb33458bcf48067dacce65474a0c209d73a25
//   - ../esquemas/retencionpago/1/operacionesconderivados/operacionesconderivados.xslt e72ed73fb0b9ae3b4143b07347b73d2de6a2f34657d8157e874b4aad85e77c5d
//   - ../esquemas/retencionpago/1/sectorfinanciero/sectorfinanciero.xslt 2d5575b989462a31a4415b54fca08e2b658daa5f45d2de6fc2347e4481b18468
//   - ../esquemas/retencionpago/1/planesderetiro11/planesderetiro11.xslt 7b7cca0c22ec8bcdc4a759c1f61c5822bd9169829e3ab718495037f90c22379f
//   - ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xslt a92c363dedf55bd4f066cd08d964a57b7ba4a1962fd1d283b3fe567634aeda83
var hojaRetenciones20 = &hojaCadena{
	nombre: "retenciones",
	raiz: []pasoCadena{
		txt("|"),
		aplicar(abs(hijo("http://www.sat.gob.mx/esquemas/retencionpago/2", "Retenciones"))),
		txt("||"),
	},
	plantillas: map[nombreXML][]pasoCadena{
		// ../esquemas/retencionpago/1/arrendamientoenfideicomiso/arrendamientoenfideicomiso.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/arrendamientoenfideicomiso", "Arrendamientoenfideicomiso"}: {
			req("Version"),
			req("PagProvEfecPorFiduc"),
			req("RendimFideicom"),
			req("DeduccCorresp"),
			opc("MontTotRet"),
			opc("MontResFiscDistFibras"),
			opc("MontOtrosConceptDistr"),
			opc("DescrMontOtrosConceptDistr"),
		},
		// ../esquemas/retencionpago/1/dividendos/dividendos.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/dividendos", "Dividendos"}: {
			req("Version"),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/dividendos", "DividOUtil"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/dividendos", "Remanente"))),
		},
		// ../esquemas/retencionpago/1/dividendos/dividendos.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/dividendos", "DividOUtil"}: {
			req("CveTipDivOUtil"),
			req("MontISRAcredRetMexico"),
			req("MontISRAcredRetExtranjero"),
			opc("MontRetExtDivExt"),
			req("TipoSocDistrDiv"),
			opc("MontISRAcredNal"),
			opc("MontDivAcumNal"),
			opc("MontDivAcumExt"),
		},
		// ../esquemas/retencionpago/1/dividendos/dividendos.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/dividendos", "Remanente"}: {
			opc("ProporcionRem"),
		},
		// ../esquemas/retencionpago/1/enajenaciondeacciones/enajenaciondeacciones.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/enajenaciondeacciones", "EnajenaciondeAcciones"}: {
			req("Version"),
			req("ContratoIntermediacion"),
			req("Ganancia"),
			req("Perdida"),
		},
		// ../esquemas/retencionpago/1/fideicomisonoempresarial/fideicomisonoempresarial.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "Fideicomisonoempresarial"}: {
			req("Version"),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "IngresosOEntradas"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "DeduccOSalidas"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "RetEfectFideicomiso"))),
		},
		// ../esquemas/retencionpago/1/fideicomisonoempresarial/fideicomisonoempresarial.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "IngresosOEntradas"}: {
			req("MontTotEntradasPeriodo"),
			req("PartPropAcumDelFideicom"),
			req("PropDelMontTot"),
			si(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "IntegracIngresos")),
				conjunto(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "IntegracIngresos")),
					req("Concepto"),
				),
			),
		},
		// ../esquemas/retencionpago/1/fideicomisonoempresarial/fideicomisonoempresarial.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "DeduccOSalidas"}: {
			req("MontTotEgresPeriodo"),
			req("PartPropDelFideicom"),
			req("PropDelMontTot"),
			si(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "IntegracEgresos")),
				conjunto(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "IntegracEgresos")),
					req("ConceptoS"),
				),
			),
		},
		// ../esquemas/retencionpago/1/fideicomisonoempresarial/fideicomisonoempresarial.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "RetEfectFideicomiso"}: {
			req("MontRetRelPagFideic"),
			req("DescRetRelPagFideic"),
		},
		// ../esquemas/retencionpago/1/intereses/intereses.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/intereses", "Intereses"}: {
			req("Version"),
			req("SistFinanciero"),
			req("RetiroAORESRetInt"),
			req("OperFinancDerivad"),
			req("MontIntNominal"),
			req("MontIntReal"),
			req("Perdida"),
		},
		// ../esquemas/retencionpago/1/intereseshipotecarios/intereseshipotecarios.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/intereseshipotecarios", "Intereseshipotecarios"}: {
			req("Version"),
			req("CreditoDeInstFinanc"),
			req("SaldoInsoluto"),
			opc("PropDeducDelCredit"),
			opc("MontTotIntNominalesDev"),
			opc("MontTotIntNominalesDevYPag"),
			opc("MontTotIntRealPagDeduc"),
			opc("NumContrato"),
		},
		// ../esquemas/retencionpago/1/pagosaextranjeros/pagosaextranjeros.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/pagosaextranjeros", "Pagosaextranjeros"}: {
			req("Version"),
			req("EsBenefEfectDelCobro"),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/pagosaextranjeros", "NoBeneficiario"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/pagosaextranjeros", "Beneficiario"))),
		},
		// ../esquemas/retencionpago/1/pagosaextranjeros/pagosaextranjeros.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/pagosaextranjeros", "NoBeneficiario"}: {
			req("PaisDeResidParaEfecFisc"),
			req("ConceptoPago"),
			req("DescripcionConcepto"),
		},
		// ../esquemas/retencionpago/1/pagosaextranjeros/pagosaextranjeros.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/pagosaextranjeros", "Beneficiario"}: {
			req("RFC"),
			req("CURP"),
			req("NomDenRazSocB"),
			req("ConceptoPago"),
			req("DescripcionConcepto"),
		},
		// ../esquemas/retencionpago/1/premios/premios.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/premios", "Premios"}: {
			req("Version"),
			req("EntidadFederativa"),
			req("MontTotPago"),
			req("MontTotPagoGrav"),
			req("MontTotPagoExent"),
		},
		// ../esquemas/retencionpago/1/operacionesconderivados/operacionesconderivados.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/operacionesconderivados", "Operacionesconderivados"}: {
			req("Version"),
			req("MontGanAcum"),
			req("MontPerdDed"),
		},
		// ../esquemas/retencionpago/1/sectorfinanciero/sectorfinanciero.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/sectorfinanciero", "SectorFinanciero"}: {
			req("Version"),
			req("IdFideicom"),
			opc("NomFideicom"),
			req("DescripFideicom"),
		},
		// ../esquemas/retencionpago/1/planesderetiro11/planesderetiro11.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/planesderetiro11", "Planesderetiro"}: {
			req("Version"),
			req("SistemaFinanc"),
			opc("MontTotAportAnioInmAnterior"),
			req("MontIntRealesDevengAniooInmAnt"),
			req("HuboRetirosAnioInmAntPer"),
			opc("MontTotRetiradoAnioInmAntPer"),
			opc("MontTotExentRetiradoAnioInmAnt"),
			opc("MontTotExedenteAnioInmAnt"),
			req("HuboRetirosAnioInmAnt"),
			opc("MontTotRetiradoAnioInmAnt"),
			opc("NumReferencia"),
			cada(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/planesderetiro11", "AportacionesODepositos")),
				aplicar(rel()),
			),
		},
		// ../esquemas/retencionpago/1/planesderetiro11/planesderetiro11.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/planesderetiro11", "AportacionesODepositos"}: {
			req("TipoAportacionODeposito"),
			req("MontAportODep"),
			opc("RFCFiduciaria"),
		},
		// ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ServiciosPlataformasTecnologicas"}: {
			req("Version"),
			req("Periodicidad"),
			req("NumServ"),
			req("MonTotServSIVA"),
			req("TotalIVATrasladado"),
			req("TotalIVARetenido"),
			req("TotalISRRetenido"),
			req("DifIVAEntregadoPrestServ"),
			req("MonTotalporUsoPlataforma"),
			opc("MonTotalContribucionGubernamental"),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "Servicios"))),
		},
		// ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "Servicios"}: {
			cada(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "DetallesDelServicio")),
				aplicar(rel()),
			),
		},
		// ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "DetallesDelServicio"}: {
			req("FormaPagoServ"),
			req("TipoDeServ"),
			opc("SubTipServ"),
			opc("RFCTerceroAutorizado"),
			req("FechaServ"),
			req("PrecioServSinIVA"),
			si(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ImpuestosTrasladadosdelServicio")),
				aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ImpuestosTrasladadosdelServicio"))),
			),
			si(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ContribucionGubernamental")),
				aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ContribucionGubernamental"))),
			),
			si(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ComisionDelServicio")),
				aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ComisionDelServicio"))),
			),
		},
		// ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ImpuestosTrasladadosdelServicio"}: {
			req("Base"),
			req("Impuesto"),
			req("TipoFactor"),
			req("TasaCuota"),
			req("Importe"),
		},
		// ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ContribucionGubernamental"}: {
			req("ImpContrib"),
			req("EntidadDondePagaLaContribucion"),
		},
		// ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ComisionDelServicio"}: {
			opc("Base"),
			opc("Porcentaje"),
			req("Importe"),
		},
		// ../esquemas/retencionpago/2/retenciones.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/2", "Retenciones"}: {
			req("Version"),
			req("NoCertificado"),
			opc("FolioInt"),
			req("FechaExp"),
			req("LugarExpRetenc"),
			req("CveRetenc"),
			opc("DescRetenc"),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/2", "CfdiRetenRelacionados"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/2", "Emisor"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/2", "Receptor"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/2", "Periodo"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/2", "Totales"))),
			aplicar(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/2", "Complemento"))),
		},
		// ../esquemas/retencionpago/2/retenciones.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/2", "CfdiRetenRelacionados"}: {
			req("TipoRelacion"),
			req("UUID"),
		},
		// ../esquemas/retencionpago/2/retenciones.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/2", "Emisor"}: {
			req("RfcE"),
			req("NomDenRazSocE"),
			req("RegimenFiscalE"),
		},
		// ../esquemas/retencionpago/2/retenciones.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/2", "Receptor"}: {
			req("NacionalidadR"),
			si(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/2", "Nacional")),
				conjunto(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/2", "Nacional")),
					req("RfcR"),
				),
				conjunto(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/2", "Nacional")),
					req("NomDenRazSocR"),
				),
				conjunto(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/2", "Nacional")),
					opc("CurpR"),
				),
				conjunto(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/2", "Nacional")),
					req("DomicilioFiscalR"),
				),
			),
			si(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/2", "Extranjero")),
				conjunto(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/2", "Extranjero")),
					opc("NumRegIdTribR"),
				),
				conjunto(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/2", "Extranjero")),
					req("NomDenRazSocR"),
				),
			),
		},
		// ../esquemas/retencionpago/2/retenciones.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/2", "Periodo"}: {
			req("MesIni"),
			req("MesFin"),
			req("Ejercicio"),
		},
		// ../esquemas/retencionpago/2/retenciones.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/2", "Totales"}: {
			req("MontoTotOperacion"),
			req("MontoTotGrav"),
			req("MontoTotExent"),
			req("MontoTotRet"),
			opc("UtilidadBimestral"),
			opc("ISRCorrespondiente"),
			cada(rel(hijo("http://www.sat.gob.mx/esquemas/retencionpago/2", "ImpRetenidos")),
				aplicar(rel()),
				opc("BaseRet"),
				opc("ImpuestoRet"),
				req("MontoRet"),
				req("TipoPagoRet"),
			),
		},
		// ../esquemas/retencionpago/2/retenciones.xslt
		{"http://www.sat.gob.mx/esquemas/retencionpago/2", "Complemento"}: {
			cada(rel(hijo("", "*")),
				aplicar(rel()),
			),
		},
	},
	integradas: map[nombreXML]bool{
		{"http://www.sat.gob.mx/esquemas/retencionpago/2", "ImpRetenidos"}: true,
	},
}
