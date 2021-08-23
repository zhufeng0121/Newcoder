package main

import "fmt"

type Node struct {
	pre  *Node
	next *Node
	key   string
	value string
}

func NewNode(key string,value string) *Node {
	return &Node{
		key: key,
		value : value,
	}
}

type LRUCache struct {
	head     *Node
	end      *Node
	limit    int
	cacheMap map[string]*Node

}

func NewLRUCache(limit int) *LRUCache {
	return &LRUCache{
		limit: limit,
		cacheMap: map[string]*Node{},
	}
}

func (lru *LRUCache) get(key string) string {
	node := lru.cacheMap[key]
	if node == nil {
		return ""
	}
	lru.refreshNode(node)
	return node.value
}

func (lru *LRUCache) put(key, value string) {
	node := lru.cacheMap[key]
	if node == nil {
		//若key不存在，插入key-value
		if len(lru.cacheMap) >= lru.limit {
			oldest := lru.removeNode(lru.head)
			delete(lru.cacheMap,oldest)
		}
		node = NewNode(key,value)
		lru.addNode(node)
		lru.cacheMap[key] = node

	} else {
		node.value = value
		lru.refreshNode(node)
	}
}

//刷新被访问节点的位置
func (lru *LRUCache) refreshNode(node *Node) {
	//尾节点不需要移动
	if node == lru.end {
		return
	}
	lru.removeNode(node)
	lru.addNode(node)
}
//移出节点
func (lru *LRUCache) removeNode(node *Node) string{
	if node == lru.end {
		//移除尾结点
		lru.end = lru.end.pre
	} else if node == lru.head {
		//移除头结点
		lru.head = lru.head.next
	} else {
		//移除中间节点
		node.pre.next = node.next
		node.next.pre = node.pre

	}
	return node.key

}
//尾部插入节点
func (lru *LRUCache) addNode(node *Node) {
	if lru.end != nil {
		lru.end.next = node
		node.pre = lru.end
		node.next = nil
	}
	lru.end = node
	if lru.head == nil {
		lru.head = node
	}
}

func main() {
	l := NewLRUCache(5)
	l.put("001","用户1信息")
	l.put("002","用户1信息")
	l.put("003","用户1信息")
	l.put("004","用户1信息")
	l.put("005","用户1信息")
	l.get("002")
	l.put("004","用户2信息更新")
	l.put("006","用户6信息")
	fmt.Println(l.get("001"))
	fmt.Println(l.get("006"))
}

