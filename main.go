package main

import (
	"time"

	"github.com/SALEH846/stream-large-files-over-tcp/stream"
)

func main() {
	// 1
	// Simple approach
	// ---------------------
	// go func() {
	// 	time.Sleep(4 * time.Second)
	// 	simple.SendFile(4000)
	// }()

	// server := &simple.FileServer{}
	// server.Start()

	// 2
	// Streaming approach
	// ----------------------
	go func() {
		time.Sleep(3 * time.Second)
		stream.SendFile(400000)
	}()

	server := &stream.FileServer{}
	server.Start()
}
