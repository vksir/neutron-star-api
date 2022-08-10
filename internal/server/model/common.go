package model

type Ok struct{}

type Err struct {
	Detail string `json:"detail"`
}
