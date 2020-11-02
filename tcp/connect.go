package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

//https://gocn.vip/topics/10626
func main() {
	count := 10
	for i := 0; i < count; i++ {
		resp, err := http.Get("https://www.oschina.net")
		if err != nil {
			panic(err)
		}

		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}
}
