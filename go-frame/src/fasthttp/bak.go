package main

import (
	"log"
	"sync"
	"runtime"

	"github.com/valyala/fasthttp"
)

var m *sync.RWMutex

func main() {

	runtime.GOMAXPROCS(2)

	m = new(sync.RWMutex)

	//cache := map[string]string{}
	var cache = make(map[string]string, 1000000)

	requestHandler := func(ctx *fasthttp.RequestCtx) {

		switch string(ctx.Path()) {
		case "/get":
			func(ctx *fasthttp.RequestCtx) {
				m.RLock()

				b := fasthttp.AcquireByteBuffer()

				args := ctx.QueryArgs()
				key := args.PeekBytes([]byte("key"))
				if v, ok := cache[string(key)]; ok {
					str := "value=" + v
					b.WriteString(str)
					ctx.Write(b.B)
				} else {
					ctx.Write([]byte("unknow error"))
				}
				fasthttp.ReleaseByteBuffer(b)

				m.RUnlock()
			}(ctx)

		case "/set":
			func(ctx *fasthttp.RequestCtx) {
				m.Lock()

				b := fasthttp.AcquireByteBuffer()
				args := ctx.QueryArgs()
				key := args.PeekBytes([]byte("key"))
				value := args.PeekBytes([]byte("value"))
				cache[string(key)] = string(value)
				b.WriteString("ok")
				ctx.Write(b.B)
				fasthttp.ReleaseByteBuffer(b)

				m.Unlock()
			}(ctx)

		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}

	if err := fasthttp.ListenAndServe("127.0.0.1:8080", requestHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
