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
		return fmt.Errorf("Error on write to clipboard : %s ", err)
	}
	p, err := syscall.UTF16PtrFromString(user)
	klog.V(9).Infof("get %s window process : %d", user, p)
	if err != nil {
		return fmt.Errorf("Error in get user chat window : %s ", err)
	}
	h2 := win.FindWindow(nil, p)
	klog.V(9).Infof("get HWND of %s : %v", user, h2)
	var ok = false
	for i := 0; i < 20; i++ {
		ok = SetForegroundWindow(h2)
		if ok {
			klog.V(9).Infof("set process foregroudwindow %d times ok", i)
			break
		}
		klog.Warningf("set process foregroudwindow %d times failed", i)
	}
	if ok {
		robotgo.KeyTap("v", "ctrl")
		robotgo.KeyTap("enter")
		klog.Infof("Success to send msg to user : %s", user)
		//return nil
	} else {
		return fmt.Errorf("Error in set window foreground %s in 20 times ", user)
	}
	return nil
}
