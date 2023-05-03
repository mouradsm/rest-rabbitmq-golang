package requests

import "go.mongodb.org/mongo-driver/bson/primitive"

var EmptyCreateCliente = ClienteCreateRequest{}

type ClienteCreateRequest struct {
	Id         primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name       string             `json:"name" binding:"required,min=2,max=100"`
	CPF        string             `json:"cpf"`
	Nascimento string             `json:"nascimento"`
	Endereco   Endereco           `json:"endereco"`
}

type Endereco struct {
	DeptId primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Rua    string             `json:"rua"`
	Bairro string             `json:"bairro"`
	Cidade string             `json:"cidade"`
	Estado string             `json:"estado"`
}

type ClienteUpdateRequest struct {
	Name     string   `json:"name" binding:"required,min=2,max=100"`
	CPF      string   `json:"nascimento"`
	Endereco Endereco `json:"endereco"`
}
