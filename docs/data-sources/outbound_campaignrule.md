---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "genesyscloud_outbound_campaignrule Data Source - terraform-provider-genesyscloud"
subcategory: ""
description: |-
  Genesys Cloud outbound campaign rule data source. Select a campaign rule by name
---

# genesyscloud_outbound_campaignrule (Data Source)

Genesys Cloud outbound campaign rule data source. Select a campaign rule by name

## Example Usage

```terraform
data "genesyscloud_outbound_campaignrule" "rule" {
  name = "Campaign Rule X"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Campaign Rule name.

### Read-Only

- `id` (String) The ID of this resource.
