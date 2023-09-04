package main

import "os"

func commandExit(*config, ...string) error {
	os.Exit(0)
	return nil
}
