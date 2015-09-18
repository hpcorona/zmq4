// +build !windows

package zmq4

/*
#include <zmq.h>
#include "zmq4.h"
*/
import "C"

/*
Sets the scheduling policy for internal context’s thread pool.

This option requires ZeroMQ version 4.1.3, and is not available on Windows.

Supported values for this option can be found in sched.h file, or at
http://man7.org/linux/man-pages/man2/sched_setscheduler.2.html

This option only applies before creating any sockets on the context.

Default value   -1

Returns ErrorNotImplemented413 with ZeroMQ version < 4.1.3

Returns ErrorNotImplementedWindows on Windows
*/
func (ctx *Context) SetThreadSchedPolicy(n int) error {
	if minor < 1 || (minor == 1 && patch < 3) {
		return ErrorNotImplemented413
	}
	return setOption(ctx, C.ZMQ_THREAD_SCHED_POLICY, n)
}

/*
Sets scheduling priority for internal context’s thread pool.

This option requires ZeroMQ version 4.1.3, and is not available on Windows.

Supported values for this option depend on chosen scheduling policy.
Details can be found in sched.h file, or at
http://man7.org/linux/man-pages/man2/sched_setscheduler.2.html

This option only applies before creating any sockets on the context.

Default value   -1

Returns ErrorNotImplemented413 with ZeroMQ version < 4.1.3

Returns ErrorNotImplementedWindows on Windows
*/
func (ctx *Context) SetThreadPriority(n int) error {
	if minor < 1 || (minor == 1 && patch < 3) {
		return ErrorNotImplemented413
	}
	return setOption(ctx, C.ZMQ_THREAD_PRIORITY, n)
}
