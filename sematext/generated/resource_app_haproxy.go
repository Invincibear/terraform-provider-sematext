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

// resourceAppHaproxy is the resource class that handles sematext_app_haproxy
func resourceAppHaproxy() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaApp("HAProxy"),
		CreateContext: resourceOperationCreateAppHaproxy,
		ReadContext:   resourceOperationReadAppHaproxy,
		UpdateContext: resourceOperationUpdateAppHaproxy,
		DeleteContext: resourceOperationDeleteAppHaproxy,
		Importer:      resourceOperationImportAppHaproxy(),
	}
}

// resourceOperationCreateAppHaproxy creates the sematext_app_haproxy resource.
func resourceOperationCreateAppHaproxy(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "HAProxy"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppHaproxy reads the sematext_app_haproxy resource from Sematext Cloud.
func resourceOperationReadAppHaproxy(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "HAProxy"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppHaproxy updates Sematext Cloud from the sematext_app_haproxy resource.
func resourceOperationUpdateAppHaproxy(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "HAProxy"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppHaproxy marks a sematext_app_haproxy resource as retired.
func resourceOperationDeleteAppHaproxy(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "HAProxy"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}

// resourceOperationImportAppHaproxy imports a sematext_app_haproxy resource into the state file.
func resourceOperationImportAppHaproxy() *schema.ResourceImporter {
	apptype := "HAProxy"
	return sematext.ResourceOperationImportApp(apptype)
}
