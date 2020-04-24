// +build !windows
// +build !cgo,osusergo

package shell

// fallback, if not using cgo
func getCurrentUserShell() string {
	return ""
}
