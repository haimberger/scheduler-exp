package test

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/haimberger/compare"
)

var update = flag.Bool("update", false, "update .golden files")

// LoadInput loads input data from a file.
func LoadInput(inputFile string, i interface{}) error {
	// construct path to input file
	inputPath := path.Join("testdata", fmt.Sprintf("%s.json", inputFile))

	// read input file
	input, err := ioutil.ReadFile(inputPath)
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
	jsonStr, err := json.MarshalIndent(actual, "", "  ")
	if err != nil {
		return err
	}
	jsonStr = []byte(fmt.Sprintf("%s\n", jsonStr))

	// construct path to golden file
	goldenPath := path.Join("testdata", fmt.Sprintf("%s.json", goldenFile))

	// read golden file, which contains the expected output
	expected, err := ioutil.ReadFile(goldenPath)
	if err != nil {
		return err
	}

	// check if the actual result matches the contents of the golden file
	jd := compare.JSONDiffer{BasicEqualer: compare.TolerantBasicEqualer{}}
	diff, err := jd.Compare(jsonStr, expected)
	if err != nil {
		return err
	}
	if diff.Modified() {
		if *update {
			// overwrite the golden file with the actual result
			return ioutil.WriteFile(goldenPath, jsonStr, 0666)
		}
		diffStr, err := diff.Format(false)
		if err != nil {
			return err
		}
		return fmt.Errorf("result doesn't match expected value\n%s", diffStr)
	}
	return nil
}
