package NewsScanner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

//googleApiResp defines the response struct for Google fact checker APIs
type googleApiResp struct {
	claims        []claim `json:"claim"`
	nextPageToken string  `json:"nextPageToken`
}

//claim is an entity under googleApiResp
type claim struct {
	text        string        `json:"text"`
	clamant     string        `json:"clamant"`
	claimDate   string        `json:"claimDate"`
	claimReview []claimReview `json:"claimReview`
}

//claimReview is an entity under googleApiResp.claim
type claimReview struct {
	publisher     publisher `json:"publisher"`
	url           string    `json:"url"`
	title         string    `json:"title"`
	reviewDate    string    `json:"reviewDate"`
	textualRating string    `json:"textualRating"`
	languageCode  string    `json:"languageCode"`
}

//publisher is an entity under claimReview, represents the fact publisher
type publisher struct {
	name string `json:"name"`
	site string `json:"site"`
}

//GoogleScanner implements INewsScanner using google APIs in the backend
type GoogleScanner struct {
	key        string
	baseApi    string
	maxAgeDays int
}

//NewGoogleScanner returns a new GoogleScanner
func NewGoogleScanner() *GoogleScanner {
	return &GoogleScanner{
		baseApi:    "https://factchecktools.googleapis.com/v1alpha1/claims:search",
		maxAgeDays: 10,
	}
}

//SetMaxAgeDays sets the maximum age(in days) for the returned search resutls
func (g *GoogleScanner) SetMaxAgeDays(maxAgeDays int) {
	g.maxAgeDays = maxAgeDays
}

//SetApiKey sets key to be used to access the APIs
func (g *GoogleScanner) SetApiKey(key string) {
	g.key = key
}

//IsFake tells whether the supplied news is fake
func (g *GoogleScanner) IsFake(news string) (isFake bool, err error) {
	resp, err := client.Get(fmt.Sprintf("%s?query=%s&languageCode=en-US&maxAgeDays=%d&key=%s", g.baseApi, news, g.maxAgeDays, g.key))
	if err != nil {
		return
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var respStruct googleApiResp

	err = json.Unmarshal(respBody, &respStruct)
	if err != nil {
		return
	}

	//fakeCount: claims suggesting the news to be false
	//totalCount: total facts/claims found for the news
	var fakeCount, totalCount int
	for _, c := range respStruct.claims {
		for _, cr := range c.claimReview {
			totalCount++
			if strings.Contains(strings.ToLower(cr.textualRating), "false") {
				fakeCount++
			}
		}
	}

	//if claims suggesting news to be fake >= 50% claims, news is fake
	if 2*fakeCount >= totalCount {
		return true, nil
	}
	return false, nil
}
