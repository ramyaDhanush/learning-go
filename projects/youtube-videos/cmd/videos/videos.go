package main

import (
	"encoding/json"
	"io/ioutil"
)

type video struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Imageurl    string `json:"imageurl"`
	Url         string `json:"url"`
}

func getVideos() (videos []video) {
	// read file
	fileBytes, err := ioutil.ReadFile("./videos.json")
	if err != nil {
		panic(err)
	}
	
	// convert bytearray to slice
	err = json.Unmarshal(fileBytes, &videos)
	
	if err!=nil {
		panic(err)
	}
	
	return videos
}

func saveVideos(videos []video) {
	videoBytes, err := json.Marshal(videos)
	
	if err != nil {
		panic(err)
	}
	
	err = ioutil.WriteFile("./videos.json", videoBytes, 0644)
	if err != nil {
		panic(err)
	}
}