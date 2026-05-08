package brasilapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/halissontorres/go-expert-consulta-cep/configs/config"
)

var erroCepInvalido = errors.New("CEP inválido")

const endpoint = "https://brasilapi.com.br/api/cep/v1/%s"

func BuscarCep(cep string) (*BrasilApi, error) {

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

	var brasilApi *BrasilApi
	if err = json.Unmarshal(body, &brasilApi); err != nil {
		return nil, err
	}

	return brasilApi, nil
}
