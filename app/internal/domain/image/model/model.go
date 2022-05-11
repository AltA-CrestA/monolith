package model

type Image struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Size  string `json:"size"`
	Bytes []byte `json:"bytes"`
}
