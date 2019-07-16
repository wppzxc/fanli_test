package utils

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/go-vgo/robotgo"
	"github.com/lxn/win"
	"k8s.io/klog"
	"syscall"
)

func SendMessage(msg string, users []string) error {
	var sendErr error = nil
	for _, user := range users {
		if err := clipboard.WriteAll(msg); err != nil {
			klog.Errorf("Error on write to clipboard : %s", err)
			//fmt.Println("Error on write to clipboard : ", err)
		}
		p, err := syscall.UTF16PtrFromString(user)
		klog.V(9).Infof("get %s window process : %d", user, p)
		if err != nil {
			klog.Errorf("Error in get user chat window : %s", err)
			continue
		}
		h2 := win.FindWindow(nil, p)
		klog.V(9).Infof("get HWND of %s : %v", user, h2)
		var re = false
		for i := 0; i < 20; i++ {
			re = SetForegroundWindow(h2)
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
			//return nil
		} else {
			klog.Errorf("Error in set window foreground %s in 20 times ", user)
			sendErr = fmt.Errorf("Error in set window foreground %s in 20 times ", user)
			continue
		}
	}
	return sendErr
}
