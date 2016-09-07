package main

import (
	"log"
	"sync"
	"strings"

	"github.com/valyala/fasthttp"
)

func main() {

	var counter = struct{
		sync.RWMutex
		m map[string]string
	}{m: make(map[string]string, 10000)}

	//m = new(sync.RWMutex)

	//var cache = make(map[string]string, 10000)
	//cache := map[string]string{}

	requestHandler := func(ctx *fasthttp.RequestCtx) {

		switch string(ctx.Path()) {
		case "/get":
			func(ctx *fasthttp.RequestCtx) {
				counter.RLock()

				args := ctx.QueryArgs().String()
				s := strings.Split(args, "&")
				key := s[0][4:]
				if v, ok := counter.m[key]; ok {
					str := "value=" + v
					ctx.WriteString(str)
				} else {
					ctx.WriteString("unknow error")
				}
				counter.RUnlock()
			}(ctx)

		case "/set":
			func(ctx *fasthttp.RequestCtx) {
				counter.Lock()
				args := ctx.QueryArgs().String()
				s := strings.Split(args, "&")
				key := s[0][4:]
				value := s[1][6:]
				counter.m[key] = value
				ctx.WriteString("ok")
				counter.Unlock()
			}(ctx)

		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}

	if err := fasthttp.ListenAndServe("127.0.0.1:8080", requestHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
