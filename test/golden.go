package test

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

var update = flag.Bool("update", false, "update .golden files")

// LoadInput loads input data from a file.
func LoadInput(inputFile string, i interface{}) error {
	// read input file
	input, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// unmarshal the contents of the input file
	if err = json.Unmarshal(input, i); err != nil {
		return err
	}

	return nil
}

// CompareResults compares actual test results with expected results,
// which are loaded from a .golden file. If the results are different
// and the update flag is set, the .golden file will be overwritten
// with the actual test results. If the flag isn't set, an error will
// be returned.
func CompareResults(actual interface{}, goldenFile string) error {
	// convert value to JSON string
	jsonStr, err := json.Marshal(actual)
	if err != nil {
		return err
	}

	// read golden file, which contains the expected output
	expected, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		return err
	}

	// check if the actual result matches the contents of the golden file
	if !bytes.Equal(jsonStr, expected) {
		if *update {
			// overwrite the golden file with the actual result
			if err = ioutil.WriteFile(goldenFile, jsonStr, 0666); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("expected %s, got %s", expected, actual)
		}
	}
	return nil
}
