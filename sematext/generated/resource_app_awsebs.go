package generated

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_App.go.template
	Then run generate/generate.sh
*/

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/invincibear/terraform-provider-sematext/sematext"
)

// resourceAppAwsebs is the resource class that handles sematext_app_awsebs
func resourceAppAwsebs() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaApp("AWS EBS"),
		CreateContext: resourceOperationCreateAppAwsebs,
		ReadContext:   resourceOperationReadAppAwsebs,
		UpdateContext: resourceOperationUpdateAppAwsebs,
		DeleteContext: resourceOperationDeleteAppAwsebs,
		Importer:      resourceOperationImportAppAwsebs(),
	}
}

// resourceOperationCreateAppAwsebs creates the sematext_app_awsebs resource.
func resourceOperationCreateAppAwsebs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS EBS"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppAwsebs reads the sematext_app_awsebs resource from Sematext Cloud.
func resourceOperationReadAppAwsebs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS EBS"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppAwsebs updates Sematext Cloud from the sematext_app_awsebs resource.
func resourceOperationUpdateAppAwsebs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS EBS"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppAwsebs marks a sematext_app_awsebs resource as retired.
func resourceOperationDeleteAppAwsebs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS EBS"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}

// resourceOperationImportAppAwsebs imports a sematext_app_awsebs resource into the state file.
func resourceOperationImportAppAwsebs() *schema.ResourceImporter {
	apptype := "AWS EBS"
	return sematext.ResourceOperationImportApp(apptype)
}
