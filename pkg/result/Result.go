package result

import (
	"time"
)

type Result struct {
	Sucess    bool        `json:"sucess"`
	Message   string      `json:"message"`
	Code      int         `json:"code"`
	Result    interface{} `json:"result"`
	Timestamp time.Time   `json:"timestamp"`
}

func (res *Result) OK() Result {
	res.Sucess = true
	res.Message = "操作成功"
	res.Code = 200
	res.Result = nil
	res.Timestamp = time.Now()
	return *res
}

func (res *Result) Ok(message string) Result {
	res.Sucess = true
	res.Message = message
	res.Code = 200
	res.Result = nil
	res.Timestamp = time.Now()
	return *res
}

func (res *Result) SUCESS(param interface{}) Result {
	res.Sucess = true
	res.Message = "操作成功"
	res.Code = 200
	res.Result = param
	res.Timestamp = time.Now()
	return *res
}

func (res *Result) ERROR() Result {
	res.Sucess = false
	res.Message = "操作失败"
	res.Code = 500
	res.Result = nil
	res.Timestamp = time.Now()
	return *res
}

func (res *Result) Error(message string) Result {
	res.Sucess = false
	res.Message = message
	res.Code = 500
	res.Result = nil
	res.Timestamp = time.Now()
	return *res
}

func (res *Result) Fail(param interface{}) Result {
	res.Sucess = false
	res.Message = "操作失败"
	res.Code = 500
	res.Result = param
	res.Timestamp = time.Now()
	return *res
}
