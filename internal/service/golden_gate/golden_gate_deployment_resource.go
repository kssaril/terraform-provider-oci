// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"
)

func GoldenGateDeploymentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("90m"),
			Update: tfresource.GetTimeoutDuration("60m"),
			Delete: tfresource.GetTimeoutDuration("30m"),
		},
		Create: createGoldenGateDeployment,
		Read:   readGoldenGateDeployment,
		Update: updateGoldenGateDeployment,
		Delete: deleteGoldenGateDeployment,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu_core_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"deployment_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_auto_scaling_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"license_model": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"deployment_backup_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fqdn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_public": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"maintenance_window": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"day": {
							Type:     schema.TypeString,
							Required: true,
						},
						"start_hour": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ogg_data": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"admin_password": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
						"admin_username": {
							Type:     schema.TypeString,
							Required: true,
						},
						"deployment_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"certificate": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ogg_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"deployment_diagnostic_data": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"bucket": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"diagnostic_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"object": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_diagnostic_end": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_diagnostic_start": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"deployment_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_healthy": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_latest_version": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_storage_utilization_limit_exceeded": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_sub_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"next_maintenance_action_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"next_maintenance_description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_golden_gate.LifecycleStateInactive),
					string(oci_golden_gate.LifecycleStateActive),
				}, true),
			},
			"storage_utilization_in_bytes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_next_maintenance": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_upgrade_required": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createGoldenGateDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	var stopDeployment = false
	if state, ok := sync.D.GetOkExists("state"); ok {
		if oci_golden_gate.LifecycleStateEnum(strings.ToUpper(state.(string))) == oci_golden_gate.LifecycleStateInactive {
			stopDeployment = true
		}
	}

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if stopDeployment {
		return sync.stopDeployment()
	}

	return nil
}

func readGoldenGateDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

func updateGoldenGateDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	var stopDeployment = false
	var startDeployment = false

	if _, ok := sync.D.GetOkExists("state"); ok && sync.D.HasChange("state") {
		oldStateStr, newStateStr := sync.D.GetChange("state")
		oldState := oci_golden_gate.LifecycleStateEnum(strings.ToUpper(oldStateStr.(string)))
		newState := oci_golden_gate.LifecycleStateEnum(strings.ToUpper(newStateStr.(string)))

		if newState == oci_golden_gate.LifecycleStateInactive {
			stopDeployment = true
		} else if oldState != oci_golden_gate.LifecycleStateFailed && oldState != oci_golden_gate.LifecycleStateNeedsAttention &&
			newState == oci_golden_gate.LifecycleStateActive {
			startDeployment = true
		}
	}

	/* Note, that "update" allowed only on running deployment, so if there is a "start", we need to start first, then update.
	If there is a "stop", we need to update first, then stop the deployment.
	Upgrade is much faster if the deployment is not running, therefore
	if we need to stop/start the deployment and upgrade it at the same time it worth to change the execution order.

	Here are the possible cases:
	a. Deployment has to be started (current state is Inactive):
	  1. let's do the upgrade first (if needed)
	  2. then start
	  3. then update
	b. Deployment has to be stopped (current state is Active):
	  1. update first
	  2. then stop
	  3. then upgrade (if needed)
	c. No stop/start needed:
	  1. upgrade first (if needed)
	  2. then update
	*/

	if startDeployment {
		if err := upgradeGoldenGateDeploymentIfNeeded(d, m); err != nil {
			return err
		}
		if err := sync.startDeployment(); err != nil {
			return err
		}
		return tfresource.UpdateResource(d, sync)

	} else if stopDeployment {
		if err := tfresource.UpdateResource(d, sync); err != nil {
			return err
		}
		if err := sync.stopDeployment(); err != nil {
			return err
		}
		return upgradeGoldenGateDeploymentIfNeeded(d, m)

	} else {
		if err := upgradeGoldenGateDeploymentIfNeeded(d, m); err != nil {
			return err
		}
		return tfresource.UpdateResource(d, sync)
	}
}

func upgradeGoldenGateDeploymentIfNeeded(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	// add support to upgrade
	oggVersionKeyFormat := fmt.Sprintf("%s.%d.%s", "ogg_data", 0, "ogg_version")
	if _, ok := sync.D.GetOkExists(oggVersionKeyFormat); ok && sync.D.HasChange(oggVersionKeyFormat) {
		oldVersion, newVersion := sync.D.GetChange(oggVersionKeyFormat)
		if newVersion != "" && oldVersion != newVersion {
			if err := sync.upgradeToSpecificVersion(newVersion); err != nil {
				return err
			}
		}
	}
	return nil
}

func deleteGoldenGateDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type GoldenGateDeploymentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_golden_gate.GoldenGateClient
	Res                    *oci_golden_gate.Deployment
	DisableNotFoundRetries bool
}

func (s *GoldenGateDeploymentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GoldenGateDeploymentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_golden_gate.LifecycleStateCreating),
		string(oci_golden_gate.LifecycleStateInProgress),
	}
}

func (s *GoldenGateDeploymentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_golden_gate.LifecycleStateActive),
		string(oci_golden_gate.LifecycleStateNeedsAttention),
		string(oci_golden_gate.LifecycleStateSucceeded),
	}
}

func (s *GoldenGateDeploymentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_golden_gate.LifecycleStateDeleting),
	}
}

func (s *GoldenGateDeploymentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_golden_gate.LifecycleStateDeleted),
	}
}

func (s *GoldenGateDeploymentResourceCrud) Create() error {
	request := oci_golden_gate.CreateDeploymentRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
		tmp := cpuCoreCount.(int)
		request.CpuCoreCount = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if deploymentBackupId, ok := s.D.GetOkExists("deployment_backup_id"); ok {
		tmp := deploymentBackupId.(string)
		request.DeploymentBackupId = &tmp
	}

	if deploymentType, ok := s.D.GetOkExists("deployment_type"); ok {
		request.DeploymentType = oci_golden_gate.DeploymentTypeEnum(deploymentType.(string))
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fqdn, ok := s.D.GetOkExists("fqdn"); ok {
		tmp := fqdn.(string)
		request.Fqdn = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
		tmp := isAutoScalingEnabled.(bool)
		request.IsAutoScalingEnabled = &tmp
	}

	if isPublic, ok := s.D.GetOkExists("is_public"); ok {
		tmp := isPublic.(bool)
		request.IsPublic = &tmp
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
		request.LicenseModel = oci_golden_gate.LicenseModelEnum(licenseModel.(string))
	}

	if maintenanceWindow, ok := s.D.GetOkExists("maintenance_window"); ok {
		if tmpList := maintenanceWindow.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_window", 0)
			tmp, err := s.mapToCreateMaintenanceWindowDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenanceWindow = &tmp
		}
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	if oggData, ok := s.D.GetOkExists("ogg_data"); ok {
		if tmpList := oggData.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ogg_data", 0)
			tmp, err := s.mapToCreateOggDeploymentDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.OggData = &tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.CreateDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDeploymentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate"), oci_golden_gate.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GoldenGateDeploymentResourceCrud) getDeploymentFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_golden_gate.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	deploymentId, err := goldenGateDeploymentWaitForWorkRequest(workId, "deployment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] operation failed: %v for identifier: %v\n", workId, deploymentId)
		return err
	}
	s.D.SetId(*deploymentId)

	return s.Get()
}

func goldenGateDeploymentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "golden_gate", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_golden_gate.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func goldenGateDeploymentWaitForWorkRequest(wId *string, entityType string, action oci_golden_gate.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_golden_gate.GoldenGateClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "golden_gate")
	retryPolicy.ShouldRetryOperation = goldenGateDeploymentWorkRequestShouldRetryFunc(timeout)

	response := oci_golden_gate.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_golden_gate.OperationStatusInProgress),
			string(oci_golden_gate.OperationStatusAccepted),
		},
		Target: []string{
			string(oci_golden_gate.OperationStatusSucceeded),
			string(oci_golden_gate.OperationStatusFailed),
			string(oci_golden_gate.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_golden_gate.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_golden_gate.OperationStatusFailed || response.Status == oci_golden_gate.OperationStatusCanceled {
		return nil, getErrorFromDeploymentWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDeploymentWorkRequest(client *oci_golden_gate.GoldenGateClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_golden_gate.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_golden_gate.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *GoldenGateDeploymentResourceCrud) Get() error {
	request := oci_golden_gate.GetDeploymentRequest{}

	tmp := s.D.Id()
	request.DeploymentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.GetDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Deployment
	return nil
}

func (s *GoldenGateDeploymentResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_golden_gate.UpdateDeploymentRequest{}

	if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
		tmp := cpuCoreCount.(int)
		request.CpuCoreCount = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	tmp := s.D.Id()
	request.DeploymentId = &tmp

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fqdn, ok := s.D.GetOkExists("fqdn"); ok {
		tmp := fqdn.(string)
		request.Fqdn = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
		tmp := isAutoScalingEnabled.(bool)
		request.IsAutoScalingEnabled = &tmp
	}

	if isPublic, ok := s.D.GetOkExists("is_public"); ok {
		tmp := isPublic.(bool)
		request.IsPublic = &tmp
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
		request.LicenseModel = oci_golden_gate.LicenseModelEnum(licenseModel.(string))
	}

	if maintenanceWindow, ok := s.D.GetOkExists("maintenance_window"); ok {
		if tmpList := maintenanceWindow.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_window", 0)
			tmp, err := s.mapToUpdateMaintenanceWindowDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenanceWindow = &tmp
		}
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	if oggData, ok := s.D.GetOkExists("ogg_data"); ok {
		if tmpList := oggData.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ogg_data", 0)
			tmp, err := s.mapToUpdateOggDeploymentDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.OggData = &tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.UpdateDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDeploymentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate"), oci_golden_gate.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *GoldenGateDeploymentResourceCrud) Delete() error {
	request := oci_golden_gate.DeleteDeploymentRequest{}

	tmp := s.D.Id()
	request.DeploymentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.DeleteDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := goldenGateDeploymentWaitForWorkRequest(workId, "deployment",
		oci_golden_gate.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GoldenGateDeploymentResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CpuCoreCount != nil {
		s.D.Set("cpu_core_count", *s.Res.CpuCoreCount)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DeploymentBackupId != nil {
		s.D.Set("deployment_backup_id", *s.Res.DeploymentBackupId)
	}

	if s.Res.DeploymentDiagnosticData != nil {
		s.D.Set("deployment_diagnostic_data", []interface{}{DeploymentDiagnosticDataToMap(s.Res.DeploymentDiagnosticData)})
	} else {
		s.D.Set("deployment_diagnostic_data", nil)
	}

	s.D.Set("deployment_type", s.Res.DeploymentType)

	if s.Res.DeploymentUrl != nil {
		s.D.Set("deployment_url", *s.Res.DeploymentUrl)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Fqdn != nil {
		s.D.Set("fqdn", *s.Res.Fqdn)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAutoScalingEnabled != nil {
		s.D.Set("is_auto_scaling_enabled", *s.Res.IsAutoScalingEnabled)
	}

	if s.Res.IsHealthy != nil {
		s.D.Set("is_healthy", *s.Res.IsHealthy)
	}

	if s.Res.IsLatestVersion != nil {
		s.D.Set("is_latest_version", *s.Res.IsLatestVersion)
	}

	if s.Res.IsPublic != nil {
		s.D.Set("is_public", *s.Res.IsPublic)
	}

	if s.Res.IsStorageUtilizationLimitExceeded != nil {
		s.D.Set("is_storage_utilization_limit_exceeded", *s.Res.IsStorageUtilizationLimitExceeded)
	}

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("lifecycle_sub_state", s.Res.LifecycleSubState)

	if s.Res.MaintenanceWindow != nil {
		s.D.Set("maintenance_window", []interface{}{MaintenanceWindowToMap(s.Res.MaintenanceWindow)})
	} else {
		s.D.Set("maintenance_window", nil)
	}

	s.D.Set("next_maintenance_action_type", s.Res.NextMaintenanceActionType)

	if s.Res.NextMaintenanceDescription != nil {
		s.D.Set("next_maintenance_description", *s.Res.NextMaintenanceDescription)
	}

	nsgIds := []interface{}{}
	for _, item := range s.Res.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))

	if s.Res.OggData != nil {
		s.D.Set("ogg_data", []interface{}{OggDeploymentToMap(s.Res.OggData, s.D)})
	} else {
		s.D.Set("ogg_data", nil)
	}

	if s.Res.PrivateIpAddress != nil {
		s.D.Set("private_ip_address", *s.Res.PrivateIpAddress)
	}

	if s.Res.PublicIpAddress != nil {
		s.D.Set("public_ip_address", *s.Res.PublicIpAddress)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StorageUtilizationInBytes != nil {
		s.D.Set("storage_utilization_in_bytes", strconv.FormatInt(*s.Res.StorageUtilizationInBytes, 10))
	}

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfNextMaintenance != nil {
		s.D.Set("time_of_next_maintenance", s.Res.TimeOfNextMaintenance.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TimeUpgradeRequired != nil {
		s.D.Set("time_upgrade_required", s.Res.TimeUpgradeRequired.String())
	}

	return nil
}

func (s *GoldenGateDeploymentResourceCrud) mapToUpdateMaintenanceWindowDetails(fieldKeyFormat string) (oci_golden_gate.UpdateMaintenanceWindowDetails, error) {
	result := oci_golden_gate.UpdateMaintenanceWindowDetails{}

	if day, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "day")); ok {
		result.Day = oci_golden_gate.DayEnum(day.(string))
	}

	if startHour, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "start_hour")); ok {
		tmp := startHour.(int)
		result.StartHour = &tmp
	}

	return result, nil
}

func (s *GoldenGateDeploymentResourceCrud) mapToCreateMaintenanceWindowDetails(fieldKeyFormat string) (oci_golden_gate.CreateMaintenanceWindowDetails, error) {
	result := oci_golden_gate.CreateMaintenanceWindowDetails{}

	if day, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "day")); ok {
		result.Day = oci_golden_gate.DayEnum(day.(string))
	}

	if startHour, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "start_hour")); ok {
		tmp := startHour.(int)
		result.StartHour = &tmp
	}

	return result, nil
}

func MaintenanceWindowToMap(obj *oci_golden_gate.MaintenanceWindow) map[string]interface{} {
	result := map[string]interface{}{}

	result["day"] = string(obj.Day)

	if obj.StartHour != nil {
		result["start_hour"] = int(*obj.StartHour)
	}

	return result
}

func (s *GoldenGateDeploymentResourceCrud) mapToCreateOggDeploymentDetails(fieldKeyFormat string) (oci_golden_gate.CreateOggDeploymentDetails, error) {
	result := oci_golden_gate.CreateOggDeploymentDetails{}

	if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok {
		tmp := adminPassword.(string)
		result.AdminPassword = &tmp
	}

	if adminUsername, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_username")); ok {
		tmp := adminUsername.(string)
		result.AdminUsername = &tmp
	}

	if certificate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate")); ok {
		tmp := certificate.(string)
		result.Certificate = &tmp
	}

	if deploymentName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deployment_name")); ok {
		tmp := deploymentName.(string)
		result.DeploymentName = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if oggVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ogg_version")); ok {
		tmp := oggVersion.(string)
		result.OggVersion = &tmp
	}

	return result, nil
}

func (s *GoldenGateDeploymentResourceCrud) mapToUpdateOggDeploymentDetails(fieldKeyFormat string) (oci_golden_gate.UpdateOggDeploymentDetails, error) {
	result := oci_golden_gate.UpdateOggDeploymentDetails{}

	if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok {
		tmp := adminPassword.(string)
		result.AdminPassword = &tmp
	}

	if adminUsername, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_username")); ok {
		tmp := adminUsername.(string)
		result.AdminUsername = &tmp
	}

	if certificate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate")); ok {
		tmp := certificate.(string)
		result.Certificate = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	return result, nil
}

func OggDeploymentToMap(obj *oci_golden_gate.OggDeployment, resourceData *schema.ResourceData) map[string]interface{} {
	result := map[string]interface{}{}

	if oggData, ok := resourceData.GetOkExists("ogg_data"); ok {
		if tmpList := oggData.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ogg_data", 0)
			if adminPassword, ok := resourceData.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok {
				tmp := adminPassword.(string)
				result["admin_password"] = &tmp
			}

			if key, ok := resourceData.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
				tmp := key.(string)
				result["key"] = &tmp
			}
		}
	}

	if obj.AdminUsername != nil {
		result["admin_username"] = string(*obj.AdminUsername)
	}

	if obj.Certificate != nil {
		result["certificate"] = string(*obj.Certificate)
	}

	if obj.DeploymentName != nil {
		result["deployment_name"] = string(*obj.DeploymentName)
	}

	if obj.OggVersion != nil {
		result["ogg_version"] = string(*obj.OggVersion)
	}

	return result
}

func DeploymentDiagnosticDataToMap(obj *oci_golden_gate.DeploymentDiagnosticData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BucketName != nil {
		result["bucket"] = string(*obj.BucketName)
	}

	result["diagnostic_state"] = string(obj.DiagnosticState)

	if obj.NamespaceName != nil {
		result["namespace"] = string(*obj.NamespaceName)
	}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	if obj.TimeDiagnosticEnd != nil {
		result["time_diagnostic_end"] = obj.TimeDiagnosticEnd.String()
	}

	if obj.TimeDiagnosticStart != nil {
		result["time_diagnostic_start"] = obj.TimeDiagnosticStart.String()
	}

	return result
}

func GoldenGateDeploymentSummaryToMap(obj oci_golden_gate.DeploymentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CpuCoreCount != nil {
		result["cpu_core_count"] = int(*obj.CpuCoreCount)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["deployment_type"] = string(obj.DeploymentType)

	if obj.DeploymentUrl != nil {
		result["deployment_url"] = string(*obj.DeploymentUrl)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Fqdn != nil {
		result["fqdn"] = string(*obj.Fqdn)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsAutoScalingEnabled != nil {
		result["is_auto_scaling_enabled"] = bool(*obj.IsAutoScalingEnabled)
	}

	if obj.IsLatestVersion != nil {
		result["is_latest_version"] = bool(*obj.IsLatestVersion)
	}

	if obj.IsPublic != nil {
		result["is_public"] = bool(*obj.IsPublic)
	}

	if obj.IsStorageUtilizationLimitExceeded != nil {
		result["is_storage_utilization_limit_exceeded"] = bool(*obj.IsStorageUtilizationLimitExceeded)
	}

	result["license_model"] = string(obj.LicenseModel)

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["lifecycle_sub_state"] = string(obj.LifecycleSubState)

	if obj.PrivateIpAddress != nil {
		result["private_ip_address"] = string(*obj.PrivateIpAddress)
	}

	if obj.PublicIpAddress != nil {
		result["public_ip_address"] = string(*obj.PublicIpAddress)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.StorageUtilizationInBytes != nil {
		result["storage_utilization_in_bytes"] = strconv.FormatInt(*obj.StorageUtilizationInBytes, 10)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.TimeUpgradeRequired != nil {
		result["time_upgrade_required"] = obj.TimeUpgradeRequired.String()
	}

	return result
}

func (s *GoldenGateDeploymentResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_golden_gate.ChangeDeploymentCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DeploymentId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.ChangeDeploymentCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, changeWorkRequestErr := goldenGateDeploymentWaitForWorkRequest(workId, "deployment",
		oci_golden_gate.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
	return changeWorkRequestErr
}

func (s *GoldenGateDeploymentResourceCrud) upgradeToSpecificVersion(oggVersion interface{}) error {
	upgradeDeploymentRequest := oci_golden_gate.UpgradeDeploymentRequest{}
	upgradeDetails := oci_golden_gate.UpgradeDeploymentSpecificReleaseDetails{}

	oggVersionTmp := oggVersion.(string)
	upgradeDetails.OggVersion = &oggVersionTmp

	idTmp := s.D.Id()
	upgradeDeploymentRequest.DeploymentId = &idTmp
	upgradeDeploymentRequest.UpgradeDeploymentDetails = upgradeDetails

	upgradeDeploymentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.UpgradeDeployment(context.Background(), upgradeDeploymentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, upgradeWorkRequestErr := goldenGateDeploymentWaitForWorkRequest(workId, "deployment",
		oci_golden_gate.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
	if upgradeWorkRequestErr != nil {
		return upgradeWorkRequestErr
	}

	// set changed parameters
	if err := s.getAndSaveStateChanges(); err != nil {
		return err
	}

	if s.Res.OggData != nil {
		if oggData, ok := s.D.GetOkExists("ogg_data"); ok {
			oggDataMap := oggData.([]interface{})[0].(map[string]interface{})
			oggDataMap["ogg_version"] = *s.Res.OggData.OggVersion
			s.D.Set("ogg_data", []interface{}{oggDataMap})
		}
	}

	if s.Res.IsLatestVersion != nil {
		s.D.Set("is_latest_version", *s.Res.IsLatestVersion)
	}

	return nil
}

func (s *GoldenGateDeploymentResourceCrud) startDeployment() error {
	startDeploymentRequest := oci_golden_gate.StartDeploymentRequest{}

	idTmp := s.D.Id()
	startDeploymentRequest.DeploymentId = &idTmp
	startDeploymentRequest.StartDeploymentDetails = oci_golden_gate.DefaultStartDeploymentDetails{}
	startDeploymentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.StartDeployment(context.Background(), startDeploymentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, startWorkRequestErr := goldenGateDeploymentWaitForWorkRequest(workId, "deployment",
		oci_golden_gate.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
	if startWorkRequestErr != nil {
		return startWorkRequestErr
	}
	// set changed parameters
	return s.getAndSaveStateChanges()
}

func (s *GoldenGateDeploymentResourceCrud) stopDeployment() error {
	stopDeploymentRequest := oci_golden_gate.StopDeploymentRequest{}

	idTmp := s.D.Id()
	stopDeploymentRequest.DeploymentId = &idTmp
	stopDeploymentRequest.StopDeploymentDetails = oci_golden_gate.DefaultStopDeploymentDetails{}
	stopDeploymentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.StopDeployment(context.Background(), stopDeploymentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, stopWorkRequestErr := goldenGateDeploymentWaitForWorkRequest(workId, "deployment",
		oci_golden_gate.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
	if stopWorkRequestErr != nil {
		return stopWorkRequestErr
	}

	// set changed parameters
	return s.getAndSaveStateChanges()
}

func (s *GoldenGateDeploymentResourceCrud) getAndSaveStateChanges() error {
	if e := s.Get(); e != nil {
		return e
	}
	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("lifecycle_sub_state", s.Res.LifecycleSubState)
	return nil
}
