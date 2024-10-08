package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := io.ReadAll(resp.Body)
		status, code := resp.Status, resp.StatusCode
		fmt.Println(status, " ", code)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reeading %s: %v\n", url, err)
			os.Exit(1)
		}
		//fmt.Println("%s", b)
	}
}
