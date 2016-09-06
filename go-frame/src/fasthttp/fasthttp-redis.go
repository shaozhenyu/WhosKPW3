package main

import (
	"log"
	"sync"

	"libs/cache"

	"github.com/valyala/fasthttp"
)

var m *sync.RWMutex

func main() {

	m = new(sync.RWMutex)

	RedisHost := "127.0.0.1:6379"
	rds, err := cache.New(RedisHost, 0, 100)
	if err != nil {
		return
	}

	requestHandler := func(ctx *fasthttp.RequestCtx) {

		switch string(ctx.Path()) {
		case "/get":
			func(ctx *fasthttp.RequestCtx) {
				m.RLock()

				args := ctx.QueryArgs()
				key :=  args.Peek("key")
				value, err := rds.Get(string(key))
				if err != nil {
					ctx.WriteString("unknow error")
					m.RUnlock()
					return
				}
				str := ""
				str = "value=" + string(value)
				ctx.WriteString(str)
				m.RUnlock()
			}(ctx)

		case "/set":
			func(ctx *fasthttp.RequestCtx) {
				m.Lock()
				args := ctx.QueryArgs()
				key :=  args.Peek("key")
				value := args.Peek("value")
				if err := rds.Set(string(key), value); err != nil {
					ctx.WriteString("unknow error")
					m.Unlock()
					return
				}
				ctx.WriteString("ok")
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
