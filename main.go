package main

import (
	"file-management-api/httpserv"
	"file-management-api/infrastructure"
)

func init() {
	infrastructure.InitConfig()
}

func main() {
	infrastructure.InitAppConfig()
	infrastructure.InitStorage()
	httpserv.Run()
}
