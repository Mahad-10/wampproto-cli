package main

import (
	"github.com/alecthomas/kingpin/v2"
)

type Call struct {
	call          *kingpin.CmdClause
	callRequestID *int64
	callURI       *string
	callArgs      *[]string
	callKwargs    *map[string]string
	callOption    *map[string]string
}