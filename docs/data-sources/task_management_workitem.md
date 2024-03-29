---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "genesyscloud_task_management_workitem Data Source - terraform-provider-genesyscloud"
subcategory: ""
description: |-
  Genesys Cloud task management workitem data source. Select an task management workitem by name
---

# genesyscloud_task_management_workitem (Data Source)

Genesys Cloud task management workitem data source. Select an task management workitem by name

## Example Usage

```terraform
data "genesyscloud_task_management_workitem" "example_workitem" {
  name = "My Workitem"

  // Requires either or both of the following fields
  workbin_id  = genesyscloud_routing_workbin.example.id
  worktype_id = genesyscloud_routing_worktype.example.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Task management workitem name

### Optional

- `workbin_id` (String) Id of the workbin where the desired workitem is.
- `worktype_id` (String) Id of the worktype of the desired workitem.

### Read-Only

- `id` (String) The ID of this resource.
