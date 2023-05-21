package alias_client

import (
	"log"

	tls_client "github.com/bogdanfinn/tls-client"
)

func CreateClient(username string, password string, args ...string) *AliasSession {
	jar := tls_client.NewCookieJar()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),
		tls_client.WithClientProfile(tls_client.Safari_IOS_16_0),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar),
		tls_client.WithCharlesProxy("127.0.0.1", "8887"),
	}
	if args != nil {
		options = append(options, tls_client.WithProxyUrl(args[0]))
	}

	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &AliasSession{
		Client:   client,
		Username: username,
		Password: password,
	}
}
