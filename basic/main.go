package main

import (
	"basic/basic2"
	"runtime"
)

func main() {
	basic2.ForLoop()
	basic2.IfControl(256)

	os := runtime.GOOS
	basic2.SwitchContorol(os)
}
