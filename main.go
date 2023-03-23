package main

import "belajar-gin/routers"

func main() {

	var PORT = ":4000"

	routers.StartServer().Run(PORT)
}
