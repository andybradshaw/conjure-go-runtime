package internal

import (
	"context"
	"net"
)

// DialFunc is a convenience type alias for the standard library net.Wrap function
type DialFunc func(ctx context.Context, network, address string) (net.Conn, error)

// A DialWrapper takes a delegate dial function and returns a corresponding dial function. This is useful for
// interacting with the underlying connection
type DialWrapper func(delegate DialFunc) DialFunc
