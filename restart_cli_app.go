package main

import (
	"os/exec"
	"syscall"
)

var cmdGlo *exec.Cmd

func restartCliApp(cmdPri *exec.Cmd, convCfgStr string) *exec.Cmd {

	stopCliApp(cmdPri)
	cmdPri, _ = startCliApp(convCfgStr)

	return cmdPri
}

func stopCliApp(cmdPri *exec.Cmd) {
	
	err := cmdPri.Process.Kill()
	ChkErr(err)
}

func startCliApp(convCfgStr string) (*exec.Cmd, bool) {

	convCfg := JsonStrToStru(convCfgStr)
	
	cmdPri := exec.Command(
		"bin/poetry_cli.exe",
		
		"--interval",
		convCfg.Interval,
		"--font",
		convCfg.Font,
		"--fontsize",
		convCfg.Fontsize,
		"--imgloc",
		convCfg.Imgloc,
		"--color",
		convCfg.Color,
		"--xoffset",
		convCfg.XOffset,
		"--yoffset",
		convCfg.YOffset,
	)

	// Hide window while spawning
	cmdPri.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	
	go func() {
		err := cmdPri.Start()
		ChkErr(err)
		
		cmdPri.Wait()
	}()

	return cmdPri, convCfg.IfDae
}


func StopCliAppWithoutArgv() {

	err := cmdGlo.Process.Kill()
	ChkErr(err)
}
