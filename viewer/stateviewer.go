package viewer

import (
	"github.com/jayjanssen/myq-tools2/loader"
)

// A StateViewer represents the output of data from a State into a (usually) constrained width with a header and one or more lines of output per State
type StateViewer interface {
	// Get name of the view
	GetName() string

	// Single line help for this view
	GetShortHelp() string

	// A list of sources that this view requires
	GetSources() ([]*loader.Source, error)

	// Header for this view, unclear if state is needed
	GetHeader(loader.StateReader) []string

	// Data for this view based on the state
	GetData(loader.StateReader) []string
}
