package environment

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"os"
	"strings"
)

func dataSourceEnvironmentVariable() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEnvironmentVariableRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default": &schema.Schema{
				Description: "The default value to return if the variable value is empty",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"fail_if_empty": &schema.Schema{
				Description: "If true, an error will be generated if the variable value and its default value are empty",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"normalize_file_path": &schema.Schema{
				Description: "Treat the value as a file system path and quote any backslash path separators, as needed to fix up Windows paths",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func dataSourceEnvironmentVariableRead(d *schema.ResourceData, _ interface{}) error {
	name := d.Get("name").(string)
	if name != "" {
		d.SetId(name)
		defaultValue := d.Get("default").(string)
		failIfEmpty := d.Get("fail_if_empty").(bool)
		value, err := getEnvironmentVariableValue(name, defaultValue, failIfEmpty)
		if err != nil {
			return err
		}

		normalizeFilePath := d.Get("normalize_file_path").(bool)
		// No need searching for backslashes in a path unless the platform is Windows...
		if os.PathSeparator == '\\' && normalizeFilePath {
			value = replaceUnquotedBackslashes(value)
		}
		d.Set("value", value)
	} else {
		return fmt.Errorf("the environment variable name was not specified")
	}
	return nil
}

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
func getEnvironmentVariableValue(name string, defaultValue string, failIfEmpty bool) (string, error) {
	variableValue := os.Getenv(name)
	if variableValue != "" {
		return variableValue, nil
	} else if defaultValue != "" {
		return defaultValue, nil
	} else if failIfEmpty {
		return "", fmt.Errorf("the environment variable %v value was empty", name)
	} else {
		return "", nil
	}
}

func replaceUnquotedBackslashes(path string) string {
	var newPath strings.Builder

	if !strings.Contains(path, "\\") {
		return path
	}

	length := len(path)
	var lastRuneWasQuote bool
	for pos, char := range path {
		if char != '\\' {
			newPath.WriteRune(char)
			lastRuneWasQuote = false
			continue
		}

		// If the last character written was a backslash, the backslash is already quoted.  As such,
		// write this backslash without quoting and reset lastRuneWasQuote
		if lastRuneWasQuote {
			newPath.WriteRune(char)
			lastRuneWasQuote = false
			continue
		}

		// If the next character is a backslash, treat this one as a quote.
		if pos+1 < length && path[pos+1] == '\\' {
			newPath.WriteRune(char)
			lastRuneWasQuote = true
			continue
		}

		// Otherwise, quote the backslash
		newPath.WriteRune(char)
		newPath.WriteRune(char)
		lastRuneWasQuote = false
	}
	return newPath.String()
}
