package provider

// func resourceScaffolding() *schema.Resource {
// 	return &schema.Resource{
// 		// This description is used by the documentation generator and the language server.
// 		Description: "Sample resource in the Terraform provider scaffolding.",

// 		CreateContext: resourceScaffoldingCreate,
// 		ReadContext:   resourceScaffoldingRead,
// 		UpdateContext: resourceScaffoldingUpdate,
// 		DeleteContext: resourceScaffoldingDelete,

// 		Schema: map[string]*schema.Schema{
// 			"sample_attribute": {
// 				// This description is used by the documentation generator and the language server.
// 				Description: "Sample attribute.",
// 				Type:        schema.TypeString,
// 				Optional:    true,
// 			},
// 		},
// 	}
// }

// func resourceScaffoldingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	// use the meta value to retrieve your client from the provider configure method
// 	// client := meta.(*apiClient)

// 	idFromAPI := "my-id"
// 	d.SetId(idFromAPI)

// 	// write logs using the tflog package
// 	// see https://pkg.go.dev/github.com/hashicorp/terraform-plugin-log/tflog
// 	// for more information
// 	tflog.Trace(ctx, "created a resource")

// 	return diag.Errorf("not implemented")
// }

// func resourceScaffoldingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	// use the meta value to retrieve your client from the provider configure method
// 	// client := meta.(*apiClient)

// 	return diag.Errorf("not implemented")
// }

// func resourceScaffoldingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	// use the meta value to retrieve your client from the provider configure method
// 	// client := meta.(*apiClient)

// 	return diag.Errorf("not implemented")
// }

// func resourceScaffoldingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	// use the meta value to retrieve your client from the provider configure method
// 	// client := meta.(*apiClient)

// 	return diag.Errorf("not implemented")
// }
