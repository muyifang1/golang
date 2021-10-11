package httpserver

import "net/http"

type Routable interface {
	Route(method string, pattern string, handlerFunc handlerFunc)
}

// Handler 组合的方式给http.Handler扩展Route方法
type Handler interface {
	//http.Handler // Handler的实现可以直接调用 http.Handler中的方法
	ServeHTTP(c *Context)
	Routable
	// 注意！！！Golang中没有重写，子类无法重写父类方法，最好做个实验
}

// 这行的目的是Go自检测 确保肯定实现了 HandlerBasedOnMap接口
var _ Handler = &HandlerBasedOnMap{}

type HandlerBasedOnMap struct {
	// key = method + url
	handlers map[string]func(ctx *Context)
	// handlers sync.Map // todo 可以做成 线程安全的sync。map
}

// Route 注册路由
func (h *HandlerBasedOnMap) Route(method string,
	pattern string,
	handleFunc handlerFunc) {
	key := h.key(method, pattern)
	h.handlers[key] = handleFunc
}

func (h *HandlerBasedOnMap) ServeHTTP(c *Context) {

	key := h.key(c.R.Method, c.R.URL.Path)

	if handler, ok := h.handlers[key]; ok {
		handler(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("Not Found"))
	}
}

func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	return method + "#" + pattern
}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		handlers: make(map[string]func(ctx *Context), 5), // 预估5，超过自动扩容
	}
}
