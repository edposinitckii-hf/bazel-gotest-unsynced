package main

import "github.com/edposinitckii-hf/bazel-gotest-unsynced/lib/log"

func main() {
	l, err := log.NewLogger()
	if err != nil {
		panic(err)
	}

	l.Info("sample app")
}
