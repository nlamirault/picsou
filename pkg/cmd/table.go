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

package cmd

import (
	"io"
	"strconv"
	"strings"

	"github.com/leekchan/accounting"
	"github.com/olekukonko/tablewriter"

	"github.com/nlamirault/picsou/providers/coinmarketcap"
)

func getAccounting(currency string) accounting.Accounting {
	var ac accounting.Accounting
	switch currency {
	case "EUR":
		ac = accounting.Accounting{Symbol: "€", Precision: 4}
	default:
		ac = accounting.Accounting{Symbol: "$", Precision: 4}
	}
	return ac
}

func getPercentColor(value string) string {
	var percent string
	if strings.HasPrefix(value, "-") {
		percent = RedOut(value)
	} else {
		percent = GreenOut(value)
	}
	return percent
}

func getMoney(ac accounting.Accounting, value string) string {
	money, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return value
	}
	return ac.FormatMoney(money)
}

func DisplayCoins(out io.Writer, coins []coinmarketcap.Coin, currency string) error {
	ac := getAccounting(currency)
	table := tablewriter.NewWriter(out)
	table.SetHeader([]string{
		"Rank",
		"Symbol",
		"Coin",
		"Price",
		"24 Hour Volume",
		"Market Cap",
		"1 Hour",
		"24 Hour",
		"7 Days",
		"Last Updated"})
	table.SetRowLine(true)
	table.SetAutoWrapText(false)

	for _, coin := range coins {
		table.Append([]string{
			YellowOut(coin.Rank),
			BlueOut(coin.Symbol),
			BlueOut(coin.Name),
			getMoney(ac, coinmarketcap.GetPrice(coin, currency)),
			getMoney(ac, coinmarketcap.Two4HVolume(coin, currency)),
			getMoney(ac, coinmarketcap.MarketCap(coin, currency)),
			getPercentColor(coin.PercentChange1H),
			getPercentColor(coin.PercentChange24H),
			getPercentColor(coin.PercentChange7D),
			coin.LastUpdated,
		})
	}
	table.Render()
	return nil
}
