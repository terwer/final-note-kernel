package detail

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

	postid := r.URL.Query().Get("id")
	var ret = controller.GetArticleAction(&postid)

	// 及时关闭链接
	service.DisconnectDB()

	fmt.Fprintf(w, *ret)
}
