package domain

import "github.com/brunoOchoa.com/api-REST-FULL/requests"

type ClienteService interface {
	GetAllClientes() ([]requests.ClienteCreateRequest, error)
	GetCliente(string) (requests.ClienteCreateRequest, error)
	CreateCliente(requests.ClienteCreateRequest) (requests.ClienteCreateRequest, error)
	UpdateCliente(string, requests.ClienteUpdateRequest) error
	DeleteCliente(string) error
}
