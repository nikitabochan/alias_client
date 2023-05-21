package alias_client

import (
	"errors"
	"regexp"

	http "github.com/bogdanfinn/fhttp"
)

type Proxy struct {
	ID       string
	Host     string
	Username string
	Password string
}

func ExtractOrderNumber(input string) string {
	re := regexp.MustCompile(`Order #(\d+)`)
	matches := re.FindStringSubmatch(input)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}
func checkStatusCode(statusCode int, desiredStatusCode []int) bool {
	for _, eachStatusCode := range desiredStatusCode {
		if eachStatusCode == statusCode {
			return true
		}
	}
	return false
}
func HandleResponse(response *http.Response, desiredStatusCode []int) (bool, error) {
	statusCode := response.StatusCode
	if checkStatusCode(statusCode, desiredStatusCode) {
		return false, nil
	}
	switch statusCode {
	case 403:
		return true, errors.New("AccessDenied")
	case 429:
		return true, errors.New("RateLimited")
	case 500, 501, 502, 503, 504:
		return true, errors.New("ServerOverload")
	case 404:
		return true, errors.New("NotFound")
	case 400:
		return true, errors.New("BadRequest")
	case 401:
		return true, errors.New("Unauthorized")
	default:
		return true, errors.New("InvalidResponseCode")
	}
}
