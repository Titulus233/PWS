package main

import (
	"runtime"
)

type local struct {
	Os_type string
}

func main() {

	l_device := local{}
	l_device.Os_type = runtime.GOOS

}
