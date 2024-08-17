package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	// 创建一个传递 string 的 channel
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// 创建一个新的 goroutine
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(urlStr string, ch chan<- string) {
	start := time.Now()

	//proxyStr := os.Getenv("http_proxy")
	//if proxyStr == "" {
	//	fmt.Println("No proxy set")
	//	return
	//}
	//proxyUrl, err := url.Parse(proxyStr)
	//if err != nil {
	//	fmt.Println("Error parsing proxy URL")
	//	return
	//}
	//
	//client := &http.Client{
	//	Transport: &http.Transport{
	//		Proxy: http.ProxyURL(proxyUrl),
	//	},
	//}
	resp, err := http.Get(urlStr)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint("while reading %s: %v", urlStr, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, urlStr)
}
