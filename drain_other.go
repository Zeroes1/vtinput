//go:build !windows

package vtinput

// DrainInput is a no-op outside Windows: a terminal is a byte stream the
// reader drains by reading it, and there is no console input buffer holding
// records behind the reader's back.
func DrainInput() {}
