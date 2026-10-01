//go:build windows

package vtinput

import (
	"os"
	"time"

	"golang.org/x/sys/windows"
)

// DrainInput drops input events still sitting in the console input buffer:
// records the terminal produced while the application was shutting down and
// nobody read. The second pass after a short wait catches reports that were
// in flight when the first pass ran -- by then the mode restore and its
// mouse-off announcement have already been written, so nothing new should
// arrive behind them.
func DrainInput() {
	h := windows.Handle(os.Stdin.Fd())
	if err := windows.FlushConsoleInputBuffer(h); err != nil {
		Log("VTINPUT: input drain failed: %v", err)
		return
	}
	time.Sleep(15 * time.Millisecond)
	if err := windows.FlushConsoleInputBuffer(h); err != nil {
		Log("VTINPUT: second input drain failed: %v", err)
	}
}
