package cmd

import "errors"

var errConfigNotFound = errors.New("could not find a konnect.yml configuration file")
var errNoHostSelected = errors.New("no host was selected")
var errHostRequired = errors.New("please specify one host")
var errHostsRequired = errors.New("please specify one or more hosts")
var errTaskRequired = errors.New("Please specify a task")
