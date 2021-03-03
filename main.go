package main

import (
	"fmt"
	_ "tinyURL/docs"
	"tinyURL/initialize"
	"tinyURL/router"
)

var aboutMe = `
┌┬┐┬┌┐┌┬ ┬╦ ╦╦═╗╦  
 │ ││││└┬┘║ ║╠╦╝║  
 ┴ ┴┘└┘ ┴ ╚═╝╩╚═╩═╝
`

// @title 短链接服务接口文档
// @version 0.0.1
// @description 该服务用以生成短链接

// @contact.name lomogo
// @contact.email dhdpat@163.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	initialize.Config()
	initialize.Logger()
	initialize.DataBase()
	initialize.Routers()
	fmt.Println(aboutMe)
	router.RunAppServer()
}
