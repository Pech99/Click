package main

import (
	"os"
	"runtime"
	"strings"
	"time"

	VK "github.com/Pech99/Click/VKconv"
	"github.com/micmonay/keybd_event"
)

// .\click.exe <ctrl> <alt> <canc>
func main() {

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	var keys []int

	for _, key := range os.Args[1:] {

		key = strings.ToUpper(key)

		switch key {
		case "<CTRL>":
			kb.HasCTRL(true)
		case "<ALT>":
			kb.HasALT(true)
		case "<SHIFT>":
			kb.HasSHIFT(true)
		case "<RCTR>":
			kb.HasCTRLR(true)
		case "<RSHIFT>":
			kb.HasSHIFTR(true)
		case "<ALTGR>":
			kb.HasALTGR(true)
		case "<Super>":
			kb.HasSuper(true)
		default:
			keys = append(keys, VK.Conv[key])
		}
	}

	if len(keys) > 0 {
		kb.SetKeys(keys...)
	}

	err = kb.Launching()
	if err != nil {
		panic(err)
	}

	kb.Press()
	time.Sleep(10 * time.Millisecond)
	kb.Release()

	return
}
