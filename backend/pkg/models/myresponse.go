package models

type MyResponse struct {
	Message    interface{} `json:"message"`
	Success    interface{} `json:"success"`
	StatusCode interface{} `json:"status_code"`
	Data       interface{} `json:"result"`
}

type PaginationResponse struct {
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"result"`
	Offset     interface{} `json:"offset"`
	Page       interface{} `json:"page"`
	NextPage   interface{} `json:"next_page"`
	TotalPage  interface{} `json:"total_page"`
	AllItem    interface{} `json:"total_item"`
	NowShow    interface{} `json:"now_show_item"`
}

func SetMyResponse(msg string, code int, success bool, data interface{}) interface{} {
	return MyResponse{
		Message:    msg,
		StatusCode: code,
		Data:       data,
		Success:    success,
	}
}

func ResponseSuccess(msg string, data ...interface{}) interface{} {
	var setData interface{}
	if len(data) == 0 {
		setData = nil
	} else {
		setData = data[0]
	}
	return MyResponse{
		Message:    msg,
		Data:       setData,
		StatusCode: 200,
		Success:    true,
	}
}

func ResponseFail(msg string) interface{} {
	return MyResponse{
		Message:    msg,
		StatusCode: 400,
		Success:    false,
	}
}

func ResponseErr(err ...error) interface{} {
	var msg = "Internal server error"
	if len(err) != 0 {
		msg = err[0].Error()
	}
	return MyResponse{
		Message:    msg,
		StatusCode: 500,
		Success:    false,
	}
}
