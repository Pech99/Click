package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
)

// .\click.exe <ctrl> <alt> <canc>

func main() {

	fmt.Println(os.Args)

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	VK.conv["<A>"]

	kb.SetKeys(keybd_event.VK_A, keybd_event.VK_B)
	kb.HasSHIFT(true)

	err = kb.Launching()
	if err != nil {
		panic(err)
	}

	kb.Press()
	time.Sleep(10 * time.Millisecond)
	kb.Release()
}
