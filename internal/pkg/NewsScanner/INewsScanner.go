package NewsScanner

import (
	"net/http"
	"time"
)

var (
	client *http.Client
)

//init initializes http client, using init() is a way of implementing singleton
func init() {
	client = &http.Client{
		Timeout: 30 * time.Second,
	}
}

//INewsScanner is an interface which defines a fake news scanner
type INewsScanner interface {
	IsFake(string) (bool, error)
	SetMaxAgeDays(int)
	SetApiKey(string)
}
