package main

import (
	"os"
	"runtime"
	"strconv"
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

		if key[:3] == "<#S" && key[len(key)-1] == '>' {
			ms, err := strconv.Atoi(key[3 : len(key)-1])
			if err == nil {
				time.Sleep(time.Duration(ms) * time.Millisecond)
			}
		}

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
		case "<SUPER>":
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
		case "<-SUPER>":
			VK.Super(false)
		default:
			VK.Press(VK.Conv[key])
		}
		time.Sleep(5 * time.Millisecond)
	}

	VK.Release()
}
