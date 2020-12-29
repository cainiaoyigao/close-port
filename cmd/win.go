package cmdhandle

import (
	"bytes"
	"close-port/utils"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type winHandle struct {
	Handle
}

func NewWinHandle() *winHandle {
	win := &winHandle{}
	win.Handle.PortInUse = win.PortInUse
	win.Handle.GetAppName = win.GetName
	win.Handle.KillApp = win.KillApp
	return win
}

func (wh winHandle) PortInUse(portNumber int) int {
	res := -1
	var outBytes bytes.Buffer
	pth := fmt.Sprintf("0.0.0.0:%d ", portNumber)
	cmdStr := fmt.Sprintf("netstat -ano -p tcp | findstr %s", pth)
	cmd := exec.Command("cmd", "/c", cmdStr)
	cmd.Stdout = &outBytes
	cmd.Run()
	resStr := outBytes.String()
	r := regexp.MustCompile(`\s\d+\s`).FindAllString(resStr, -1)
	if len(r) > 0 {
		pid, err := strconv.Atoi(strings.TrimSpace(r[0]))
		if err == nil {
			res = pid
		}
	}
	return res
}

func (wh winHandle) GetName(pid int) string {
	//var outBytes bytes.Buffer
	//var stderr bytes.Buffer
	cmdStr := fmt.Sprintf("tasklist | findstr %d", pid)
	//fmt.Println(cmds)
	//cmd := exec.Command("tasklist",  cmds...)
	//cmd.Stdout = &outBytes
	//cmd.Stderr = &stderr
	//err := cmd.Run()

	cmd := exec.Command("cmd", "/c", cmdStr)
	out, err := cmd.CombinedOutput()
	stderrs := utils.ConvertByte2String(out, "GB18030")
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderrs)
		return "不存在"
	}
	resStr := out
	decodeBytes := utils.ConvertByte2String(resStr, "GB18030")
	fmt.Println(decodeBytes)

	resArray := regexp.MustCompile(`\s+`).Split(decodeBytes, -1)
	for i, v := range resArray {
		if strconv.Itoa(pid) == v {
			return resArray[i-1]
		}
	}
	return "不存在"
}

func (wh winHandle) KillApp(pid int) bool {
	var outBytes bytes.Buffer
	cmdStr := fmt.Sprintf("taskkill  /pid %d /F", pid)
	cmd := exec.Command("cmd", "/c", cmdStr)
	cmd.Stdout = &outBytes
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return true
}
