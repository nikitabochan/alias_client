package alias_client

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	http "github.com/bogdanfinn/fhttp"
)

func (session *AliasSession) Login() (LoginResponse, error) {
	var loginResp LoginResponse
	req, err := http.NewRequest("POST", "https://sell-api.goat.com/api/v1/unstable/users/login", strings.NewReader(fmt.Sprintf(`{"grantType":"password","username":"%s","password":"%s"}`, session.Username, session.Password)))
	if err != nil {
		return loginResp, err
	}
	req.Header = http.Header{
		"content-type":      {"application/json"},
		"accept":            {"application/json"},
		"accept-encoding":   {"gzip, deflate, br"},
		"accept-language":   {"de-de"},
		"x-emb-st":          {fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond))},
		"user-agent":        {"alias/1.18.2 (iPhone; iOS 14.3; Scale/3.00) Locale/de"},
		"x-emb-id":          {"ABCA048E120349B4A1B0ABD8FCA7E617"},
		http.HeaderOrderKey: {"content-type", "x-px-authorization", "accept", "accept-encoding", "accept-language", "x-emb-st", "content-length", "user-agent", "x-emb-id", "cookie"},
	}
	resp, err := session.Client.Do(req)
	if err != nil {
		return loginResp, err
	}
	defer resp.Body.Close()
	_, siteErr := HandleResponse(resp, []int{200}) // Put all of the status codes you want to handle in the list.
	if siteErr != nil {
		panic(siteErr)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return loginResp, err
	}
	_ = json.Unmarshal(body, &loginResp)
	session.LoginResponse = &loginResp
	return loginResp, nil
}
