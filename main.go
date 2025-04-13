package main

import (
	router "github.com/yuki-yamamura/echo-with-openapi/internal/presentation"
)

func main() {
	r := router.NewRouter()
	r.Logger.Fatal(r.Start(":1323"))
}
