package viacep

import "fmt"

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (ba ViaCep) String() string {
	return fmt.Sprintf(
		"[ViaCep] Cep: %s, Logradouro: %s, Complemento: %s, Unidade: %s, Bairro: %s, Localidade: %s, Uf: %s, Estado: %s, Regiao: %s, Ibge: %s, Gia: %s, Ddd: %s, Siafi: %s",
		ba.Cep, ba.Logradouro, ba.Complemento, ba.Unidade, ba.Bairro, ba.Localidade, ba.Uf, ba.Estado, ba.Regiao, ba.Ibge, ba.Gia, ba.Ddd, ba.Siafi,
	)
}

type Option func(*ViaCep)

func NewViaCep(options ...Option) *ViaCep {
	ba := &ViaCep{}
	for _, option := range options {
		option(ba)
	}
	return ba
}

func WithCep(cep string) Option {
	return func(ba *ViaCep) {
		ba.Cep = cep
	}
}

func WithLogradouro(logradouro string) Option {
	return func(ba *ViaCep) {
		ba.Logradouro = logradouro
	}
}

func WithComplemento(complemento string) Option {
	return func(ba *ViaCep) {
		ba.Complemento = complemento
	}
}

func WithUnidade(unidade string) Option {
	return func(ba *ViaCep) {
		ba.Unidade = unidade
	}
}

func WithBairro(bairro string) Option {
	return func(ba *ViaCep) {
		ba.Bairro = bairro
	}
}

func WithLocalidade(localidade string) Option {
	return func(ba *ViaCep) {
		ba.Localidade = localidade
	}
}

func WithUf(uf string) Option {
	return func(ba *ViaCep) {
		ba.Uf = uf
	}
}

func WithEstado(estado string) Option {
	return func(ba *ViaCep) {
		ba.Estado = estado
	}
}

func WithRegiao(regiao string) Option {
	return func(ba *ViaCep) {
		ba.Regiao = regiao
	}
}

func WithIbge(ibge string) Option {
	return func(ba *ViaCep) {
		ba.Ibge = ibge
	}
}

func WithGia(gia string) Option {
	return func(ba *ViaCep) {
		ba.Gia = gia
	}
}

func WithDdd(ddd string) Option {
	return func(ba *ViaCep) {
		ba.Ddd = ddd
	}
}

func WithSiafi(siafi string) Option {
	return func(ba *ViaCep) {
		ba.Siafi = siafi
	}
}
