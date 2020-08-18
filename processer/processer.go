package processer

import (
	"math"
)

/*******************************************/
/*******Processer Module Definition*********/
/*******************************************/

// IProcesser is flow process logic interface.
type IProcesser interface {
	Use(middleware ...HandlerFunc) *Processer
	Do(e interface{})
}

/*******************************************/
/************Handler Structure**************/
/*******************************************/

// HandlerFunc is used as middleware by flow process logic.
type HandlerFunc func(*Context)

// HandlersChain connects HandlerFunc.
type HandlersChain []HandlerFunc

// Last returns last HandlerFunc in HandlersChain.
func (c HandlersChain) Last() HandlerFunc {
	if l := len(c); l > 0 {
		return c[l-1]
	}
	return nil
}

/*******************************************/
/************Context Structure**************/
/*******************************************/

const abortIndex int8 = math.MaxInt8 / 2

// Context is the flow processing container of events.
type Context struct {
	handlers HandlersChain
	index    int8
	Event    interface{}
}

// makeContext retuns Context pointer.
func makeContext(h HandlersChain, e interface{}) *Context {
	c := &Context{
		handlers: h,
		index:    -1,
		Event:    e,
	}
	return c
}

// Reset is used to reset the context for the next data.
func (c *Context) Reset(e interface{}) {
	c.index = -1
	c.Event = e
}

// Next execute next HandleFunc in HandleChain.
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

// Abort uses for stop flow process logic.
func (c *Context) Abort() {
	c.index = abortIndex
}

// IsAborted returns a bool variable to determine whether flow process logic stops.
func (c *Context) IsAborted() bool {
	return c.index >= abortIndex
}

/*******************************************/
/***********Processer Structure*************/
/*******************************************/

// Processer is IProcesser interface implmentot.
type Processer struct {
	handlers HandlersChain
	context  *Context
}

// Use attaches a middleware to the Processer.
func (p *Processer) Use(middleware ...HandlerFunc) *Processer {
	p.handlers = append(p.handlers, middleware...)
	return p
}

// Do Start flow processing.
func (p *Processer) Do(e interface{}) {
	if p.context == nil {
		p.context = makeContext(p.handlers, e)
	} else {
		p.context.Reset(e)
	}
	p.context.Next()
}
