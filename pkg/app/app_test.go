package app

import (
	"github.com/wpp/fanli_test/pkg/types"
	"testing"
)

func TestStartProcess(t *testing.T) {
	conf := types.Config{
		Uname: "bachinanfei@163.com",
		Password: "Aa123456",
		ToEmail: "bachinanfei@qq.com",
		ToWeChat: "哥哥",
	}
	AppRun(conf)
}