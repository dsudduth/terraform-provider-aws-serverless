package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
)

func resourceServerlessLambda() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServerlessLambdaCreate,
		ReadContext:   resourceServerlessLambdaRead,
		UpdateContext: resourceServerlessLambdaUpdate,
		DeleteContext: resourceServerlessLambdaDelete,
		Schema: map[string]*schema.Schema{
			"source": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Source of the Lambda function to package.",
			},
			"output_path": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The output of the archive file. If  not specified, a folder will be created.",
			},
		},
	}
}

func resourceServerlessLambdaCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warnings and errors are collected in a slice type
	var diags diag.Diagnostics
	err := os.MkdirAll(d.Get("output_path").(string), 0755)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to create output directory",
			Detail:   err.Error(),
		})
	}

	return diags
}

func resourceServerlessLambdaRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warnings and errors are collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourceServerlessLambdaUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warnings and errors are collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourceServerlessLambdaDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warnings and errors are collected in a slice type
	var diags diag.Diagnostics

	return diags
}
