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

package coinmarketcap

import (
	"fmt"
	"strings"
)

const (
	defaultCurrency = "USD"
	defaultLimit    = 100
)

type Coin struct {
	Name      string
	Symbol    string
	Price     float64
	Currency  string
	Volume24  float64
	Maxsupply float64
}

var cryptoCurrencies = map[string]string{
	"42":        "42-coin",
	"300":       "300-token",
	"611":       "sixeleven",
	"808":       "808coin",
	"888":       "octocoin",
	"1337":      "1337coin",
	"BTC":       "bitcoin",
	"ETH":       "ethereum",
	"BCH":       "bitcoin-cash",
	"XRP":       "ripple",
	"DASH":      "dash",
	"LTC":       "litecoin",
	"XEM":       "nem",
	"MIOTA":     "iota",
	"XMR":       "monero",
	"ETC":       "ethereum-classic",
	"NEO":       "neo",
	"OMG":       "omisego",
	"BCC":       "bitconnect",
	"LSK":       "lisk",
	"QTUM":      "qtum",
	"ZEC":       "zcash",
	"USDT":      "tether",
	"STRAT":     "stratis",
	"WAVES":     "waves",
	"ARK":       "ark",
	"STEEM":     "steem",
	"BCN":       "bytecoin-bcn",
	"PAY":       "tenx",
	"MAID":      "maidsafecoin",
	"GNT":       "golem-network-tokens",
	"BAT":       "batcoin",
	"EOS":       "eos",
	"DCR":       "decred",
	"REP":       "augur",
	"XLM":       "stellar",
	"KMD":       "komodo",
	"BTS":       "bitshares",
	"HSR":       "hshare",
	"VERI":      "veritaseum",
	"PIVX":      "pivx",
	"MTL":       "metal",
	"ICN":       "icoin",
	"FCT":       "factom",
	"SC":        "siacoin",
	"GAS":       "gas",
	"DGD":       "digixdao",
	"GBYTE":     "byteball",
	"CVC":       "civic",
	"PPT":       "populous",
	"DGB":       "digibyte",
	"ARDR":      "ardor",
	"NXS":       "nexus",
	"GXS":       "gxshares",
	"GAME":      "gamecredits",
	"DOGE":      "dogecoin",
	"BTCD":      "bitcoindark",
	"SNGLS":     "singulardtv",
	"MCAP":      "mcap",
	"GNO":       "gnosis-gno",
	"DCN":       "dentacoin",
	"ZRX":       "0x",
	"BNT":       "bancor",
	"LKK":       "lykke",
	"BLOCK":     "blocknet",
	"MCO":       "monaco",
	"AE":        "aeternity",
	"BTM":       "bitmark",
	"FUN":       "funfair",
	"SNT":       "status",
	"BNB":       "binance-coin",
	"EDG":       "edgeless",
	"SYS":       "syscoin",
	"XVG":       "verge",
	"BDL":       "bitdeal",
	"FRST":      "firstcoin",
	"NXT":       "nxt",
	"ANT":       "aragon",
	"BQX":       "bitquence",
	"UBQ":       "ubiq",
	"PART":      "particl",
	"RISE":      "rise",
	"IOC":       "iocoin",
	"WINGS":     "wings",
	"NAV":       "nav-coin",
	"LINK":      "chainlink",
	"STORJ":     "storj",
	"MGO":       "mobilego",
	"CFI":       "cofound-it",
	"WTC":       "walton",
	"VTC":       "vertcoin",
	"PLR":       "pillar",
	"RLC":       "rlc",
	"TNT":       "tierion",
	"LRC":       "loopring",
	"NLG":       "gulden",
	"FAIR":      "faircoin",
	"ETP":       "metaverse",
	"TRIG":      "triggers",
	"XRL":       "rialto",
	"MLN":       "melon",
	"XEL":       "elastic",
	"CLOAK":     "cloakcoin",
	"XZC":       "zcoin",
	"NLC2":      "nolimitcoin",
	"ADK":       "aidos-kuneen",
	"ADX":       "adx-net",
	"TKN":       "tokencard",
	"TRST":      "trust",
	"PPC":       "peercoin",
	"NXC":       "nexium",
	"MTH":       "monetha",
	"QRL":       "quantum-resistant-ledger",
	"RDD":       "reddcoin",
	"PTOY":      "patientory",
	"DLT":       "agrello-delta",
	"TAAS":      "taas",
	"DICE":      "etheroll",
	"1ST":       "firstblood",
	"XCP":       "counterparty",
	"VIA":       "viacoin",
	"DCT":       "decent",
	"OK":        "okcash",
	"MONA":      "monacoin",
	"SNM":       "sonm",
	"EMC":       "emercoin",
	"XAUR":      "xaurum",
	"DNT":       "district0x",
	"CRW":       "crown",
	"BMC":       "blackmoon-crypto",
	"TCC":       "the-champcoin",
	"SAFEX":     "safe-exchange-coin",
	"MSP":       "mothership",
	"LEO":       "leocoin",
	"ION":       "ion",
	"XAS":       "asch",
	"ADT":       "adtoken",
	"HMQ":       "humaniq",
	"NMC":       "namecoin",
	"CLAM":      "clams",
	"BAY":       "bitbay",
	"SAN":       "santiment",
	"PLBT":      "polybius",
	"LUN":       "lunyr",
	"UNO":       "unobtanium",
	"SPRTS":     "sprouts",
	"POE":       "poet",
	"VSL":       "vslice",
	"SKY":       "skycoin",
	"EXP":       "expanse",
	"ROUND":     "round",
	"DMD":       "diamond",
	"MYST":      "mysterium",
	"NMR":       "numeraire",
	"CDT":       "coindash",
	"MUE":       "monetaryunit",
	"SIB":       "sibcoin",
	"SPR":       "spreadcoin",
	"POT":       "potcoin",
	"ZEN":       "zencash",
	"AGRS":      "agoras-tokens",
	"RBY":       "rubycoin",
	"RADS":      "radium",
	"XDN":       "digitalnote",
	"QAU":       "quantum",
	"XPA":       "xplay",
	"LBC":       "library-credit",
	"SHIFT":     "shift",
	"CMP":       "compcoin",
	"BURST":     "burst",
	"BLK":       "blackcoin",
	"EMC2":      "einsteinium",
	"STX":       "stox",
	"MOON":      "mooncoin",
	"ENRG":      "energycoin",
	"ATCC":      "atc-coin",
	"CREDO":     "credo",
	"OMNI":      "omni",
	"GAM":       "gambit",
	"SLS":       "salus",
	"NEBL":      "neblio",
	"GUP":       "guppy",
	"MXT":       "martexcoin",
	"GRC":       "gridcoin",
	"B@":        "bankcoin",
	"SWT":       "swarm-city",
	"TIME":      "chronobank",
	"MDA":       "moeda-loyalty-points",
	"AMP":       "synereo",
	"GOLOS":     "golos",
	"ECN":       "e-coin",
	"WGR":       "wagerr",
	"IXT":       "ixledger",
	"AEON":      "aeon",
	"WCT":       "waves-community-token",
	"GRS":       "groestlcoin",
	"IFT":       "investfeed",
	"DTB":       "databits",
	"NEOS":      "neoscoin",
	"LMC":       "lomocoin",
	"BCAP":      "bcap",
	"PST":       "primas",
	"OAX":       "openanx",
	"AVT":       "aventus",
	"EDR":       "e-dinar-coin",
	"FLO":       "florincoin",
	"VOX":       "voxels",
	"NET":       "netcoin",
	"PPY":       "peerplays-ppy",
	"VRC":       "vericoin",
	"XRB":       "raiblocks",
	"PRO":       "propy",
	"SMART":     "smartcash",
	"KORE":      "korecoin",
	"NVC":       "novacoin",
	"HVN":       "hive",
	"BCY":       "bitcrystals",
	"OBITS":     "obits",
	"BITCNY":    "bitcny",
	"INCNT":     "incent",
	"XSPEC":     "spectrecoin",
	"TOA":       "toacoin",
	"PINK":      "pinkcoin",
	"ECOB":      "ecobit",
	"MUSIC":     "musicoin",
	"TIX":       "tickets",
	"PLU":       "pluton",
	"XC":        "xcurrency",
	"FTC":       "feathercoin",
	"BET":       "betacoin",
	"PZM":       "prizm",
	"SOAR":      "soarcoin",
	"SPHR":      "sphere",
	"XVC":       "vcash",
	"CRB":       "creditbit",
	"SLR":       "solarcoin",
	"ECC":       "eccoin",
	"XWC":       "whitecoin",
	"BSD":       "bitsend",
	"PEPECASH":  "pepe-cash",
	"PDC":       "project-decorum",
	"IOP":       "internet-of-people",
	"UNY":       "unity-ingot",
	"NOTE":      "dnotes",
	"APX":       "apx",
	"SEQ":       "sequence",
	"BITB":      "bitbean",
	"MYB":       "mybit-token",
	"XBC":       "bitcoin-plus",
	"XBY":       "xtrabytes",
	"QWARK":     "qwark",
	"EMB":       "embercoin",
	"CAT":       "catcoin",
	"ERC":       "europecoin",
	"ABY":       "applebyte",
	"TFL":       "trueflip",
	"XST":       "stealthcoin",
	"EXCL":      "exclusivecoin",
	"BLITZ":     "blitzcash",
	"HEAT":      "heat-ledger",
	"JINN":      "jinn",
	"PASC":      "pascal-coin",
	"DAR":       "darcrus",
	"AUR":       "auroracoin",
	"BELA":      "belacoin",
	"DAXX":      "daxxcoin",
	"CURE":      "curecoin",
	"COVAL":     "circuits-of-value",
	"NVST":      "nvo",
	"EQT":       "equitrader",
	"DENT":      "dent",
	"DBIX":      "dubaicoin-dbix",
	"GLD":       "goldcoin",
	"NDC":       "neverdie",
	"FLDC":      "foldingcoin",
	"BTX":       "bitcointx",
	"RAIN":      "condensate",
	"SWIFT":     "bitswift",
	"SNRG":      "synergy",
	"PTC":       "pesetacoin",
	"CVCOIN":    "cvcoin",
	"POSW":      "posw-coin",
	"GCR":       "global-currency-reserve",
	"HTML5":     "htmlcoin",
	"SYNX":      "syndicate",
	"ATMS":      "atmos",
	"SIGT":      "signatum",
	"SNC":       "suncontract",
	"DYN":       "dynamic",
	"OCT":       "oraclechain",
	"TIPS":      "fedoracoin",
	"PUT":       "putincoin",
	"HUSH":      "hush",
	"ZRC":       "zrcoin",
	"SBD":       "steem-dollars",
	"HPC":       "happycoin",
	"XMY":       "myriad",
	"CV2":       "colossuscoin-v2",
	"GEO":       "geocoin",
	"BITUSD":    "bitusd",
	"DOPE":      "dopecoin",
	"CSC":       "casinocoin",
	"BTA":       "bata",
	"NTRN":      "neutron",
	"ZCL":       "zclassic",
	"REX":       "real-estate-tokens",
	"PBT":       "primalbase",
	"BRX":       "breakout-stake",
	"OCL":       "oceanlab",
	"2GIVE":     "2give",
	"CHC":       "chaincoin",
	"VISIO":     "visio",
	"TX":        "transfercoin",
	"MBRS":      "embers",
	"THC":       "hempcoin",
	"WBB":       "wild-beast-block",
	"ZEIT":      "zeitcoin",
	"VRM":       "veriumreserve",
	"OPT":       "opus",
	"AC":        "asiacoin",
	"DIME":      "dimecoin",
	"BRK":       "breakout",
	"TKS":       "tokes",
	"CNT":       "centurion",
	"ESP":       "espers",
	"CRAVE":     "crave",
	"EGC":       "evergreencoin",
	"NSR":       "nushares",
	"START":     "startcoin",
	"XPM":       "primecoin",
	"XMCC":      "monacocoin",
	"TRC":       "terracoin",
	"CADASTRAL": "bitland",
	"CREA":      "creativecoin",
	"TRUST":     "trustplus",
	"BLU":       "bluecoin",
	"XTO":       "tao",
	"VTR":       "vtorrent",
	"CAGE":      "cagecoin",
	"CANN":      "cannabiscoin",
	"MEME":      "memetic",
	"ICOO":      "ico-openledger",
	"STA":       "starta",
	"WTT":       "giga-watt-token",
	"B3":        "b3coin",
	"ONION":     "deeponion",
	"NAUT":      "nautiluscoin",
	"LINDA":     "linda",
	"LGD":       "legends-room",
	"ADZ":       "adzcoin",
	"FYN":       "fundyourselfnow",
	"PKB":       "parkbyte",
	"HUC":       "huntercoin",
	"PING":      "cryptoping",
	"ZENI":      "zennies",
	"RIC":       "riecoin",
	"XMG":       "magi",
	"SKIN":      "skincoin",
	"MINT":      "mintcoin",
	"ADL":       "adelphoi",
	"EFL":       "e-gulden",
	"ADST":      "adshares",
	"YBC":       "ybcoin",
	"ALT":       "altcoin-alt",
	"CBX":       "cryptogenic-bullion",
	"INPAY":     "inpay",
	"ARC":       "arcade-token",
	"HYP":       "hyperstake",
	"CCRB":      "cryptocarbon",
	"MNE":       "minereum",
	"FUCK":      "fucktoken",
	"SMLY":      "smileycoin",
	"LNK":       "link-platform",
	"BASH":      "luckchain",
	"XP":        "xp",
	"VASH":      "vpncoin",
	"PIE":       "piecoin",
	"CPC":       "capricoin",
	"IXC":       "ixcoin",
	"MER":       "mercury",
	"ADC":       "audiocoin",
	"GRWI":      "growers-international",
	"VIVO":      "vivo",
	"IFC":       "infinitecoin",
	"DP":        "digitalprice",
	"QRK":       "quark",
	"INSN":      "insanecoin-insn",
	"DRACO":     "dt-token",
	"IFLT":      "inflationcoin",
	"HNC":       "huncoin",
	"NTO":       "fujinto",
	"ZOI":       "zoin",
	"INFX":      "influxcoin",
	"EBST":      "eboostcoin",
	"BYC":       "bytecent",
	"ONX":       "onix",
	"USNBT":     "nubits",
	"RNS":       "renos",
	"XBL":       "billionaire-token",
	"SXC":       "sexcoin",
	"MRT":       "miners-reward-token",
	"LDOGE":     "litedoge",
	"WDC":       "worldcoin",
	"DOT":       "dotcoin",
	"XHI":       "hicoin",
	"EAC":       "earthcoin",
	"HTC":       "hitcoin",
	"SUMO":      "sumokoin",
	"DDF":       "digital-developers-fund",
	"ITI":       "iticoin",
	"FIMK":      "fimkrypto",
	"MAX":       "maxcoin",
	"RMC":       "remicoin",
	"NETKO":     "netko",
	"MGC":       "gulfcoin",
	"XVS":       "vsync",
	"BLOCKPAY":  "blockpay",
	"VSX":       "vsync-vsx",
	"ENT":       "eternity",
	"WGO":       "wavesgo",
	"HERO":      "sovereign-hero",
	"TYCHO":     "tychocoin",
	"PND":       "pandacoin-pnd",
	"UIS":       "unitus",
	"KEK":       "kekcoin",
	"UNIFY":     "unify",
	"BRO":       "bitradio",
	"MEC":       "megacoin",
	"NYC":       "newyorkcoin",
	"DNR":       "denarius-dnr",
	"CARBON":    "carboncoin",
	"TOKEN":     "swaptoken",
	"UNB":       "unbreakablecoin",
	"ICE":       "idice",
	"GRE":       "greencoin",
	"PROC":      "procurrency",
	"FCN":       "fantomcoin",
	"MOIN":      "moin",
	"CRM":       "cream",
	"EMP":       "emoneypower",
	"LINX":      "linx",
	"I0C":       "i0coin",
	"XVP":       "virtacoinplus",
	"KRS":       "krypstal",
	"DAS":       "das",
	"ZER":       "zero",
	"VUC":       "virta-unique-coin",
	"KRB":       "karbowanec",
	"FAL":       "falcoin",
	"RLT":       "roulettetoken",
	"POST":      "postcoin",
	"RBIES":     "rubies",
	"SMOKE":     "smoke",
	"EOT":       "eot-token",
	"DCY":       "dinastycoin",
	"XCXT":      "coinonatx",
	"ZCC":       "zccoin",
	"HBN":       "hobonickels",
	"BAS":       "bitasean",
	"NEWB":      "newbium",
	"GCC":       "guccionecoin",
	"VLT":       "veltor",
	"WOMEN":     "women",
	"BITS":      "bitstar",
	"TRK":       "truckcoin",
	"TRUMP":     "trumpcoin",
	"BUCKS":     "swagbucks",
	"PR":        "prototanium",
	"FXE":       "futurexe",
	"GCN":       "gcoin",
	"ATOM":      "atomic-coin",
	"TIT":       "titcoin",
	"LTB":       "litebar",
	"EL":        "elcoin-el",
	"MAO":       "mao-zedong",
	"PXC":       "phoenixcoin",
	"XPTX":      "platinumbar",
	"BITBTC":    "bitbtc",
	"SCORE":     "scorecoin",
	"XLR":       "solaris",
	"HONEY":     "honey",
	"BUN":       "bunnycoin",
	"BITSILVER": "bitsilver",
	"TEK":       "tekcoin",
	"DALC":      "dalecoin",
	"XRA":       "ratecoin",
	"ZUR":       "zurcoin",
	"BTCS":      "bitcoin-scrypt",
	"LANA":      "lanacoin",
	"YOC":       "yocoin",
	"BITGOLD":   "bitgold",
	"BTWTY":     "bit20",
	"KLC":       "kilocoin",
	"GLT":       "globaltoken",
	"PUTIC":     "putin-classic",
	"ECO":       "ecocoin",
	"CJ":        "cryptojacks",
	"MSCN":      "master-swiscoin",
	"BSTY":      "globalboost-y",
	"XTC":       "tilecoin",
	"LCP":       "litecoin-plus",
	"BITEUR":    "biteur",
	"GB":        "goldblocks",
	"RUP":       "rupee",
	"KURT":      "kurrent",
	"4CHN":      "chancoin",
	"GPU":       "gpu-coin",
	"PLNC":      "plncoin",
	"XCO":       "x-coin",
	"ATX":       "artex-coin",
	"TSTR":      "tristar-coin",
	"BLAS":      "blakestar",
	"BRIA":      "briacoin",
	"DRXNE":     "droxne",
	"CMPCO":     "campuscoin",
	"RBT":       "rimbit",
	"VRS":       "veros",
	"BOAT":      "doubloon",
	"CNC":       "chncoin",
	"KRONE":     "kronecoin",
	"BENJI":     "benjirolls",
	"SOJ":       "sojourn",
	"NANOX":     "project-x",
	"APW":       "applecoin-apw",
	"LTCU":      "litecoin-ultra",
	"LVPS":      "levoplus",
	"ALTC":      "antilitecoin",
	"LBTC":      "litebitcoin",
	"EBT":       "ebittree-coin",
	"ULA":       "ulatech",
	"HMC":       "harmonycoin-hmc",
	"REE":       "reecoin",
	"ABN":       "abncoin",
	"UNITY":     "supernet-unity",
	"ETT":       "encryptotel",
	"STCN":      "stakecoin-stcn",
	"ETBS":      "ethbits",
	"VSM":       "voise",
	"BPC":       "bitpark-coin",
	"AHT":       "bowhead",
	"YASH":      "yashcoin",
	"JNS":       "janus",
	"NKA":       "incakoin",
	"FNC":       "fincoin",
	"RUSTBITS":  "rustbits",
	"LTBC":      "ltbcoin",
	"TES":       "teslacoin",
	"FST":       "fastcoin",
	"LOG":       "woodcoin",
	"MBI":       "monster-byte",
	"STRC":      "starcredits",
	"ZET":       "zetacoin",
	"USC":       "ultimate-secure-cash",
	"BRIT":      "britcoin",
	"SDC":       "shadowcash",
	"CRYPT":     "cryptcoin",
	"BITZ":      "bitz",
	"METAL":     "metalcoin",
	"JET":       "jetcoin",
	"RIYA":      "etheriya",
	"CASINO":    "casino",
	"CDN":       "canada-ecoin",
	"XCN":       "cryptonite",
	"SHORTY":    "shorty",
	"NOBL":      "noblecoin",
	"COE":       "coeval",
	"ORB":       "orbitcoin",
	"SUPER":     "supercoin",
	"TROLL":     "trollcoin",
	"BTB":       "bitbar",
	"FJC":       "fujicoin",
	"GLC":       "globalcoin",
	"TAG":       "tagcoin",
	"DFT":       "draftcoin",
	"MAC":       "machinecoin",
	"DVC":       "devcoin",
	"UFO":       "ufo-coin",
	"UTC":       "ultracoin",
	"MZC":       "mazacoin",
	"8BIT":      "8bit",
	"FC2":       "fuelcoin",
	"KOBO":      "kobocoin",
	"DSH":       "dashcoin",
	"BTSR":      "btsr",
	"RC":        "russiacoin",
	"MALC":      "malcoin",
	"SMC":       "smartcoin",
	"SHDW":      "shadow-token",
	"UNIC":      "unicoin",
	"ANC":       "anoncoin",
	"FUNK":      "the-cypherfunks",
	"PIGGY":     "piggycoin",
	"AMBER":     "ambercoin",
	"FUNC":      "funcoin",
	"PAK":       "pakcoin",
	"RAREPEPEP": "rare-pepe-party",
	"TALK":      "btctalkcoin",
	"CCN":       "cannacoin",
	"V":         "version",
	"DEM":       "deutsche-emark",
	"020":       "o2olondoncoin",
	"CFT":       "cryptoforecast",
	"TRI":       "triangles",
	"GAIA":      "gaia",
	"GOOD":      "goodomy",
	"FLT":       "fluttercoin",
	"LOT":       "lottocoin",
	"XJO":       "joulecoin",
	"AU":        "aurumcoin",
	"FLY":       "flycoin",
	"JIN":       "jin-coin",
	"STS":       "stress",
	"Q2C":       "qubitcoin",
	"SLM":       "slimcoin",
	"MTM":       "mtmgaming",
	"VIDZ":      "purevidz",
	"UNIT":      "universal-currency",
	"FRN":       "francs",
	"MNM":       "mineum",
	"VAL":       "valorbit",
	"HODL":      "hodlcoin",
	"PSB":       "pesobit",
	"CAP":       "bottlecaps",
	"SWING":     "swing",
	"SLG":       "sterlingcoin",
	"KED":       "darsek",
	"EUC":       "eurocoin",
	"CHESS":     "chesscoin",
	"EMD":       "emerald",
	"ACOIN":     "acoin",
	"XPY":       "paycoin2",
	"FRC":       "freicoin",
	"BLC":       "blakecoin",
	"YAC":       "yacoin",
	"KUSH":      "kushcoin",
	"USDE":      "usde",
	"WAY":       "wayguide",
	"UNI":       "universe",
	"CTO":       "crypto",
	"ICON":      "iconic",
	"UNITS":     "gameunits",
	"TTC":       "tittiecoin",
	"KIC":       "kibicoin",
	"QCN":       "quazarcoin",
	"RBX":       "ripto-bux",
	"XRE":       "revolvercoin",
	"VC":        "virtualcoin",
	"SPEX":      "sproutsextreme",
	"FLAX":      "flaxscript",
	"BXT":       "bittokens",
	"GRT":       "grantcoin",
	"XPD":       "petrodollar",
	"BOLI":      "bolivarcoin",
	"MOJO":      "mojocoin",
	"PASL":      "pascal-lite",
	"OHM":       "ohm-wallet",
	"HMP":       "hempcoin-hmp",
	"CYP":       "cypher",
	"DUO":       "parallelcoin",
	"NYAN":      "nyancoin",
	"C2":        "coin2-1",
	"SAC":       "sacoin",
	"TGC":       "tigercoin",
	"CUBE":      "digicube",
	"BIGUP":     "bigup",
	"DGC":       "digitalcoin",
	"EVIL":      "evil-coin",
	"BBP":       "biblepay",
	"J":         "joincoin",
	"BERN":      "berncash",
	"WMC":       "wmcoin",
	"SOIL":      "soilcoin",
	"XGR":       "goldreserve",
	"GUN":       "guncoin",
	"CRX":       "chronos",
	"MAD":       "satoshimadness",
	"PRC":       "prcoin",
	"PXI":       "prime-xi",
	"IMS":       "independent-money-system",
	"VEC2":      "vector",
	"SPACE":     "spacecoin",
	"UNIBURST":  "uniburst",
	"ICOB":      "icobid",
	"CNNC":      "cannation",
	"BTPL":      "bitcoin-planet",
	"SLING":     "sling",
	"BTCR":      "bitcurrency",
	"DLC":       "dollarcoin",
	"HKG":       "hacker-gold",
	"ZNY":       "bitzeny",
	"RUPX":      "rupaya",
	"URC":       "unrealcoin",
	"XCT":       "c-bit",
	"ARCO":      "aquariuscoin",
	"ISL":       "islacoin",
	"SCRT":      "secretcoin",
	"CACH":      "cachecoin",
	"ECA":       "electra",
	"E4ROW":     "ether-for-the-rest-of-the-world",
	"MAY":       "theresa-may-coin",
	"CON":       "paycon",
	"ELE":       "elementrem",
	"WYV":       "wyvern",
	"BIP":       "bipcoin",
	"HAL":       "halcyon",
	"ANTI":      "antibitcoin",
	"RED":       "redcoin",
	"UET":       "useless-ethereum-token",
	"MST":       "mustangcoin",
	"STV":       "sativacoin",
	"DRM":       "dreamcoin",
	"SPT":       "spots",
	"ADCN":      "asiadigicoin",
	"ARI":       "aricoin",
	"AGLC":      "agrolifecoin",
	"CORG":      "corgicoin",
	"DIBC":      "dibcoin",
	"MARS":      "marscoin",
	"FIRE":      "firecoin",
	"SRC":       "securecoin",
	"ASAFE2":    "allsafe",
	"BOST":      "boostcoin",
	"CNO":       "coin",
	"BSTAR":     "blackstar",
	"ARG":       "argentum",
	"GAP":       "gapcoin",
	"EVO":       "evotion",
	"CXT":       "coinonat",
	"ALL":       "allion",
	"DBTC":      "debitcoin",
	"BCF":       "bitcoinfast",
	"PONZI":     "ponzicoin",
	"CREVA":     "crevacoin",
	"FRK":       "franko",
	"BUMBA":     "bumbacoin",
	"MEOW":      "kittehcoin",
	"JWL":       "jewels",
	"ERY":       "eryllium",
	"XBTC21":    "bitcoin-21",
	"$$$":       "money",
	"GPL":       "gold-pressed-latinum",
	"WORM":      "healthywormcoin",
	"TSE":       "tattoocoin",
	"VIP":       "vip-tokens",
	"MAR":       "marijuanacoin",
	"CPN":       "compucoin",
	"QTL":       "quatloo",
	"DRS":       "digital-rupees",
	"MCRN":      "macron",
	"PHS":       "philosopher-stones",
	"URO":       "uro",
	"BLRY":      "billarycoin",
	"MND":       "mindcoin",
	"BVC":       "beavercoin",
	"GP":        "goldpieces",
	"POP":       "popularcoin",
	"GBC":       "gbcgoldcoin",
	"ARB":       "arbit",
	"MTLMC3":    "metal-music-coin",
	"SONG":      "songcoin",
	"FLVR":      "flavorcoin",
	"XCRE":      "creatio",
	"LEA":       "leacoin",
	"STEPS":     "steps",
	"BIOS":      "bios-crypto",
	"ZYD":       "zayedcoin",
	"RPC":       "ronpaulcoin",
	"WARP":      "warp",
	"CMT":       "comet",
	"NEVA":      "nevacoin",
	"SOON":      "sooncoin",
	"MILO":      "milocoin",
	"PHO":       "photon",
	"G3N":       "genstake",
	"HXX":       "hexx",
	"PULSE":     "pulse",
	"GTC":       "global-tour-coin",
	"VTA":       "virtacoin",
	"TRADE":     "tradecoin-v2",
	"CESC":      "cryptoescudo",
	"OFF":       "cthulhu-offerings",
	"RIDE":      "ride-my-car",
	"CAB":       "cabbage",
	"CWXT":      "cryptoworldx-token",
	"COAL":      "bitcoal",
	"VPRC":      "vaperscoin",
	"NRO":       "neuro",
	"TAJ":       "tajcoin",
	"TAGR":      "tagrcoin",
	"WBC":       "wallet-builders-coin",
	"GBT":       "gamebet-coin",
	"DES":       "destiny",
	"BTQ":       "bitquark",
	"KNC":       "kingn-coin",
	"ZMC":       "zetamicron",
	"DLISK":     "dappster",
	"ORLY":      "orlycoin",
	"AMMO":      "ammo-rewards",
	"FRAZ":      "frazcoin",
	"IMX":       "impact",
	"BLZ":       "blazecoin",
	"PRX":       "printerium",
	"OS76":      "osmiumcoin",
	"LUNA":      "luna-coin",
	"QBK":       "qibuck-asset",
	"LTCR":      "litecred",
	"RSGP":      "rsgpcoin",
	"HVCO":      "high-voltage",
	"EGO":       "ego",
	"XBTS":      "beatcoin",
	"CTIC2":     "coimatic-2",
	"VLTC":      "vault-coin",
	"ZNE":       "zonecoin",
	"BNX":       "bnrtxcoin",
	"FUZZ":      "fuzzballs",
	"SOCC":      "socialcoin-socc",
	"XOC":       "xonecoin",
	"LIR":       "letitride",
	"CASH":      "cashcoin",
	"DOLLAR":    "dollar-online",
	"ACP":       "anarchistsprime",
	"CRT":       "crtcoin",
	"BSC":       "bowscoin",
	"IMPS":      "impulsecoin",
	"BIOB":      "biobar",
	"PEX":       "posex",
	"CCM100":    "ccminer",
	"WEX":       "wexcoin",
	"JOBS":      "jobscoin",
	"TOR":       "torcoin-tor",
	"DPAY":      "dpay",
	"SANDG":     "save-and-gain",
	"SLEVIN":    "slevin",
	"PX":        "px",
	"MGM":       "magnum",
	"IBANK":     "ibank",
	"BQC":       "bbqcoin",
	"SCS":       "speedcash",
	"ARGUS":     "argus",
	"SFC":       "solarflarecoin",
	"CONX":      "concoin",
	"SDP":       "sydpak",
	"DIX":       "dix-asset",
	"XRC":       "rawcoin2",
	"1CR":       "1credit",
	"OCEAN":     "burstocean",
	"REV":       "revenu",
	"SH":        "shilling",
	"VOLT":      "bitvolt",
	"GEERT":     "geertcoin",
	"U":         "ucoin",
	"NODC":      "nodecoin",
	"JIO":       "jio-token",
	"ENV":       "environ",
	"LEX":       "lex4all",
	"SLFI":      "selfiecoin",
	"P7C":       "p7coin",
	"MNC":       "mantracoin",
	"DRAGON":    "btcdragon",
	"XNG":       "enigma",
	"MUG":       "mikethemug",
	"PIZZA":     "pizzacoin",
	"PWR":       "powercoin",
	"CF":        "californium",
	"ELS":       "elysium",
	"FDC":       "future-digital-currency",
	"MI":        "xiaomicoin",
	"DGCS":      "digital-credits",
	"DMB":       "digital-money-bits",
	"CALC":      "caliphcoin",
	"XEN":       "xenixcoin",
	"FEDS":      "fedorashare",
	"INF":       "infchain",
	"ATB":       "atbcoin",
	"GBG":       "golos-gold",
	"DMC":       "dynamiccoin",
	"KEXCOIN":   "kexcoin",
	"FRGC":      "fargocoin",
	"MANA":      "decentraland",
	"PURA":      "pura",
	"ELIX":      "elixir",
	"ETHD":      "ethereum-dark",
	"VEN":       "vechain",
	"BGC":       "bagcoin",
	"CTR":       "centra",
	"ABC":       "alphabitcoinfund",
	"SJCX":      "storjcoin-x",
	"DFS":       "dfscoin",
	"BUZZ":      "buzzcoin",
	"XID":       "international-diamond",
	"SIGMA":     "sigmacoin",
	"TRX":       "tronix",
	"THS":       "techshares",
	"NDAO":      "neurodao",
	"DEUS":      "deuscoin",
	"BQ":        "bitqy",
	"WA":        "wa-space",
	"YOYOW":     "yoyow",
	"REGA":      "regacoin",
	"TERA":      "teracoin",
	"CLUB":      "clubcoin",
	"AURS":      "aureus",
	"BTDX":      "bitcloud",
	"MG":        "mind-gene",
	"WIC":       "wi-coin",
	"KAYI":      "kayi",
	"VOYA":      "voyacoin",
	"9COIN":     "9coin",
	"IND":       "indorse-token",
	"ATMC":      "atmcoin",
	"BLX":       "blockchain-index",
	"SUR":       "suretly",
	"ASC":       "asiccoin",
	"OX":        "ox-fina",
	"APC":       "alpacoin",
	"PAC":       "paccoin",
	"AMS":       "amsterdamcoin",
	"MRJA":      "ganjacoin",
	"TER":       "terranova",
	"UGT":       "ug-token",
	"IQT":       "iquant",
	"SCL":       "nexus-social",
	"LDCN":      "landcoin",
	"MRNG":      "morningstar-payments",
	"BTU":       "bitcoin-unlimited",
	"MTNC":      "masternodecoin",
	"NBIT":      "netbit",
	"XBG":       "btcgold",
	"SHND":      "stronghands",
	"ZBC":       "zilbercoin",
	"EDRC":      "edrcoin",
	"PCS":       "pabyosi-coin-special",
	"XOT":       "internet-of-things",
	"BITCF":     "first-bitcoin-capital",
	"PRES":      "president-trump",
	"CYDER":     "cyder",
	"PRN":       "protean",
	"BRAT":      "brat",
	"GARY":      "president-johnson",
	"TESLA":     "teslacoilcoin",
	"SKULL":     "pirate-blocks",
	"XQN":       "quotient",
	"TYC":       "tyrocoin",
	"MARX":      "marxcoin",
	"BLN":       "bolenum",
	"MAGN":      "magnetcoin",
	"ACC":       "adcoin",
	"NAMO":      "namocoin",
	"X2":        "x2",
	"BITOK":     "bitok",
	"UR":        "ur",
	"FC":        "facecoin",
	"ACN":       "avoncoin",
	"COUPE":     "coupecoin",
	"FBC":       "fibocoins",
	"RUBIT":     "rublebit",
	"QORA":      "qora",
	"TODAY":     "todaycoin",
	"GUC":       "goldunioncoin",
	"DTF":       "digitalfund",
	"EMV":       "ethereum-movie-venture",
	"UTA":       "utacoin",
	"SNAKE":     "snakeeyes",
	"ANTX":      "antimatter",
	"UNRC":      "universalroyalcoin",
	"ANI":       "animecoin",
	"ASN":       "aseancoin",
	"MOTO":      "motocoin",
	"BRAIN":     "braincoin",
	"RBBT":      "rabbitcoin",
	"HBC":       "hbcoin",
	"WOW":       "wowcoin",
	"GRN":       "granitecoin",
	"XTD":       "xtd-coin",
	"XLC":       "leviarcoin",
	"RHFC":      "rhfcoin",
	"ZSE":       "zsecoin",
	"AV":        "avatarcoin",
	"GAY":       "gaycoin",
	"FRWC":      "frankywillcoin",
	"SYNC":      "sync",
	"FID":       "bitfid",
	"FEI":       "fuda-energy",
	"XSTC":      "safe-trade-coin",
	"DASHS":     "dashs",
	"FAZZ":      "fazzcoin",
	"FFC":       "fireflycoin",
	"ELC":       "elacoin",
	"TCOIN":     "t-coin",
	"QBC":       "quebecoin",
	"EXL":       "excelcoin",
	"CYC":       "cycling-coin",
	"SFE":       "safecoin",
	"AXIOM":     "axiom",
	"PAYP":      "paypeer",
	"SKR":       "sakuracoin",
	"AMIS":      "amis",
	"CC":        "cybercoin",
	"OPAL":      "opal",
	"MONETA":    "moneta2",
	"FLASH":     "flash",
	"HALLO":     "halloween-coin",
	"HCC":       "happy-creator-coin",
	"UNC":       "uncoin",
	"DUB":       "dubstep",
	"SAK":       "sharkcoin",
	"VGC":       "vegascoin",
	"YES":       "yescoin",
	"GSR":       "geysercoin",
	"IRL":       "irishcoin",
	"PRM":       "prismchain",
	"LEPEN":     "lepen",
	"COXST":     "coexistcoin",
	"DON":       "donationcoin",
	"FUTC":      "futcoin",
	"BUB":       "bubble",
	"BEST":      "bestchain",
	"XDE2":      "xde-ii",
	"BIT":       "first-bitcoin",
	"EFYT":      "ergo",
	"WEC":       "wowecoin",
	"ETX":       "etherx",
	"OCOW":      "ocow",
	"PRIMU":     "primulon",
	"XRY":       "royalties",
	"LAZ":       "lazaruscoin",
	"BLAZR":     "blazercoin",
	"PDG":       "pinkdog",
	"MBL":       "mobilecash",
	"MEN":       "peoplecoin",
	"CHEAP":     "cheapcoin",
	"TELL":      "tellurion",
	"TOP":       "topcoin",
	"BIRDS":     "birds",
	"KASHH":     "kashhcoin",
	"FONZ":      "fonziecoin",
	"WINK":      "wink",
	"PI":        "picoin",
	"TOPAZ":     "topaz",
	"POKE":      "pokecoin",
	"PEC":       "peacecoin",
	"CLINT":     "clinton",
	"BXC":       "bitcedi",
	"XAU":       "xaucoin",
	"ACES":      "aces",
	"AIB":       "advanced-internet-blocks",
	"MRC":       "microcoin",
	"BTG":       "bitgem",
	"MONEY":     "moneycoin",
	"DBG":       "digital-bullion-gold",
	"MAVRO":     "mavro",
	"MMXVI":     "mmxvi",
	"CBD":       "cbd-crystals",
	"DISK":      "darklisk",
	"NXX":       "nexxus",
	"NBE":       "bitcentavo",
	"VTY":       "victoriouscoin",
	"SHA":       "shacoin",
	"LTH":       "lathaan",
	"GMX":       "goldmaxcoin",
	"ROYAL":     "royalcoin",
	"SKC":       "skeincoin",
	"TRICK":     "trickycoin",
	"RCN":       "rcoin",
	"KARMA":     "karmacoin",
	"XVE":       "the-vegan-initiative",
	"TCR":       "thecreed",
	"DCRE":      "deltacredits",
	"TURBO":     "turbocoin",
	"RYCN":      "royalcoin-2",
	"BAC":       "bitalphacoin",
	"ZENGOLD":   "zengold",
	"GAIN":      "ugain",
	"GOLF":      "golfcoin",
	"AXF":       "axfunds",
	"TLE":       "tattoocoin-limited",
	"OP":        "operand",
	"OPES":      "opescoin",
	"OMC":       "omicron",
	"IVZ":       "invisiblecoin",
	"PSY":       "psilocybin",
	"GML":       "gameleaguecoin",
	"GBRC":      "global-business-revolution",
	"STRB":      "superturbostake",
	"SPORT":     "sportscoin",
	"SHELL":     "shellcoin",
	"QBT":       "cubits",
	"GYC":       "gycoin",
	"SOUL":      "soulcoin",
	"LKC":       "linkedcoin",
	"STEX":      "stex",
	"WSX":       "wearesatoshi",
	"EGG":       "eggcoin",
	"NTCC":      "neptune-classic",
	"RICHX":     "richcoin",
	"PCN":       "peepcoin",
	"CME":       "cashme",
	"TEAM":      "teamup",
	"BGR":       "bongger",
	"MIU":       "miyucoin",
	"HYPER":     "hyper",
	"YOG":       "yogold",
	"HLB":       "lepaoquan",
	"QRZ":       "quartz-qrz",
}

func getCryptoCurrency(cryptocurrency string) (string, error) {
	if value, ok := cryptoCurrencies[strings.ToUpper(cryptocurrency)]; ok {
		return value, nil
	}

	return "", fmt.Errorf("%s is not a valid cryptocurrency", cryptocurrency)
}
