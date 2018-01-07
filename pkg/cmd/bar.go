// Copyright (C) 2017-2018 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/leekchan/accounting"
	//	"github.com/olekukonko/tablewriter"
	//	"github.com/nlamirault/picsou/pkg/coins"
)

const (
	barChar = "∎"
)

// type WalletCoin struct {
// 	Name    string
// 	Percent float64
// 	Money   float64
// }

// func DisplayWalletTable(out io.Writer, wallet map[string]coins.Wallet, ac accounting.Accounting) {

// 	table := tablewriter.NewWriter(out)
// 	table.SetHeader([]string{
// 		"Symbol",
// 		"Money",
// 		"Percent",
// 		"Vue"})
// 	table.SetRowLine(true)
// 	table.SetAutoWrapText(false)

// 	for _, coin := range coins {
// 		table.Append([]string{
// 			YellowOut(coin.Rank),
// 			BlueOut(coin.Symbol),
// 			BlueOut(coin.Name),
// 			GetMoney(ac, coinmarketcap.GetPrice(coin, currency)),
// 			GetMoney(ac, coinmarketcap.Two4HVolume(coin, currency)),
// 			GetMoney(ac, coinmarketcap.MarketCap(coin, currency)),
// 			getPercentColor(coin.PercentChange1H),
// 			getPercentColor(coin.PercentChange24H),
// 			getPercentColor(coin.PercentChange7D),
// 			coin.LastUpdated,
// 		})
// 	}
// 	table.Render()
// 	return nil
// }

func DisplayWalletBars(out io.Writer, name string, money float64, walletTotal float64, ac accounting.Accounting) {

	percent := fmt.Sprintf("%.0f", (money*100)/walletTotal)
	fmt.Fprintf(out, "%s: %s, [%s%%]\t", GreenOut(name), GetMoney(ac, fmt.Sprintf("%f", money)), percent)

	var barLen int
	barLen, err := strconv.Atoi(percent)
	if err != nil {
		return
	}

	var bars string
	switch {
	case 0 < barLen && barLen < 30:
		bars = BlueOut(barChar)
	case 30 < barLen && barLen < 50:
		bars = GreenOut(barChar)
	case 50 < barLen && barLen < 70:
		bars = YellowOut(barChar)
	case 70 < barLen:
		bars = RedOut(barChar)
	}

	fmt.Fprintf(out, "%v\n", strings.Repeat(bars, barLen))
}
