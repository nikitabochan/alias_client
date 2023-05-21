package main

import (
	"alias_client"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	//if you want to use a proxy, pass it as an argument to CreateClient:
	// alias_client.CreateClient("username", "password", "http://user:pass@host:port")
	session := alias_client.CreateClient("username", "password")

	// Login to alias
	loginResponse, err := session.Login()
	if err != nil {
		panic(err)
	}
	fmt.Println(loginResponse)
	fmt.Println("==============")
	// Get Earnings (Balance of the account)
	earningsResponse, err := session.GetEarnings()
	fmt.Println(earningsResponse)
	fmt.Println("==============")
	// Get SaleIDs grouped by cashout ids. those are needed to get the actual sale data at a later point
	salesData, err := session.GetSales()
	for key, value := range salesData {
		fmt.Println("Cashout:", key)
		saleIDs := []string{}
		for _, sale := range value {
			saleIDs = append(saleIDs, sale.ID)
		}
		fmt.Println("SaleIDS:", strings.Join(saleIDs, ","))
	}

	// Get sale data from a specific sale id
	// For example we want to get the sale data of the first cashout([0]).

	// get the first key in map
	cashout := reflect.ValueOf(salesData).MapKeys()[0].Interface().(alias_client.Cashout)
	fmt.Println("==============")
	fmt.Println("Cashout:", cashout)
	for _, sale := range salesData[cashout] {
		saleID := alias_client.ExtractOrderNumber(sale.Type)
		saleResponse, err := session.GetSale(saleID)
		if err != nil {
			panic(err)
		}
		fmt.Println("Item", saleResponse.PurchaseOrder.Listing.Product.Name, saleResponse.PurchaseOrder.Listing.Size, saleResponse.PurchaseOrder.LocalizedAmountMadeCents.AmountCents)
	}

}
