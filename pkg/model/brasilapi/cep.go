package brasilapi

import "fmt"

type BrasilApi struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func (ba BrasilApi) String() string {
	return fmt.Sprintf(
		"BrasilApi{Cep: %s, State: %s, City: %s, Neighborhood: %s, Street: %s, Service: %s\n}",
		ba.Cep, ba.State, ba.City, ba.Neighborhood, ba.Street, ba.Service,
	)
}

type Option func(*BrasilApi)

func NewBrasilApi(options ...Option) *BrasilApi {
	ba := &BrasilApi{}
	for _, option := range options {
		option(ba)
	}
	return ba
}

func WithCep(cep string) Option {
	return func(ba *BrasilApi) {
		ba.Cep = cep
	}
}

func WithState(state string) Option {
	return func(ba *BrasilApi) {
		ba.State = state
	}
}

func WithCity(city string) Option {
	return func(ba *BrasilApi) {
		ba.City = city
	}
}

func WithNeighborhood(neighborhood string) Option {
	return func(ba *BrasilApi) {
		ba.Neighborhood = neighborhood
	}
}

func WithStreet(street string) Option {
	return func(ba *BrasilApi) {
		ba.Street = street
	}
}

func WithService(service string) Option {
	return func(ba *BrasilApi) {
		ba.Service = service
	}
}
