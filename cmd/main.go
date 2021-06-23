package main

import "github.com/fastrodev/serverless/internal"

func main() {
	internal.CreateApp().Listen(9000)
}
