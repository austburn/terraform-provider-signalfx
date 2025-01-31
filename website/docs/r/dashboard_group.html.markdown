---
layout: "signalfx"
page_title: "SignalFx: signalfx_resource"
sidebar_current: "docs-signalfx-resource-dashboard-group"
description: |-
  Allows Terraform to create and manage SignalFx Dashboard Groups
---

# Resource: signalfx_dashboard_group

In the SignalFx web UI, a [dashboard group](https://developers.signalfx.com/v2/docs/dashboard-group-model) is a collection of dashboards.

**NOTE:** Dashboard groups cannot be accessed directly, but just via a dashboard contained in them. This is the reason why make show won't show any of yours dashboard groups.

## Example Usage

```terraform
resource "signalfx_dashboard_group" "mydashboardgroup0" {
    name = "My team dashboard group"
    description = "Cool dashboard group"
}
```

## Argument Reference

The following arguments are supported in the resource block:

* `name` - (Required) Name of the dashboard group.
* `description` - (Required) Description of the dashboard group.
* `teams` - (Optional) Team IDs to associate the dashboard group to.
