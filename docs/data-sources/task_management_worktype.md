---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "genesyscloud_task_management_worktype Data Source - terraform-provider-genesyscloud"
subcategory: ""
description: |-
  Genesys Cloud task management worktype data source. Select a task management worktype by name
---

# genesyscloud_task_management_worktype (Data Source)

Genesys Cloud task management worktype data source. Select a task management worktype by name

## Example Usage

```terraform
data "genesyscloud_task_management_worktype" "worktype_sample" {
  name = "My Worktype"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Task management worktype name

### Read-Only

- `id` (String) The ID of this resource.
