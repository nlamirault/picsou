// Copyright (C) 2017-2018 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package coinmarketcap

import (
	"encoding/json"
	"fmt"
	// "strconv"
	"strings"

	"github.com/golang/glog"
)

const (
	apiEndpoint = "https://api.coinmarketcap.com/v1"
)

type Client struct {
	BaseURL string
}

func NewClient() (*Client, error) {
	return &Client{
		BaseURL: apiEndpoint,
	}, nil
}

func (client Client) GetCoins(currency string, limit int64) ([]Coin, error) {
	var url = fmt.Sprintf("%s/ticker/", client.BaseURL)
	return makeAPICall(url, currency, limit)
}

func (client Client) GetCoin(cryptoCurrency string, currency string, limit int64) ([]Coin, error) {
	glog.V(2).Infof("Get coin: %s using currency %s, %d", cryptoCurrency, currency, limit)
	var url = fmt.Sprintf("%s/ticker/", client.BaseURL)

	cryptoCurrencyName, err := getCryptoCurrencyName(cryptoCurrency)
	if err != nil {
		return nil, err
	}
	url = fmt.Sprintf("%s%s/", url, cryptoCurrencyName)
	return makeAPICall(url, currency, limit)
}

func makeAPICall(url string, currency string, limit int64) ([]Coin, error) {
	params := []string{}
	if limit > 0 {
		params = append(params, fmt.Sprintf("limit=%v", limit))
	}

	if currency != DefaultCurrency {
		params = append(params, fmt.Sprintf("convert=%s", strings.ToLower(currency)))
	}

	apiURL := fmt.Sprintf("%s?%s", url, strings.Join(params, "&"))
	glog.V(2).Infof("API Url: %s", apiURL)
	response, err := fetchCoin(apiURL)
	if err != nil {
		return nil, err
	}

	return readCoinData(response, currency)

}

func readCoinData(response []byte, currency string) ([]Coin, error) {
	glog.V(2).Infof("HTTP response: %s", string(response))
	var coins []Coin
	err := json.Unmarshal(response, &coins)
	if err != nil {
		return nil, err
	}
	glog.V(2).Infof("Response : %s", coins)
	return coins, nil
}
