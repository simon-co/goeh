package goeh

import (
	"encoding/json"
	"fmt"
	"github.com/simon-co/goeh/internal/calltrace"
)

type AppErr struct {
	File      string
	Operation string
	Line      int
	Message   string
	Cause     error
	trace     []string
}

// return a Json formatted error Message
func (e *AppErr) Error() string {
	jsonByteSlice, _ := json.MarshalIndent(e, "", "  ")
	return string(jsonByteSlice)
}

// used in errors.Is check
func (e *AppErr) Is(target error) bool {
	t, ok := target.(*AppErr)
	if !ok {
		return false
	}
	return e.Message == t.Message
}

// used in errors.Has
func (e *AppErr) UnWrap() error {
	return e.Cause
}

// formats error trace string and adds to Error trace
func (e *AppErr) addTrace(line int, operation string, file string) {
	trace := fmt.Sprintf("file: %s, operation: %s, line: %d;", file, operation, line)
	e.trace = append(e.trace, trace)
}

func (e *AppErr) String() string {
	return e.Error()
}

// Parses the supplied error.  If error is of type AppErr, then
// a calltrace is added to the error and it is returned.  Else, a
// *AppErr is created and returned.
func Parse(e error) *AppErr {
	trace, _ := calltrace.Full(2)
	ae, ok := e.(*AppErr)
	if !ok {
		ae := &AppErr{
			File:      trace.File,
			Operation: trace.Function,
			Line:      trace.Line,
			Message:   e.Error(),
			Cause:     e,
		}
		ae.addTrace(ae.Line, ae.Operation, ae.File)
		return ae
	} else {
		ae.addTrace(trace.Line, trace.Function, trace.File)
		return ae
	}
}

// Parses the supplied error to the supplied callstack depth
func ParseToDepth(e error, depth int) *AppErr {
	trace, _ := calltrace.Full(depth)
	ae, ok := e.(*AppErr)
	if !ok {
		ae := &AppErr{
			File:      trace.File,
			Operation: trace.Function,
			Line:      trace.Line,
			Message:   e.Error(),
			Cause:     e,
		}
		ae.addTrace(ae.Line, ae.Operation, ae.File)
		return ae
	} else {
		ae.addTrace(trace.Line, trace.Function, trace.File)
		return ae
	}
}
