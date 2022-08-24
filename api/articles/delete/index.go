package delete

import (
	"fmt"
	"github.com/terwer/final-note-kernel/controller"
	"github.com/terwer/final-note-kernel/service"
	"github.com/terwer/final-note-kernel/starter"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	starter.SetupConf()

	// 主逻辑
	service.ConnectDB()

	postid := r.URL.Query().Get("id")
	var ret = controller.RemoveArticleAction(&postid)

	// 及时关闭链接
	service.DisconnectDB()

	fmt.Fprintf(w, *ret)
}
