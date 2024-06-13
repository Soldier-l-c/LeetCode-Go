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
	if this.Emprty() {
		return nil
	}
	node := this.Tail.Pre
	this.RemoveNode(node)
	return node
}

func (this *DNodeList) Last() *DNode {
	if this.Emprty() {
		return nil
	}
	return this.Tail.Pre
}

func (this *DNodeList) Emprty() bool {
	return this.Head.Suf == this.Tail
}
