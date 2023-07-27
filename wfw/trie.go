package wfw

import "strings"

type node struct {
	pattern  string  // 非空表示成功匹配(非路径节点)
	part     string  // 路由中的一部分
	children []*node // 子节点
	isWild   bool    // 是否是通配符[*:]
}

// 返回第一个匹配的子节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 返回所有匹配的子节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 递归插入pattern
func (n *node) insert(u int, parts []string, pattern string) {
	if u == len(parts) {
		n.pattern = pattern
	} else {
		part := parts[u]
		child := n.matchChild(part)
		if child == nil {
			child = &node{
				part:   part,
				isWild: part[0] == ':' || part[0] == '*',
			}
			n.children = append(n.children, child)
		}
		child.insert(u+1, parts, pattern)
	}
}

func (n *node) search(u int, parts []string) *node {
	if u == len(parts) || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[u]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(u+1, parts)
		if result != nil {
			return result
		}
	}
	return nil
}
