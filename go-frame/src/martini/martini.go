package main

import (
	"sync"
	"net/http"

	"github.com/go-martini/martini"
)

var mutex *sync.RWMutex

func main() {
	m := martini.Classic()

	mutex = new(sync.RWMutex)

	cache := map[string]string{}

	m.Get("/get", func(req *http.Request) (int, string) {
		mutex.RLock()
		key := req.URL.Query().Get("key")

		if value, ok := cache[key]; ok {
			str := "value=" + value
			mutex.RUnlock()
			return 200, str
		} else {
			mutex.RUnlock()
			return 400, "key not found"
		}
	})

	m.Get("/set", func(req *http.Request) (int, string) {
		mutex.Lock()
		key := req.URL.Query().Get("key")
		value := req.URL.Query().Get("value")
		cache[key] = value
		mutex.Unlock()
		return 200, "ok"
	})

	m.RunOnAddr(":8080")
}

