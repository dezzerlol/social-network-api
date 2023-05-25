package models

type Image struct {
	Id     int64  `json:"id"`
	Url    string `json:"url"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

type ImageJsonb struct {
	Images []Image `json:"images"`
}
