package main

import (
	"fmt"
	"net/http"
	_ "qor-test/config"
	"qor-test/config/admin"
	"qor-test/config/routes"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", routes.Router())

	admin.Admin.MountTo("/admin", mux)

	fmt.Printf("Listening on: %v\n", "1337")
	if err := http.ListenAndServe("localhost:1337", mux); err != nil {
		panic(err)
	}
}
