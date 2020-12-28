package cmd

import (
	"bytes"
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
	return &winHandle{}
}

func (wh winHandle) PortInUse(portNumber int) int {
	res := -1
	var outBytes bytes.Buffer
	pth := fmt.Sprintf("0.0.0.0:%d", portNumber)
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
	var outBytes bytes.Buffer
	cmdStr := fmt.Sprintf("tasklist | findstr %d", pid)
	cmd := exec.Command("cmd", "/c", cmdStr)
	cmd.Stdout = &outBytes
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	resStr := outBytes.String()
	fmt.Println(resStr)
	resArray := regexp.MustCompile(`[^ ]+`).FindAllString(resStr, -1)
	return resArray[0]
}

func (wh winHandle) KillApp(pid int) string {
	var outBytes bytes.Buffer
	cmdStr := fmt.Sprintf("taskkill  /pid %d /F", pid)
	cmd := exec.Command("cmd", "/c", cmdStr)
	cmd.Stdout = &outBytes
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	resStr := outBytes.String()
	return resStr
}
