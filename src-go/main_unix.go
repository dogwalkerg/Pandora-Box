//go:build linux || darwin
// +build linux darwin

package main

import "github.com/metacubex/mihomo/log"

func runPowerShellScript() {
	// 在 Linux 和 macOS 上只记录日志，不执行任何操作
	log.Debugln("Skipping PowerShell script on non-Windows platform")
}
