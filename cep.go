package brasilapi

// ZipCode models the data of a zip code
type ZipCode struct {
	Cep      *string
	State    *string
	City     *string
	Service  *string
	Location *Location
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

// ZipCodeErr models the data of a error
type ZipCodeErr struct {
	Name    *string `json:"name"`
	Message *string `json:"message"`
	TypeErr *string `json:"type"`
}

// Error implements error interface
func (e *ZipCodeErr) Error() string {
	return *e.Message
}

// GetZipCode return the data of a zip code
func GetZipCode(zipcode string) (zc *ZipCode, err error) {
	var (
		client = newHttpClient("https://brasilapi.com.br/api/cep/v2/")
		reqErr = new(ZipCodeErr)
	)

	if err = client.get(zipcode, zc, reqErr); err != nil {
		return nil, err
	}

	return
}
