package main

import (
	"github.com/lingyao2333/mo-zero/core/load"
	"github.com/lingyao2333/mo-zero/core/logx"
	"github.com/lingyao2333/mo-zero/tools/goctl/cmd"
)

func main() {
	logx.Disable()
	load.Disable()
	cmd.Execute()
}
