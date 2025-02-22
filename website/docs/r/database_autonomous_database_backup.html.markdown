---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_backup"
sidebar_current: "docs-oci-resource-database-autonomous_database_backup"
description: |-
  Provides the Autonomous Database Backup resource in Oracle Cloud Infrastructure Database service
---

# oci_database_autonomous_database_backup
This resource provides the Autonomous Database Backup resource in Oracle Cloud Infrastructure Database service.

Creates a new Autonomous Database backup for the specified database based on the provided request parameters.


## Example Usage

```hcl
resource "oci_database_autonomous_database_backup" "test_autonomous_database_backup" {
	#Required
	autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id

	#Optional
	display_name = var.autonomous_database_backup_display_name
	is_long_term_backup = var.autonomous_database_backup_is_long_term_backup
	retention_period_in_days = var.autonomous_database_backup_retention_period_in_days
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database backup.
* `display_name` - (Optional) The user-friendly name for the backup. The name does not have to be unique.
* `is_long_term_backup` - (Optional) Indicates whether the backup is long-term
* `retention_period_in_days` - (Optional) (Updatable) Retention period, in days, for long-term backups


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `autonomous_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_size_in_tbs` - The size of the database in terabytes at the time the backup was taken. 
* `db_version` - A valid Oracle Database version for Autonomous Database.
* `display_name` - The user-friendly name for the backup. The name does not have to be unique.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database backup.
* `is_automatic` - Indicates whether the backup is user-initiated or automatic.
* `is_restorable` - Indicates whether the backup can be used to restore the associated Autonomous Database.
* `key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store.
* `key_store_wallet_name` - The wallet name for Oracle Key Vault.
* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. 
* `lifecycle_details` - Additional information about the current lifecycle state.
* `retention_period_in_days` - Retention period, in days, for long-term backups
* `size_in_tbs` - The backup size in terrabytes (TB).
* `state` - The current state of the backup.
* `time_available_till` - Timestamp until when the backup will be available
* `time_ended` - The date and time the backup completed.
* `time_started` - The date and time the backup started.
* `type` - The type of backup.
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Autonomous Database Backup
	* `update` - (Defaults to 20 minutes), when updating the Autonomous Database Backup
	* `delete` - (Defaults to 20 minutes), when destroying the Autonomous Database Backup


## Import

AutonomousDatabaseBackups can be imported using the `id`, e.g.

```
$ terraform import oci_database_autonomous_database_backup.test_autonomous_database_backup "id"
```

