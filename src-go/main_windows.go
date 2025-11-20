//go:build windows
// +build windows

package main

import (
	"os/exec"
	"syscall"
)

// setWindowsHideWindow 只在 Windows 平台设置隐藏窗口属性
func setWindowsHideWindow(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}
