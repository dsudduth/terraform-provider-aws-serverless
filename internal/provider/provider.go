package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"aws_serverless_lambda": resourceServerlessLambda(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}
