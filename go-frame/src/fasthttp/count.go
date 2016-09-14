package main

import (
	"log"
	"sync"
	"runtime"

	"github.com/valyala/fasthttp"
	"github.com/valyala/bytebufferpool"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU() + 1)

	var counter = struct{
		sync.RWMutex
		m map[string]string
	}{m: make(map[string]string, 1000000)}

	requestHandler := func(ctx *fasthttp.RequestCtx) {

		b := bytebufferpool.Get()

		switch string(ctx.Path()) {
		case "/get":
			func(ctx *fasthttp.RequestCtx) {
				counter.RLock()

				args := ctx.QueryArgs()
				key := args.Peek("key")
				if v, ok := counter.m[string(key)]; ok {
					b.WriteString(v)
					ctx.Write(b.B)
				} else {
					ctx.WriteString("unknow error")
				}

				counter.RUnlock()
			}(ctx)

		case "/set":
			func(ctx *fasthttp.RequestCtx) {
				counter.Lock()

				args := ctx.QueryArgs()
				key := args.Peek("key")
				value := args.Peek("value")
				counter.m[string(key)] = string(value)

				b.WriteString("ok")
				ctx.Write(b.B)

				counter.Unlock()
			}(ctx)

		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}

		bytebufferpool.Put(b)
	}

	if err := fasthttp.ListenAndServe("127.0.0.1:3000", requestHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
