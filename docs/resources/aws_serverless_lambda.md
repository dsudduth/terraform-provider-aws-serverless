---
page_title: "aws_serverless_lambda Resource - terraform-provider-aws-serverless"
subcategory: ""
description: |-
  The aws_serverless_lambda resources locally manages the packaging of Lambda functions.
---

# Resource aws_serverless_lambda

## Schema

### Required

- `source` (String) Source of the Lambda function to package.

### Optional

- `output_path` (String) The output of the archive file. If  not specified, a folder will be created.

### Read-Only

- `id` (String) The ID of this resource.
