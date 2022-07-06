package battlescribe

import (
	"encoding/xml"
	"errors"
	"log"
	"net/http"
)

type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Title   string   `xml:"title"`
	Link    []struct {
		Text string `xml:",chardata"`
		Rel  string `xml:"rel,attr"`
		Type string `xml:"type,attr"`
		Href string `xml:"href,attr"`
	} `xml:"link"`
	Author struct {
		Text string `xml:",chardata"`
		Name string `xml:"name"`
		URI  string `xml:"uri"`
	} `xml:"author"`
	Subtitle struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
	} `xml:"subtitle"`
	ID      string `xml:"id"`
	Updated string `xml:"updated"`
	Entry   []struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:",chardata"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
			Href string `xml:"href,attr"`
		} `xml:"link"`
		ID        string `xml:"id"`
		Updated   string `xml:"updated"`
		Published string `xml:"published"`
		Summary   struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"summary"`
	} `xml:"entry"`
}

const (
	atomAllURL = "https://battlescribedata.appspot.com/repos/feeds/all.atom"
)

var (
	ErrFeedCouldNotGetFeed    = errors.New("could not get atom feed")
	ErrFeedCouldNotDecodeFeed = errors.New("could not decode atom feed")
)

// GetFeedData gets the data in the BSData Atom feed
func GetFeedData() (*Feed, error) {
	var feed Feed
	res, err := http.Get(atomAllURL)
	if err != nil {
		log.Println("could not get atom feed:", err.Error())
		return nil, ErrFeedCouldNotGetFeed
	}

	defer res.Body.Close()

	dec := xml.NewDecoder(res.Body)
	err = dec.Decode(&feed)
	if err != nil {
		log.Println("could not decode atom feed:", err.Error())
		return nil, ErrFeedCouldNotDecodeFeed
	}

	return &feed, nil
}
