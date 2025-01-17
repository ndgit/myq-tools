package viewer

import (
	"fmt"
	"testing"

	"github.com/jayjanssen/myq-tools2/loader"
)

// Funcs to get some test columns
func getTestCol() defaultCol {
	sources := []loader.SourceName{"status"}
	return defaultCol{
		Name:        "cons",
		Description: "Connections per second",
		Sources:     sources,
		Length:      4,
	}
}

func TestColGetShortHelp(t *testing.T) {
	col := getTestCol()

	help := col.GetShortHelp()
	if help != "cons: Connections per second" {
		t.Error("Bad short help!")
	}
}

func TestColGetDetailedHelp(t *testing.T) {
	col := getTestCol()

	help := col.GetDetailedHelp()
	if len(help) != 1 {
		t.Errorf("detailed help unexpected line count: %d", len(help))
	}

	if help[0] != "cons: Connections per second" {
		t.Error("bad detailed help")
	}
}

func TestColGetSources(t *testing.T) {
	loader.LoadDefaultSources()

	col := getTestCol()
	sources, err := col.GetSources()

	if err != nil {
		t.Error(err)
	}

	fmt.Printf("sources: %v\n", sources)

	if len(sources) != 1 {
		t.Errorf("Got the wrong number of sources: %d", len(sources))
	}
}

func getTestState() loader.StateReader {
	sp := loader.NewState()
	prevss := loader.NewSampleSet()

	cursamp := loader.NewSample()
	sp.GetCurrentWriter().SetSample(`status`, cursamp)

	prevsamp := loader.NewSample()
	prevss.SetSample(`status`, prevsamp)

	sp.SetPrevious(prevss)

	cursamp.Data[`connections`] = `105`
	prevsamp.Data[`connections`] = `100`

	return sp
}

func TestColGetHeader(t *testing.T) {
	col := getTestCol()
	state := getTestState()
	headers := col.GetHeader(state)

	// Expect one line header
	if len(headers) != 1 {
		t.Errorf("Header more than 1 line: %d", len(headers))
	}

	header := headers[0]
	if len(header) != col.Length {
		t.Errorf("Got header of length: %d, expected: %d", len(header), col.Length)
	}

	if header != "cons" {
		t.Errorf("Expected header to be 'cons', not: %s", header)
	}

	// Test too long name
	col.Name = "consss"
	headers = col.GetHeader(state)
	// Expect one line header
	if len(headers) != 1 {
		t.Errorf("Header more than 1 line: %d", len(headers))
	}

	header = headers[0]
	if len(header) != col.Length {
		t.Errorf("Got header of length: %d, expected: %d", len(header), col.Length)
	}

	if header != "cons" {
		t.Errorf("Expected header to be 'cons', not: %s", header)
	}

}

func TestColGetBlankLine(t *testing.T) {
	col := getTestCol()
	line := col.GetBlankLine()
	if line != `    ` {
		t.Errorf(`unexpected blank line: '%s'`, line)
	}
}
