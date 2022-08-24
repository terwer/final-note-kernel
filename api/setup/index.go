package setup

import (
	"fmt"
	"github.com/terwer/final-note-kernel/service"
	"github.com/terwer/final-note-kernel/starter"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	starter.SetupConf()

	// 主逻辑
	service.ConnectDB()

	// 及时关闭链接
	service.DisconnectDB()

	fmt.Fprintf(w, "<h1>初始化完毕</h1>")
}
