package main

import (
	"os"
	"time"

	"github.com/halissontorres/go-expert-consulta-cep/configs/config"
	"github.com/halissontorres/go-expert-consulta-cep/pkg/api"
	"github.com/halissontorres/go-expert-consulta-cep/pkg/brasilapi"
	"github.com/halissontorres/go-expert-consulta-cep/pkg/viacep"
)

func main() {
	cep := "13330-250" // CEP DA FullCycle :)
	l := config.GetLogger("")

	timeout, err := time.ParseDuration(os.Getenv("CONSULTA_CEP_TIMEOUT"))
	if err != nil {
		timeout = config.ConsultaCepTimeOut
	}

	ch := make(chan *api.Resposta, 2)

	go func() {
		via, err := viacep.BuscarCep(cep)
		ch <- api.NewResposta(api.WithProvedor("ViaCEP"), api.WithValor(via), api.WithError(err))
	}()

	go func() {
		brasil, err := brasilapi.BuscarCep(cep)
		ch <- api.NewResposta(api.WithProvedor("BrasilAPI"), api.WithValor(brasil), api.WithError(err))
	}()

	select {
	case r := <-ch:
		if r.Err != nil {
			l.Error("%s retornou erro: %v\n", r.Provedor, r.Err)
			return
		}
		l.Info(r.Valor.String())

	case <-time.After(timeout):
		l.Warn("Tempo excedido: sem resposta dos serviços")
	}
}
