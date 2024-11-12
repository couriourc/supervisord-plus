package config

import "strings"

var (
	GroupNamespace         = "group:"
	ProgramNamespace       = "program:"
	EventListenerNamespace = "eventlistener:"
	UpdaterNamespace       = "updater:"
)

// IsProgram return true if this section is for event listener
func IsProgram(c string) bool {
	return strings.HasPrefix(c, ProgramNamespace)
}

// IsEventListener return true if this section is for event listener
func IsEventListener(c string) bool {
	return strings.HasPrefix(c, EventListenerNamespace)
}

// IsUpdater return true if this section is for updater
func IsUpdater(c string) bool {
	return strings.HasPrefix(c, UpdaterNamespace)
}

// IsGroup return true if this section is for updater
func IsGroup(c string) bool {
	return strings.HasPrefix(c, GroupNamespace)
}
