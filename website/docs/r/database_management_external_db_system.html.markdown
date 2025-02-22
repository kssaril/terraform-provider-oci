---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_db_system"
sidebar_current: "docs-oci-resource-database_management-external_db_system"
description: |-
  Provides the External Db System resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_external_db_system
This resource provides the External Db System resource in Oracle Cloud Infrastructure Database Management service.

Creates an external DB system and its related resources.


## Example Usage

```hcl
resource "oci_database_management_external_db_system" "test_external_db_system" {
	#Required
	compartment_id = var.compartment_id
	db_system_discovery_id = oci_database_management_db_system_discovery.test_db_system_discovery.id

	#Optional
	database_management_config {
		#Required
		license_model = var.external_db_system_database_management_config_license_model
	}
	display_name = var.external_db_system_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the external DB system resides.
* `database_management_config` - (Optional) The details required to enable Database Management for an external DB system.
	* `license_model` - (Required) The Oracle license model that applies to the external database. 
* `db_system_discovery_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system discovery.
* `display_name` - (Optional) (Updatable) The user-friendly name for the DB system. The name does not have to be unique.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_management_config` - The details required to enable Database Management for an external DB system.
	* `license_model` - The Oracle license model that applies to the external database. 
* `db_system_discovery_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system discovery.
* `discovery_agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent used during the discovery of the DB system.
* `display_name` - The user-friendly name for the DB system. The name does not have to be unique.
* `home_directory` - The Oracle Grid home directory in case of cluster-based DB system and Oracle home directory in case of single instance-based DB system. 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system.
* `is_cluster` - Indicates whether the DB system is a cluster DB system or not.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current lifecycle state of the external DB system resource.
* `time_created` - The date and time the external DB system was created.
* `time_updated` - The date and time the external DB system was last updated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the External Db System
	* `update` - (Defaults to 20 minutes), when updating the External Db System
	* `delete` - (Defaults to 20 minutes), when destroying the External Db System


## Import

ExternalDbSystems can be imported using the `id`, e.g.

```
$ terraform import oci_database_management_external_db_system.test_external_db_system "id"
```

