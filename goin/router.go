package goin

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{roots: make(map[string]*node), handlers: make(map[string]HandlerFunc)}
}

func parsePattern(pattern string) []string {
	parts := strings.Split(pattern, "/")
	res := make([]string, 0)

	for _, part := range parts {
		if part != "" {
			res = append(res, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return res
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s\n", method, pattern)
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method, pattern string) (*node, map[string]string) {
	parts := parsePattern(pattern)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	found := root.search(parts, 0)
	if found == nil {
		return nil, nil
	}
	for index, part := range parsePattern(found.pattern) {
		if part[0] == ':' {
			params[part[1:]] = parts[index]
		}
		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(parts[index:], "/")
			break
		}
	}
	return found, params
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + c.Path
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
