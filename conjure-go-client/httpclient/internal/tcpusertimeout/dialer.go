package tcpusertimeout

import (
	"time"

	"github.com/palantir/conjure-go-runtime/v2/conjure-go-client/httpclient/internal"
)

func NewDialWrapper(timeout time.Duration) internal.DialWrapper {
	return newDialWrapper(timeout).Wrap
}
