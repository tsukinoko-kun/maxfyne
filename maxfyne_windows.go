//go:build windows

package maxfyne

import (
	"errors"
	"reflect"
	"time"
	"unsafe"

	"fyne.io/fyne/v2"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func getWindow(w fyne.Window) (*glfw.Window, bool) {
	v := reflect.ValueOf(w).Elem()
	viewportField := v.FieldByName("viewport")
	glfwWin, ok := reflect.NewAt(viewportField.Type(), unsafe.Pointer(viewportField.UnsafeAddr())).Elem().Interface().(*glfw.Window)
	return glfwWin, ok
}

func Maximize(w fyne.Window) error {
	for i := 0; i < 200; i++ {
		window, ok := getWindow(w)
		if !ok || window == nil {
			<-time.After(100 * time.Millisecond)
			continue
		}

		window.Maximize()

		return nil
	}

	return errors.New("Failed to get glfw.Window from fyne.Window")
}
