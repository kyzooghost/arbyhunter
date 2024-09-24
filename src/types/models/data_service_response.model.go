package types

type DataServiceResponse struct {
	RequestId string
	Code      int
	Message   string
	Data      interface{}
}
