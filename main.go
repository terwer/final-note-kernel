package final_note_kernel

import (
	"fmt"
	"github.com/88250/gulu"
	"github.com/terwer/final-note-kernel/model"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Logger
var logger *gulu.Logger

// Init The only one init function in pipe.
func Init() {
	rand.Seed(time.Now().UTC().UnixNano())

	gulu.Log.SetLevel("debug")
	logger = gulu.Log.NewLogger(os.Stdout)

	model.LoadConf()
	replaceServerConf()

	fmt.Println("init")
}

func replaceServerConf() {
	path := "theme/sw.min.js.tpl"
	if gulu.File.IsExist(path) {
		data, err := ioutil.ReadFile(path)
		if nil != err {
			logger.Fatal("read file [" + path + "] failed: " + err.Error())
		}
		content := string(data)
		content = strings.Replace(content, "http://server.tpl.json", model.Conf.Server, -1)
		content = strings.Replace(content, "http://staticserver.tpl.json", model.Conf.StaticServer, -1)
		content = strings.Replace(content, "${StaticResourceVersion}", model.Conf.StaticResourceVersion, -1)
		writePath := strings.TrimSuffix(path, ".tpl")
		if err = ioutil.WriteFile(writePath, []byte(content), 0644); nil != err {
			logger.Fatal("replace sw.min.js in [" + path + "] failed: " + err.Error())
		}
	}

	if gulu.File.IsExist("console/dist/") {
		err := filepath.Walk("console/dist/", func(path string, f os.FileInfo, err error) error {
			if strings.HasSuffix(path, ".tpl") {
				data, err := ioutil.ReadFile(path)
				if nil != err {
					logger.Fatal("read file [" + path + "] failed: " + err.Error())
				}
				content := string(data)
				content = strings.Replace(content, "http://server.tpl.json", model.Conf.Server, -1)
				content = strings.Replace(content, "http://staticserver.tpl.json", model.Conf.StaticServer, -1)
				writePath := strings.TrimSuffix(path, ".tpl")
				if err = ioutil.WriteFile(writePath, []byte(content), 0644); nil != err {
					logger.Fatal("replace server conf in [" + writePath + "] failed: " + err.Error())
				}
			}

			return err
		})
		if nil != err {
			logger.Fatal("replace server conf in [theme] failed: " + err.Error())
		}
	}
}
