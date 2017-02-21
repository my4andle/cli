// Copyright © 2016 Paul Allen <paul@cloudcoreo.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package content

const (
	//CmdPlanUse plan command
	CmdPlanUse = "plan SUBCOMMAND"

	//CmdPlanFinalizeUse plan command
	CmdPlanFinalizeUse = "finalize [flags]"

	//CmdPlanShort short description
	CmdPlanShort = "Manage Plans"

	//CmdPlanLong long description
	CmdPlanLong = `Manage Plans.`

	//CmdPlanListShort short description
	CmdPlanListShort = "List all plans"

	//CmdPlanListLong long description
	CmdPlanListLong = `List all plans.`

	//CmdPlanInitShort short description
	CmdPlanInitShort = "Init a plan"

	//CmdPlanInitLong long description
	CmdPlanInitLong = `Init a plan.`

	//CmdPlanCreateShort short description
	CmdPlanCreateShort = "Finalize a plan"

	//CmdPlanCreateLong long description
	CmdPlanCreateLong = `Finalize a plan.`

	//CmdPlanDeleteShort short description
	CmdPlanDeleteShort = "Delete a plan"

	//CmdPlanDeleteLong long description
	CmdPlanDeleteLong = `Delete a plan.`

	//CmdPlanShowShort short description
	CmdPlanShowShort = "Show a plan"

	//CmdPlanShowLong long description
	CmdPlanShowLong = `Show a plan.`

	//CmdPlanDisableShort short description
	CmdPlanDisableShort = "Disable a plan"

	//CmdPlanDisableLong long description
	CmdPlanDisableLong = `Disable a plan.`

	//CmdPlanEnableShort short description
	CmdPlanEnableShort = "Enable a plan"

	//CmdPlanEnableLong long description
	CmdPlanEnableLong = `Enable a plan.`

	//CmdPlanRunShort short description
	CmdPlanRunShort = "Run a plan"

	//CmdPlanRunLong long description
	CmdPlanRunLong = `Run a plan.`

	//CmdFlagPlanIDLong flag
	CmdFlagPlanIDLong = "plan-id"

	//CmdFlagPlanIDDescription flag description
	CmdFlagPlanIDDescription = "Coreo plan id"

	//CmdFlagBranchLong flag
	CmdFlagBranchLong = "branch"

	//CmdFlagBranchDescription flag description
	CmdFlagBranchDescription = "Git branch for plan"

	//CmdFlagGitRevisionLong commit id flag
	CmdFlagGitRevisionLong = "revision"

	//CmdFlagGitRevisionDescription flag description
	CmdFlagGitRevisionDescription = "Git revision for plan"

	//CmdFlagCloudRegionLong cloud region flag
	CmdFlagCloudRegionLong = "region"

	//CmdFlagCloudRegionDescription flag description
	CmdFlagCloudRegionDescription = "Cloud region, e.g. AWS 'us-east-1'"

	//CmdFlagIntervalLong interval flag
	CmdFlagIntervalLong = "interval"

	//CmdFlagIntervalDescription flag description
	CmdFlagIntervalDescription = "Refresh rate value with any increment in mintues (between 1 and 525547)"

	//CmdFlagJSONFileLong JSON file flag
	CmdFlagJSONFileLong = "file"

	//CmdFlagJSONFileShort JSON file flag
	CmdFlagJSONFileShort = "f"

	//CmdFlagJSONFileDescription JSON file description
	CmdFlagJSONFileDescription = "Plan config JSON file"

	//InfoPlanDeleted information
	InfoPlanDeleted = "[Done] Plan was deleted.\n"

	//InfoPlanJSONFileCreated plan json file creation
	InfoPlanJSONFileCreated = "[Done] Plan %s.json file created in directory %s. Update variables in this json file and finalize this plan using `coreo plan finalize` command.\n"

	//InfoUsingPlanID informational using Plan id
	InfoUsingPlanID = "[ OK ] Using Plan ID %s\n"

	//ErrorPlanIDRequired error message
	ErrorPlanIDRequired = "Plan ID is required for this command. Use flag '--plan-id'\n"

	//ErrorPlanCreateJSONFileRequired error message
	ErrorPlanCreateJSONFileRequired = "Plan JSON file is required for this command. Use flag '--file/-f'\n"
)
