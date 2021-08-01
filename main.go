package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"regexp"
)

func main() {
	defaultUrl := "http://localhost"
	defaultCode := "200"

	url := flag.String("url", defaultUrl, "This should be an URL. Can contain a path like 'http://localhost:8080/x/app'.")
	code := flag.String("code", defaultCode, "This should be a regexp in \"\" or clear integer of expected return code, e.g. \"[2,3]0\\d\", \"200\"")

	flag.Parse()
	log.Printf("[ INFO ] Checking %s for code %s", *url, *code)
	resp, err := http.Get(*url)
	if err != nil {
		log.Fatalf("[ FATAL ] Check failed: %s", err)
	}

	status, err := regexp.MatchString(*code, resp.Status)
	if status != true {
		log.Printf("[ FATAL ] Check failed: %s %s", *url, resp.Status)
		os.Exit(1)
	}
	log.Printf("[ INFO ] Check passed: %s %s", *url, resp.Status)

}