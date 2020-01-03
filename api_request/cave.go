package api_request

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

const XML_CAVE_URL = "http://rss.cavelis.net/index_live.xml"

type caveLiveData struct {
	Entries []Entry `xml:"entry"`
}

type Entry struct {
	Title  string `xml:"title"`
	Author author `xml:"author"`
}

type author struct {
	Name string `xml:"name"`
}

func GetCaveLiveData() []Entry {
	resp, _ := http.Get(XML_CAVE_URL)
	xmlData, _ := ioutil.ReadAll(resp.Body)
	var caveLiveData caveLiveData
	xml.Unmarshal(xmlData, &caveLiveData)
	return caveLiveData.Entries
}
