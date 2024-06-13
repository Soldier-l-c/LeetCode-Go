package Code

type LFUCache struct {
	MaxCap   int
	CurSize  int
	MinIndex int
	Nodes    map[int]*DNode
	NodeList map[int]*DNodeList
}

func Constructor1(capacity int) LFUCache {
	var lfu_cache LFUCache
	lfu_cache.MaxCap = capacity
	lfu_cache.Nodes = make(map[int]*DNode, 100)
	lfu_cache.NodeList = make(map[int]*DNodeList, 100)
	lfu_cache.MinIndex = 1
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
		this.Elimination()

		v = new(DNode)
		this.Nodes[key] = v
		this.MinIndex = 1
		this.CurSize++
	}

	v.Key = key
	v.Val = vaue

	this.RemoveFromNodeList(v.ref, v)
	v.ref++
	this.ADDToNodeList(v.ref, v)
}

func (this *LFUCache) RemoveFromNodeList(ref int, node *DNode) {
	if 0 == ref || node == nil {
		return
	}

	node_list := this.NodeList[ref]
	node_list.RemoveNode(node)
	if node_list.Emprty() {
		if ref == this.MinIndex {
			this.MinIndex = ref + 1
		}
		delete(this.NodeList, ref)
	}
}

func (this *LFUCache) ADDToNodeList(ref int, node *DNode) {
	v, ok := this.NodeList[ref]
	if !ok {
		v = new(DNodeList)
		this.NodeList[ref] = v
		v.Init()
	}
	v.AddNode(node)
}

func (this *LFUCache) Elimination() {
	if this.CurSize == this.MaxCap {
		this.CurSize--
		node := this.NodeList[this.MinIndex].Last()
		this.RemoveFromNodeList(this.MinIndex, node)
		delete(this.Nodes, node.Key)
	}
}
