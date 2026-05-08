package viacep

import "testing"

func TestViaCepApi_String(t *testing.T) {
	t.Parallel()

	vc := NewViaCep(WithCep("01001000"))
	if vc == nil {
		t.Fatal("expected non-nil BrasilApi")
	}
	t.Log(vc.String())
}

func TestViaCepApi_Cep(t *testing.T) {
	t.Parallel()

	vc := NewViaCep(WithCep("01001000"))

	if vc.Cep != "01001000" {
		t.Fatal("expected 01001000 ViaCep CEP")
	}
}
