package cmd

import "errors"

var errConfigNotFound = errors.New("could not find a konnect.yml configuration file")
var errNoHostSelected = errors.New("no host was selected")
