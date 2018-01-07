# Picsou

[![License Apache 2][badge-license]](LICENSE)
[![GitHub version](https://badge.fury.io/gh/nlamirault%2Fpicsou.svg)](https://badge.fury.io/gh/nlamirault%2Fpicsou)

* Master : [![pipeline status](https://gitlab.com/nicolas-lamirault/picsou/badges/master/pipeline.svg)](https://gitlab.com/nicolas-lamirault/picsou/commits/master)

* Develop : [![pipeline status](https://gitlab.com/nicolas-lamirault/picsou/badges/develop/pipeline.svg)](https://gitlab.com/nicolas-lamirault/picsou/commits/develop)


This tools is a CLI which cant display crypto currencies informations.
Available providers :

* [x] Coinmarketcap

![List](assets/images/picsou-list-0.2.0.png)

![Portfolio](assets/images/picsou-portfolio-0.2.0.png)


## Installation

You can download the binaries :

* Architecture i386 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/picsou-0.3.0_linux_386) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/picsou-0.3.0_darwin_386) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/picsou-0.3.0_freebsd_386) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/picsou-0.3.0_netbsd_386) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/picsou-0.3.0_openbsd_386) / [windows](https://bintray.com/artifact/download/nlamirault/oss/picsou-0.3.0_windows_386.exe) ]
* Architecture amd64 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/picsou-0.3.0_linux_amd64) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/picsou-0.3.0_darwin_amd64) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/picsou-0.3.0_freebsd_amd64) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/picsou-0.3.0_netbsd_amd64) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/picsou-0.3.0_openbsd_amd64) / [windows](https://bintray.com/artifact/download/nlamirault/oss/picsou-0.3.0_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/nlamirault/oss/picsou-0.3.0_linux_arm) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/picsou-0.3.0_freebsd_arm) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/picsou-0.3.0_netbsd_arm) ]


## Usage

* CLI help:

        $ picsou help

* List crypto currencies:

        $ picsou  crypto list
        +------+--------+-------------------------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | RANK | SYMBOL |          COIN           |  EUR PRICE   |    24 HOUR VOLUME    |      MARKET CAP       | 1 HOUR | 24 HOUR | 7 DAYS | LAST UPDATED |
        +------+--------+-------------------------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | 1    | BTC    | Bitcoin                 | €11,473.2375 | €10,629,779,512.0000 | €192,297,908,015.0000 | 0.57   | -11.11  | -30.48 |   1514129055 |
        +------+--------+-------------------------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | 2    | ETH    | Ethereum                | €560.7617    | €2,073,011,060.4500  | €54,136,718,951.0000  | 0.2    | -11.26  | -7.54  |   1514129059 |
        +------+--------+-------------------------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | 3    | BCH    | Bitcoin Cash            | €2,366.9523  | €1,755,375,736.9500  | €39,938,326,226.0000  | -0.82  | -17.23  | 50.8   |   1514128788 |
        +------+--------+-------------------------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | 4    | XRP    | Ripple                  | €0.8334      | €650,536,873.5250    | €32,286,485,475.0000  | 0.75   | -16.14  | 34.27  |   1514129041 |
        +------+--------+-------------------------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | 5    | LTC    | Litecoin                | €226.2513    | €1,009,923,492.6000  | €12,319,633,746.0000  | -0.11  | -11.59  | -16.28 |   1514129042 |
        +------+--------+-------------------------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | 6    | ADA    | Cardano                 | €0.3180      | €64,642,203.3125     | €8,243,988,863.0000   | 1.1    | -13.6   | -25.57 |   1514128797 |
        +------+--------+-------------------------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | 7    | MIOTA  | IOTA                    | €2.8779      | €228,179,705.0500    | €7,999,282,807.0000   | 0.87   | -16.3   | -5.7   |   1514128782 |
        +------+--------+-------------------------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+


* Check a specific crypto currency:

        $ picsou  crypto get --name BTC
        +------+--------+---------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | RANK | SYMBOL |  COIN   |  EUR PRICE   |    24 HOUR VOLUME    |      MARKET CAP       | 1 HOUR | 24 HOUR | 7 DAYS | LAST UPDATED |
        +------+--------+---------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | 1    | BTC    | Bitcoin | €11,518.5176 | €10,707,270,051.5000 | €193,056,828,081.0000 | 1.07   | -10.75  | -30.21 |   1514128755 |
        +------+--------+---------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+

* Check your wallet currencies:

        $ picsou portfolio status --config picsou.example.toml
        +------+--------+----------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | RANK | SYMBOL |   COIN   |    PRICE     |    24 HOUR VOLUME    |      MARKET CAP       | 1 HOUR | 24 HOUR | 7 DAYS | LAST UPDATED |
        +------+--------+----------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | 1    | BTC    | Bitcoin  | €13,960.6724 | €13,163,199,920.0000 | €234,369,841,477.0000 | 0.49   | -0.74   | 23.03  |   1515337760 |
        +------+--------+----------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | 28   | DOGE   | Dogecoin | €0.0152      | €210,006,070.5600    | €1,710,898,885.0000   | 9.51   | 20.59   | 110.39 |   1515337741 |
        +------+--------+----------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | 3    | ETH    | Ethereum | €930.7918    | €4,305,370,487.6000  | €90,128,470,218.0000  | 1.28   | 8.2     | 51.39  |   1515337748 |
        +------+--------+----------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | 7    | LTC    | Litecoin | €246.3179    | €996,449,911.6000    | €13,462,335,299.0000  | 0.82   | -2.11   | 31.08  |   1515337741 |
        +------+--------+----------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+

* Portfolio :

        $ picsou portfolio get --config picsou.example.toml
        +--------+-----------+---------+------------------------------------------------------+
        | SYMBOL |   MONEY   | PERCENT |                         VUE                          |
        +--------+-----------+---------+------------------------------------------------------+
        | BTC    | €181.4887 |      52 | ∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎ |
        +--------+-----------+---------+------------------------------------------------------+
        | DOGE   | €2.2779   |       1 | ∎                                                    |
        +--------+-----------+---------+------------------------------------------------------+
        | ETH    | €134.9648 |      39 | ∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎              |
        +--------+-----------+---------+------------------------------------------------------+
        | LTC    | €30.2971  |       9 | ∎∎∎∎∎∎∎∎∎                                            |
        +--------+-----------+---------+------------------------------------------------------+

        $ picsou portfolio status --config picsou.example.toml
        +------+--------+----------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | RANK | SYMBOL |   COIN   |    PRICE     |    24 HOUR VOLUME    |      MARKET CAP       | 1 HOUR | 24 HOUR | 7 DAYS | LAST UPDATED |
        +------+--------+----------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | 1    | BTC    | Bitcoin  | €13,936.4008 | €13,693,684,524.0000 | €233,962,372,803.0000 | 0.2    | -1.02   | 22.77  |   1515338060 |
        +------+--------+----------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | 28   | DOGE   | Dogecoin | €0.0153      | €212,094,926.4200    | €1,724,964,711.0000   | 9.36   | 21.36   | 111.98 |   1515338041 |
        +------+--------+----------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | 3    | ETH    | Ethereum | €925.8960    | €4,307,340,479.0000  | €89,654,402,143.0000  | 0.54   | 7.58    | 50.52  |   1515338049 |
        +------+--------+----------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+
        | 7    | LTC    | Litecoin | €246.4501    | €996,998,516.8000    | €13,469,558,630.0000  | 0.77   | -2.06   | 31.07  |   1515338041 |
        +------+--------+----------+--------------+----------------------+-----------------------+--------+---------+--------+--------------+

With this example portfolio:

```toml
# Picsou - Portfolio configuration file

currency = "EUR"

[portfolio]
BTC = "0.013"
DOGE = "150"
ETH = "0.145"
LTC = "0.123"
```




## Development

* Initialize environment

        $ make init

* Build tool :

        $ make build

* Launch unit tests :

        $ make test

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>

[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat
