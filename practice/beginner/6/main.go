package main

import (
	"GoWorkspace/education/practice/beginner/6/data"
	"sync"
)

func main()  {
	posts := data.GetPosts()
	var wgPosts sync.WaitGroup
	for i := 0; i < len(posts); i++ {
		wgPosts.Add(1)
		go func(post *data.Post) {
			data.GetComments(post)
			defer wgPosts.Done()
		}(&posts[i])
	}
	wgPosts.Wait()
}