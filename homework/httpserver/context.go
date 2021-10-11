package httpserver

import (
	"encoding/json"
	"io"
	"net/http"
)

// Context 设计一个上下文结构体
type Context struct{
	W http.ResponseWriter // 因为是接口所以不适用指针接收
	Resp http.Response
	R *http.Request // 结构体所以使用指针接收
}

func (c *Context) ReadJson(data interface{}) error {
	// read body
	// Unmarshall json
	r := c.R
	body,err := io.ReadAll(r.Body)

	if err != nil{
		return err
	}

	err = json.Unmarshal(body,data)
	if err != nil{
		return err
	}

	return nil
}

func (c *Context) WriteJson(code int, resp interface{}) error {

	c.W.WriteHeader(code)

	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = c.W.Write(respJson)

	return err
}

func (c *Context) OKJson(resp interface{}) error {
	return c.WriteJson(http.StatusOK,resp)
}

func (c *Context) CreatedJson(resp interface{}) error {
	return c.WriteJson(http.StatusCreated,resp)
}

func (c *Context) SystemErrJson(resp interface{}) error {
	return c.WriteJson(http.StatusInternalServerError,resp)
}

func (c *Context) BadRequestJson(resp interface{}) error {
	return c.WriteJson(http.StatusBadRequest,resp)
}

func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		W:writer,
		R:request,
	}
}
