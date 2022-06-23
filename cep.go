package brasilapi

const (
	// urlBrasilAPI base url from BrasilApi
	urlBrasilAPI = "https://brasilapi.com.br/api/cep/v2/"
)

// ZipCode models the data of a zip code
type ZipCode struct {
	Cep      *string   `json:"cep"`
	State    *string   `json:"state"`
	City     *string   `json:"city"`
	Service  *string   `json:"service"`
	Location *Location `json:"location"`
}

// Location models the data of a location
type Location struct {
	Type        *string      `json:"type"`
	Coordenates *Coordenates `json:"coordenates"`
}

// Coordenates models the data of a coordenates
type Coordenates struct {
	Longitude *string `json:"longitude"`
	Latitude  *string `json:"latitude"`
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
	zc = new(ZipCode)

	var (
		client = newHttpClient(urlBrasilAPI)
		reqErr = new(ZipCodeErr)
	)

	if err = client.get(zipcode, zc, reqErr); err != nil {
		return nil, err
	}

	return
}
