package app

import (
	"github.com/wpp/fanli_test/pkg/signals"
	"github.com/wpp/fanli_test/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog"
	"time"
)

func Start(conf *types.Config, fs []func()) {
	defer klog.Flush()
	stopCh := signals.SetupSignalHandler()

	klog.Info("Starting fanli ... ")
	for _, f := range fs {
		go wait.Until(f, time.Duration(conf.Fanli.RefreshInterval)*time.Second, stopCh)
	}

	klog.Info("Started fanli")
	<-stopCh
	klog.Info("Shutting down fanli")
}
