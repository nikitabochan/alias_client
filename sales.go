package alias_client

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	http "github.com/bogdanfinn/fhttp"
)

func (session *AliasSession) GetSales() (map[Cashout][]Item, error) {
	var sales []Item
	var data [][]Item
	salesData := make(map[Cashout][]Item)
	index := 1
	lastCashout := Cashout{}
	for {
		body, err := session.getSales(index)
		if err != nil {
			return nil, err
		}
		if !strings.Contains(string(body), `{"items":[`) {
			if len(sales) > 0 {
				data = append(data, sales)
				salesData[lastCashout] = sales
			}
			break
		}
		var salesResp SalesResponse
		_ = json.Unmarshal(body, &salesResp)
		for _, item := range salesResp.Items {
			if item.TransactionType == "CASHOUT" {
				cashoutResp, err := session.GetCashout(item.ID)
				if err != nil {
					return nil, err
				}
				amount, _ := strconv.ParseFloat(cashoutResp.AmountCents, 64)
				cashoutUSD := amount / 100.0

				amount, _ = strconv.ParseFloat(cashoutResp.CashOut.LocalizedTransferAmountCents.AmountCents, 64)
				currency := cashoutResp.CashOut.LocalizedTransferAmountCents.Currency
				cashoutCurrency := amount / 100.0

				amount, _ = strconv.ParseFloat(cashoutResp.CashOut.FeeRevenueCents, 64)
				feesFl := amount / 100.0

				cashout := Cashout{
					AmountUsd:       cashoutUSD,
					AmountLocalized: cashoutCurrency,
					Currency:        currency,
					Fee:             feesFl,
					Date:            cashoutResp.Date.String(),
					ID:              cashoutResp.ID,
				}
				if lastCashout.ID == "" {
					lastCashout = cashout
					continue
				} else {
					if len(sales) > 0 {
						data = append(data, sales)
						salesData[lastCashout] = sales
					}
					sales = []Item{}
					lastCashout = cashout
					continue
				}
			} else {
				if item.TransactionType == "SALE" {
					sales = append(sales, item)
				}
			}
		}
		index += 1
	}
	return salesData, nil
}

func (session *AliasSession) getSales(index int) ([]byte, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in getSales", r)
		}
	}()
	req, err := http.NewRequest("GET", fmt.Sprintf("https://sell-api.goat.com/api/v1/users/transactions?includeMetadata=1&page=%d", index), nil)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"accept":            {"application/json"},
		"x-emb-st":          {"1659263475685"},
		"x-emb-id":          {"EB2B3999EB544C3C83CB6C1828978EAE"},
		"accept-language":   {"de-de"},
		"user-agent":        {"alias/1.18.2 (iPhone; iOS 14.3; Scale/3.00) Locale/de"},
		"authorization":     {fmt.Sprintf("Bearer %s", session.LoginResponse.AuthToken.AccessToken)},
		"accept-encoding":   {"gzip, deflate, br"},
		http.HeaderOrderKey: {"accept", "x-emb-st", "cookie", "x-emb-id", "accept-language", "user-agent", "authorization", "accept-encoding"},
	}
	resp, err := session.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	_, siteErr := HandleResponse(resp, []int{200}) // Put all of the status codes you want to handle in the list.
	if siteErr != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
