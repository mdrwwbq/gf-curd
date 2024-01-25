package main

import (
	"context"
	"flag"
	"github.com/mdrwwbq/gf-curd/src"
)

var (
	moduleName string // 模块名，如：基础(basic)、系统(system)、公共(common)
	toolsName  string // 具体功能名，如：菜单(menu)
	outputPath string
	cover      bool
)

func init() {
	flag.StringVar(&moduleName, "m", "", "模块名称，如：-m basic 非必填")
	flag.StringVar(&toolsName, "t", "", "具体功能名，如：-t menu或-t userRole")
	flag.StringVar(&outputPath, "o", "./output", "输出路径，如：-o output")
	flag.BoolVar(&cover, "c", true, "覆盖写入，如：-c true覆盖（默认），-c=false不覆盖")
}
func main() {
	flag.Parse()
	ctx := context.TODO()
	h := src.NewHandle(ctx)
	h.ModuleName = moduleName
	h.ToolsName = toolsName
	h.OutputPath = outputPath
	h.Cover = cover
	h.Run()
}
