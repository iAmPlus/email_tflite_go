package app

import (
	"net/http"
	"runtime/debug"
	"runtime/pprof"
	"runtime"
	"strconv"
)

func (a *Application) GoRoutinesCount(w http.ResponseWriter, r *http.Request) {
    count := runtime.NumGoroutine()
    w.Write([]byte(strconv.Itoa(count)))
}

// respond with the stack trace of the system.
func (a *Application)  GetStackTraceHandler(w http.ResponseWriter, r *http.Request) {
	stack := debug.Stack()
	w.Write(stack)
	pprof.Lookup("goroutine").WriteTo(w, 2)
}

func (a *Application) addStackTraces() {
	a.router.HandleFunc("/_count", a.GoRoutinesCount)
	a.router.HandleFunc("/_stack", a.GetStackTraceHandler)
}