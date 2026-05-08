package viacep

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/halissontorres/go-expert-consulta-cep/configs/config"
)

var erroCepInvalido = errors.New("CEP inválido")

const endpoint = "http://viacep.com.br/ws/%s/json/"

func BuscarCep(cep string) (*ViaCep, error) {

	if cep == "" {
		return nil, erroCepInvalido
	}
	l := config.GetLogger("viacep.buscarCep")

	url := fmt.Sprintf(endpoint, cep)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			l.Error("Erro ao fechar resposta: ", err, "")
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var viacep *ViaCep
	if err = json.Unmarshal(body, &viacep); err != nil {
		return nil, err
	}

	return viacep, nil
}
