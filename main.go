package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"regexp"
)

func main() {
	defaultUrl := "http://localhost:8080"
	defaultCode := "200"

	url := flag.String("url", defaultUrl, "This should be an URL")
	code := flag.String("code", defaultCode, "This should be a regexp in \"\" or clear integer of expected return code, e.g. \"20\\d\"")

	flag.Parse()
	resp, err := http.Get(*url)
	if err != nil {
		log.Fatalln("[ CHECK FAILED ] ",err)
	}

	status, err := regexp.MatchString(*code, resp.Status)
	if status != true {
		log.Printf("[ CHECK FAILED ] %s", resp.Status)
		os.Exit(1)
	}
	log.Printf("[ CHECK OK ] %s", resp.Status)

}