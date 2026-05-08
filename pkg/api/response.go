package api

import "fmt"

type Resposta struct {
	Provedor string
	Valor    fmt.Stringer
	Err      error
}

type Option func(*Resposta)

func NewResposta(options ...Option) *Resposta {
	r := &Resposta{}
	for _, opt := range options {
		opt(r)
	}
	return r
}

func WithProvedor(provedor string) Option {
	return func(r *Resposta) {
		r.Provedor = provedor
	}
}

func WithValor(valor fmt.Stringer) Option {
	return func(r *Resposta) {
		r.Valor = valor
	}
}

func WithError(err error) Option {
	return func(r *Resposta) {
		r.Err = err
	}
}
