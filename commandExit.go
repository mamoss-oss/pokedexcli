package main

import "os"

func commandExit(*config) error {
	os.Exit(0)
	return nil
}
