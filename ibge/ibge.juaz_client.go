// Code generated by juaz. DO NOT EDIT.
// versions: v1.0.1
// source: ibge.juaz

package ibge

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/isaqueveras/juazeiro"
)

// UF defines the uf enum type
type UF string

const (
	UfAc UF = "AC"
	UfAl UF = "AL"
	UfAp UF = "AP"
	UfAm UF = "AM"
	UfBa UF = "BA"
	UfCe UF = "CE"
	UfDf UF = "DF"
	UfEs UF = "ES"
	UfGo UF = "GO"
	UfMa UF = "MA"
	UfMt UF = "MT"
	UfMs UF = "MS"
	UfMg UF = "MG"
	UfPa UF = "PA"
	UfPb UF = "PB"
	UfPr UF = "PR"
	UfPe UF = "PE"
	UfPi UF = "PI"
	UfRj UF = "RJ"
	UfRn UF = "RN"
	UfRs UF = "RS"
	UfRo UF = "RO"
	UfRr UF = "RR"
	UfSc UF = "SC"
	UfSp UF = "SP"
	UfSe UF = "SE"
	UfTo UF = "TO"
)

// String convert uf type to string
func (t UF) String() string {
	return string(t)
}

// Provider defines the provider enum type
type Provider string

const (
	ProviderGov            Provider = "gov"
	ProviderWikipedia      Provider = "wikipedia"
	ProviderDadosAbertosBr Provider = "dados-abertos-br"
)

// String convert provider type to string
func (t Provider) String() string {
	return string(t)
}

// Identifier data model for the identifier structure
type Identifier struct {
	UF         *UF `json:"uf,omitempty"`
	parameters *Params
}

// NewParams ...
func (i *Identifier) NewParams() {
	i.parameters = &Params{}
}

// WithParamProviders ...
func (i *Identifier) WithParamProviders(providers *Provider) {
	i.parameters.Providers = providers
}

// Params data model for the params structure
type Params struct {
	Providers *Provider `json:"providers,omitempty"`
}

// Response data model for the response structure
type Response struct {
	Nome       *string `json:"nome,omitempty"`
	CodigoIbge *string `json:"codigo_ibge,omitempty"`
}

// IIbgeClient defines the interface of the provided methods
type IIbgeClient interface {
	ObtainCounties(context.Context, *Identifier) (*[]Response, error)
}

type IbgeClient struct {
	cc juazeiro.ClientConnInterface
}

func NewIbgeClient(cc juazeiro.ClientConnInterface) IIbgeClient {
	return &IbgeClient{cc: cc}
}

// ObtainCounties implements the ObtainCounties method of the interface
func (c *IbgeClient) ObtainCounties(ctx context.Context, in *Identifier) (*[]Response, error) {
	out := new([]Response)
	uri := fmt.Sprintf("/ibge/municipios/v1/%v", *in.UF)
	if in.parameters != nil {
		uri += _build_params_parameters(in.parameters)
		in.parameters = nil
	}
	err := c.cc.Invoke(ctx, http.MethodGet, uri, http.StatusOK, in, out)
	return out, err
}

func _build_params_parameters(in *Params) string {
	val := &url.Values{}
	if in.Providers != nil {
		val.Add("providers", fmt.Sprintf("%v", in.Providers.String()))
	}
	return "?" + val.Encode()
}