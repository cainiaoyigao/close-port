package cmdhandle

import (
	"close-port/entity"
	"fmt"
	"strconv"
)

type Handle struct {
	PortInUse  func(portNumber int) int
	GetAppName func(pid int) string
	KillApp    func(pid int) bool
	GetFuzzy   func(name string) []entity.AppInfo
}

func (h Handle) Handing() {
	fmt.Println("\n请输入查找的端口:")
	var reqPid string
	_, err := fmt.Scan(&reqPid)
	reqId, err := strconv.Atoi(reqPid)
	if err != nil {
		panic(err)
		return
	}
	pid := h.PortInUse(reqId)
	if pid == -1 {
		fmt.Printf("未找到端口为 %d 的程序 \n", pid)
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

func (h Handle) FuzzyHanding() {
	fmt.Println("\n请输入查找内容（端口或者软件名）:")
	var name string
	_, err := fmt.Scan(&name)

	appArray := make([]entity.AppInfo, 0)
	reqId, err := strconv.Atoi(name)
	if err == nil {
		pid := h.PortInUse(reqId)
		if pid != -1 {
			appInfos := h.GetFuzzy(strconv.Itoa(pid))
			if appInfos != nil {
				appArray = append(appArray, appInfos...)
			}
		}
	}

	appInfos := h.GetFuzzy(name)
	if appInfos != nil {
		appArray = append(appArray, appInfos...)
	}
	if len(appArray) == 0 {
		fmt.Printf("未找到 %s 的内容程序 \n", name)
		return
	}
	i := 0
	appMap := make(map[int]entity.AppInfo)
	for _, v := range appArray {
		i++
		appMap[i] = v
		fmt.Printf("请选择关闭的进程 %d : %s %d \n", i, v.Name, v.Pid)
	}
	var enterStr string
	_, err = fmt.Scan(&enterStr)
	if err != nil {
		fmt.Printf("你输入的 %v 选项有误", enterStr)
		return
	}
	req, err := strconv.Atoi(enterStr)
	if err != nil {
		panic(err)
		return
	}
	appInfo, ok := appMap[req]
	if !ok {
		fmt.Printf("你选择的 %v 选项程序不存在", req)
		return
	}
	ok = h.KillApp(appInfo.Pid)
	if ok {
		fmt.Printf("选项： %v 程序 %s 成功关闭", req, appInfo.Name)
	} else {
		fmt.Printf("选项： %v 程序 %s 成功关闭 \n", req, appInfo.Name)
	}
}
