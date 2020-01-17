package service

import (
	"net/http"
)

func fetch(url string, c chan string){
	resp, err := http.Get(url)
	//we could not access that endpoint
	if err != nil {
		c <- "cant_access_resource"
		return
	}
	//the response code may be 404, 422, etc
	if resp.StatusCode > 299 {
		c <- "no_match"
	}
	if resp.StatusCode == 200 {
		c <- url
	}
}

func UsernameCheck(urls []string) []string {
	c := make(chan string)
	var links []string
	var matchingLinks []string

	for _, url := range urls {
		go fetch(url, c)
	}
	for i := 0; i < len(urls); i++ {
		links = append(links, <-c)
	}
	//Remove the "no_match" and "cant_access_resource" values from the links array:
	for _, v := range links {
		if v == "cant_access_resource" {
			continue
		}
		if v == "no_match" {
			continue
		}
		matchingLinks = append(matchingLinks, v)
	}
	return matchingLinks
}