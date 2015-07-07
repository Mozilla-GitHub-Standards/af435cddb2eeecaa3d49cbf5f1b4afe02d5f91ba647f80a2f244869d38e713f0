// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Contributor:
// - Aaron Meihm ameihm@mozilla.com
package main

import (
	"flag"
	"fmt"
	"os"
	"scribe"
)

var flagDebug bool

func showTextResults(doc scribe.Document) {
	for _, x := range doc.Tests {
		res, err := x.GetResults()
		fmt.Fprintf(os.Stdout, "result %v \"%v\"\n", x.Identifier, x.Name)
		if err != nil {
			fmt.Fprintf(os.Stdout, "\t[error] error: %v\n", err)
		} else {
			if len(res) == 0 {
				fmt.Fprintf(os.Stdout, "\t[false] no candidates found\n")
				continue
			}
			for _, y := range res {
				if y.Result {
					fmt.Fprintf(os.Stdout, "\t[true]")
				} else {
					fmt.Fprintf(os.Stdout, "\t[false]")
				}
				fmt.Fprintf(os.Stdout, " identifier: \"%v\"", y.Criteria.Identifier)
				fmt.Fprintf(os.Stdout, "\n")
			}
		}
	}
}

func main() {
	err := scribe.Bootstrap()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	var docpath string
	flag.BoolVar(&flagDebug, "d", false, "enable debugging")
	flag.StringVar(&docpath, "f", "", "path to document")
	flag.Parse()

	if flagDebug {
		scribe.SetDebug(true, os.Stderr)
	}

	if docpath == "" {
		fmt.Fprintf(os.Stderr, "error: must specify document path\n")
		os.Exit(1)
	}

	fd, err := os.Open(docpath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer fd.Close()
	doc, err := scribe.LoadDocument(fd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	err = scribe.AnalyzeDocument(doc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	showTextResults(doc)

	os.Exit(0)
}
