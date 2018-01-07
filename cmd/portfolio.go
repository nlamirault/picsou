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

package cmd

import (
	"io"
	"strconv"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/nlamirault/picsou/config"
	pkgcmd "github.com/nlamirault/picsou/pkg/cmd"
	pkgcoins "github.com/nlamirault/picsou/pkg/coins"
	"github.com/nlamirault/picsou/providers/coinmarketcap"
)

var (
	configFilename string
)

type portfolioCmd struct {
	out io.Writer
}

func newPortfolioCmd(out io.Writer) *cobra.Command {
	portfolioCmd := &portfolioCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "portfolio",
		Short: "Manage portfolio.",
		Long:  "Manage portfolio.",
		RunE:  nil,
	}

	getPortfolioCmd := &cobra.Command{
		Use:   "get",
		Short: "Display portfolio.",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := config.LoadFileConfig(configFilename)
			if err != nil {
				return err
			}
			client, err := coinmarketcap.NewClient()
			if err != nil {
				return err
			}
			return portfolioCmd.getPortfolio(client, conf)
		},
	}

	statusPortfolioCmd := &cobra.Command{
		Use:   "status",
		Short: "Display portfolio currencies status.",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := config.LoadFileConfig(configFilename)
			if err != nil {
				return err
			}
			client, err := coinmarketcap.NewClient()
			if err != nil {
				return err
			}
			return portfolioCmd.portfolioStatus(client, conf)
		},
	}

	cmd.PersistentFlags().StringVar(&configFilename, "config", "picsou.toml", "Configuration filename")
	cmd.AddCommand(getPortfolioCmd)
	cmd.AddCommand(statusPortfolioCmd)
	return cmd
}

// type walletCoin struct {
// 	Name    string
// 	Percent float64
// 	Money   float64
// }

type portfolio struct {
	coins []coinmarketcap.Coin
	// wallet map[string]coins.Wallet
	// total  float64
	wallet *pkgcoins.Wallet
}

func (cmd portfolioCmd) getPortfolio(client *coinmarketcap.Client, conf *config.Configuration) error {
	glog.V(1).Infof("Get crypto currencies portfolio: %s", conf)

	ac := pkgcmd.GetAccounting(conf.Currency)

	portfolio, err := cmd.retrievePortofolio(client, conf.Currency, conf.Portfolio)
	if err != nil {
		return err
	}
	return pkgcmd.DisplayWalletTable(cmd.out, portfolio.wallet, ac)
}

func (cmd portfolioCmd) portfolioStatus(client *coinmarketcap.Client, conf *config.Configuration) error {
	glog.V(1).Infof("Crypto currencies portfolio status: %s", conf)

	portfolio, err := cmd.retrievePortofolio(client, conf.Currency, conf.Portfolio)
	if err != nil {
		return err
	}
	return pkgcmd.DisplayCoins(cmd.out, portfolio.coins, conf.Currency)
}

func (cmd portfolioCmd) retrievePortofolio(client *coinmarketcap.Client, currency string, currencies map[string]string) (*portfolio, error) {
	portfolio := &portfolio{
		coins: []coinmarketcap.Coin{},
		wallet: &pkgcoins.Wallet{
			Coins: map[string]pkgcoins.Coin{},
		},
	}

	for name, owned := range currencies {
		coin, err := client.GetCoin(name, currency, 1)
		if err != nil {
			return nil, err
		}
		portfolio.coins = append(portfolio.coins, coin...)
		nb, err := strconv.ParseFloat(owned, 64)
		if err != nil {
			return nil, err
		}
		price, err := strconv.ParseFloat(coinmarketcap.GetPrice(coin[0], currency), 64)
		if err != nil {
			return nil, err
		}
		portfolio.wallet.Coins[name] = pkgcoins.Coin{
			Name:  name,
			Money: nb * price,
		}
		portfolio.wallet.Total = portfolio.wallet.Total + portfolio.wallet.Coins[name].Money
	}
	glog.V(2).Infof("Portfolio: %s", portfolio)
	return portfolio, nil
}
