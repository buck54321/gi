// Copyright 2018 The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package windriver

import (
	"fmt"
	"log"
	"syscall"
	"time"
	"unsafe"

	"github.com/goki/gi/oswin/mimedata"
	"github.com/goki/gi/oswin/driver/internal/win32"
)

// implements clipboard support for Windows
// https://github.com/jtanx/libclipboard/blob/master/src/clipboard_win32.c
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms649016(v=vs.85).aspx
// formats:
// https://msdn.microsoft.com/en-us/library/windows/desktop/ff729168(v=vs.85).aspx
// https://github.com/AllenDang/w32

type clipImpl struct {
}

var theClip = clipImpl{}

// ClipRetries determines how many times to retry in opening the clipboard
var ClipRetries = 5

// ClipRetrySleep determines how long to sleep between retries
var ClipRetrySleep = 5 * time.Millisecond

func (ci *clipImpl) OpenClipboard() bool {
	if len(theApp.winlist) == 0 {
		return false
	}
	win := theApp.winlist[0]
	for retry := 0; retry < ClipRetries; retry++ {
		got := win32.OpenClipboard(win.hwnd)
		if got {
			return true
		}
		time.Sleep(ClipRetrySleep)
	}
	return false
}

func (ci *clipImpl) Read(types []string) mimedata.Mimes {
	if types == nil {
		return nil
	}
	if !ci.OpenClipboard() {
		return nil
	}
	defer win32.CloseClipboard()

	for _, typ := range types {
		if typ == mimedata.TextPlain || typ == mimedata.TextAny || typ == mimedata.AppJSON {
			hData := win32.GetClipboardData(win32.CF_UNICODETEXT)
			if hData == 0 {
				log.Printf("clip.Board.Read couldn't get clip data\n")
				return nil
			}
			wd := win32.GlobalLock(hData)
			if wd == nil {
				log.Printf("clip.Board.Read couldn't lock clip data\n")
				return nil
			}
			defer win32.GlobalUnlock(hData)
			txt := syscall.UTF16ToString((*[1 << 20]uint16)(unsafe.Pointer(wd))[:])
			// todo: verify txt format for JSON etc
			return mimedata.NewMime(typ, []byte(txt))
		}
	}
	return nil
}

func (ci *clipImpl) Write(data mimedata.Mimes, clearFirst bool) error {
	// clearFirst not relevant
	if !ci.OpenClipboard() {
		return fmt.Errorf("clip.Board.Write could not open clipboard\n")
	}
	defer win32.CloseClipboard()

	if !win32.EmptyClipboard() {
		return fmt.Errorf("clip.Board.Write could not empty clipboard\n")
	}

	for _, d := range data {
		if d.Type == mimedata.TextPlain || d.Type == mimedata.TextAny || d.Type == mimedata.AppJSON {
			wc, err := syscall.UTF16FromString(string(d.Data))
			if err != nil {
				return err
			}
			sz := uintptr(len(wc)*2)
			hData := win32.GlobalAlloc(win32.GMEM_MOVEABLE, sz)
			wd := win32.GlobalLock(hData)
			if wd == nil {
				log.Printf("clip.Board.Write couldn't lock clip data\n")
				return nil
			}
			win32.CopyMemory(uintptr(unsafe.Pointer(wd)), uintptr(unsafe.Pointer(&wc[0])), sz)
			win32.GlobalUnlock(hData)

			hRes := win32.SetClipboardData(win32.CF_UNICODETEXT, hData)
			if hRes == 0 {
				win32.GlobalFree(hData)
			}
			break // only 1
		}
	}
	return nil
}

func (ci *clipImpl) Clear() {
	if !ci.OpenClipboard() {
		return
	}
	win32.EmptyClipboard()
	win32.CloseClipboard()
}