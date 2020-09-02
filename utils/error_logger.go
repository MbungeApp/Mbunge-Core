package utils

import (
	"fmt"
	"github.com/MbungeApp/mbunge-core/config"
	"runtime"
)

func HandleError(err error) (b bool) {
	if err != nil {
		// notice that we're using 1, so it will actually log where
		// the error happened, 0 = this function, we don't want that.
		_, fn, line, _ := runtime.Caller(1)
		config.ErrorReporter(fmt.Sprintf("[error] %s:%d %v", fn, line, err))
		b = true
	}
	return
}
