package app

import (
	"github.com/astaxie/beego/validation"
	"k8s.io/klog"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		klog.Info(err.Key, err.Message)
	}

	return
}
