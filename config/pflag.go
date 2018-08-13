package config

import (
	"github.com/amy911/amy911/onfail"
	"github.com/amy911/env911/app"

	"github.com/ogier/pflag"
)

type PFlag struct {
	pflag.FlagSet
}

func (pflag PFlag) Count(name string, usage string) {
	panic("github.com/ogier/pflag does not support the Count* functions")
}

func (pflag PFlag) CountP(name, shorthand string, usage string) {
	panic("github.com/ogier/pflag does not support the Count* functions")
}

func (pflag PFlag) CountVar(p *int, name string, usage string) {
	panic("github.com/ogier/pflag does not support the Count* functions")
}

func (pflag PFlag) CountVarP(p *int, name, shorthand string, usage string) {
	panic("github.com/ogier/pflag does not support the Count* functions")
}

// TODO: make my own CC0-1.0-licensed POSIX-compliant flag replacement library that does