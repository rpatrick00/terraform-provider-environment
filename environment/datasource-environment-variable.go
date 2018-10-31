package environment

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"os"
)

func dataSourceEnvironmentVariable() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEnvironmentVariableRead,

		Schema: map[string]*schema.Schema {
			"name": &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default": &schema.Schema {
				Description: "The default value to return if the variable value is empty",
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
				Default:  "",
			},
			"fail_if_empty": &schema.Schema {
				Description:
					"If true, an error will be generated if the variable value and its default value are empty",
				Type:     schema.TypeBool,
				Optional: true,
				Computed: false,
				Default:  false,
			},
		},
	}
}

func dataSourceEnvironmentVariableRead(d *schema.ResourceData, _ interface{}) error {
	name := d.Get("name").(string)
	if name != "" {
		d.SetId(name)
		value := os.Getenv(name)
		if value == "" {
			defaultValue := d.Get("default").(string)
			log.Printf("[INFO] env var %#v default value = %#v", name, defaultValue)
			failIfEmpty := d.Get("fail_if_empty").(bool)
			if defaultValue == "" && failIfEmpty {
				return fmt.Errorf("the environment variable %v value was empty", name)
			} else {
				d.Set("value", defaultValue)
			}
		} else {
			log.Printf("[INFO] Setting %#v = %#v", name, value)
			d.Set("value", value)
		}
	} else {
		return fmt.Errorf("the environment variable name was not specified")
	}
	return nil
}
