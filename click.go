package main

import (
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/Pech99/Click/VK"
)

// .\click.exe <ctrl> <alt> <canc>
//complile: go build -ldflags "-H windowsgui"
//hiden console

func main() {

	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	for _, key := range os.Args[1:] {

		key = strings.ToUpper(key)

		switch key {
		case "<CTRL>":
			VK.CTRL(true)
		case "<ALT>":
			VK.ALT(true)
		case "<SHIFT>":
			VK.SHIFT(true)
		case "<RCTR>":
			VK.CTRLR(true)
		case "<RSHIFT>":
			VK.SHIFTR(true)
		case "<ALTGR>":
			VK.ALTGR(true)
		case "<Super>":
			VK.Super(true)

		case "<-CTRL>":
			VK.CTRL(false)
		case "<-ALT>":
			VK.ALT(false)
		case "<-SHIFT>":
			VK.SHIFT(false)
		case "<-RCTR>":
			VK.CTRLR(false)
		case "<-RSHIFT>":
			VK.SHIFTR(false)
		case "<-ALTGR>":
			VK.ALTGR(false)
		case "<-Super>":
			VK.Super(false)
		default:
			VK.Press(VK.Conv[key])
		}
	}

	VK.Release()
}
