# alias_client

### Description
This is a go libary which helps you interact with the alias api. This libary is not using official api endpoints. This libary is using the TLS client from [@bogdanfinn](https://github.com/bogdanfinn)!

### Installation

```go
go get -u github.com/vadeex/alias_client
```

### Quick Example
```go
package main

import (
	"alias_client"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	// If you want to use a proxy, pass it as an argument to CreateClient:
	// alias_client.CreateClient("username", "password", "http://user:pass@host:port")
	session := alias_client.CreateClient("username", "password")

	// Login to alias
	loginResponse, err := session.Login()
	if err != nil {
		panic(err)
	}
	fmt.Println(loginResponse)
}
```
For a more detailed example please check [here](https://github.com/vadeex/alias_client/blob/main/example/main.go)

### Methods
Method | Description                                                                                                                                                        | Parameters                                   | Return
--- |--------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------|--------------
CreateClient | Creates a new session for browsing the API. Takes username string and a password string. Optional proxy parameter (format: "http://user:pass@host:port") | `username string`, `password string`, (`proxy string`) | `LoginResponse`, `error`
GetEarnings | Returns the balance of the account |  | `EarningsResponse`, `error`
GetSales | Returns a map with sales activity grouped by cashout(key) |  | `map[Cashout][]Item`, `error`
GetSale | Lookup a sale id. It will return the details of the sale like name and size | `saleid string` | `EarningsResponseSale`, `error`

### Help or Questions?
Add me on discord: vadex#6367
