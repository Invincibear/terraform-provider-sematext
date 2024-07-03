# <img src="https://sematext.com/wp-content/uploads/2020/09/just-octi-blue.png" valign="bottom" width="60px"/>**&nbsp;&nbsp;Terraform Provider For Sematext Cloud**

# Overview

The Sematext provider is used to interact with resources supplied by [Sematext Cloud](https://sematext.com/cloud/).


## Example Usage

```hcl
terraform {
  required_providers {
    sematext = {
      source = "sematext/sematext"
      version = ">=0.1.3"
    }
  }
}

provider "sematext" {
    sematext_api_key = "00000000-0000-0000-0000-000000000000"
    sematext_region  = "US"
}

resource "sematext_monitor_nodejs" "mymonitor" {
    name = "Node.js Monitor Example"
    billing_plan_id = 6
}
```

## Argument Reference

The following arguments are supported:

* `sematext_api_key` - (Optional) your Sematext Cloud API key. The provider will use environment variable SEMATEXT_API_KEY (see below) by default if both are set;
* `sematext_region` - (Required) desired Sematext Cloud Region  "US" or "EU";


## Authentication

There are two authentication tokens

* Sematext Cloud API access token - [available from your account](https://apps.sematext.com/ui/account/api);
* Sematext Cloud App access token - retrieved on resource creation - refer to examples on how to access this inside your Terrform scripting.


## Environment Variables

The following environment variables are required:

* SEMATEXT_REGION="US"
* SEMATEXT_API_KEY="&lt;Sematext-Cloud-Token&gt;" # If not set using sematext_api_key in the `provider {}` code block

If working with AWS Cloudwatch the following environment vars should be set:

* AWS_ACCESS_KEY_ID
* AWS_SECRET_ACCESS_KEY
* AWS_REGION


## Resources

Refer to docs/resources for guidance on individual resources.
