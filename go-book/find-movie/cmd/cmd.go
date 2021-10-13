package cmd

import "github.com/cosiner/flag"

type PosterCmd struct {
	Terms []string `args:"true" important:"1" desc:"Movie search terms"`
}

func (pc *PosterCmd) Metadata() map[string]flag.Flag {
	const (
		usage   = "poster is a tool for querying movie data from an internet database"
		version = "v0.0.1"
		desc    = "This tool is under development, documentation will follow"
	)

	return map[string]flag.Flag{
		"": {
			Usage:   usage,
			Version: version,
			Desc:    desc,
			Arglist: "[SearchTerms]",
		},
	}
}
