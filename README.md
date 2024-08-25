# maxfyne

Add the missing functionality of maximizing the Fyne windows.

## Usage

```go
package main

import (
    "errors"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
    "github.com/tsukinoko-kun/maxfyne"
)

func main() {
    a := app.New()
    w := a.NewWindow("Example")

    w.SetContent(widget.NewLabel("Hello, World!"))

	// Maximize needs to run in a goroutine
    // because it may need to wait for the window to be initialized.
    go func() {
		// you don't need to handle the error if you don't care about it
		// if Maximize fails, it will return an error
        // no crash will occur
        if err := maxfyne.Maximize(w); err != nil {
            // maximize failed
            if errors.Is(err, maxfyne.NotImplemented) {
                // unsupported platform
            } else {
                // other error
            }
        }
    }()

    w.ShowAndRun()
}
```

## Installation

```shell
go get -u github.com/tsukinoko-kun/maxfyne
```

## Works on

- [x] macOS
- [x] Windows
- [ ] Linux (Windows implementation might work here as well)

On unsupported platforms, `Maximize` will return `maxfyne.NotImplemented`.
