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
	name  string
	names []string
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
			return cryptoCmd.listCryptoCurrencies(client)
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
			return cryptoCmd.getCryptoCurrency(client, name)
		},
	}

	walletCryptoCmd := &cobra.Command{
		Use:   "wallet",
		Short: "Display specific crypto currencies",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	getCryptoCmd.PersistentFlags().StringVar(&name, "name", "", "Crypto currency name")
	walletCryptoCmd.PersistentFlags().StringSlice("cryptos", nil, "Cryptos' names")
	cmd.AddCommand(walletCryptoCmd)
	cmd.AddCommand(getCryptoCmd)
	cmd.AddCommand(listCryptoCmd)
	return cmd
}

func (cmd cryptoCmd) listCryptoCurrencies(client *coinmarketcap.Client) error {
	glog.V(1).Info("List crypto currencies")
	coins, err := client.GetCoins("EUR", 100)
	if err != nil {
		return err
	}
	return cmd.displayCoins(coins)
}

func (cmd cryptoCmd) getCryptoCurrency(client *coinmarketcap.Client, name string) error {
	glog.V(1).Infof("Get crypto currency: %s", name)
	coins, err := client.GetCoin(name, "EUR", 100)
	if err != nil {
		return err
	}
	return cmd.displayCoins(coins)
}

func (cmd cryptoCmd) displayCoins(coins []coinmarketcap.Coin) error {
	ac := accounting.Accounting{Symbol: "€", Precision: 4}
	table := tablewriter.NewWriter(cmd.out)
	table.SetHeader([]string{
		"Rank",
		"Symbol",
		"Coin",
		"EUR Price",
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
			getMoney(ac, coin.PriceEur),
			getMoney(ac, coin.Two4HVolumeEur),
			getMoney(ac, coin.MarketCapEur),
			getPercentColor(coin.PercentChange1H),
			getPercentColor(coin.PercentChange24H),
			getPercentColor(coin.PercentChange7D),
			coin.LastUpdated,
		})
	}
	table.Render()
	return nil
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
