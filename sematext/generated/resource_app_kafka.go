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

// resourceAppKafka is the resource class that handles sematext_app_kafka
func resourceAppKafka() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaApp("Kafka"),
		CreateContext: resourceOperationCreateAppKafka,
		ReadContext:   resourceOperationReadAppKafka,
		UpdateContext: resourceOperationUpdateAppKafka,
		DeleteContext: resourceOperationDeleteAppKafka,
		Importer:      resourceOperationImportAppKafka(),
	}
}

// resourceOperationCreateAppKafka creates the sematext_app_kafka resource.
func resourceOperationCreateAppKafka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Kafka"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppKafka reads the sematext_app_kafka resource from Sematext Cloud.
func resourceOperationReadAppKafka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Kafka"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppKafka updates Sematext Cloud from the sematext_app_kafka resource.
func resourceOperationUpdateAppKafka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Kafka"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppKafka marks a sematext_app_kafka resource as retired.
func resourceOperationDeleteAppKafka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Kafka"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}

// resourceOperationImportAppKafka imports a sematext_app_kafka resource into the state file.
func resourceOperationImportAppKafka() *schema.ResourceImporter {
	apptype := "Kafka"
	return sematext.ResourceOperationImportApp(apptype)
}
