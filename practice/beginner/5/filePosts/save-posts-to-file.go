package filePosts

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"
)

const (
	pathDir = "./Storage"
	pathPostsInDir = "./Storage/posts"
	writePerm = os.FileMode(0777)
)

func savePostToFile(i int, post Post)  {
	preSavePost()
	_, err := json.Marshal(post)  // Allow to save keys from structure Post.
	if err != nil {
		panic(err)
	}
	file := pathPostsInDir + "/" + strconv.Itoa(i) + ".txt"
	bufSave(file, post)
}

func bufSave(file string, p Post)  {
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Create an output stream via a buffer.
	writer := bufio.NewWriter(f)

	_, err = writer.Write([]byte(fmt.Sprintf("%+v\n", p)))  // Write struct to a file by saving the keys.*
	if err != nil {
		panic(err)
	}

	// Write post in file.
	writer.Flush()
}

func preSavePost()  {
	_, err := os.Stat(pathDir)  // Check if dir exist.
	if err != nil {
		os.MkdirAll(pathDir, writePerm)  // Create dir in path "./Storage".
	}

	_, err = os.Stat(pathPostsInDir)
	if err != nil {
		os.MkdirAll(pathPostsInDir, writePerm)  // Create dir in path "./Storage/posts".
	}
}

var postCounter int = 0

// Save 5 posts in pathPostsInDir. Remove mutex from here and main.go to save all posts.
func SavePostById(id int, m *sync.Mutex) {
	m.Lock()
	savePost := postCounter < 5
	if savePost {
		post := GetPostById(id)
		savePostToFile(postCounter+1, *post)
		postCounter++
	} else {
		os.Exit(1)
	}
	m.Unlock()
}
