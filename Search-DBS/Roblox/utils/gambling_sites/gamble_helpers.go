package gambling_sites

import (
	"github.com/RomainMichau/cloudscraper_go/cloudscraper"
)

// getRequest performs a GET request using cloudscraper
func (gs *GambleScraper) getRequest(url string) (string, error) {
	client, err := cloudscraper.Init(false, false)
	if err != nil {
		return "", err
	}
	res, err := client.Get(url, make(map[string]string), "")
	if err != nil {
		return "", err
	}
	return res.Body, nil
}
