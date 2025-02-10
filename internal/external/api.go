package external

import (
	"fmt"
	"net/http"
	"time"
)

const ExternalAPI string = "http://api.iplocation.net/?ip=%s"

// CallExternalAPI retrieves data from the external API for the given IP address.
// It returns the HTTP response and any error encountered.
func CallExternalAPI(ip string) (*http.Response, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	url := fmt.Sprintf(ExternalAPI, ip)
	resp, err := client.Get(url)
	return resp, err
}
