package api

// API wrapper interface
type Wrapper interface {
	Method(string, func(*Ctx))
	Serve(int)
}

// API context
type Context struct {
	String func(string)
	Bytes  func([]byte)
	JSON   func(interface{})
}

type Ctx = Context

// new API wrapper
var (
	currentWrapper = HTTPWrapper{}

	Method = currentWrapper.Method
	Serve  = currentWrapper.Serve
)
