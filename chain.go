package httpexpect

import (
	"fmt"
)

type chain struct {
	context AssertionContext
	handler AssertionHandler
	isFatal bool
	failbit bool
}

func newChain(context AssertionContext, handler AssertionHandler) *chain {
	return &chain{
		context: context,
		handler: handler,
		isFatal: true,
	}
}

func newDefaultChain(name string, reporter Reporter) *chain {
	return &chain{
		context: AssertionContext{
			Path: []string{
				name,
			},
		},
		handler: &DefaultAssertionHandler{
			Formatter: &DefaultFormatter{},
			Reporter:  reporter,
		},
		isFatal: true,
	}
}

func (c *chain) setFatal(isFatal bool) {
	c.isFatal = isFatal
}

func (c *chain) setRequestName(name string) {
	c.context.RequestName = name
}

func (c *chain) setRequest(req *Request) {
	c.context.Request = req
}

func (c *chain) setResponse(resp *Response) {
	c.context.Response = resp
}

func (c *chain) clone() *chain {
	ret := *c

	ret.context.Path = nil
	ret.context.Path = append(ret.context.Path, c.context.Path...)

	return &ret
}

func (c *chain) enter(name string, args ...interface{}) {
	c.context.Path = append(c.context.Path, fmt.Sprintf(name, args...))
}

func (c *chain) leave() {
	if len(c.context.Path) == 0 {
		panic("unpaired enter/leave")
	}

	if !c.failbit {
		c.handler.Success(&c.context)
	}

	c.context.Path = c.context.Path[:len(c.context.Path)-1]
}

func (c *chain) fail(failure AssertionFailure) {
	if c.failbit {
		return
	}
	c.failbit = true

	if c.isFatal {
		failure.IsFatal = true
	}

	c.handler.Failure(&c.context, &failure)
}

func (c *chain) failed() bool {
	return c.failbit
}

func (c *chain) reset() {
	c.failbit = false
}

func (c *chain) assertOK(r Reporter) {
	if c.failbit {
		r.Errorf("failbit is true, but should be false")
	}
}

func (c *chain) assertFailed(r Reporter) {
	if !c.failbit {
		r.Errorf("failbit is false, but should be true")
	}
}
