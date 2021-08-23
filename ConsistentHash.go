package main

import (
	"container/list"
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	hash       Hash // 计算hash的函数
	replicas   int  // 副本数
	keys       []int // 有序的列表，由大到小排序
	hashMap    map[int]string // 可以理解成用来记录虚拟节点和物理节点元数据关系的
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash    : fn,
		hashMap: make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}

	return m
}

func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas;i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			// 虚拟节点 <-> 实际节点
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	if m == nil {
		return ""
	}
	//根据用户输入的key值，计算出一个hash值
	hash := int(m.hash([]byte(key)))
	//查看值落在哪个
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
		
	})
	if idx == len(m.keys) {
		idx = 0
	}

	return m.hashMap[m.keys[idx]]
}

func main() {
	tmpList := list.New()

	for i := 0; i <= 10; i++ {
		tmpList.PushBack(i)
	}

	first := tmpList.PushFront(0)
	tmpList.Remove(first)

	for l := tmpList.Front(); l != nil; l = l.Next() {
		fmt.Print(l.Value, " ")
	}
}






