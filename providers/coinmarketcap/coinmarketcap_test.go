// Copyright (C) 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getCoinMarket(t *testing.T, content string) (*Client, *httptest.Server) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, content)
	}))

	client, err := NewClient()
	assert.NoError(t, err)
	client.BaseURL = testServer.URL
	return client, testServer
}

func TestGetCoinsWithDefaultCurrency(t *testing.T) {
	client, testServer := getCoinMarket(t, `[
    {
        "id": "bitcoin",
        "name": "Bitcoin",
        "symbol": "BTC",
        "rank": "1",
        "price_usd": "15451.6",
        "price_btc": "1.0",
        "24h_volume_usd": "12449900000.0",
        "market_cap_usd": "259021351440",
        "available_supply": "16763400.0",
        "total_supply": "16763400.0",
        "max_supply": "21000000.0",
        "percent_change_1h": "-0.23",
        "percent_change_24h": "9.91",
        "percent_change_7d": "-17.87",
        "last_updated": "1514273660",
        "price_eur": "13017.6176132",
        "24h_volume_eur": "10488754402.3",
        "market_cap_eur": "218219531097"
    },
    {
        "id": "ethereum",
        "name": "Ethereum",
        "symbol": "ETH",
        "rank": "2",
        "price_usd": "774.768",
        "price_btc": "0.0504444",
        "24h_volume_usd": "2387070000.0",
        "market_cap_usd": "74823618581.0",
        "available_supply": "96575515.0",
        "total_supply": "96575515.0",
        "max_supply": null,
        "percent_change_1h": "-0.06",
        "percent_change_24h": "3.09",
        "percent_change_7d": "-7.24",
        "last_updated": "1514273649",
        "price_eur": "652.724220336",
        "24h_volume_eur": "2011051572.39",
        "market_cap_eur": "63037177711.0"
    }
]`)
	defer testServer.Close()

	coins, err := client.GetCoins("EUR", 2)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(coins))
	assert.Equal(t, "13017.6176132", coins[0].PriceEur)
	assert.Equal(t, "Bitcoin", coins[0].Name)
	assert.Equal(t, "652.724220336", coins[1].PriceEur)
	assert.Equal(t, "Ethereum", coins[1].Name)
}

func TestGetCoinWithDefaultCurrency(t *testing.T) {
	client, testServer := getCoinMarket(t, `[{
      "id":"bitcoin",
      "name":"Bitcoin",
      "symbol":"BTC",
      "rank":"1",
      "price_usd":"573.137",
      "price_btc":"1.0",
      "24h_volume_usd":"72855700.0",
      "market_cap_usd":"9080883500.0",
      "available_supply":"15844176.0",
      "total_supply":"15844176.0",
      "percent_change_1h":"0.04",
      "percent_change_24h":"-0.3",
      "percent_change_7d":"-0.57",
      "last_updated":"1472762067"
   }
]`)
	defer testServer.Close()

	coins, err := client.GetCoin("BTC", "USD", 1)
	assert.NoError(t, err)
	assert.Equal(t, "573.137", coins[0].PriceUsd)
	assert.Equal(t, "Bitcoin", coins[0].Name)
	assert.Equal(t, "BTC", coins[0].Symbol)
}

func TestGetCoinWithEURCurrency(t *testing.T) {
	client, testServer := getCoinMarket(t, `[
   {
      "id":"bitcoin",
      "name":"Bitcoin",
      "symbol":"BTC",
      "rank":"1",
      "price_usd":"15480.2",
      "price_btc":"1.0",
      "24h_volume_usd":"12403600000.0",
      "market_cap_usd":"259500784680",
      "available_supply":"16763400.0",
      "total_supply":"16763400.0",
      "max_supply":"21000000.0",
      "percent_change_1h":"0.33",
      "percent_change_24h":"10.22",
      "percent_change_7d":"-17.76",
      "last_updated":"1514273061",
      "price_eur":"13041.7124554",
      "24h_volume_eur":"10449747717.2",
      "market_cap_eur":"218623442575"
   }
]`)
	defer testServer.Close()

	coins, err := client.GetCoin("BTC", "EUR", 1)
	assert.NoError(t, err)
	assert.Equal(t, "15480.2", coins[0].PriceUsd)
	assert.Equal(t, "13041.7124554", coins[0].PriceEur)
	assert.Equal(t, "Bitcoin", coins[0].Name)
	assert.Equal(t, "BTC", coins[0].Symbol)
}
