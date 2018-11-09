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

const knownName string = "HOME"      // POSIX-defined variable so should always be defined...
const unknownName string = "ABCDEFG" // Hopefully random enough it will never be set

const unquotedWindowsPathAtStart string = "\\foo"
const quotedWindowsPathAtStart string = "\\\\foo"
const unquotedWindowsPathAtEnd string = "c:\\"
const quotedWindowsPathAtEnd string = "c:\\\\"
const unquotedWindowsPath string = "c:\\foo\\bar\\baz"
const quotedWindowsPath string = "c:\\\\foo\\\\bar\\\\baz"
const unquotedMixedWindowsPath string = "c:\\foo\\\\bar\\baz"
const quotedMixedWindowsPath string = "c:\\\\foo\\\\bar\\\\baz"

const emptyString string = "\"\""

const errorUnexpectedErrorTemplate string = "getEnvironmentVariableValue(%v, %v, %v) returned an error: %v"
const errorUnexpectedEmptyValueTemplate string = "getEnvironmentVariableValue(%v, %v, %v) returned an empty value"
const errorUnexpectedResultsTemplate string = "Input value %v returned %v instead of %v"
const errorUnexpectedDefaultValueTemplate string = "getEnvironmentVariableValue(%v, %v, %v) returned the default value"

func TestGetEnvironmentVariableValueVariableSet(t *testing.T) {
	name := knownName

	testResult, err := getEnvironmentVariableValue(name, "", false)
	if err != nil {
		t.Errorf(errorUnexpectedErrorTemplate, name, emptyString, false, err)
	} else if len(testResult) == 0 {
		t.Errorf(errorUnexpectedEmptyValueTemplate, name, emptyString, false)
	}
}

func TestGetEnvironmentVariableValueVariableSetDefault(t *testing.T) {
	name := knownName
	defaultValue := "/foo/bar/baz123456"

	testResult, err := getEnvironmentVariableValue(name, defaultValue, false)
	if err != nil {
		t.Errorf(errorUnexpectedErrorTemplate, name, defaultValue, false, err)
	} else if len(testResult) == 0 {
		t.Errorf(errorUnexpectedEmptyValueTemplate, name, defaultValue, false)
	} else if testResult == defaultValue {
		t.Errorf(errorUnexpectedDefaultValueTemplate, name, defaultValue, false)
	}
}

func TestGetEnvironmentVariableValueVariableNotSet(t *testing.T) {
	name := unknownName

	testResult, err := getEnvironmentVariableValue(name, "", false)
	if err != nil {
		t.Errorf(errorUnexpectedErrorTemplate, name, emptyString, false, err)
	} else if len(testResult) != 0 {
		t.Errorf(errorUnexpectedEmptyValueTemplate, name, emptyString, false)
	}
}

func TestGetEnvironmentVariableValueVariableNotSetDefault(t *testing.T) {
	name := unknownName
	defaultValue := "testing123"

	testResult, err := getEnvironmentVariableValue(name, defaultValue, false)
	if err != nil {
		t.Errorf(errorUnexpectedErrorTemplate, name, defaultValue, false, err)
	} else if len(testResult) == 0 {
		t.Errorf(errorUnexpectedEmptyValueTemplate, name, defaultValue, false)
	} else if testResult != defaultValue {
		t.Errorf(errorUnexpectedDefaultValueTemplate, name, defaultValue, false)
	}
}

func TestGetEnvironmentVariableValueVariableNotSetDefaultFail(t *testing.T) {
	name := unknownName
	defaultValue := "foobar"

	testResult, err := getEnvironmentVariableValue(name, defaultValue, true)
	if err != nil {
		t.Errorf(errorUnexpectedErrorTemplate, name, defaultValue, true, err)
	} else if len(testResult) == 0 {
		t.Errorf(errorUnexpectedEmptyValueTemplate, name, defaultValue, true)
	} else if testResult != defaultValue {
		t.Errorf("getEnvironmentVariableValue(%v, %v, %v) did not return the default value", name, defaultValue, false)
	}
}

func TestGetEnvironmentVariableValueVariableNotSetFail(t *testing.T) {
	name := unknownName

	_, err := getEnvironmentVariableValue(name, "", true)
	if err == nil {
		t.Errorf("getEnvironmentVariableValue(%v, \"\", true) did not return an error", name)
	}
}

func TestReplaceUnquotedBackslashesBackslashOnEnd(t *testing.T) {
	testString := unquotedWindowsPathAtEnd
	expectedResultString := quotedWindowsPathAtEnd

	testResult := replaceUnquotedBackslashes(testString)
	if testResult != expectedResultString {
		t.Errorf(errorUnexpectedResultsTemplate, testString, testResult, expectedResultString)
	}
}

func TestReplaceUnquotedBackslashesQuotedBackslashOnEnd(t *testing.T) {
	testString := quotedWindowsPathAtEnd
	expectedResultString := quotedWindowsPathAtEnd

	testResult := replaceUnquotedBackslashes(testString)
	if testResult != expectedResultString {
		t.Errorf(errorUnexpectedResultsTemplate, testString, testResult, expectedResultString)
	}
}

func TestReplaceUnquotedBackslashesBackslashAtStart(t *testing.T) {
	testString := unquotedWindowsPathAtStart
	expectedResultString := quotedWindowsPathAtStart

	testResult := replaceUnquotedBackslashes(testString)
	if testResult != expectedResultString {
		t.Errorf(errorUnexpectedResultsTemplate, testString, testResult, expectedResultString)
	}
}

func TestReplaceUnquotedBackslashesQuotedBackslashAtStart(t *testing.T) {
	testString := quotedWindowsPathAtStart
	expectedResultString := quotedWindowsPathAtStart

	testResult := replaceUnquotedBackslashes(testString)
	if testResult != expectedResultString {
		t.Errorf(errorUnexpectedResultsTemplate, testString, testResult, expectedResultString)
	}
}

func TestReplaceUnquotedBackslashesBackslashs(t *testing.T) {
	testString := unquotedWindowsPath
	expectedResultString := quotedWindowsPath

	testResult := replaceUnquotedBackslashes(testString)
	if testResult != expectedResultString {
		t.Errorf(errorUnexpectedResultsTemplate, testString, testResult, expectedResultString)
	}
}

func TestReplaceUnquotedBackslashesQuotedBackslashs(t *testing.T) {
	testString := quotedWindowsPath
	expectedResultString := quotedWindowsPath

	testResult := replaceUnquotedBackslashes(testString)
	if testResult != expectedResultString {
		t.Errorf(errorUnexpectedResultsTemplate, testString, testResult, expectedResultString)
	}
}

func TestReplaceUnquotedBackslashesMixedBackslashs(t *testing.T) {
	testString := unquotedMixedWindowsPath
	expectedResultString := quotedMixedWindowsPath

	testResult := replaceUnquotedBackslashes(testString)
	if testResult != expectedResultString {
		t.Errorf(errorUnexpectedResultsTemplate, testString, testResult, expectedResultString)
	}
}
