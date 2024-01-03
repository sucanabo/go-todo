package common

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging"`
	Filter interface{} `json:"filter"`
}

func NewSuccessResponse(data, paging, filter interface{}) *successRes {
	return &successRes{data, paging, filter}
}

func NewSimpleResponse(data interface{}) *successRes {
	return NewSuccessResponse(data, nil, nil)
}
