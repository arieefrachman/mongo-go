package entity

import "time"

// Bookmark model is
type Bookmark struct{
	ID ID `json:"id"`
	Name string `json:"name"`
	Description string `json:"desc"`
	CreatedAt time.Time `json:"created_at"`
}