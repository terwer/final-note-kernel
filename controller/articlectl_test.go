package controller

import (
	"fmt"
	"github.com/terwer/final-note-kernel/service"
	"github.com/terwer/final-note-kernel/starter"
	"testing"
)

func TestAddArticleAction(t *testing.T) {
	starter.SetupConf()

	// 主逻辑
	service.ConnectDB()

	var ret = AddArticleAction()

	// 及时关闭链接
	service.DisconnectDB()

	fmt.Sprintln(ret)
}
