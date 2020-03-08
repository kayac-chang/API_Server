package model

type Response struct {
	Code  int         `json:"-"`
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}
