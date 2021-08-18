// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v46/common"
	oci_database_migration "github.com/oracle/oci-go-sdk/v46/databasemigration"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AgentRequiredOnlyResource = generateResourceFromRepresentationMap("oci_database_migration_agent", "test_agent", Required, Create, agentRepresentation2)

	AgentResourceConfig = generateResourceFromRepresentationMap("oci_database_migration_agent", "test_agent", Optional, Update, agentRepresentation2)

	agentSingularDataSourceRepresentation = map[string]interface{}{
		"agent_id": Representation{repType: Required, create: `${oci_database_migration_agent.test_agent.id}`},
	}

	agentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `TF_displayName`, update: `TF_displayName2`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"filter":         RepresentationGroup{Required, agentDataSourceFilterRepresentation}}
	agentDataSourceRepresentation2 = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"filter":         RepresentationGroup{Required, agentDataSourceFilterRepresentation}}
	agentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `agent_id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_migration_agent.test_agent.id}`}},
	}

	agentRepresentation = map[string]interface{}{
		"agent_id":       Representation{repType: Required, create: `${oci_database_migration_agent.test_agent.id}`},
		"compartment_id": Representation{repType: Optional, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `TF_displayName`, update: `TF_displayName2`},
		"stream_id":      Representation{repType: Optional, create: `${oci_streaming_stream.test_stream.id}`},
		"version":        Representation{repType: Optional, create: `version`, update: `version2`},
	}

	agentRepresentation2 = map[string]interface{}{
		"agent_id":       Representation{repType: Required, create: `${oci_database_migration_agent.test_agent.id}`},
		"compartment_id": Representation{repType: Optional, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `TF_displayName`, update: `TF_displayName2`},
		"stream_id":      Representation{repType: Optional, create: `${oci_streaming_stream.test_stream.id}`},
		"version":        Representation{repType: Optional, create: `version`, update: `version2`},
	}

	AgentResourceDependencies = generateResourceFromRepresentationMap("oci_database_migration_agent", "test_agent", Required, Create, agentRepresentation) +
		DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", Required, Create, streamRepresentation)
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationAgentResource_basic(t *testing.T) {
	t.Skip("Skip this test agent creation is an independent operation.")
	httpreplay.SetScenario("TestDatabaseMigrationAgentResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_migration_agent.test_agent"
	datasourceName := "data.oci_database_migration_agents.test_agents"
	singularDatasourceName := "data.oci_database_migration_agent.test_agent"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ //AgentResourceDependencies+
		generateResourceFromRepresentationMap("oci_database_migration_agent", "test_agent", Optional, Create, agentRepresentation), "databasemigration", "agent", t)

	ResourceTest(t, testAccCheckDatabaseMigrationAgentDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr +
				generateResourceFromRepresentationMap("oci_database_migration_agent", "test_agent", Required, Create, agentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "agent_id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr +
				generateResourceFromRepresentationMap("oci_database_migration_agent", "test_agent", Optional, Create, agentRepresentation2),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "stream_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "version", "version"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "agent_id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr +
				generateResourceFromRepresentationMap("oci_database_migration_agent", "test_agent", Optional, Create,
					representationCopyWithNewProperties(agentRepresentation2, map[string]interface{}{
						"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "stream_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "version", "version"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "agent_id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr +
				generateResourceFromRepresentationMap("oci_database_migration_agent", "test_agent", Optional, Update, agentRepresentation2),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "stream_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "version", "version2"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "agent_id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_database_migration_agents", "test_agents", Optional, Update, agentDataSourceRepresentation2) +
				compartmentIdVariableStr +
				generateResourceFromRepresentationMap("oci_database_migration_agent", "test_agent", Optional, Update, agentRepresentation2),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "agent_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "agent_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_database_migration_agent", "test_agent", Required, Create, agentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AgentResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "agent_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TF_displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "version", "version2"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AgentResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       false,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseMigrationAgentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseMigrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_migration_agent" {
			noResourceFound = false
			request := oci_database_migration.GetAgentRequest{}

			tmp := rs.Primary.ID
			request.AgentId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database_migration")

			response, err := client.GetAgent(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_migration.LifecycleStatesDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatabaseMigrationAgent") {
		resource.AddTestSweepers("DatabaseMigrationAgent", &resource.Sweeper{
			Name:         "DatabaseMigrationAgent",
			Dependencies: DependencyGraph["agent"],
			F:            sweepDatabaseMigrationAgentResource,
		})
	}
}

func sweepDatabaseMigrationAgentResource(compartment string) error {
	databaseMigrationClient := GetTestClients(&schema.ResourceData{}).databaseMigrationClient()
	agentIds, err := getAgentIds(compartment)
	if err != nil {
		return err
	}
	for _, agentId := range agentIds {
		if ok := SweeperDefaultResourceId[agentId]; !ok {
			deleteAgentRequest := oci_database_migration.DeleteAgentRequest{}

			deleteAgentRequest.AgentId = &agentId

			deleteAgentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database_migration")
			_, error := databaseMigrationClient.DeleteAgent(context.Background(), deleteAgentRequest)
			if error != nil {
				fmt.Printf("Error deleting Agent %s %s, It is possible that the resource is already deleted. Please verify manually \n", agentId, error)
				continue
			}
			waitTillCondition(testAccProvider, &agentId, agentSweepWaitCondition, time.Duration(3*time.Minute),
				agentSweepResponseFetchOperation, "database_migration", true)
		}
	}
	return nil
}

func getAgentIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "AgentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseMigrationClient := GetTestClients(&schema.ResourceData{}).databaseMigrationClient()

	listAgentsRequest := oci_database_migration.ListAgentsRequest{}
	listAgentsRequest.CompartmentId = &compartmentId
	listAgentsRequest.LifecycleState = oci_database_migration.ListAgentsLifecycleStateActive
	listAgentsResponse, err := databaseMigrationClient.ListAgents(context.Background(), listAgentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Agent list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, agent := range listAgentsResponse.Items {
		id := *agent.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "AgentId", id)
	}
	return resourceIds, nil
}

func agentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if agentResponse, ok := response.Response.(oci_database_migration.GetAgentResponse); ok {
		return agentResponse.LifecycleState != oci_database_migration.LifecycleStatesDeleted
	}
	return false
}

func agentSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseMigrationClient().GetAgent(context.Background(), oci_database_migration.GetAgentRequest{
		AgentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
