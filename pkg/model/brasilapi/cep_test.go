package brasilapi

import "testing"

func TestBrasilApi_Cep(t *testing.T) {
	t.Parallel()

	b := NewBrasilApi(WithCep("01001000"))

	if b.Cep != "01001000" {
		t.Fatal("expected 01001000 BrasilApi CEP")
	}
}

func TestBrasilApi_String(t *testing.T) {
	t.Parallel()

	ba := NewBrasilApi(WithCep("01001000"), WithCity("João Pessoa"), WithNeighborhood("Bela Vista"), WithStreet("Rua dos pinheiros"), WithService("PAC"))
	if ba == nil {
		t.Fatal("expected non-nil BrasilApi")
	}
	t.Log(ba.String())
}
