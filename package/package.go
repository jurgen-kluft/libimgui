package libimgui

// GetPackage returns the package object of 'xbase'
// How could we put the base URL github.com and the
// user 'jurgen-kluft' as a configuration var?

import (
	xbase "github.com/jurgen-kluft/xbase/package"
	"github.com/jurgen-kluft/xcode/denv"
	xentry "github.com/jurgen-kluft/xentry/package"
	xunittest "github.com/jurgen-kluft/xunittest/package"
)

func GetPackage() *denv.Package {
	name := "libimgui"

	// Dependencies
	xunittestpkg := xunittest.GetPackage()
	xentrypkg := xentry.GetPackage()
	xbasepkg := xbase.GetPackage()

	// The main (xbase) package
	mainpkg := denv.NewPackage(name)
	mainpkg.AddPackage(xunittestpkg)
	mainpkg.AddPackage(xentrypkg)
	mainpkg.AddPackage(xbasepkg)

	// library
	mainlib := denv.SetupDefaultCppLibProject(name, "github.com\\jurgen-kluft\\"+name)
	mainlib.Dependencies = append(mainlib.Dependencies, xunittestpkg.GetMainLib())

	// unittest project
	maintest := denv.SetupDefaultCppTestProject(name+"_test", "github.com\\jurgen-kluft\\"+name)
	maintest.Dependencies = append(maintest.Dependencies, xunittestpkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, xentrypkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, xbasepkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, mainlib)

	mainpkg.AddMainLib(mainlib)
	mainpkg.AddUnittest(maintest)
	return mainpkg
}
