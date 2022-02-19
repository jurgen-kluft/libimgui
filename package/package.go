package libimgui

// GetPackage returns the package object of 'xbase'
// How could we put the base URL github.com and the
// user 'jurgen-kluft' as a configuration var?

import (
	libglfw "github.com/jurgen-kluft/libglfw/package"
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
	libglfwpkg := libglfw.GetPackage()

	// The main (libimgui) package
	mainpkg := denv.NewPackage(name)
	mainpkg.AddPackage(xbasepkg)
	mainpkg.AddPackage(libglfwpkg)

	// Imgui Library
	mainlib := denv.SetupDefaultCppLibProject(name, "github.com\\jurgen-kluft\\"+name)
	mainlib.Dependencies = append(mainlib.Dependencies, libglfwpkg.GetMainLib())

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
