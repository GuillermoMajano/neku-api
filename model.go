package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tags struct {
	Data struct {
		ID     primitive.ObjectID `bson:"_id"`
		URL    string             `json:"url" bson:"url"`
		Images `json:"images" bson:"images"`
		Titles []struct {
			Type  string `json:"type" bson:"type"`
			Title string `json:"title" bson:"title"`
		} `json:"titles" bson:"titles"`
		Chapters     int    `json:"chapters" bson:"chapters"`
		Volumes      int    `json:"volumes" bson:"volumes"`
		Status       string `json:"status" bson:"status"`
		Publishing   bool   `json:"publishing" bson:"publishing"`
		Published    `json:"published" bson:"published"`
		Synopsis     string    `json:"synopsis" bson:"synopsis"`
		Score        float32   `json:"score" bson:"score"`
		Author       []Mail_ID `json:"authors" bson:"author"`
		Genres       []Mail_ID `json:"genres" bson:"genres"`
		Themes       []Mail_ID `json:"themes" bson:"themes"`
		Demographics []Mail_ID `json:"demographics" bson:"demographics"`
	} `json:"data"`
}

type Images struct {
	JPG struct {
		image_format
	}

	Webp struct {
		image_format
	}
}

type image_format struct {
	Image_URL       string `json:"image_url" bson:"image_url"`
	Small_Image_URL string `json:"small_image_url" bson:"small_image_url"`
	Large_Image_URL string `json:"large_image_url" bson:"large_image_url"`
}

type Mail_ID struct {
	Mail_ID int    `json:"mail_id" bson:"mail_id"`
	Type    string `json:"type" bson:"type"`
	Name    string `json:"name" bson:"name"`
	URL     string `json:"url" bson:"url"`
}

type Published struct {
	From string `json:"from" bson:"from"`
	To   string `json:"to" bson:"to"`
	Pro  struct {
		From Date `json:"from" bson:"from"`
		To   Date `json:"to" bson:"to"`

		String string `json:"string" bson:"string"`
	} `json:"pro" bson:"pro"`
}

type Date struct {
	Day   int `json:"day" bson:"day"`
	Month int `json:"month" bson:"month"`
	Year  int `json:"year" bson:"year"`
}
