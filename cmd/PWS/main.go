package main

import (
	//"github.com/asticode/go-astilectron"
	"runtime"
	// "net/rpc"
)

type local struct {
	Os_type string
}

func main() {

	l_device := local{}
	l_device.Os_type = runtime.GOOS

}
