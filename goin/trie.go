package goin

import "strings"

type node struct {
	// 到当前节点的路径，叶子节点才会有值
	pattern string
	// 当前节点的代表的部分
	part string
	// 子节点
	children []*node
	// 是否随意，如果是*或者:就是true，为了实现动态匹配
	isWild bool
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 支持插入和查询，对应路由的注册和匹配
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part}
		if part[0] == ':' || part[0] == '*' {
			child.isWild = true
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		if res := child.search(parts, height+1); res != nil {
			return res
		}
	}
	return nil
}
