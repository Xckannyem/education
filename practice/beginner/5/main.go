package main

import (
	"GoWorkspace/education/practice/beginner/5/filePosts"
	"fmt"
	"sync"
)

func main()  {
	postsId := filePosts.GetId()
	m := new(sync.Mutex)
	for _, post := range *postsId {
		go filePosts.SavePostById(post.Id, m)
	}
	var input string
	fmt.Scanln(&input)
}
