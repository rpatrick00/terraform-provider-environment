/*
 * Copyright 2018 Robert Patrick <rhpatrick@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package environment

import (
	"testing"
)

func TestGetEnvironmentVariableValue_VariableSet(t *testing.T) {
	name := "HOME" // POSIX-defined variable so should always be defined...

	testResult, err := getEnvironmentVariableValue(name, "", false)
	if err != nil {
		t.Errorf("getEnvironmentVariableValue(%v, \"\", false) returned an error: %v", name, err)
	} else if len(testResult) == 0 {
		t.Errorf("getEnvironmentVariableValue(%v, \"\", false) returned an empty value", name)
	}
}

func TestGetEnvironmentVariableValue_VariableSetDefault(t *testing.T) {
	name := "HOME" // POSIX-defined variable so should always be defined...
	defaultValue := "/foo/bar/baz123456"

	testResult, err := getEnvironmentVariableValue(name, defaultValue, false)
	if err != nil {
		t.Errorf("getEnvironmentVariableValue(%v, %v, false) returned an error: %v", name, defaultValue, err)
	} else if len(testResult) == 0 {
		t.Errorf("getEnvironmentVariableValue(%v, %v, false) returned an empty value", name, defaultValue)
	} else if testResult == defaultValue {
		t.Errorf("getEnvironmentVariableValue(%v, %v, false) returned the default value", name, defaultValue)
	}
}

func TestGetEnvironmentVariableValue_VariableNotSet(t *testing.T) {
	name := "ABCDEFG" // Hopefully random enough it will never be set

	testResult, err := getEnvironmentVariableValue(name, "", false)
	if err != nil {
		t.Errorf("getEnvironmentVariableValue(%v, \"\", false) returned an error: %v", name, err)
	} else if len(testResult) != 0 {
		t.Errorf("getEnvironmentVariableValue(%v, \"\", false) returned an empty value", name)
	}
}

func TestGetEnvironmentVariableValue_VariableNotSetDefault(t *testing.T) {
	name := "ABCDEFG" // Hopefully random enough it will never be set
	defaultValue := "foobar"

	testResult, err := getEnvironmentVariableValue(name, defaultValue, false)
	if err != nil {
		t.Errorf("getEnvironmentVariableValue(%v, %v, false) returned an error: %v", name, defaultValue, err)
	} else if len(testResult) == 0 {
		t.Errorf("getEnvironmentVariableValue(%v, %v, false) returned an empty value", name, defaultValue)
	} else if testResult != defaultValue {
		t.Errorf("getEnvironmentVariableValue(%v, %v, false) did not return the default value", name, defaultValue)
	}
}

func TestGetEnvironmentVariableValue_VariableNotSetDefaultFail(t *testing.T) {
	name := "ABCDEFG" // Hopefully random enough it will never be set
	defaultValue := "foobar"

	testResult, err := getEnvironmentVariableValue(name, defaultValue, true)
	if err != nil {
		t.Errorf("getEnvironmentVariableValue(%v, %v, true) returned an error: %v", name, defaultValue, err)
	} else if len(testResult) == 0 {
		t.Errorf("getEnvironmentVariableValue(%v, %v, true) returned an empty value", name, defaultValue)
	} else if testResult != defaultValue {
		t.Errorf("getEnvironmentVariableValue(%v, %v, false) did not return the default value", name, defaultValue)
	}
}

func TestGetEnvironmentVariableValue_VariableNotSetFail(t *testing.T) {
	name := "ABCDEFG" // Hopefully random enough it will never be set

	_, err := getEnvironmentVariableValue(name, "", true)
	if err == nil {
		t.Errorf("getEnvironmentVariableValue(%v, \"\", true) did not return an error", name)
	}
}

func TestReplaceUnquotedBackslashes_BackslashOnEnd(t *testing.T) {
	testString := "c:\\"
	expectedResultString := "c:\\\\"

	testResult := replaceUnquotedBackslashes(testString)
	if testResult != expectedResultString {
		t.Errorf("Input value %v returned %v instead of %v", testString, testResult, expectedResultString)
	}
}

func TestReplaceUnquotedBackslashes_QuotedBackslashOnEnd(t *testing.T) {
	testString := "c:\\\\"
	expectedResultString := "c:\\\\"

	testResult := replaceUnquotedBackslashes(testString)
	if testResult != expectedResultString {
		t.Errorf("Input value %v returned %v instead of %v", testString, testResult, expectedResultString)
	}
}

func TestReplaceUnquotedBackslashes_BackslashAtStart(t *testing.T) {
	testString := "\\foo"
	expectedResultString := "\\\\foo"

	testResult := replaceUnquotedBackslashes(testString)
	if testResult != expectedResultString {
		t.Errorf("Input value %v returned %v instead of %v", testString, testResult, expectedResultString)
	}
}

func TestReplaceUnquotedBackslashes_QuotedBackslashAtStart(t *testing.T) {
	testString := "\\\\foo"
	expectedResultString := "\\\\foo"

	testResult := replaceUnquotedBackslashes(testString)
	if testResult != expectedResultString {
		t.Errorf("Input value %v returned %v instead of %v", testString, testResult, expectedResultString)
	}
}

func TestReplaceUnquotedBackslashes_Backslashs(t *testing.T) {
	testString := "c:\\foo\\bar\\baz"
	expectedResultString := "c:\\\\foo\\\\bar\\\\baz"

	testResult := replaceUnquotedBackslashes(testString)
	if testResult != expectedResultString {
		t.Errorf("Input value %v returned %v instead of %v", testString, testResult, expectedResultString)
	}
}

func TestReplaceUnquotedBackslashes_QuotedBackslashs(t *testing.T) {
	testString := "c:\\\\foo\\\\bar\\\\baz"
	expectedResultString := "c:\\\\foo\\\\bar\\\\baz"

	testResult := replaceUnquotedBackslashes(testString)
	if testResult != expectedResultString {
		t.Errorf("Input value %v returned %v instead of %v", testString, testResult, expectedResultString)
	}
}

func TestReplaceUnquotedBackslashes_MixedBackslashs(t *testing.T) {
	testString := "c:\\foo\\\\bar\\baz"
	expectedResultString := "c:\\\\foo\\\\bar\\\\baz"

	testResult := replaceUnquotedBackslashes(testString)
	if testResult != expectedResultString {
		t.Errorf("Input value %v returned %v instead of %v", testString, testResult, expectedResultString)
	}
}
