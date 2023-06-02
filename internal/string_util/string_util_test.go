package string_util

import (
	"fmt"
	"testing"
)

func TestToSnakeCase(t *testing.T) {
	testCases := []string{
		"JapanCanadaAustralia",
		"japan-canada_australia",
		"GE Café ™ 36\" Built-In Touch Control Induction Cooktop With 5 Elements",
		"Arizona 10¼” Inseam Flat-Front Shorts",
	}
	results := []string{
		"japan_canada_australia",
		"japan_canada_australia",
		"ge_caf_36_built_in_touch_control_induction_cooktop_with_5_elements",
		"arizona_10_inseam_flat_front_shorts",
	}
	for index, testCase := range testCases {
		result := ToSnakeCase(testCase, true)

		if result == results[index] {
			t.Logf(fmt.Sprintf("%s transform to %s", testCase, results[index]))
		} else {
			t.Logf("%s expected but get %s", results[index], result)
			t.Error(fmt.Sprintf("%s fail to transformed", testCase))
		}
	}
}

func TestToKebabCase(t *testing.T) {
	testCases := []string{
		"JapanCanadaAustralia",
		"japan-canada_australia",
		"GE Café ™ 36\" Built-In Touch Control Induction Cooktop With 5 Elements",
		"Arizona 10¼” Inseam Flat-Front Shorts",
		"Test Variant 1::TestVariant2::Testvariant 3",
	}
	results := []string{
		"japan-canada-australia",
		"japan-canada-australia",
		"ge-caf-36-built-in-touch-control-induction-cooktop-with-5-elements",
		"arizona-10-inseam-flat-front-shorts",
		"test-variant-1-test-variant2-testvariant-3",
	}
	for index, testCase := range testCases {
		result := ToKebabCase(testCase, true)

		if result == results[index] {
			t.Logf(fmt.Sprintf("%s transform to %s", testCase, results[index]))
		} else {
			t.Logf("%s expected but get %s", results[index], result)
			t.Error(fmt.Sprintf("%s fail to transformed", testCase))
		}
	}
}
