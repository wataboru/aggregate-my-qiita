// Package main
package main

import (
	"aggregate-my-qiita/aggreagtemyqiita"
	"flag"
	"fmt"
	"os"
)

const (
	// ExitCodeSuccess is the exit code on success
	ExitCodeSuccess int = iota
	// ExitCodeError is the exit code when failed
	ExitCodeError
	// ExitCodeFileError is the exit code when file conrtoll failed
	ExitCodeFileError
)

var (
	params aggreagtemyqiita.Params
)

func init() {
	if err := aggreagtemyqiita.GetParameters(&params); err != nil {
		fmt.Fprintf(os.Stderr, "[ERR]: %s", err)
		os.Exit(ExitCodeFileError)
	}

	flag.StringVar(&params.UserID, "user", params.UserID, "Input Qiita user id")
	flag.StringVar(&params.UserID, "u", params.UserID, "Input Qiita user id (short)")
	flag.StringVar(&params.Token, "token", params.Token, "Input Qiita API Token")
	flag.StringVar(&params.Token, "t", params.Token, "Input Qiita API Token (short)")
	flag.Parse()
}

func run() int {
	if params.UserID == "" || params.Token == "" {
		fmt.Fprintf(os.Stderr, "[ERR]: %s", "Input userID and Token.")
		return ExitCodeError
	}

	if err := aggreagtemyqiita.WriteParameters(&params); err != nil {
		fmt.Fprintf(os.Stderr, "[ERR]: %s", err)
		return ExitCodeFileError
	}

	err := aggreagtemyqiita.Aggregate(params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERR]: %s", err)
		return ExitCodeError
	}
	return ExitCodeSuccess
}

func main() {
	os.Exit(run())
}
