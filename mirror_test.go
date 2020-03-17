package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchURLContent(t *testing.T) {
	const urlTest = "https://example.org/"
	content, e := fetchURLContent(urlTest)

	assert.Nil(t, e, "fetch content shouldn't yield an error")
	assert.NotEmpty(t, content, "the content cannot be empty")
}

func TestGetMirrorProviders(t *testing.T) {
	providers, e := getMirrorProviders([]byte(yamlTestMirrorDecode))

	assert.Nil(t, e, "should decode without error")
	assert.Equal(t, providers, []MirrorProvider{
		{"Unique Studio", "https://mirrors.hustunique.com/archlinuxcn/"},
		{"浙江大学", "https://mirrors.zju.edu.cn/archlinuxcn/"},
	}, "should decode exactly the same")
}

func TestFetchMirrorProviderList(t *testing.T) {
	providerList, e := FetchMirrorProviderList()

	assert.Nil(t, e, "should fetch without error")
	assert.NotEmpty(t, providerList, "should not fetch an empty list")
	fmt.Printf("The fetched list is: \n%v", providerList)
}

func TestGetMirrorLastUpdated(t *testing.T) {
	const testURL = "https://mirrors.tuna.tsinghua.edu.cn/archlinuxcn/"
	updated, e := getMirrorLastUpdated(testURL)
	assert.Nil(t, e, "should fetch without error")
	fmt.Printf("Mirror %s:\n - Last Update: %v", testURL, updated)
}

const yamlTestMirrorDecode = `
- provider: Unique Studio
  url: https://mirrors.hustunique.com/archlinuxcn/
  location: 湖北武汉
  protocols:
  - ipv4
  added_date: 2017-08-02
  upstream: TUNA
- provider: 浙江大学
  url: https://mirrors.zju.edu.cn/archlinuxcn/
  location: 浙江杭州
  protocols:
  - ipv4
  - ipv6
  - http
  - https
  network: CERNET
  frequency: 6h
  added_date: 2017-06-05
`
