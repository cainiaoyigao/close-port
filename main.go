package main

import "close-port/cmd"

func main() {
	handler := cmdhandle.NewWinHandle()
	handler.Handing()
}
