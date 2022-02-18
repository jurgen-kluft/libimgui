package main

import (
	pkg "github.com/jurgen-kluft/libimgui/package"
	"github.com/jurgen-kluft/xcode"
)

func main() {
	xcode.Init()
	xcode.Generate(pkg.GetPackage())
}
