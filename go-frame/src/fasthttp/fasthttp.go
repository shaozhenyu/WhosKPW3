package main

import (
	"log"
	"sync"
	"strings"

	"github.com/valyala/fasthttp"
)

var m *sync.RWMutex

func main() {

	m = new(sync.RWMutex)

	var cache = make(map[string]string, 10000)
	//cache := map[string]string{}

	requestHandler := func(ctx *fasthttp.RequestCtx) {

		switch string(ctx.Path()) {
		case "/get":
			func(ctx *fasthttp.RequestCtx) {
				m.RLock()

				args := ctx.QueryArgs().String()
				s := strings.Split(args, "&")
				key := s[0][4:]
				if v, ok := cache[key]; ok {
					str := "value=" + v
					ctx.WriteString(str)
				} else {
					ctx.WriteString("unknow error")
				}
				m.RUnlock()
			}(ctx)

		case "/set":
			func(ctx *fasthttp.RequestCtx) {
				m.Lock()
				args := ctx.QueryArgs().String()
				s := strings.Split(args, "&")
				key := s[0][4:]
				value := s[1][6:]
				cache[key] = value
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
