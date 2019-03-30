// Copyright 2019 The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build 3d

package glos

import (
	"image"
	"time"

	"github.com/go-gl/glfw/v3.0/glfw"
	"github.com/goki/gi/oswin/mouse"
)

var lastMouseClickTime time.Time
var lastMousePos image.Point
var lastMouseButton glfw.MouseButton
var lastMods int32

func glfwMods(mod glfw.ModifierKey) int32 {
	m := int32(0)
	if mod & glfw.ModShift {
		m |= 1 << uint32(key.Shift)
	}
	if mod & glfw.ModControl {
		m |= 1 << uint32(key.Control)
	}
	if mod & glfw.ModAlt {
		m |= 1 << uint32(key.Alt)
	}
	if mod & glfw.ModSuper {
		m |= 1 << uint32(key.Super)
	}
	return m
}

// physical key
func (w *windowImpl) keyEvent(gw *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mod glfw.ModifierKey) {
	em := glfwMods(mod)
	lastMods = em
	ec := glfwKeyCode(key)
	act := key.Press
	if action == glfw.Release {
		act = key.Release
	} else if action == glfw.Repeat {
		act = key.Press
	}
	ea := key.Actions(act)

	event := &key.Event{
		Code:      ec,
		Modifiers: em,
		Action:    act,
	}
	event.Init()
	w.Send(event)
}

// char input
func (w *windowImpl) charEvent(gw *glfw.Window, char rune, mods glfw.ModifierKey) {
	che := &key.ChordEvent{
		Rune: char,
	}
	sendEvent(id, che)

}

func (w *windowImpl) mouseButtonEvent(gw *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	mods := glfwMods(mod)
	lastMods = mods
	but := mouse.Left
	switch button {
	case glfw.MouseButtonMiddle:
		but = mouse.Middle
	case glfw.MouseButtonRight:
		but = mouse.Right

	}
	act := mouse.Press
	if action == glfw.Release {
		act = mouse.Release
	}
	if action == glfw.Press {
		interval := time.Now().Sub(lastMouseClickTime)
		// fmt.Printf("interval: %v\n", interval)
		if (interval / time.Millisecond) < time.Duration(mouse.DoubleClickMSec) {
			act = mouse.DoubleClick
		}
	}
	if mod&glfw.ModControl != 0 {
		but = mouse.Right
	}
	event := &mouse.Event{
		Where:     lastMousePos,
		Button:    but,
		Action:    act,
		Modifiers: mods,
	}
	event.Init()
	if act == mouse.Press {
		lastMouseClickTime = event.Time()
	}
	w.Send(event)
}

func (w *windowImpl) scrollEvent(gw *glfw.Window, xoff, yoff float64) {
	mods := lastMods
	event := &mouse.ScrollEvent{
		Event: mouse.Event{
			Where:     lastMousePos,
			Action:    mouse.Scroll,
			Modifiers: mods,
		},
		Delta: image.Point{int(-xoff), int(-yoff)},
	}
	event.Init()
	w.Send(event)
}

func (w *windowImpl) cursorPosEvent(gw *glfw.Window, x, y float64) {
	where := image.Point{int(x), int(y)}
	from := lastMousePos
	lastMousePos = where
	event := &mouse.MoveEvent{
		Event: mouse.Event{
			Where:     where,
			Button:    cmButton,
			Action:    mouse.Move,
			Modifiers: mods,
		},
		From: from,
	}
	event.Init()
	w.Send(event)
}

func (w *windowImpl) cursorEnterEvent(gw *glfw.Window, entered bool) {
}

func (w *windowImpl) dropEvent(gw *glfw.Window, names []string) {
}

func glfwKeyCode(kcode glfw.Key) key.Codes {
	switch vkcode {
	case glfw.KeyA:
		return key.CodeA
	case glfw.KeyB:
		return key.CodeB
	case glfw.KeyC:
		return key.CodeC
	case glfw.KeyD:
		return key.CodeD
	case glfw.KeyE:
		return key.CodeE
	case glfw.KeyF:
		return key.CodeF
	case glfw.KeyG:
		return key.CodeG
	case glfw.KeyH:
		return key.CodeH
	case glfw.KeyI:
		return key.CodeI
	case glfw.KeyJ:
		return key.CodeJ
	case glfw.KeyK:
		return key.CodeK
	case glfw.KeyL:
		return key.CodeL
	case glfw.KeyM:
		return key.CodeM
	case glfw.KeyN:
		return key.CodeN
	case glfw.KeyO:
		return key.CodeO
	case glfw.KeyP:
		return key.CodeP
	case glfw.KeyQ:
		return key.CodeQ
	case glfw.KeyR:
		return key.CodeR
	case glfw.KeyS:
		return key.CodeS
	case glfw.KeyT:
		return key.CodeT
	case glfw.KeyU:
		return key.CodeU
	case glfw.KeyV:
		return key.CodeV
	case glfw.KeyW:
		return key.CodeW
	case glfw.KeyX:
		return key.CodeX
	case glfw.KeyY:
		return key.CodeY
	case glfw.KeyZ:
		return key.CodeZ
	case glfw.Key1:
		return key.Code1
	case glfw.Key2:
		return key.Code2
	case glfw.Key3:
		return key.Code3
	case glfw.Key4:
		return key.Code4
	case glfw.Key5:
		return key.Code5
	case glfw.Key6:
		return key.Code6
	case glfw.Key7:
		return key.Code7
	case glfw.Key8:
		return key.Code8
	case glfw.Key9:
		return key.Code9
	case glfw.Key0:
		return key.Code0
	case glfw.KeyEnnter:
		return key.CodeReturnEnter
	case glfw.KeyEscape:
		return key.CodeEscape
	case glfw.KeyDelete:
		return key.CodeDeleteBackspace
	case glfw.KeyTab:
		return key.CodeTab
	case glfw.KeySpace:
		return key.CodeSpacebar
	case glfw.KeyMinus:
		return key.CodeHyphenMinus
	case glfw.KeyEqual:
		return key.CodeEqualSign
	case glfw.KeyLeftBracket:
		return key.CodeLeftSquareBracket
	case glfw.KeyRightBracket:
		return key.CodeRightSquareBracket
	case glfw.KeyBackslash:
		return key.CodeBackslash
	case glfw.KeySemicolon:
		return key.CodeSemicolon
	case glfw.KeyQuote:
		return key.CodeApostrophe
	case glfw.KeyGraveAccent:
		return key.CodeGraveAccent
	case glfw.KeyComma:
		return key.CodeComma
	case glfw.KeyPeriod:
		return key.CodeFullStop
	case glfw.KeySlash:
		return key.CodeSlash
	case glfw.KeyCapsLock:
		return key.CodeCapsLock
	case glfw.KeyF1:
		return key.CodeF1
	case glfw.KeyF2:
		return key.CodeF2
	case glfw.KeyF3:
		return key.CodeF3
	case glfw.KeyF4:
		return key.CodeF4
	case glfw.KeyF5:
		return key.CodeF5
	case glfw.KeyF6:
		return key.CodeF6
	case glfw.KeyF7:
		return key.CodeF7
	case glfw.KeyF8:
		return key.CodeF8
	case glfw.KeyF9:
		return key.CodeF9
	case glfw.KeyF10:
		return key.CodeF10
	case glfw.KeyF11:
		return key.CodeF11
	case glfw.KeyF12:
		return key.CodeF12
	// 70: PrintScreen
	// 71: Scroll Lock
	// 72: Pause
	// 73: Insert
	case glfw.KeyHome:
		return key.CodeHome
	case glfw.KeyPageUp:
		return key.CodePageUp
	case glfw.KeyForwardDelete:
		return key.CodeDeleteForward
	case glfw.KeyEnd:
		return key.CodeEnd
	case glfw.KeyPageDown:
		return key.CodePageDown
	case glfw.KeyRightArrow:
		return key.CodeRightArrow
	case glfw.KeyLeftArrow:
		return key.CodeLeftArrow
	case glfw.KeyDownArrow:
		return key.CodeDownArrow
	case glfw.KeyUpArrow:
		return key.CodeUpArrow
	case glfw.KeyKPClear:
		return key.CodeKeypadNumLock
	case glfw.KeyKPDivide:
		return key.CodeKeypadSlash
	case glfw.KeyKPMultiply:
		return key.CodeKeypadAsterisk
	case glfw.KeyKPMinus:
		return key.CodeKeypadHyphenMinus
	case glfw.KeyKPPlus:
		return key.CodeKeypadPlusSign
	case glfw.KeyKPEnter:
		return key.CodeKeypadEnter
	case glfw.KeyKP1:
		return key.CodeKeypad1
	case glfw.KeyKP2:
		return key.CodeKeypad2
	case glfw.KeyKP3:
		return key.CodeKeypad3
	case glfw.KeyKP4:
		return key.CodeKeypad4
	case glfw.KeyKP5:
		return key.CodeKeypad5
	case glfw.KeyKP6:
		return key.CodeKeypad6
	case glfw.KeyKP7:
		return key.CodeKeypad7
	case glfw.KeyKP8:
		return key.CodeKeypad8
	case glfw.KeyKP9:
		return key.CodeKeypad9
	case glfw.KeyKP0:
		return key.CodeKeypad0
	case glfw.KeyKPDecimal:
		return key.CodeKeypadFullStop
	case glfw.KeyKPEquals:
		return key.CodeKeypadEqualSign
	case glfw.KeyF13:
		return key.CodeF13
	case glfw.KeyF14:
		return key.CodeF14
	case glfw.KeyF15:
		return key.CodeF15
	case glfw.KeyF16:
		return key.CodeF16
	case glfw.KeyF17:
		return key.CodeF17
	case glfw.KeyF18:
		return key.CodeF18
	case glfw.KeyF19:
		return key.CodeF19
	case glfw.KeyF20:
		return key.CodeF20
	// 116: Keyboard Execute
	case glfw.KeyHelp:
		return key.CodeHelp
	// 118: Keyboard Menu
	// 119: Keyboard Select
	// 120: Keyboard Stop
	// 121: Keyboard Again
	// 122: Keyboard Undo
	// 123: Keyboard Cut
	// 124: Keyboard Copy
	// 125: Keyboard Paste
	// 126: Keyboard Find
	case glfw.KeyMute:
		return key.CodeMute
	case glfw.KeyVolumeUp:
		return key.CodeVolumeUp
	case glfw.KeyVolumeDown:
		return key.CodeVolumeDown
	// 130: Keyboard Locking Caps Lock
	// 131: Keyboard Locking Num Lock
	// 132: Keyboard Locking Scroll Lock
	// 133: Keyboard Comma
	// 134: Keyboard Equal Sign
	// ...: Bunch of stuff
	case glfw.KeyControl:
		return key.CodeLeftControl
	case glfw.KeyShift:
		return key.CodeLeftShift
	case glfw.KeyOption:
		return key.CodeLeftAlt
	case glfw.KeyCommand:
		return key.CodeLeftGUI
	case glfw.KeyRightControl:
		return key.CodeRightControl
	case glfw.KeyRightShift:
		return key.CodeRightShift
	case glfw.KeyRightOption:
		return key.CodeRightAlt
	default:
		return key.CodeUnknown
	}
}