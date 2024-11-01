package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id      int       `json:"ID"`
	Content string    `json:"Content"`
	Author  Author    `json:"Author"`
	Comment []Comment `json:"Comments`
}
type Author struct {
	Id   int    `json:"ID"`
	Name string `json:"Name"`
}
type Comment struct {
	Id      int    `json:"ID"`
	Content string `json:"Content"`
	Author  string `json:"Author"`
}

func main() {
	post := &Post{
		Id:      1,
		Content: "Marshalling struct data into json",
		Author: Author{
			Id:   1,
			Name: "Yousuf",
		},
		Comment: []Comment{
			Comment{
				Id:      2,
				Content: "Easy way!",
				Author:  "Mugera",
			},
			Comment{
				Id:      2,
				Content: "That's easy way!",
				Author:  "Luqman",
			},
		},
	}

	jsondata, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		fmt.Println("Error while marshalling data!")
		return
	}
	/*showingJsondata,*/ err = ioutil.WriteFile("post.json", jsondata, 0644)
	if err != nil {
		fmt.Println("Error while transfering json data into file with permission!", err)
		return
	}
	//fmt.Println(showingJsondata)
}
