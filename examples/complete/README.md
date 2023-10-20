<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

No requirements.

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_appmesh"></a> [appmesh](#module\_appmesh) | ../.. | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_name"></a> [name](#input\_name) | Name to use for the service mesh. Must be between 1 and 255 characters in length | `string` | n/a | yes |
| <a name="input_spec_egress_filter_type"></a> [spec\_egress\_filter\_type](#input\_spec\_egress\_filter\_type) | Egress filter type. By default, the type is DROP\_ALL. Valid values are ALLOW\_ALL and DROP\_ALL | `string` | `"DROP_ALL"` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of custom tags to be attached to this resource | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | ID of the service mesh. |
| <a name="output_arn"></a> [arn](#output\_arn) | ARN of the service mesh |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
