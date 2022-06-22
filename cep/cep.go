package cep

// CEP models the data of a zip code
type CEP struct {
	Cep     *string
	State   *string
	City    *string
	Service *string
}

// Location models the data of a location
type Location struct {
	Type        *string
	Coordenates *Coordenates
}

// Coordenates models the data of a coordenates
type Coordenates struct {
	Longitude *float64
	Latitude  *float32
}
