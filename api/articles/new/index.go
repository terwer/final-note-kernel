package new

import (
	"fmt"
	kernel "github.com/terwer/final-note-kernel"
	"github.com/terwer/final-note-kernel/controller"
	"github.com/terwer/final-note-kernel/service"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	kernel.Init()

	// 主逻辑
	service.ConnectDB()

	var ret = controller.AddArticleAction()

	// 及时关闭链接
	service.DisconnectDB()

	fmt.Fprintf(w, *ret)
}
