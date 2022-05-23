---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "alzlib_archetypes Data Source - terraform-provider-alzlib"
subcategory: ""
description: |-
  Archetypes data from the provider
---

# alzlib_archetypes (Data Source)

Archetypes data from the provider

## Example Usage

```terraform
data "alzlib_archetypes" "test" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `archetypes` (Map of Object) (see [below for nested schema](#nestedatt--archetypes))
- `id` (Number) The ID of this resource.

<a id="nestedatt--archetypes"></a>
### Nested Schema for `archetypes`

Read-Only:

- `name` (String)
- `policy_definitions` (Map of Object) (see [below for nested schema](#nestedobjatt--archetypes--policy_definitions))

<a id="nestedobjatt--archetypes--policy_definitions"></a>
### Nested Schema for `archetypes.policy_definitions`

Read-Only:

- `display_name` (String)
- `name` (String)

