package decimal

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestParse(t *testing.T) {
	casos := []struct {
		in, original, normal string
		scale                int32
		sign                 int
	}{
		{"0", "0", "0", 0, 0},
		{"0.00", "0.00", "0", 2, 0},
		{"-0", "-0", "0", 0, 0},
		{"-0.50", "-0.50", "-0.5", 2, -1},
		{"1160.00", "1160.00", "1160", 2, 1},
		{"1234567890123.123456", "1234567890123.123456", "1234567890123.123456", 6, 1},
		{"0001.500", "0001.500", "1.5", 3, 1},
		{"+3.", "+3.", "3", 0, 1},
		{".5", ".5", "0.5", 1, 1},
		{"0.000001", "0.000001", "0.000001", 6, 1},
		{"100", "100", "100", 0, 1},
		{"99999999999999999999999999.99", "99999999999999999999999999.99", "99999999999999999999999999.99", 2, 1},
		// whiteSpace="collapse" del XSD: se recortan los extremos.
		{" 1.0 ", "1.0", "1", 1, 1},
		{"\t\n2.50\r\n", "2.50", "2.5", 2, 1},
	}
	for _, c := range casos {
		d, err := Parse(c.in)
		if err != nil {
			t.Errorf("Parse(%q): %v", c.in, err)
			continue
		}
		if d.Original() != c.original {
			t.Errorf("Parse(%q).Original() = %q, quiero %q", c.in, d.Original(), c.original)
		}
		if d.String() != c.normal {
			t.Errorf("Parse(%q).String() = %q, quiero %q", c.in, d.String(), c.normal)
		}
		if d.Scale() != c.scale {
			t.Errorf("Parse(%q).Scale() = %d, quiero %d", c.in, d.Scale(), c.scale)
		}
		if d.Sign() != c.sign {
			t.Errorf("Parse(%q).Sign() = %d, quiero %d", c.in, d.Sign(), c.sign)
		}
		if d.IsZero() != (c.sign == 0) {
			t.Errorf("Parse(%q).IsZero() = %v", c.in, d.IsZero())
		}
	}
}

func TestParseRechaza(t *testing.T) {
	for _, in := range []string{
		"", " ", "1e3", "1E3", "1.5e-2", "NaN", "nan", "Inf", "-Inf", "1,000.00", "1.000,00",
		"1_000", "1.2.3", ".", "-", "+", "+-1", "--1", "1 000", "0x10", "١٢", "$100", "12a",
	} {
		if d, err := Parse(in); err == nil {
			t.Errorf("Parse(%q) = %q, quería error", in, d)
		} else if !errors.Is(err, ErrFormato) {
			t.Errorf("Parse(%q): el error %v no envuelve ErrFormato", in, err)
		}
	}
}

func TestCmp(t *testing.T) {
	casos := []struct {
		a, b string
		cmp  int
	}{
		{"1.50", "1.5", 0},
		{"0", "-0.00", 0},
		{"1160.00", "1160.01", -1},
		{"-1", "-2", 1},
		{"0.1", "0.09", 1},
		{"753130.199999", "753130.20", -1},
		{"1234567890123.123456", "1234567890123.123457", -1},
	}
	for _, c := range casos {
		if got := MustParse(c.a).Cmp(MustParse(c.b)); got != c.cmp {
			t.Errorf("Cmp(%s, %s) = %d, quiero %d", c.a, c.b, got, c.cmp)
		}
		if got := MustParse(c.b).Cmp(MustParse(c.a)); got != -c.cmp {
			t.Errorf("Cmp(%s, %s) = %d, quiero %d", c.b, c.a, got, -c.cmp)
		}
	}
}

func TestAritmetica(t *testing.T) {
	casos := []struct {
		a, b, suma, resta, prod string
	}{
		{"1000.00", "160.00", "1160", "840", "160000"},
		{"0.1", "0.2", "0.3", "-0.1", "0.02"},
		{"1500.00", "0.160000", "1500.16", "1499.84", "240"},
		{"-0.50", "0.50", "0", "-1", "-0.25"},
		{"1234567890123.123456", "0.000001", "1234567890123.123457", "1234567890123.123455", "1234567.890123123456"},
	}
	for _, c := range casos {
		a, b := MustParse(c.a), MustParse(c.b)
		if got := a.Add(b).String(); got != c.suma {
			t.Errorf("%s + %s = %s, quiero %s", c.a, c.b, got, c.suma)
		}
		if got := a.Sub(b).String(); got != c.resta {
			t.Errorf("%s - %s = %s, quiero %s", c.a, c.b, got, c.resta)
		}
		if got := a.Mul(b).String(); got != c.prod {
			t.Errorf("%s × %s = %s, quiero %s", c.a, c.b, got, c.prod)
		}
	}
	// Las operaciones no tocan a los operandos.
	a := MustParse("1.00")
	_ = a.Add(MustParse("2")).Mul(MustParse("3"))
	if a.Original() != "1.00" || a.String() != "1" {
		t.Errorf("Add mutó al operando: %q", a.Original())
	}
	// El valor cero sirve de acumulador.
	var acc Decimal
	for _, s := range []string{"100.00", "0.16", "-0.01"} {
		acc = acc.Add(MustParse(s))
	}
	if acc.String() != "100.15" || acc.Original() != "100.15" {
		t.Errorf("acumulado = %q", acc.Original())
	}
}

func TestRoundYStringFixed(t *testing.T) {
	casos := []struct {
		in     string
		places int32
		out    string
	}{
		{"2.5", 2, "2.50"},
		{"2.345", 2, "2.35"},
		{"2.344999", 2, "2.34"},
		{"-2.345", 2, "-2.35"},
		{"0.005", 2, "0.01"},
		{"-0.004", 2, "0.00"},
		{"1160", 2, "1160.00"},
		{"1.5", 0, "2"},
	}
	for _, c := range casos {
		if got := MustParse(c.in).StringFixed(c.places); got != c.out {
			t.Errorf("StringFixed(%s, %d) = %s, quiero %s", c.in, c.places, got, c.out)
		}
	}
}

func TestJSON(t *testing.T) {
	type monto struct {
		Total Decimal  `json:"total"`
		Desc  *Decimal `json:"desc,omitempty"`
	}
	m := monto{Total: MustParse("1160.00")}
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != `{"total":"1160.00"}` {
		t.Errorf("Marshal = %s", b)
	}

	var back monto
	if err := json.Unmarshal([]byte(`{"total":"-0.50","desc":"0.10"}`), &back); err != nil {
		t.Fatal(err)
	}
	if back.Total.Original() != "-0.50" || back.Desc == nil || back.Desc.String() != "0.1" {
		t.Errorf("Unmarshal = %+v", back)
	}
	// Número JSON sin exponente: se acepta.
	if err := json.Unmarshal([]byte(`{"total":12.30}`), &back); err != nil || back.Total.Original() != "12.30" {
		t.Errorf("Unmarshal número: %v %q", err, back.Total.Original())
	}
	for _, malo := range []string{`{"total":"1e3"}`, `{"total":1e3}`, `{"total":"abc"}`, `{"total":""}`, `{"total":true}`} {
		if err := json.Unmarshal([]byte(malo), &back); err == nil {
			t.Errorf("Unmarshal(%s) debía fallar", malo)
		}
	}
	// Valor cero.
	var cero Decimal
	if b, _ := json.Marshal(cero); string(b) != `"0"` {
		t.Errorf("Marshal(cero) = %s", b)
	}
}
