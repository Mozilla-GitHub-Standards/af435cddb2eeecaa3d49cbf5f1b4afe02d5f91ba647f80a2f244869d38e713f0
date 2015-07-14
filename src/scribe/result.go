// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Contributor:
// - Aaron Meihm ameihm@mozilla.com

package scribe

import (
	"fmt"
	"strings"
)

type TestResult struct {
	// The name of the test.
	Name string `json:"name"`

	// True if error encountered during evaluation.
	IsError bool `json:"iserror"`
	// Error associated with test.
	Error string `json:"error"`

	// Master result of test.
	MasterResult bool `json:"masterresult"`
	// True if some source evaluations resulted in the test being true.
	HasTrueResults bool `json:"hastrueresults"`

	Results []TestSubResult `json:"results"`
}

type TestSubResult struct {
	// The result of evaluation for an identifier source.
	Result bool `json:"result"`
	// The identifier for the source.
	Identifier string `json:"identifier"`
}

// Return test results for a given test. Returns an error if an
// error occured during test preparation or execution.
func GetResults(d Document, name string) (TestResult, error) {
	t, err := d.getTest(name)
	if err != nil {
		return TestResult{}, err
	}
	if t.err != nil {
		return TestResult{}, t.err
	}
	ret := TestResult{}
	ret.Name = t.Name
	if t.err != nil {
		ret.Error = fmt.Sprintf("%v", t.err)
	}
	ret.MasterResult = t.masterResult
	ret.HasTrueResults = t.hasTrueResults
	for _, x := range t.results {
		nr := TestSubResult{}
		nr.Result = x.result
		nr.Identifier = x.criteria.identifier
		ret.Results = append(ret.Results, nr)
	}
	return ret, nil
}

// A helper function to convert TestResult r into a human readable result
// suitable for display.
func HumanResult(r TestResult) string {
	lns := make([]string, 0)
	lns = append(lns, fmt.Sprintf("result for \"%v\"", r.Name))
	if r.MasterResult {
		lns = append(lns, "\tmaster result: true")
	} else {
		buf := "\tmaster result: false"
		if r.HasTrueResults {
			buf = buf + ", has true results, failure caused by dependency"
		}
		lns = append(lns, buf)
	}
	if r.Error != "" {
		buf := fmt.Sprintf("[error] error: %v", r.Error)
		lns = append(lns, buf)
	}
	for _, x := range r.Results {
		buf := fmt.Sprintf("\t[%v] identifier: \"%v\"", x.Result, x.Identifier)
		lns = append(lns, buf)
	}
	return strings.Join(lns, "\n") + "\n"
}
