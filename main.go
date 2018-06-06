package main

import (
	"github.com/secmsg/go-api/api"
)

func main() {
	api.Method("/dev/hello", func(ctx *api.Ctx) {
		ctx.String("Hello!")
	})
	api.Serve(8000)
}
