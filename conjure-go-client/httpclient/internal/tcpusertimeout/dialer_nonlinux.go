// +build !linux

package tcpusertimeout

import (
	"time"

	"github.com/palantir/conjure-go-runtime/v2/conjure-go-client/httpclient/internal"
)

type dialer struct{}

func newDialWrapper(_ time.Duration) dialer {
	return dialer{}
}

func (d dialer) Wrap(delegate internal.DialFunc) internal.DialFunc {
	return delegate
}
