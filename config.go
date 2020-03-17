package main

import (
	"flag"
	"time"
)

// Default URL for mirror list.
//
// This is set to the GitHub repository of archlinuxcn:
// https://github.com/archlinuxcn/mirrorlist-repo
const mirrorListURLDefault = "https://raw.githubusercontent.com/archlinuxcn/mirrorlist-repo/master/mirrors.yaml"

// Default update interval of the mirror list.
// The default value is 1 week (7 days).
const mirrorListUpdateIntervalMinutesDefault = 7 * 24 * time.Hour

// Default status check interval.
// The default value is 10 minutes.
const checkIntervalDefault = 10 * time.Minute

var mirrorListURL = flag.String(
	"-mirror-list-url",
	mirrorListURLDefault,
	"url for the mirror list",
)

var mirrorListUpdateIntervalMinutes = flag.Duration(
	"-list-update-interval",
	mirrorListUpdateIntervalMinutesDefault,
	"interval of updating mirror list",
)

var checkInterval = flag.Duration(
	"-check-interval",
	checkIntervalDefault,
	"interval of updating mirror statuses",
)
