package work

import "close-port/cmd"

type workHandle struct {
}

func NewWorkHandle() *workHandle {
	return &workHandle{}
}

func (wh *workHandle) LoopWork() {
	handler := cmdhandle.NewWinHandle()
	handler.Handing()
}
