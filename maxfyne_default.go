//go:build !darwin && !windows

package maxfyne

import "fyne.io/fyne/v2"

func Maximize(_ fyne.Window) error {
	return NotImplemented
}
