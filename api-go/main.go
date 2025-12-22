package main

import (
	_ "teaching-open/internal/logic/sys"
	_ "teaching-open/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"teaching-open/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
