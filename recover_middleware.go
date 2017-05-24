package recover

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

// RecoverMiddleware work when your app occure panic
type RecoverMiddleware struct {
	writeErrorResponseFunc WriteErrorResponseFunc
	logger                 LoggerInterface
	errorHandlerFunc       ErrorHandlerFunc
}

func getDefaultLogger() LoggerInterface {
	return log.New(os.Stdout, "[PANIC-RECOVER]", 0)
}

// Set DefaultModule(Logger and ErrorResponseWriter)
func Default() *RecoverMiddleware {
	return &RecoverMiddleware{DefaultWriteErrorResponseFunc, getDefaultLogger(), func(err interface{}) {}}
}

func Custom(writer WriteErrorResponseFunc, logger LoggerInterface, errorHandler ErrorHandlerFunc) *RecoverMiddleware {
	if writer == nil {
		writer = DefaultWriteErrorResponseFunc
	}
	if errorHandler == nil {
		errorHandler = func(err interface{}) {}
	}
	if logger == nil {
		logger = getDefaultLogger()
	}

	return &RecoverMiddleware{writer, logger, errorHandler}
}

func (r *RecoverMiddleware) CustomErrorResponseWriter(writer WriteErrorResponseFunc) {
	r.writeErrorResponseFunc = writer
}

func (r *RecoverMiddleware) CustomLogger(logger LoggerInterface) {
	r.logger = logger
}

func (r *RecoverMiddleware) CustomErrorHandler(f ErrorHandlerFunc) {
	r.errorHandlerFunc = f
}

func (r *RecoverMiddleware) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	defer func(r *RecoverMiddleware, rw http.ResponseWriter, req *http.Request) {
		err := recover()
		if err == nil {
			return
		}
		r.writeErrorResponseFunc(rw, req, err)
		traces := ""
		depth := 1
		for i := 20; i > 0; i-- {
			counter, file, line, ok := runtime.Caller(i)
			if !ok {
				continue
			}
			fn := runtime.FuncForPC(counter)
			traces += fmt.Sprintf("[%d]: %s: %s line:%d \n", depth, fn.Name(), file, line)
			depth++
		}
		r.logger.Printf("[PANIC] %s \n%s", err, traces)
		func(r *RecoverMiddleware, cerr interface{}) {
			defer func(r *RecoverMiddleware) {
				if panicErr := recover(); panicErr != nil {
					r.logger.Printf("[PANIC] RecoverPanic %v \n %v", err)
				}
			}(r)
			r.errorHandlerFunc(cerr)
		}(r, err)
	}(r, rw, req)
	next(rw, req)
}
