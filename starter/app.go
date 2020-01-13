package starter

import (
	"configer"
	//"github.com/fatih/color"
	_ "github.com/joho/godotenv/autoload"
	"github.com/micro/go-micro/util/log"
)

func init() {

	//color.Red(fmt.Sprintf("%s:服务启动", app.NewConfig().GetName()))
	log.Logf("[Init] 服务启动：%s\n", configer.AppConfig().GetName())

	//fmt.Printf("%+v\n", app.DB)
}
