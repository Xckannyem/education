package data

import (
	"GoWorkspace/education/practice/trainee/3/db"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int 	  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Comment struct {
	PostId int    `json:"postId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

var name = "beginner_6"

func GetPosts() []Post {
	url := "https://jsonplaceholder.typicode.com/posts?userId=7"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Error getting posts request: ", err.Error())
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err.Error())
	}

	var postStruct []Post
	err = json.Unmarshal(bodyBytes, &postStruct)
	if err != nil {
		log.Fatal("Unable to unmarshal response body: ", err.Error())
	}

	return postStruct
}

func GetComments(p *Post) {
	url := "https://jsonplaceholder.typicode.com/comments?postId=" + strconv.Itoa(p.Id)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Error getting posts request: ", err.Error())
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err.Error())
	}

	var commentStruct []Comment
	err = json.Unmarshal(bodyBytes, &commentStruct)
	if err != nil {
		log.Fatal("Unable to unmarshal response body: ", err.Error())
	}

	savePosts(p)

	var wgComments sync.WaitGroup

	for i := 0; i < len(commentStruct); i++ {
		wgComments.Add(1)
		go func(comment *Comment) {
			saveComments(comment)
			defer wgComments.Done()
		}(&commentStruct[i])
	}
	wgComments.Wait()
}

func savePosts(post *Post) {
	database, err := db.InitDB(name)
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
	database.Create(post)
}

func saveComments(comment *Comment) {
	database, err := db.InitDB(name)
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
	database.Create(comment)
}
