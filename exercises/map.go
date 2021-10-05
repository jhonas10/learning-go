package main

import "fmt"

func main() {
	website := map[string]map[string]string {
		"Google": map[string]string {
			"name":"Google",
			"type":"Search",
		},
		"YouTube": map[string]string {
			"name":"YouTube",
			"type":"video",
		},
	}
	
	fmt.Println(website["Google"]["name"])
    fmt.Println(website["Google"]["type"])

}
