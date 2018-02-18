package test

import "testing"

func TestLoadData(t *testing.T) {
	type testCase struct {
		name      string
		in        string
		out       string
		expectErr bool
	}
	tcs := []testCase{
		{name: "basic", in: "basic", out: "basic"},
		{name: "misformatted", in: "misformatted", expectErr: true},
		{name: "nonexistent", in: "nonexistent", expectErr: true},
	}
	for _, tc := range tcs {
		var v map[string]string
		err := LoadInput(tc.in, &v)
		if tc.expectErr {
			if err == nil {
				t.Errorf("[%s] expected an error, but didn't get one; got %v instead", tc.name, v)
			}
		} else if err != nil {
			t.Errorf("[%s] got error: %v", tc.name, err)
		} else if err = CompareResults(v, tc.out); err != nil {
			t.Errorf("[%s] got error: %v", tc.name, err)
		}
	}
}
