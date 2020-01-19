package provider

import "username_across_platforms/server/client"

type checkInterface interface {
	CheckUrl(string, chan string)
}
type checker struct {}

var Checker checkInterface = &checker{}

func (check *checker) CheckUrl(url string, c chan string){
	resp, err := client.ClientCall.GetValue(url)
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