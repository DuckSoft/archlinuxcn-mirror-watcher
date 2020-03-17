package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// MirrorProvider type.
type MirrorProvider struct {
	// provider name of the mirror source.
	// eg: 中国科学院开源软件协会
	Provider string `yaml:"provider" json:"provider"`

	// url of the mirror.
	// eg: http://mirrors.opencas.org/archlinuxcn/
	URL string `yaml:"url" json:"url"`
}

// Fetch the URL and returns the content, if successful.
func fetchURLContent(url string) (content []byte, e error) {
	// FIXME(DuckSoft): note that this function put in blind trust of content providers.
	//                  this would not be immune of things like "archive bomb".
	//                  take consideration before using this function.
	result, e := http.Get(url)
	if e != nil {
		return nil, e
	}

	content, e = ioutil.ReadAll(result.Body)
	if e != nil {
		return nil, e
	}

	return content, nil
}

// Decode & parse the content to get the mirror providers.
func getMirrorProviders(content []byte) (providers []MirrorProvider, e error) {
	e = yaml.Unmarshal(content, &providers)
	if e != nil {
		return nil, e
	}

	return providers, nil
}

// Fetch the mirror provider list from the configured URL.
func FetchMirrorProviderList() (providers []MirrorProvider, e error) {
	content, e := fetchURLContent(*mirrorListURL)
	if e != nil {
		return nil, e
	}

	providers, e = getMirrorProviders(content)
	if e != nil {
		return nil, e
	}

	return providers, nil
}

// Check the mirror last update time.
func getMirrorLastUpdated(mirrorURL string) (lastUpdated time.Time, e error) {
	response, e := http.Get(mirrorURL + "/lastupdate")
	if e != nil {
		return
	}

	body, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return
	}

	timestamp, e := strconv.Atoi(strings.TrimSpace(string(body)))
	if e != nil {
		return
	}

	return time.Unix(int64(timestamp), 0), nil
}
