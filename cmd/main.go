package main

import (
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("POST", "http://www.baidu.com?a=2&b=2", nil)
	if nil != err {
		log.Fatal(err)
	}

	log.Println(req.URL.Path)
	log.Println(req.URL.Path)
	log.Println(req.URL.RawPath)
	log.Println(req.URL.RawQuery)
	log.Println(req.URL.Query())

	v := req.URL.Query()
	v.Set("f", "3")

	log.Println(v.Encode())
	req.URL.RawQuery = v.Encode()
	log.Println(req.URL)

	client := http.Client{}
	reply, err := client.Do(req)
	log.Println(reply)
}
