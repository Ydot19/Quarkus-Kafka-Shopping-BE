package restclient

type bearer struct {
	token string `env:"POLYGON_API_KEY"`
}

type BaseUrl string

func (c BaseUrl) String() string {
	return string(c)
}

// TODO: use environment variables to determine the correct url to use
const (
	DevBaseUrl  BaseUrl = "https://api.polygon.io"
	ProdBaseUrl BaseUrl = "https://api.polygon.io"
)

const Version string = "v3"
