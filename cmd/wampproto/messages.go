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

type Register struct {
	register     *kingpin.CmdClause
	regRequestID *int64
	regProcedure *string
	regOptions   *map[string]string
}

type Registered struct {
	registered          *kingpin.CmdClause
	registeredRequestID *int64
	registrationID      *int64
}

type Invocation struct {
	invocation        *kingpin.CmdClause
	invRequestID      *int64
	invRegistrationID *int64
	invDetails        *map[string]string
	invArgs           *[]string
	invKwArgs         *map[string]string
}

type Yield struct {
	yield          *kingpin.CmdClause
	yieldRequestID *int64
	yieldOptions   *map[string]string
	yieldArgs      *[]string
	yieldKwArgs    *map[string]string
}
