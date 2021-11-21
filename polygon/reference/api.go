package reference

import "time"

type GetTickerEndpointFilterParameters interface {
	SetType(tickerType string)
	GetType()
	SetMarket(market string)
	GetMarket()
	SetCUSIP(cusip string)
	GetCUSIP()
	SetCIK(cik string)
	GetCIK()
	SetPointInTimeDate(dt time.Time)
	GetPointInTimeDate()
	SetSearchTerm(term string)
	GetSearchTerm()
	SetOrder(isAscendingOrder bool)
	GetOrder()
}

type GetTickerEndpoint interface {
	Fetch()
	*GetTickerEndpointFilterParameters
	constructQueryParameters()
	GetNext()
	GetPrevious()
}

type RestApi interface {
	GetTickerEndpoint
}
