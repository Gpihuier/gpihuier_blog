package main

import (
	"github.com/Gpihuier/gpihuier_blog/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	// 启动程序
	initialize.RunSever()
}
