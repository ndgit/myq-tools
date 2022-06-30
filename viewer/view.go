package viewer

import "github.com/jayjanssen/myq-tools2/loader"

func (v View) GetName() string {
	return v.Name
}

// Single line help for the view
func (v View) GetShortHelp() string {
	return ""
}

// Detailed multi-line help for the view
func (v View) GetDetailedHelp() []string {
	return []string{""}
}

// A list of sources that this view requires
func (v View) GetSources() ([]*loader.Source, error) {
	return []*loader.Source{
		&loader.Source{},
	}, nil
}

// Header for this view, unclear if state is needed
func (v View) GetHeader(sr loader.StateReader) []string {
	return []string{""}
}

// Data for this view based on the state
func (v View) GetData(sr loader.StateReader) []string {
	return []string{""}
}

// Live views output the current timestamp whereas Runtime views output the delta in Uptime since the first state
func (v View) SetLive() {

}

func (v View) SetRuntime() {

}
