package Code

type LRUCache struct {
	NodeList DNodeList
	Nodes    map[int]*DNode
	MaxCapa  int
	CurSize  int
}

func Constructor(capacity int) LRUCache {
	var lrc_cache LRUCache
	lrc_cache.NodeList.Init()
	lrc_cache.MaxCapa = capacity
	lrc_cache.Nodes = make(map[int]*DNode)
	return lrc_cache
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.Nodes[key]
	res := -1
	if ok {
		res = v.Val
		this.NodeList.MoveToHead(v)
	}
	return res
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.Nodes[key]
	if ok {
		v.Val = value
		this.NodeList.MoveToHead(v)
	} else {
		if this.CurSize == this.MaxCapa {
			delete(this.Nodes, this.NodeList.RemoveLast().Key)
			this.CurSize--
		}

		node := new(DNode)
		node.Val = value
		node.Key = key
		this.Nodes[key] = node
		this.NodeList.AddNode(node)
		this.CurSize++
	}
}
