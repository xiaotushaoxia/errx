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

func WithStack(err error) error {
	if err == nil {
		return nil
	}
	if _, found := firstStackError(err); found {
		return err
	}
	return errors.WithStack(err)
}

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
	var ss = []string{"Error: "+err.Error()}
	stacks := GetStack(err)
	if len(stacks) > 0 {
		ss = append(ss, "Stacks:")
		for _, s := range stacks {
			ss = append(ss, fmt.Sprintf("%+s:%d", s, s))
		}
	} else {
		ss = append(ss, "Stack: empty")
	}
	return strings.Join(ss, "\n")
}

func GetStack(err error) errors.StackTrace {
	if err == nil {
		return nil
	}
	st, found := firstStackError(err)
	if !found {
		return nil
	}
	var result errors.StackTrace
	for _, f := range st.StackTrace() {
		result = append(result, f)
	}
	return result
}
