---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_db_nodes"
sidebar_current: "docs-oci-datasource-database_management-external_db_nodes"
description: |-
  Provides the list of External Db Nodes in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_db_nodes
This data source provides the list of External Db Nodes in Oracle Cloud Infrastructure Database Management service.

Lists the external DB nodes in the specified external DB system.

## Example Usage

```hcl
data "oci_database_management_external_db_nodes" "test_external_db_nodes" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.external_db_node_display_name
	external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to only return the resources that match the entire display name.
* `external_db_system_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system.


## Attributes Reference

The following attributes are exported:

* `external_db_node_collection` - The list of external_db_node_collection.

### ExternalDbNode Reference

The following attributes are exported:

* `additional_details` - The additional details of the external DB node defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `component_name` - The name of the external DB node.
* `cpu_core_count` - The number of CPU cores available on the DB node.
* `display_name` - The user-friendly name for the external DB node. The name does not have to be unique.
* `external_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external connector.
* `external_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system that the DB node is a part of.
* `host_name` - The host name for the DB node.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB node.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `memory_size_in_gbs` - The total memory in gigabytes (GB) on the DB node.
* `state` - The current lifecycle state of the external DB node.
* `time_created` - The date and time the external DB node was created.
* `time_updated` - The date and time the external DB node was last updated.

