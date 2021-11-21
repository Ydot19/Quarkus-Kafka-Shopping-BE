package models

type TickerResults struct {
	TickerIsActive                      bool   `json:"active"`
	CentralIndexKey                     string `json:"cik"`
	FinancialInstrumentGlobalIdentifier string `json:"composite_figi"`
	CurrencyName                        string `json:"currency_name"`
	LastUpdatedUTC                      string `json:"last_updated_utc"`
	Locale                              string `json:"locale"`
	MarketType                          string `json:"market"`
	PrimaryExchange                     string `json:"primary_exchange"`
	ShareClassFigi                      string `json:"share_class_figi"`
	TickerAbbr                          string `json:"ticker"`
	Type                                string `json:"type"`
}

type Ticker struct {
	Count     int32           `json:"count"`
	NextUrl   string          `json:"next_url"`
	RequestId string          `json:"request_id"`
	Results   []TickerResults `json:"results"`
	Status    string          `json:"status"`
}
