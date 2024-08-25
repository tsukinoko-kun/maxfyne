//go:build darwin

package maxfyne

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import <Cocoa/Cocoa.h>
//
// static inline CGRect GetVisibleFrame() {
//     NSScreen *screen = [NSScreen mainScreen];
//     return [screen visibleFrame];
// }
import "C"
import "fyne.io/fyne/v2"

func getUsableScreenSize() (width float32, height float32) {
	// Get the main screen's visible frame
	visibleFrame := C.GetVisibleFrame()

	// Return the usable screen width and height
	return float32(visibleFrame.size.width), float32(visibleFrame.size.height)
}

func Maximize(w fyne.Window) error {
	w.Resize(fyne.NewSize(getUsableScreenSize()))
	return nil
}
