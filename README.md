# grs-dev-tools

Small tools useful during Groestlcoin development.

##  base58grs

Converts address/xpub/WIF between Bitcoin and Groestlcoin formats. Examples:
```
$ base58grs Fj62rBJi8LvbmWu2jzkaUX1NFXLEqDLoZM
BTC: 1EvKQGaLZrF4Kusurtm721D3bN4HCSvVVV
GRS: Fj62rBJi8LvbmWu2jzkaUX1NFXLEqDLoZM

$ base58grs 1EvKQGaLZrF4Kusurtm721D3bN4HCSvVVV
BTC: 1EvKQGaLZrF4Kusurtm721D3bN4HCSvVVV
GRS: Fj62rBJi8LvbmWu2jzkaUX1NFXLEqDLoZM

$ base58grs 31inaRqambLsd9D7Ke4USZmGEVd3PHkh7P
BTC: 31inaRqambLsd9D7Ke4USZmGEVd3LVt8yd
GRS: 31inaRqambLsd9D7Ke4USZmGEVd3PHkh7P

$ base58grs KxqNwzatFKUTLRmb26kw1bi13TSFHvmyde4Y4YQxz1VRQYVof15z
BTC: KxqNwzatFKUTLRmb26kw1bi13TSFHvmyde4Y4YQxz1VRQYYyUtWV
GRS: KxqNwzatFKUTLRmb26kw1bi13TSFHvmyde4Y4YQxz1VRQYVof15z

$ base58grs xprv9zHDfLCJPTf5UrS16CrJ56WzSSoAYhJriX8Lfsco3TtPhG2DkwkVXjaDxZKU5USfmq5xjp1CZhpSrpHAPFwZWN75egb19TxWmMMmkd3csxP
BTC: xprv9zHDfLCJPTf5UrS16CrJ56WzSSoAYhJriX8Lfsco3TtPhG2DkwkVXjaDxZKU5USfmq5xjp1CZhpSrpHAPFwZWN75egb19TxWmMMmkcvNENM
GRS: xprv9zHDfLCJPTf5UrS16CrJ56WzSSoAYhJriX8Lfsco3TtPhG2DkwkVXjaDxZKU5USfmq5xjp1CZhpSrpHAPFwZWN75egb19TxWmMMmkd3csxP

$ base58grs xpub6DGa4qjCDqDNhLWUCEPJSETizUdexA2i5k3wUG2QboRNa4MNJV4k5XthorGcogStY5K5iJ6NHtsznNK599ir8PmA3d1jqEoZHsixDTddNA9
BTC: xpub6DGa4qjCDqDNhLWUCEPJSETizUdexA2i5k3wUG2QboRNa4MNJV4k5XthorGcogStY5K5iJ6NHtsznNK599ir8PmA3d1jqEoZHsixDUBw9Fz
GRS: xpub6DGa4qjCDqDNhLWUCEPJSETizUdexA2i5k3wUG2QboRNa4MNJV4k5XthorGcogStY5K5iJ6NHtsznNK599ir8PmA3d1jqEoZHsixDTddNA9
```
