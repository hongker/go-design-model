package composit

import (
	"fmt"
	"strings"
)

// Router 路由
type Router struct {
	uri string
	nodes map[string]*Router
}

func (router *Router) AddNode(node *Router) {
	router.nodes[node.uri] = node
}
func (router *Router) RemoveNode(node *Router) {
	delete(router.nodes, node.uri)
}
func (router *Router) GetNode(uri string) *Router{
	return router.nodes[uri]
}
func (router *Router) Handle(uri string){
	// 判断当前路由是否匹配
	if router.uri == uri {
		fmt.Printf("found %s,processing..\n", uri)
		return
	}

	// 判断该uri是否可能是子路由
	if strings.HasPrefix(uri, router.uri) {
		childUri := strings.TrimPrefix(uri, router.uri)
		// 对于路由分组，则进行递归遍历
		for _, child := range router.nodes {
			child.Handle(childUri)
		}
	}



}

func NewRouter(uri string) *Router {
	return &Router{
		uri:   uri,
		nodes: make(map[string]*Router, 0),
	}
}