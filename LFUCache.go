package Code

type LFUCache struct {
	MaxCap   int
	CurSize  int
	Nodes    map[int]*DNode
	NodeList map[int]*DNodeList
}

func Constructor1(capacity int) LFUCache {
	var lfu_cache LFUCache
	lfu_cache.MaxCap = capacity
	lfu_cache.Nodes = make(map[int]*DNode, 100)
	lfu_cache.NodeList = make(map[int]*DNodeList, 100)
	return lfu_cache
}

func (this *LFUCache) Get(key int) int {
	v, ok := this.Nodes[key]
	if !ok {
		return -1
	}

	this.RemoveFromNodeList(v.ref, v)

	v.ref++
	this.ADDToNodeList(v.ref, v)

	return v.Val
}

func (this *LFUCache) Put(key, vaue int) {
	v, ok := this.Nodes[key]
	if !ok {
		v = new(DNode)
		this.Nodes[key] = v

		if this.CurSize == this.MaxCap {
			this.CurSize--

			min_index := -1
			var node_list *DNodeList

			for k, list := range this.NodeList {
				if min_index == -1 || min_index > k {
					min_index = k
					node_list = list
				}
			}

			node := node_list.RemoveLast()
			delete(this.Nodes, node.Key)
			if node_list.Emprty() {
				delete(this.NodeList, min_index)
			}
		}
		this.CurSize++
	}

	v.Key = key
	v.Val = vaue

	this.RemoveFromNodeList(v.ref, v)
	v.ref++
	this.ADDToNodeList(v.ref, v)
}

func (this *LFUCache) RemoveFromNodeList(ref int, node *DNode) {
	if 0 == ref {
		return
	}

	node_list := this.NodeList[ref]
	node_list.RemoveNode(node)
	if node_list.Emprty() {
		delete(this.NodeList, ref)
	}
}

func (this *LFUCache) ADDToNodeList(ref int, node *DNode) {
	v, ok := this.NodeList[ref]
	if ok {
		v.AddNode(node)
	} else {
		v = new(DNodeList)
		this.NodeList[ref] = v
		v.Init()
		v.AddNode(node)
	}
}
