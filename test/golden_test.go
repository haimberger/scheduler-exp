package test

import (
	"fmt"
	"testing"
)

func TestLoadInput(t *testing.T) {
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
		if err := testLoadInput(tc.in, tc.out, tc.expectErr); err != nil {
			t.Errorf("[%s] %v", tc.name, err)
		}
	}
}

func testLoadInput(in, out string, expectErr bool) error {
	var v map[string]string
	err := LoadInput(in, &v)
	if expectErr {
		if err == nil {
			return fmt.Errorf("expected an error, but didn't get one; got %v instead", v)
		}
		return nil
	}
	if err != nil {
		return err
	}
	return CompareResults(&v, out)
}
