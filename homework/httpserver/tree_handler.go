package httpserver

import (
	"net/http"
	"strings"
)

//var ErrorInvalidRouterPattern = error.New("invalid router pattern")

type HandlerBasedOnTree struct {
	root *node
}

type node struct {
	path     string
	children []*node

	// 如果这是叶子节点
	// 那么匹配上之后可以调用该方法
	handler handlerFunc
}

func (h *HandlerBasedOnTree) ServeHTTP(c *Context) {
	handler, found := h.findRouter(c.R.URL.Path)
	if !found {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("Note Found"))
		return
	}
	handler(c)
}

func (h *HandlerBasedOnTree) findRouter(path string) (handlerFunc, bool) {
	// trim掉前后的 /, 然后按照/切分
	paths := strings.Split(strings.Trim(path, "/"), "/")
	cur := h.root
	for _, p := range paths {
		// 从子节点里边找一个匹配到了当前path的节点
		matchChild, found := h.findMatchChild(cur, p)
		// 没有找到节点
		if !found {
			return nil, false
		}
		cur = matchChild
	}
	// 循环完毕说明找到了
	if cur.handler == nil {
		return nil, false
	}
	return cur.handler, true
}

func (h *HandlerBasedOnTree) Route(method string,
	pattern string,
	handlerFunc handlerFunc) {
	// trim掉前后的 '/'
	pattern = strings.Trim(pattern, "/")
	// 切割成切片
	paths := strings.Split(pattern, "/")

	cur := h.root // 当前节点
	// 遍历
	for index, path := range paths {
		matchChild, ok := h.findMatchChild(cur, path)
		if ok {
			cur = matchChild
		} else {
			h.createSubTree(cur, paths[index:], handlerFunc)
			return
		}
	}
	cur.handler = handlerFunc
}

// validatePattern 校验*，如果存在，必须在最后一个，且前面必须是/
//func (h *HandlerBasedOnTree) validatePattern(pattern string) error {
//	pos := strings.Index(pattern, "*")
//
//	// 找到了*
//	if pos>0{
//		if pos !=len(pattern)-1 {
//			return ErrorInvalidRouterPattern
//		}
//		if pattern[pos-1] != '/' {
//			return ErrorInvalidRouterPattern
//		}
//	}
//	return nil
//}

func (h *HandlerBasedOnTree) createSubTree(root *node,
	paths []string, handlerFn handlerFunc) {
	cur := root
	for _, path := range paths {
		nn := newNode(path)
		cur.children = append(cur.children, nn)
		cur = nn
	}
	cur.handler = handlerFn
}

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 2),
	}
}

// todo 接收使用node 更好
func (h *HandlerBasedOnTree) findMatchChild(root *node,
	path string) (*node, bool) {
	var wildcardNode *node

	// 遍历children
	for _, child := range root.children {
		// 注意这里要防止用户输入 *
		if child.path == path &&
			child.path != "*" {
			return child, true
		}
		// 命中了通配符的，扔要继续循环结束，防止看看后面还有没有更详细的
		if child.path == "*" {
			wildcardNode = child
		}
	}

	return wildcardNode, wildcardNode != nil
}
