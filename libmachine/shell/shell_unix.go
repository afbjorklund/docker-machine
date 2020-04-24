// +build !windows
// +build cgo,!osusergo

package shell

/*
#include <pwd.h>
#include <unistd.h>

static char *get_current_user_shell() {
    struct passwd *pwd = getpwuid(geteuid());
    return (pwd != NULL) ? pwd->pw_shell : "";
}
*/
import "C"

func getCurrentUserShell() string {
	return C.GoString(C.get_current_user_shell())
}
