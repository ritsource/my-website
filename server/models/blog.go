package models

import (
	"fmt"
	"encoding/json"
)

// Blog - Blog model type
type Blog struct {
	Title string `json:"title"`
	Description string `json:"description"`
	HTML string `json:"html"`
	Markdown string `json:"markdown"`
	ImageURL string `json:"image_url"`
	IsPublic bool `json:"is_public"`
	IsDeleted bool `json:"is_deleted"`
}

// ToJSON - ...
func (b Blog) ToJSON() ([]byte, error) {
	var bData []byte
	var err error
	
	bData, err = json.Marshal(b)
	if err != nil {
		fmt.Println("Error: toJSON error:", err)
	}

	return bData, err
}

// Blogs - ...
type Blogs []Blog

// ToJSON - ...
func (bs Blogs) ToJSON() ([]byte, error) {
	var bData []byte
	var err error
	
	bData, err = json.Marshal(bs)
	if err != nil {
		fmt.Println("Error: toJSON error:", err)
	}

	return bData, err
}