package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
)

func hammer() {
	client := &http.Client{}

	//resp, err := client.Get("http://onramp.400lbs.com/")
	req, err := http.NewRequest("GET", "http://onramp.400lbs.com/", nil)
	req.Header.Add("x-ots", `1492719972123`)
	resp, err := client.Do(req)


	//resp, err := http.Get("http://onramp.400lbs.com/")

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	xrotsheader := resp.Header.Get("x-rots")

	fmt.Printf("%s", xrotsheader)
	fmt.Printf("%s", err)
	fmt.Printf("%s", body)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
	fmt.Printf("go")
	go hammer()
	fmt.Printf("done")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
