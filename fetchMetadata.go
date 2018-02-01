package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var metaUrl string = "http://s3.amazonaws.com/ec2metadata/ec2-metadata"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	resp, err := http.Get(metaUrl)
	if err != nil {
		fmt.Fprint(os.Stderr, "fetch: %v\n", err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", metaUrl, err)
		os.Exit(1)
	}
	err = ioutil.WriteFile("/tmp/ec2-metadata", b, 0755)
	check(err)

}
