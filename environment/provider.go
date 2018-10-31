package environment

import 	"github.com/hashicorp/terraform/helper/schema"

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},

		ResourcesMap: map[string]*schema.Resource{},

		DataSourcesMap: map[string]*schema.Resource{
			"environment_variable": dataSourceEnvironmentVariable(),
		},
	}
}

