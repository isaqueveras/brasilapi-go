// Code generated by juaz. DO NOT EDIT.
// versions: v1.0.1
// source: ddd.juaz

package ddd

import (
	"context"
	"net/http"
	"fmt"

	"github.com/isaqueveras/juazeiro"
)

// Identifier data model for the identifier structure
type Identifier struct {
	Ddd *int64 `json:"ddd,omitempty"`
}

// Response data model for the response structure
type Response struct {
	State *string `json:"state,omitempty"`
	Cities []*string `json:"cities,omitempty"`
}

// IDddClient defines the interface of the provided methods
type IDddClient interface {
	Obtain(context.Context, *Identifier) (*Response, error)
}

type DddClient struct {
	cc juazeiro.ClientConnInterface
}

func NewDddClient(cc juazeiro.ClientConnInterface) IDddClient {
	return &DddClient{cc: cc}
}

// Obtain implements the Obtain method of the interface
func (c *DddClient) Obtain(ctx context.Context, in *Identifier) (*Response, error) {
	out := new(Response)
	uri := fmt.Sprintf("/ddd/v1/%v", *in.Ddd)
	err := c.cc.Invoke(ctx, http.MethodGet, uri, http.StatusOK, in, out)
	return out, err
}
