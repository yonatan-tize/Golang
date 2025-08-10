package main

import (
	"fmt"
	"net/url"
)

func main() {

	rawURL := "https://example.com:4000/path/home?query=param#fragmnet"
	url, err := url.Parse(rawURL)

	if err != nil {
		fmt.Println("Error while parsing url")
		return
	}

	fmt.Println(url.Scheme)
	fmt.Println(url.Host)
	fmt.Println(url.Port())
	fmt.Println(url.Path)
	fmt.Println(url.RawQuery)
	fmt.Println(url.Fragment)

}
