package image

type Image struct {
	Id     string `json:"id"`
	Url    string `json:"url"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

type ImageJsonb struct {
	Images []Image `json:"images"`
}
