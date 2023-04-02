package errx

import "github.com/pkg/errors"

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func firstStackError(err error) (st stackTracer, found bool) {
	if err == nil {
		return nil, false
	}
	var ss stackTracer
	var er = err
	for {
		stTemp, ok := er.(stackTracer)
		if ok {
			ss = stTemp
		}
		ner := errors.Unwrap(er)
		if ner == er {
			break
		}
		er = ner
	}
	if ss == nil {
		return nil, false
	}
	return ss, true
}
