//go:build linux
// +build linux

package main

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"mydocker/container"
)

// Run 执行具体 command
/*
这里的Start方法是真正开始执行由NewParentProcess构建好的command的调用，它首先会clone出来一个namespace隔离的
进程，然后在子进程中，调用/proc/self/exe,也就是调用自己，发送init参数，调用我们写的init方法，
去初始化容器的一些资源。
*/
func Run(tty bool, cmd string) {
	parent := container.NewParentProcess(tty, cmd)
	log.Info("command开始执行")
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	log.Info("parent.Wait() begin-----------", os.Getpid())
	_ = parent.Wait()
	log.Info("parent.Wait() end-----------", os.Getpid())
	time.Sleep(time.Second * 1)
	os.Exit(-1)
}
