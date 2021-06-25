package main

import "github.com/fastrodev/serverless/internal"

func main() {
	internal.CreateApp(false).Listen(9000)
}
