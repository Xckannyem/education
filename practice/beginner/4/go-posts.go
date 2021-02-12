package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type postId []struct{
	Id     int    `json:"id"`
}

var counter = 0

func getPostById(id int)  {
	url := "https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(id)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	
	// Read the response body.
	bodyBytes, _ := ioutil.ReadAll(res.Body)

	// Convert response body to Post struct.
	var postStruct Post
	json.Unmarshal(bodyBytes, &postStruct)

	if counter < 5 {
		fmt.Printf("%+v\n", postStruct)
	}
	counter ++
}

func main()  {
	url := "https://jsonplaceholder.typicode.com/posts/"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(res.Body)

	var postIdStruct postId
	json.Unmarshal(bodyBytes, &postIdStruct)

	for postStruct := range postIdStruct {
		go getPostById(postStruct)
	}
	var input string
	fmt.Scanln(&input)
}
