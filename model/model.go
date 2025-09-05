package model

type AddRequest struct {
	Number1 int64 `json:"number1,omitempty"`
	Number2 int64 `json:"number2,omitempty"`
}

type SubstractRequest struct {
	Number1 int64 `json:"number1,omitempty"`
	Number2 int64 `json:"number2,omitempty"`
}
