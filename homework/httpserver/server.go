package main

import (
	"fmt"
	"net/http"
	"os"
)

type Server interface {
	Routable
	Start(address string) error
}

var _ Server = &sdkHttpServer{}

// sdkHttpServer 基于net/http
type sdkHttpServer struct {
	Name    string
	handler Handler
	root    Filter
}

// Route 注册路由
func (s *sdkHttpServer) Route(method string,
	pattern string,
	handleFunc handlerFunc) {
	//key := s.handler.key(method, pattern)
	//s.handler.handlers[key] = handleFunc
	s.handler.Route(method, pattern, handleFunc)
}

// Start 启动服务
func (s *sdkHttpServer) Start(address string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter,
		request *http.Request) {
		c := NewContext(writer, request)
		s.root(c)
	})
	return http.ListenAndServe(address, nil)
}

func NewSdkHttpServer(name string, builders ...FilterBuilder) Server {
	handler := NewHandlerBasedOnMap()

	var root Filter = handler.ServeHTTP

	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}

	return &sdkHttpServer{
		Name:    name,
		handler: handler,
		root:    root,
	}
}

//func home(w http.ResponseWriter, r *http.Request){
//	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
//}

// SignUp 注册路由
func SignUp(ctx *Context) {
	req := &signUpReq{}

	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}

	ctx.TransferHeader()

	// 返回一个 resp 作为json
	resp := &commonResponse{
		//BizCode: sLen,
		//Data: "this is common response data",
		Data: req.Email + "#" + req.Password + "#" + req.ConfirmedPassword,
	}

	err = ctx.CreatedJson(resp)
	//err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("写入响应失败 %v", err)
	}
}

// signUpReq 模拟一个注册信息结构体
type signUpReq struct {
	Email             string `json:"email"` // Tag用法，可以在运行期通过反射拿到该信息
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

// commonResponse 定义一个返回信息结构体
type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func healthz(ctx *Context) {
	resp := &commonResponse{
		Data: "this is test Server healthz!",
	}

	ctx.TransferHeader()

	err := ctx.OKJson(resp)

	if err != nil {
		fmt.Printf("写入响应失败 %v", err)
	}
}

func (ctx *Context) TransferHeader() {
	for headerKey, values := range ctx.R.Header {
		ctx.W.Header().Add(headerKey, values[0])
	}
	verFromOsEnv := os.Getenv("VERSION")
	fmt.Printf("从环境变量中取得 VERSION的值 ： %s ,并写入Header中", verFromOsEnv)
	fmt.Println()
	ctx.W.Header().Add("VERSION", verFromOsEnv)
}
