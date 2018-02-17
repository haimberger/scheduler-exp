package test

import "testing"

func TestLoadData(t *testing.T) {
	type testCase struct {
		name       string
		inputFile  string
		goldenFile string
		expectErr  bool
	}
	tcs := []testCase{
		{name: "basic", inputFile: "testdata/basic.input", goldenFile: "testdata/LoadInput.golden"},
		{name: "misformatted", inputFile: "testdata/misformatted.input", expectErr: true},
		{name: "nonexistent", inputFile: "testdata/nonexistent.input", expectErr: true},
	}
	for _, tc := range tcs {
		var v map[string]string
		err := LoadInput(tc.inputFile, &v)
		if tc.expectErr {
			if err == nil {
				t.Errorf("[%s] expected an error, but didn't get one; got %v instead", tc.name, v)
			}
		} else if err != nil {
			t.Errorf("[%s] got error: %v", tc.name, err)
		} else if err = CompareResults(v, tc.goldenFile); err != nil {
			t.Errorf("[%s] got error: %v", tc.name, err)
		}
	}
}
