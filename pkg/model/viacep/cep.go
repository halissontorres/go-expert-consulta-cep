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
		"ViaCep{Cep: %s, Logradouro: %s, Complemento: %s, Unidade: %s, Bairro: %s, Localidade: %s, Uf: %s, Estado: %s, Regiao: %s, Ibge: %s, Gia: %s, Ddd: %s, Siafi: %s\n}",
		ba.Cep, ba.Logradouro, ba.Complemento, ba.Unidade, ba.Bairro, ba.Localidade, ba.Uf, ba.Estado, ba.Regiao, ba.Ibge, ba.Gia, ba.Ddd, ba.Siafi,
	)
}

type Option func(*ViaCep)
