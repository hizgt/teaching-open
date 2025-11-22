package main

import (
	_ "api-go/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"api-go/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
