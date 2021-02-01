package main

import (
	"fmt"
	"tinyURL/initialize"
	"tinyURL/router"
)

var aboutMe = `
┌┬┐┬┌┐┌┬ ┬╦ ╦╦═╗╦  
 │ ││││└┬┘║ ║╠╦╝║  
 ┴ ┴┘└┘ ┴ ╚═╝╩╚═╩═╝
`

func main() {
	initialize.Config()
	initialize.DataBase()
	initialize.Routers()
	fmt.Println(aboutMe)
	router.RunAppServer()
}
