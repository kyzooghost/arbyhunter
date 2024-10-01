package types

type UserResponse struct {
	RequestId string
	Code      int
	Message   string
	Data      interface{}
}
