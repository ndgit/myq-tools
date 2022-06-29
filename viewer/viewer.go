package viewer

import (
	"github.com/jayjanssen/myq-tools2/loader"
)

// A StateViewer represents the output of data from a State into a (usually) constrained width with a header and one or more lines of output per State
type StateViewer interface {
	// Single line help for the view
	GetShortHelp() string

	// Detailed multi-line help for the view
	GetHelp() []string

	// A list of sources that this view requires
	GetSources() []loader.Source

	// Header for this view, unclear if state is needed
	GetHeader(state *loader.State) []string

	// Data for this view based on the state
	GetData(state *loader.State) []string
}
