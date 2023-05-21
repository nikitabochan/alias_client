# alias_client
Alias client to interact with the alias api

### Description
This is a go libary which helps you interact with the alias api. This libary is not using official api endpoints.

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
For a more detailed example please check here

### Help or Questions?
Add me on discord: vadex#6367