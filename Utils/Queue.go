package Utils

type Queue []interface{}

func (self *Queue) Enqueue(x interface{}) {
	*self = append(*self, x)
}

func (self *Queue) Dequeue() interface{} {
	h := *self
	var el interface{}
	l := len(h)
	el, *self = h[0], h[1:l]
	return el
}

func (self *Queue) GetFirst() interface{}{
	var el interface{}
	el=(*self)[0]
	return el
}

func (self *Queue) GetLast() interface{}{
	var el interface{}
	el=(*self)[len(*self)-1]
	return el
}


func (self *Queue) IsEmpty() bool {
	size:=len(*self)
	if size==0{
		return true
	}
	return false
}

func (self *Queue) Size() int {
	return len(*self)
}
