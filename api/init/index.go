package init

import (
	"fmt"
	kernel "github.com/terwer/final-note-kernel"
	"github.com/terwer/final-note-kernel/service"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	kernel.Init()

	// 主逻辑
	service.ConnectDB()

	// 及时关闭链接
	service.DisconnectDB()

	fmt.Fprintf(w, "<h1>初始化完毕</h1>")
}
