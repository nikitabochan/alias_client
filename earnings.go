package alias_client

import (
	"encoding/json"
	"fmt"
	"io"

	http "github.com/bogdanfinn/fhttp"
)

func (session *AliasSession) GetEarnings() (*EarningsResponse, error) {
	var earningsResp EarningsResponse
	req, err := http.NewRequest("GET", "https://sell-api.goat.com/api/v1/users/earnings", nil)
	if err != nil {
		return &earningsResp, err
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
		return &earningsResp, err
	}
	defer resp.Body.Close()
	_, siteErr := HandleResponse(resp, []int{200}) // Put all of the status codes you want to handle in the list.
	if siteErr != nil {
		panic(siteErr)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &earningsResp, err
	}
	_ = json.Unmarshal(body, &earningsResp)
	return &earningsResp, nil
}
