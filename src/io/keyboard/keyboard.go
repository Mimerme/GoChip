package keyboard

import (
	"github.com/faiface/pixel/pixelgl"
)

//Keyboard mappings

const (
	key_1 = pixelgl.Key1
	key_2 = pixelgl.Key2
	key_3 = pixelgl.Key3
	key_4 = pixelgl.KeyQ
	key_5 = pixelgl.KeyW
	key_6 = pixelgl.KeyE
	key_7 = pixelgl.KeyA
	key_8 = pixelgl.KeyS
	key_9 = pixelgl.KeyD
	key_A = pixelgl.KeyZ
	key_B = pixelgl.KeyC
	key_C = pixelgl.Key4
	key_D = pixelgl.KeyR
	key_E = pixelgl.KeyF
	key_F = pixelgl.KeyV
	key_0 = pixelgl.KeyX
)

var window *pixelgl.Window

func Initialize(w *pixelgl.Window) {
	window = w
}

func Check_Keys(keys *[16]byte) {
	if window.Pressed(key_0) {
		keys[0x0] = 1
	} else {
		keys[0x0] = 0
	}
	if window.Pressed(key_1) {
		keys[0x1] = 1
	} else {
		keys[0x1] = 0
	}
	if window.Pressed(key_2) {
		keys[0x2] = 1
	} else {
		keys[0x2] = 0
	}
	if window.Pressed(key_3) {
		keys[0x3] = 1
	} else {
		keys[0x3] = 0
	}
	if window.Pressed(key_4) {
		keys[0x4] = 1
	} else {
		keys[0x4] = 0
	}
	if window.Pressed(key_5) {
		keys[0x5] = 1
	} else {
		keys[0x5] = 0
	}
	if window.Pressed(key_6) {
		keys[0x6] = 1
	} else {
		keys[0x6] = 0
	}
	if window.Pressed(key_7) {
		keys[0x7] = 1
	} else {
		keys[0x7] = 0
	}
	if window.Pressed(key_8) {
		keys[0x8] = 1
	} else {
		keys[0x8] = 0
	}
	if window.Pressed(key_9) {
		keys[0x9] = 1
	} else {
		keys[0x9] = 0
	}
	if window.Pressed(key_A) {
		keys[0xA] = 1
	} else {
		keys[0xA] = 0
	}
	if window.Pressed(key_B) {
		keys[0xB] = 1
	} else {
		keys[0xB] = 0
	}
	if window.Pressed(key_C) {
		keys[0xC] = 1
	} else {
		keys[0xC] = 0
	}
	if window.Pressed(key_D) {
		keys[0xD] = 1
	} else {
		keys[0xD] = 0
	}
	if window.Pressed(key_E) {
		keys[0xE] = 1
	} else {
		keys[0xE] = 0
	}
	if window.Pressed(key_F) {
		keys[0xF] = 1
	} else {
		keys[0xF] = 0
	}
}
