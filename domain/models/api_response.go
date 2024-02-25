package models

type ApiResponse[T any] struct {
	Message   string `json:"message"`
	Succeeded bool   `json:"succeeded"`
	Data      T      `json:"data"`
}
