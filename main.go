package main

import (
	_ "gf-admin/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-admin/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
