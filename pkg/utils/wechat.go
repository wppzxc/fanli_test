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
	if err != nil {
		klog.Errorf("Error in get user chat window : %s", err)
		//fmt.Println("Error in get user chat window : ", err)
		return err
	}
	h2 := win.FindWindow(nil, p)
	re := win.SetForegroundWindow(h2)
	if re {
		robotgo.KeyTap("v", "ctrl")
		robotgo.KeyTap("enter")
		klog.Infof("Success to send msg to user : %s", user)
		//fmt.Println("Success to send msg to user : ", user)
		return nil
	} else {
		return fmt.Errorf("Error on get user : %s window ", user)
	}
}

func CheckWeChat() bool {
	_, err := robotgo.FindIds("WeChat")
	if err != nil {
		return false
	}
	klog.Info("Check WeChat OK !")
	//fmt.Println("Check WeChat OK !")
	return true
}
