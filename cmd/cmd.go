package cmdhandle

import (
	"fmt"
	"strconv"
)

type Handle struct {
	PortInUse  func(portNumber int) int
	GetAppName func(pid int) string
	KillApp    func(pid int) bool
}

func (h Handle) Handing() {
	fmt.Println("请输入查找的端口:")
	var reqPid string
	_, err := fmt.Scan(&reqPid)
	reqId, err := strconv.Atoi(reqPid)
	if err != nil {
		panic(err)
		return
	}
	pid := h.PortInUse(reqId)
	if pid == -1 {
		fmt.Printf("未找到端口为 %d 的出程序 \n", pid)
		return
	}
	name := h.GetAppName(pid)
	fmt.Printf("pid : %d , appName：%s \n", pid, name)
	fmt.Printf("是否关闭进程 %s（1 是 0 否）\n", name)
	var enterStr string
	_, err = fmt.Scan(&enterStr)
	if err != nil {
		panic(err)
		return
	}
	req, err := strconv.Atoi(enterStr)
	if err != nil {
		panic(err)
		return
	}
	if req == 1 {
		ok := h.KillApp(pid)
		if ok {
			fmt.Printf("pid： %v 程序成功关闭", pid)
		}
	} else {
		fmt.Printf("pid： %v 未关闭 \n", pid)
	}
}
