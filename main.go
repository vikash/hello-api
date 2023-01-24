package main

import "github.com/vikash/gofr/pkg/gofr"

func main() {
	app := gofr.New()
	app.GET("/hello", func(c *gofr.Context) (interface{}, error) {
		return "hello world!", nil
	})
	app.Run()
}
