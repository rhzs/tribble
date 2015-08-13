package main

import (
	"github.com/mitchellh/cli"
	"tribble/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"stat": func() (cli.Command, error) {
			return &command.StatCommand{
				Meta: *meta,
			}, nil
		},
		"cpu": func() (cli.Command, error) {
			return &command.CpuCommand{
				Meta: *meta,
			}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name: Name,
			}, nil
		},
	}
}
