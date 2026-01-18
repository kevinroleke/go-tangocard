package tangocard

type GetExchangeRatesRequest struct {
	BaseCurrency string `json:"baseCurrency"`
	RewardCurrency string `json:"rewardCurrency"`
}

type ExchangeRate struct {
	LastModifiedDate string `json:"lastModifiedDate"`
	RewardCurrency string `json:"rewardCurrency"`
	BaseCurrency string `json:"baseCurrency"`
	BaseFx float64 `json:"baseFx"`
}

type GetExchangeRatesResponse struct {
	ExchangeRates []ExchangeRate `json:"exchangeRates"`
}
