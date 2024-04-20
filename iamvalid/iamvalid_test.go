package iamvalid

import (
	"reflect"
	"testing"
)

func TestVerifyJSON_ErrorsOnInvalid(t *testing.T) {
	cases := []string{
		"test_json/no_policy_name.json",
		"test_json/invalid_policy_name.json",
		"test_json/too_long_policy_name.json",
		"test_json/invalid_policy_version.json",
		"test_json/no_statement.json",
		"test_json/invalid_effect.json",
		"test_json/no_actions.json",
	}

	for _, tcase := range cases {
		t.Run(tcase, func(t *testing.T) {
			_, err := IsValid(tcase)
			if err == nil {
				t.Errorf("should get an error %s", err)
			}
		})
	}
}

func TestVerifyJSON_IsCorrectForValidInput(t *testing.T) {
	cases := []struct {
		file string
		want []bool
	}{
		{"test_json/valid_policy.json", []bool{true}},
		{"test_json/valid_policy_asterisk.json", []bool{false}},
		{"test_json/valid_policy_multiple_statements.json", []bool{true, true, false}},
	}

	for _, tcase := range cases {
		t.Run(tcase.file, func(t *testing.T) {
			got, err := IsValid(tcase.file)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, tcase.want) {
				t.Errorf("%s expected: %t but got: %t", tcase.file, tcase.want, got)
			}
		})
	}
}
