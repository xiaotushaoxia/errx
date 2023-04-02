package errx

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

var (
	Errorf       = errors.Errorf
	New          = errors.New
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef

	Cause  = errors.Cause
	Is     = errors.Is
	As     = errors.As
	Unwrap = errors.Unwrap
)

func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	if _, found := firstStackError(err); found {
		return WithMessage(err, msg)
	}
	return errors.Wrap(err, msg)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return Wrap(err, fmt.Sprintf(format, args...))
}

func Format(err error) string {
	if err == nil {
		return fmt.Sprint(err)
	}
	var ss []string
	ss = append(ss, "Error: "+err.Error())

	st, found := firstStackError(err)
	if !found {
		ss = append(ss, "Stack: empty")
	} else {
		for i, f := range st.StackTrace() {
			if i == 0 { // skip stack of Wrap
				ss = append(ss, "Stack:")
				continue
			}
			ss = append(ss, fmt.Sprintf("%+s:%d", f, f))
		}
	}
	return strings.Join(ss, "\n")
}
