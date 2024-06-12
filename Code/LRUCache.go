package Code

type DNode struct {
	Pre *DNode
	Suf *DNode
	Val int
	Key int
	ref int
}

type DNodeList struct {
	Head *DNode
	Tail *DNode
}

func (this *DNodeList) Init() {
	this.Head = new(DNode)
	this.Tail = new(DNode)
	this.Head.Suf = this.Tail
	this.Tail.Pre = this.Head
}

func (this *DNodeList) AddNode(node *DNode) {
	suf := this.Head.Suf
	this.Head.Suf = node
	node.Pre = this.Head
	node.Suf = suf
	suf.Pre = node
}

func (this *DNodeList) RemoveNode(node *DNode) {
	suf := node.Suf
	pre := node.Pre
	suf.Pre = pre
	pre.Suf = suf
}

func (this *DNodeList) MoveToHead(node *DNode) {
	this.RemoveNode(node)
	this.AddNode(node)
}

func (this *DNodeList) RemoveLast() *DNode {
	node := this.Tail.Pre
	this.RemoveNode(node)
	return node
}

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
