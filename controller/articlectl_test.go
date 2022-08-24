package controller

import (
	"fmt"
	kernel "github.com/terwer/final-note-kernel"

	"github.com/terwer/final-note-kernel/service"
	"testing"
)

func TestAddArticleAction(t *testing.T) {
	kernel.Init()

	// 主逻辑
	service.ConnectDB()

	var ret = AddArticleAction()

	// 及时关闭链接
	service.DisconnectDB()

	fmt.Sprintln(ret)
}
