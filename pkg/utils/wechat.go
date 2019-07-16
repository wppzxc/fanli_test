package utils

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/go-vgo/robotgo"
	"github.com/lxn/win"
	"k8s.io/klog"
	"syscall"
)

func SendMessage(msg string, user string) error {
	if err := clipboard.WriteAll(msg); err != nil {
		klog.Errorf("Error on write to clipboard : %s", err)
		//fmt.Println("Error on write to clipboard : ", err)
	}
	p, err := syscall.UTF16PtrFromString(user)
	klog.V(9).Infof("get %s window process : %d", user, p)
	if err != nil {
		klog.Errorf("Error in get user chat window : %s", err)
		return err
	}
	h2 := win.FindWindow(nil, p)
	klog.V(9).Infof("get HWND of %s : %s", user, h2)
	var re = false
	for i := 0; i < 20; i++ {
		re = win.SetForegroundWindow(h2)
		if re {
			klog.V(9).Infof("set process foregroudwindow %d times ok", i)
			break
		}
		klog.Warningf("set process foregroudwindow %d times failed", i)
	}
	if re {
		robotgo.KeyTap("v", "ctrl")
		robotgo.KeyTap("enter")
		klog.Infof("Success to send msg to user : %s", user)
		return nil
	} else {
		return fmt.Errorf("Error in set window foreground %d in 10 times ", h2)
	}
}

func CheckProcess(processName string) bool {
	_, err := robotgo.FindIds(processName)
	if err != nil {
		return false
	}
	klog.Infof("Check %s OK !", processName)
	return true
}
