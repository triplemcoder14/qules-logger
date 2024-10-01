package main

import (
	"sync"

	"github.com/triplemcoder14/qules-logger/client/internal/server"
	"github.com/triplemcoder14/qules-logger/internal/config"
)

func main() {
	config := config.GetClient()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server.Run(":" + config.Port)
		wg.Done()
	}()
	wg.Wait()
}
