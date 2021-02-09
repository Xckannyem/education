package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	url := "https://jsonplaceholder.typicode.com/posts"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	text := string(body)
	fmt.Println(text)
}
