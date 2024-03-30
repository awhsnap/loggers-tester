package main

import awhlog "github.com/awhsnap/loggers"

func main() {
	awhlog.LogInfo("Testing INFO message")
	awhlog.LogError("Testing ERROR message")
	awhlog.LogWarn("Testing WARN message")
}
