package main

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */
func LRU( operators [][]int ,  k int ) []int {
	lruCache := NewLRUCache(k)
	res := make([]int, 0)
	for i := 0; i < len(operators);i++ {
		if operators[i][0] == 1 {
			//node := NewNode(operators[i][1],operators[i][2])
			lruCache.set(operators[i][1],operators[i][2])
		} else {
			res = append(res,lruCache.get(operators[i][1]))
		}
	}
	return res
	// write code here
}
//思路是采用哈希链表，先构造Node，再构造Map结构
type Node struct {
	key int
	value int
	pre * Node
	next *Node
}
// LRU结构 链表头代表最长不访问的节点， 表尾代表最长访问的节点
type LRUCache struct {
	head *Node
	end  *Node
	limit  int
	cache map[int]*Node
}

func NewNode(key,value int) *Node {
	return &Node{key : key, value: value}
}


func NewLRUCache(limit int) *LRUCache {
	return &LRUCache{limit : limit, cache: map[int]*Node{}}
}


//LRU结构的 get操作， 对于key存在的，返回value,同时将其值移向链表尾，key不存在返回-1
func (lru *LRUCache) get(key int) int{
	node, ok := lru.cache[key]
	if !ok {
		return -1
	}
	lru.refreshNode(node)
	return node.value

}

//LRU结构的Set操作，对于key存在的，直接更新其value，并移至表尾，key不存在的，若没有
//超过limit容量的限制，则直接建立新节点，并添加到表尾，否则，删除表头节点，并将新建节点添加
//至表尾
func (lru* LRUCache) set(key, value int) {
	node := lru.cache[key]
	if node == nil {
		if len(lru.cache) >= lru.limit {
			oldest := lru.removeNode(lru.head)
			delete(lru.cache, oldest)
		}
		node = NewNode(key,value)
		lru.addNode(node)
		lru.cache[key]= node
	} else {
		node.value = value
		lru.refreshNode(node)
	}

}


func (lru *LRUCache) refreshNode(node *Node) {
	if node == lru.end {
		return
	}
	lru.removeNode(node)
	lru.addNode(node)
}
//返回移出节点的key
func (lru *LRUCache) removeNode(node *Node) int{
	//移出尾节点
	if node == lru.end {
		lru.end = lru.end.pre
	} else if node == lru.head {
		//移除头结点
		lru.head = lru.head.next
	} else {
		//如果是中间节点
		node.pre.next = node.next
		node.next.pre = node.pre

	}
	return node.key

}
//在双向链表的尾部添加节点
func (lru *LRUCache) addNode(node *Node) {
	if lru.end != nil {
		lru.end.next = node
		node.pre = lru.end
		node.next = nil
	}
	lru.end = node
	//再判断一下头结点是否为空
	if lru.head == nil {
		lru.head = node
	}
}