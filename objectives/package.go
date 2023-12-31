package objectives

import (
	"encoding/json"
	"github.com/go-ai-agent/core/httpx"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"reflect"
	"sync/atomic"
)

type pkg struct{}

var (
	PkgUrl         = runtime.ParsePkgUrl(reflect.TypeOf(any(pkg{})).PkgPath())
	PkgUri         = PkgUrl.Host + PkgUrl.Path
	GuidancePath   = PkgUrl.Path + "/guidance/entry"
	ConstraintPath = PkgUrl.Path + "/constraint/entry"

	started int64
)

// IsPkgStarted - returns status of startup
func IsPkgStarted() bool {
	return atomic.LoadInt64(&started) != 0
}

func DoHandler(req *http.Request) (*http.Response, error) {
	recorder := httpx.NewRecorder()
	//status := sloHandler[runtime.BypassError](recorder, req)
	//var err error
	//if status.IsErrors() {
	//	err = status.Errors()[0]
	//}
	return recorder.Result(), nil
}

func GuidanceHandler(w http.ResponseWriter, r *http.Request) {
	guidanceHandler[runtime.LogError](w, r)
}

func guidanceHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) *runtime.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewHttpStatusCode(http.StatusBadRequest)
	}
	switch r.Method {
	case "GET":
		buf, status := marshalGuidance[E](GetGuidance())
		httpx.WriteResponse[E](w, buf, status)
		return status
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
	return runtime.NewStatusOK()
}

func marshalGuidance[E runtime.ErrorHandler](entry []Guidance) ([]byte, *runtime.Status) {
	buf, err := json.Marshal(entry)
	if err != nil {
		var e E
		return nil, e.Handle(nil, "marshal", err)
	}
	return buf, runtime.NewStatusOK()
}

func unmarshalGuidance[E runtime.ErrorHandler](buf []byte) ([]Guidance, *runtime.Status) {
	var entry []Guidance

	err := json.Unmarshal(buf, &entry)
	if err != nil {
		var e E
		return nil, e.Handle(nil, "unmarshal", err)
	}
	return entry, runtime.NewStatusOK()
}

func ConstraintHandler(w http.ResponseWriter, r *http.Request) {
	constraintHandler[runtime.LogError](w, r)
}

func constraintHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) *runtime.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewHttpStatusCode(http.StatusBadRequest)
	}
	switch r.Method {
	case "GET":
		buf, status := marshalConstraint[E](GetConstraint())
		httpx.WriteResponse[E](w, buf, status)
		return status
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
	return runtime.NewStatusOK()
}

func marshalConstraint[E runtime.ErrorHandler](entry []Constraint) ([]byte, *runtime.Status) {
	buf, err := json.Marshal(entry)
	if err != nil {
		var e E
		return nil, e.Handle(nil, "marshal", err)
	}
	return buf, runtime.NewStatusOK()
}

func unmarshalConstraint[E runtime.ErrorHandler](buf []byte) ([]Constraint, *runtime.Status) {
	var entry []Constraint

	err := json.Unmarshal(buf, &entry)
	if err != nil {
		var e E
		return nil, e.Handle(nil, "unmarshal", err)
	}
	return entry, runtime.NewStatusOK()
}
