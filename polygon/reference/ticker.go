package reference

import (
	"github.com/Ydot19/Stonks/polygon/restclient"
	"time"
)

type GetTickerFilterParameters struct {
	ticketType  string
	market      string
	cusip       string
	cik         string
	pointInTime time.Time
	searchTerm  string
	orderAsc    bool
}

func CreateGetTickerFilterParameters() *FilterGetTickerEndpoint {
	var ret FilterGetTickerEndpoint = &GetTickerFilterParameters{}
	return &ret
}

func (tf *GetTickerFilterParameters) SetType(tickerType string) {

}

func (tf *GetTickerFilterParameters) GetType() {

}

func (tf *GetTickerFilterParameters) SetMarket(market string) {

}

func (tf *GetTickerFilterParameters) GetMarket() {

}

func (tf *GetTickerFilterParameters) SetCUSIP(cusip string) {

}

func (tf *GetTickerFilterParameters) GetCUSIP() {

}

func (tf *GetTickerFilterParameters) SetCIK(cik string) {

}

func (tf *GetTickerFilterParameters) GetCIK() {

}

func (tf *GetTickerFilterParameters) SetPointInTimeDate(dt time.Time) {

}

func (tf *GetTickerFilterParameters) GetPointInTimeDate() {

}

func (tf *GetTickerFilterParameters) SetSearchTerm(term string) {}

func (tf *GetTickerFilterParameters) GetSearchTerm() {}

func (tf *GetTickerFilterParameters) SetOrder(isAscendingOrder bool) {

}

func (tf *GetTickerFilterParameters) GetOrder() {

}

type GetTicker struct {
	next_url         string
	previous_url     string
	client           *restclient.RestClient
	FilterParameters *FilterGetTickerEndpoint
}

func CreateGetTicker(rc *restclient.RestClient) *GetTicker {
	filters := CreateGetTickerFilterParameters()
	getTicker := GetTicker{
		next_url:         "",
		previous_url:     "",
		client:           rc,
		FilterParameters: filters,
	}
	return &getTicker
}
