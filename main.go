package main

import (
	"close-port/work"
)

func main() {
	workHandler := work.NewWorkHandle()
	for {
		workHandler.LoopWork()
	}
}
