package cmd

import (
	"fmt"
	"strconv"
)

type Handle struct {
	PortInUse  func(portNumber int) int
	GetAppName func(pid int) string
	KillApp    func(pid int) string
}

func (h Handle) Handing() {
	fmt.Println("请输入查找的端口:")
	var reqPid string
	_, err := fmt.Scan(&reqPid)
	reqId, err := strconv.Atoi(reqPid)
	if err != nil {
		panic(err)
	}
	pid := h.PortInUse(reqId)
	if pid == -1 {
		fmt.Printf("未找到端口为 %d 的出程序 \n", pid)
		return
	}
	name := h.GetAppName(pid)
	fmt.Printf("pid %d : %s appName ：\n", pid, name)
	fmt.Printf("是否关闭进程 %s（1 是 0 否）\n", name)
	var enterStr string
	_, err = fmt.Scan(&enterStr)
	if err != nil {
		panic(err)
	}
	req, err := strconv.Atoi(enterStr)
	if err != nil {
		panic(err)
	}
	if req == 1 {
		killStr := h.KillApp(pid)
		fmt.Printf("输出： %s", killStr)
	}
}
