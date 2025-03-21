package mockutils

import (
	"fmt"

	"github.com/onsi/ginkgo/v2"
	"go.uber.org/mock/gomock"
)

type MockSequence struct {
	first, last *gomock.Call // may be both nil (= the default value of mockSequence) to mean empty sequence
}

// like gomock.InOrder, but can be nested.
func InOrder(calls ...any) MockSequence {
	var first, last *gomock.Call
	// This implementation does a few more assignments and checks than strictly necessary, but it is O(N) and reasonably easy to read, so, whatever.
	for i := range calls {
		var elem MockSequence
		switch e := calls[i].(type) {
		case MockSequence:
			elem = e
		case *gomock.Call:
			elem = MockSequence{e, e}
		default:
			ginkgo.Fail(fmt.Sprintf("Invalid inOrder parameter %#v", e))
		}

		if elem.first == nil {
			continue
		}

		if first == nil {
			first = elem.first
		} else if last != nil {
			elem.first.After(last)
		}

		last = elem.last
	}

	return MockSequence{first, last}
}
