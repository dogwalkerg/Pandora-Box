//go:build windows
// +build windows

package main

import (
	"os/exec"
	"syscall"
)

func setWindowsAttributes(cmd *exec.Cmd) {
	// 只在Windows平台设置HideWindow属性
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}
