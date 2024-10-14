package main

import (
	"gindemo/routers"
)

func main() {
	r := routers.SetupRouter()
	r.Run(":1234")
}
