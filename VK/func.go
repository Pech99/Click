package VK

import "syscall"

var dll = syscall.NewLazyDLL("user32.dll")
var procKeyBd = dll.NewProc("keybd_event")

func Release() {
	upKey(_VK_ALT)
	upKey(_VK_SHIFT)
	upKey(_VK_CTRL)
	upKey(_VK_RSHIFT)
	upKey(_VK_RCONTROL)
	upKey(_VK_LWIN)
}

//HasALT If key ALT pressed
func ALT(b bool) {
	if b {
		downKey(_VK_ALT)
	} else {
		upKey(_VK_ALT)
	}
}

//HasCTRL If key CTRL pressed
func CTRL(b bool) {
	if b {
		downKey(_VK_CTRL)
	} else {
		upKey(_VK_CTRL)
	}
}

//HasRCTRL If key CTRL pressed
func RCTRL(b bool) {
	if b {
		downKey(_VK_RCONTROL)
	} else {
		upKey(_VK_RCONTROL)
	}
}

//HasSHIFT If key SHIFT pressed
func SHIFT(b bool) {
	if b {
		downKey(_VK_SHIFT)
	} else {
		upKey(_VK_SHIFT)
	}
}

//HasSHIFTR If key SHIFT pressed
func RSHIFT(b bool) {
	if b {
		downKey(_VK_RSHIFT)
	} else {
		upKey(_VK_RSHIFT)
	}
}

//HasALTGR If key ALTGR pressed
func ALTGR(b bool) {
	if b {
		downKey(_VK_ALT)
		downKey(_VK_CTRL)
	} else {
		upKey(_VK_ALT)
		upKey(_VK_CTRL)
	}
}

//HasSuper If key Super pressed
func Super(b bool) {
	if b {
		downKey(_VK_LWIN)
	} else {
		upKey(_VK_LWIN)
	}
}

//HasCTRLR If key CTRLR pressed
//
//This is currently not supported on macOS
func CTRLR(b bool) {
	if b {
		downKey(_VK_LCONTROL)
	} else {
		upKey(_VK_LCONTROL)
	}
}

//HasSHIFTR If key SHIFTR pressed
//
//This is currently not supported on macOS
func SHIFTR(b bool) {
	if b {
		downKey(_VK_ALT)
	} else {
		upKey(_VK_ALT)
	}
}

func Press(key int) {
	downKey(key)
}

func downKey(key int) {
	flag := 0
	if key < 0xFFF { // Detect if the key code is virtual or no
		flag |= _KEYEVENTF_SCANCODE
	} else {
		key -= 0xFFF
	}
	vkey := key + 0x80
	procKeyBd.Call(uintptr(key), uintptr(vkey), uintptr(flag), 0)
}

func upKey(key int) {
	flag := _KEYEVENTF_KEYUP
	if key < 0xFFF {
		flag |= _KEYEVENTF_SCANCODE
	} else {
		key -= 0xFFF
	}
	vkey := key + 0x80
	procKeyBd.Call(uintptr(key), uintptr(vkey), uintptr(flag), 0)
}
