package processer

import (
	"testing"
)

func TestUseAndDo(t *testing.T) {
	tests := []struct{ a, e int }{
		{0, 1},
		{0, 2},
		{0, 3},
		{0, 4},
		{0, 5},
	}

	var p IProcesser = &Processer{}
	p.Use(func(c *Context) {
		tests[0].a = 1
		c.Next()
	})
	p.Use(func(c *Context) {
		tests[1].a = 2
		c.Next()
	})
	p.Use(func(c *Context) {
		tests[2].a = 3
		c.Next()
	})
	p.Use(func(c *Context) {
		tests[3].a = 4
		c.Next()
	})
	p.Use(func(c *Context) {
		tests[4].a = 5
		c.Next()
	})
	p.Do(nil)

	for i := range tests {
		if tests[i].a != tests[i].e {
			t.Errorf("the %d handler fail. want %d got %d", i, tests[i].e, tests[i].a)
		}
	}
}

func BenchmarkProcesser(b *testing.B) {
	var p IProcesser = &Processer{}

	for i := 0; i < 5; i++ {
		p.Use(func(c *Context) {
			c.Event = nil
			c.Next()
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Do(nil)
	}
}
