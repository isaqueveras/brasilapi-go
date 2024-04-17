// Code generated by juaz. DO NOT EDIT.
// versions: v1.0.1
// source: cep.juaz

package cep

import (
	"context"
	"net/http"
	"fmt"

	"github.com/isaqueveras/juazeiro"
)

// Identifier data model for the identifier structure
type Identifier struct {
	Cep *string `json:"cep,omitempty"`
}

// Response data model for the response structure
type Response struct {
	Cep *string `json:"cep,omitempty"`
	State *string `json:"state,omitempty"`
	City *string `json:"city,omitempty"`
	Service *string `json:"service,omitempty"`
	Location *Location `json:"location,omitempty"`
}

// Location data model for the location structure
type Location struct {
	Type *string `json:"type,omitempty"`
	Coordenates *Coordenates `json:"coordenates,omitempty"`
}

// Coordenates data model for the coordenates structure
type Coordenates struct {
	Longitude *string `json:"longitude,omitempty"`
	Latitude *string `json:"latitude,omitempty"`
}

// ICepClient defines the interface of the provided methods
type ICepClient interface {
	GetZipCode(context.Context, *Identifier) (*Response, error)
}

type CepClient struct {
	cc juazeiro.ClientConnInterface
}

func NewCepClient(cc juazeiro.ClientConnInterface) ICepClient {
	return &CepClient{cc: cc}
}

// GetZipCode implements the GetZipCode method of the interface
func (c *CepClient) GetZipCode(ctx context.Context, in *Identifier) (*Response, error) {
	out := new(Response)
	uri := fmt.Sprintf("/cep/v2/%v", *in.Cep)
	err := c.cc.Invoke(ctx, http.MethodGet, uri, http.StatusOK, in, out)
	return out, err
}