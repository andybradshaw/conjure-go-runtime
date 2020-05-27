// +build linux

package tcpusertimeout

import (
	"context"
	"net"
	"syscall"
	"time"

	"github.com/palantir/conjure-go-runtime/v2/conjure-go-client/httpclient/internal"
	werror "github.com/palantir/witchcraft-go-error"
	"golang.org/x/sys/unix"
)

type dialer struct {
	timeout time.Duration
}

func newDialWrapper(timeout time.Duration) dialer {
	return dialer{
		timeout: timeout,
	}
}

func (d dialer) Wrap(delegate internal.DialFunc) internal.DialFunc {
	return func(ctx context.Context, network string, address string) (net.Conn, error) {
		conn, err := delegate(ctx, network, address)
		if err != nil {
			return nil, err
		}
		// TODO(abradshaw): close connection on failures? safe to return conn and err?
		if err := setTCPUserTimeout(conn, d.timeout); err != nil {
			return nil, err
		}
		return conn, nil
	}
}

// setTCPUserTimeout sets the TCP user timeout on a connection's socket
func setTCPUserTimeout(conn net.Conn, timeout time.Duration) error {
	tcpconn, ok := conn.(*net.TCPConn)
	if !ok {
		// not a TCP connection. exit early
		return nil
	}
	rawConn, err := tcpconn.SyscallConn()
	if err != nil {
		return werror.Wrap(err, "error getting raw connection")
	}
	err = rawConn.Control(func(fd uintptr) {
		err = syscall.SetsockoptInt(int(fd), syscall.IPPROTO_TCP, unix.TCP_USER_TIMEOUT, int(timeout/time.Millisecond))
	})
	if err != nil {
		return werror.Wrap(err, "error setting TCP_USER_TIMEOUT option on socket")
	}

	return nil
}
