// Copyright (C) 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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

	"github.com/golang/glog"
	"github.com/leekchan/accounting"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	pkgcmd "github.com/nlamirault/picsou/pkg/cmd"
	"github.com/nlamirault/picsou/providers/coinmarketcap"
)

var (
	name     string
	coins    []string
	limit    int64
	currency string
)

type cryptoCmd struct {
	out io.Writer
}

func newCryptoCmd(out io.Writer) *cobra.Command {
	cryptoCmd := &cryptoCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "crypto",
		Short: "Manage crypto currencies.",
		Long:  "Manage crypto currencies. See List, Get, ... subcommands.",
		RunE:  nil,
	}

	listCryptoCmd := &cobra.Command{
		Use:   "list",
		Short: "Display crypto currencies",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := coinmarketcap.NewClient()
			if err != nil {
				return err
			}
			return cryptoCmd.listCryptoCurrencies(client, currency, limit)
		},
	}

	getCryptoCmd := &cobra.Command{
		Use:   "get",
		Short: "Display specific crypto currency",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(name) == 0 {
				return fmt.Errorf("missing crypto currency name")
			}
			client, err := coinmarketcap.NewClient()
			if err != nil {
				return err
			}
			return cryptoCmd.getCryptoCurrency(client, currency, name)
		},
	}

	walletCryptoCmd := &cobra.Command{
		Use:   "wallet",
		Short: "Display specific crypto currencies",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(coins) == 0 {
				return fmt.Errorf("missing crypto currencies names")
			}
			client, err := coinmarketcap.NewClient()
			if err != nil {
				return err
			}
			return cryptoCmd.getCryptoWallet(client, coins, currency)
		},
	}
	listCryptoCmd.PersistentFlags().Int64Var(&limit, "limit", coinmarketcap.DefaultLimit, "Return a maximum of crypto currencies")
	listCryptoCmd.PersistentFlags().StringVar(&currency, "currency", coinmarketcap.DefaultCurrency, "Default currency to used")
	getCryptoCmd.PersistentFlags().StringVar(&name, "name", "", "Crypto currency name")
	getCryptoCmd.PersistentFlags().StringVar(&currency, "currency", coinmarketcap.DefaultCurrency, "Default currency to used")
	walletCryptoCmd.PersistentFlags().StringSliceVar(&coins, "coins", nil, "Coins' names")
	walletCryptoCmd.PersistentFlags().StringVar(&currency, "currency", coinmarketcap.DefaultCurrency, "Default currency to used")
	cmd.AddCommand(walletCryptoCmd)
	cmd.AddCommand(getCryptoCmd)
	cmd.AddCommand(listCryptoCmd)
	return cmd
}

func (cmd cryptoCmd) listCryptoCurrencies(client *coinmarketcap.Client, currency string, result int64) error {
	glog.V(1).Info("List crypto currencies")
	coins, err := client.GetCoins(currency, result)
	if err != nil {
		return err
	}
	return cmd.displayCoins(coins, currency)
}

func (cmd cryptoCmd) getCryptoCurrency(client *coinmarketcap.Client, name string, currency string) error {
	glog.V(1).Infof("Get crypto currency: %s", name)
	coin, err := client.GetCoin(name, currency, 1)
	if err != nil {
		return err
	}
	return cmd.displayCoins(coin, currency)
}

func (cmd cryptoCmd) getCryptoWallet(client *coinmarketcap.Client, names []string, currency string) error {
	glog.V(1).Infof("Get crypto wallet: %s", coins)
	coins := []coinmarketcap.Coin{}
	for _, name := range names {
		coin, err := client.GetCoin(name, currency, 1)
		if err != nil {
			return err
		}
		coins = append(coins, coin...)
	}
	return cmd.displayCoins(coins, currency)
}

func (cmd cryptoCmd) displayCoins(coins []coinmarketcap.Coin, currency string) error {
	ac := getAccounting(currency)
	table := tablewriter.NewWriter(cmd.out)
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
			pkgcmd.YellowOut(coin.Rank),
			pkgcmd.BlueOut(coin.Symbol),
			pkgcmd.BlueOut(coin.Name),
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

func getMoney(ac accounting.Accounting, value string) string {
	money, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return value
	}
	return ac.FormatMoney(money)
}

func getPercentColor(value string) string {
	var percent string
	if strings.HasPrefix(value, "-") {
		percent = pkgcmd.RedOut(value)
	} else {
		percent = pkgcmd.GreenOut(value)
	}
	return percent
}
