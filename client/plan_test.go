package client

import (
	"net/http"
	"testing"

	"fmt"

	"github.com/jharlap/httpstub"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

const compositeJSONPayloadForPlan = `[
	{
		"name": "audit-aws-s3",
		"gitUrl": "git@github.com:CloudCoreo/audit-aws-s3.git",
		"hasCustomDashboard": false,
		"createdAt": "2016-11-28T06:10:53.903Z",
		"gitKeyId": "0",
		"teamId": "teamID",
		"id": "compositeID",
		"links": [
			{
				"ref": "self",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/composites/compositeID"
			},
			{
				"ref": "team",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/teams/teamID"
			},
			{
				"ref": "gitKey",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/gitkeys/0"
			},
			{
				"ref": "plans",
				"method": "GET",
				"href": "%s/api/composites/compositeID/plans"
			}
		]
	}
]`

const compositeJSONPayloadForPlanMissingPlanLinks = `[
	{
		"name": "audit-aws-s3",
		"gitUrl": "git@github.com:CloudCoreo/audit-aws-s3.git",
		"hasCustomDashboard": false,
		"createdAt": "2016-11-28T06:10:53.903Z",
		"gitKeyId": "0",
		"teamId": "teamID",
		"id": "compositeID",
		"links": [
			{
				"ref": "self",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/composites/compositeID"
			},
			{
				"ref": "team",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/teams/teamID"
			},
			{
				"ref": "gitKey",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/gitkeys/0"
			}
		]
	}
]`

const planJSONPayloadSingle = `{

		"isDraft": true,
		"links": [
			{
				"ref": "self",
				"method": "GET",
				"href": "%s/api/plans/planid"
			},
			{
				"ref": "team",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/teams/teamID"
			},
			{
				"ref": "composite",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/composites/compositeID"
			},
			{
				"ref": "cloudAccount",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/cloudaccounts/cloudAccountID"
			},
			{
				"ref": "planconfig",
				"method": "GET",
				"href": "%s/api/plans/planID/planconfig"
			},
			{
				"ref": "runnow",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/plans/planID/runnow"
			}
		],
		"id": "planID"
	}`

const planJSONPayload = `[{
		"defaultPanelRepo": "git@github.com:CloudCoreo/default-panel.git",
		"defaultPanelDirectory": "panel",
		"defaultPanelBranch": "master",
		"name": "Audit-S3",
		"iamUserAccessKeyId": "IAMUSERACCESSKEYID",
		"iamUserId": "iamUserId",
		"iamUserSecretAccessKey": "iamUserSecretAccessKey",
		"snsSubscriptionArn": "arn:aws:sns:us-west-1:12312123:coreo-asi-planID:f50ac4e8-82a8-4291-a0c3-d9e299f79d8d",
		"sqsArn": "arn:aws:sqs:us-west-1:12312123:coreo-asi-planID",
		"sqsUrl": "sqsUrl",
		"topicArn": "arn:aws:sns:us-west-1:12312123:coreo-asi-planID",
		"defaultRegion": "us-east-1",
		"refreshInterval": 1,
		"revision": "HEAD",
		"branch": "master",
		"enabled": false,
		"links": [
			{
				"ref": "self",
				"method": "GET",
				"href": "%s/api/plans/planID"
			},
			{
				"ref": "team",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/teams/teamID"
			},
			{
				"ref": "composite",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/composites/compositeID"
			},
			{
				"ref": "cloudAccount",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/cloudaccounts/cloudAccountID"
			},
			{
				"ref": "planconfig",
				"method": "GET",
				"href": "%s/api/plans/planID/planconfig"
			},
			{
				"ref": "panel",
				"method": "GET",
				"href": "%s/api/plans/planID/panel"
			},
			{
				"ref": "runnow",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/plans/planID/runnow"
			}
		],
		"id": "planID"
	}]`

const planJSONPayloadMissingLink = `[{
		"defaultPanelRepo": "git@github.com:CloudCoreo/default-panel.git",
		"defaultPanelDirectory": "panel",
		"defaultPanelBranch": "master",
		"name": "Audit-S3",
		"iamUserAccessKeyId": "IAMUSERACCESSKEYID",
		"iamUserId": "iamUserId",
		"iamUserSecretAccessKey": "iamUserSecretAccessKey",
		"snsSubscriptionArn": "arn:aws:sns:us-west-1:12312123:coreo-asi-planID:f50ac4e8-82a8-4291-a0c3-d9e299f79d8d",
		"sqsArn": "arn:aws:sqs:us-west-1:12312123:coreo-asi-planID",
		"sqsUrl": "sqsUrl",
		"topicArn": "arn:aws:sns:us-west-1:12312123:coreo-asi-planID",
		"defaultRegion": "us-east-1",
		"refreshInterval": 1,
		"revision": "HEAD",
		"branch": "master",
		"enabled": false,
		"id": "planID"
	}]`

const planJSONPayloadPlanEnabled = `{
		"defaultPanelRepo": "git@github.com:CloudCoreo/default-panel.git",
		"defaultPanelDirectory": "panel",
		"defaultPanelBranch": "master",
		"name": "Audit-S3",
		"iamUserAccessKeyId": "IAMUSERACCESSKEYID",
		"iamUserId": "iamUserId",
		"iamUserSecretAccessKey": "iamUserSecretAccessKey",
		"snsSubscriptionArn": "arn:aws:sns:us-west-1:12312123:coreo-asi-planID:f50ac4e8-82a8-4291-a0c3-d9e299f79d8d",
		"sqsArn": "arn:aws:sqs:us-west-1:12312123:coreo-asi-planID",
		"sqsUrl": "sqsUrl",
		"topicArn": "arn:aws:sns:us-west-1:12312123:coreo-asi-planID",
		"defaultRegion": "us-east-1",
		"refreshInterval": 1,
		"revision": "HEAD",
		"branch": "master",
		"enabled": true,
		"id": "planID"
	}`

const planJSONPayloadPlanDisabled = `{
		"defaultPanelRepo": "git@github.com:CloudCoreo/default-panel.git",
		"defaultPanelDirectory": "panel",
		"defaultPanelBranch": "master",
		"name": "Audit-S3",
		"iamUserAccessKeyId": "IAMUSERACCESSKEYID",
		"iamUserId": "iamUserId",
		"iamUserSecretAccessKey": "iamUserSecretAccessKey",
		"snsSubscriptionArn": "arn:aws:sns:us-west-1:12312123:coreo-asi-planID:f50ac4e8-82a8-4291-a0c3-d9e299f79d8d",
		"sqsArn": "arn:aws:sqs:us-west-1:12312123:coreo-asi-planID",
		"sqsUrl": "sqsUrl",
		"topicArn": "arn:aws:sns:us-west-1:12312123:coreo-asi-planID",
		"defaultRegion": "us-east-1",
		"refreshInterval": 1,
		"revision": "HEAD",
		"branch": "master",
		"enabled": false,
		"id": "planID"
	}`

const planConfigPayload = `{
  "gitRevision": "HEAD",
  "links": [
	{
		"ref": "self",
		"method": "GET",
		"href": "%s/api/planconfigs/planConfigID"
	},
	{
		"ref": "plan",
		"method": "GET",
		"href": "%s/api/plans/planID"
	}
  ],
  "gitBranch": "master",
  "variables": {
  },
  "planId": "planID",
  "id": "ID"
}`

func TestGetPlansSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(planJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlans(context.Background(), "teamID", "compositeID")
	assert.Nil(t, err, "GetPlans shouldn't return error.")
}

func TestGetPlansFailureInvalidCompositeID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(planJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlans(context.Background(), "teamID", "invalidCompositeID")
	assert.NotNil(t, err, "GetPlans should return error.")
}

func TestGetPlansFailureInvalidMissingPlanLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(planJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(compositeJSONPayloadForPlanMissingPlanLinks).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlans(context.Background(), "teamID", "compositeID")
	assert.NotNil(t, err, "GetPlans should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestGetPlansFailureInvalidPlanResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlans(context.Background(), "teamID", "compositeID")
	assert.NotNil(t, err, "GetPlans should return error.")
}

func TestGetPlansFailureNoPlansFound(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(`[]`).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlans(context.Background(), "teamID", "compositeID")
	assert.NotNil(t, err, "GetPlans should return error.")
	assert.Equal(t, "No plans found under team team ID teamID and composite ID compositeID.", err.Error())
}

func TestGetPlanByIDSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(planJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlanByID(context.Background(), "teamID", "compositeID", "planID")
	assert.Nil(t, err, "GetPlanByID shouldn't return error.")
}

func TestGetPlanByIDFailurePlanNotFound(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(planJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlanByID(context.Background(), "teamID", "compositeID", "invalidPlanID")
	assert.NotNil(t, err, "GetPlanByID should return error.")
	assert.Equal(t, "No plan with ID invalidPlanID found under team ID teamID and composite ID compositeID.", err.Error())
}

func TestGetPlanByIDFailureInvalidPlanResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlanByID(context.Background(), "teamID", "compositeID", "invalidPlanID")
	assert.NotNil(t, err, "GetPlanByID should return error.")
}

func TestEnablePlanSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanEnabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.EnablePlan(context.Background(), "teamID", "compositeID", "planID")
	assert.Nil(t, err, "EnablePlan shouldn't return error.")
}

func TestEnablePlanFailureInvalidPlanResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanEnabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.EnablePlan(context.Background(), "teamID", "compositeID", "planID")
	assert.NotNil(t, err, "EnablePlan should return error.")
}

func TestEnablePlanFailureInvalidPlanID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanEnabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.EnablePlan(context.Background(), "teamID", "compositeID", "invalidPlanID")
	assert.NotNil(t, err, "EnablePlan should return error.")
	assert.Equal(t, "Failed to enable plan ID invalidPlanID found under team ID teamID and composite ID compositeID.", err.Error())
}

func TestEnablePlanFailureToUpdatePlan(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanDisabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.EnablePlan(context.Background(), "teamID", "compositeID", "invalidPlanID")
	assert.NotNil(t, err, "EnablePlan should return error.")
	assert.Equal(t, "Failed to enable plan ID invalidPlanID found under team ID teamID and composite ID compositeID.", err.Error())
}

func TestEnablePlanFailureToUpdatePlanBadRequest(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanDisabled).WithStatus(http.StatusBadRequest)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.EnablePlan(context.Background(), "teamID", "compositeID", "planID")
	assert.NotNil(t, err, "EnablePlan should return error.")
}

func TestEnablePlanFailureMissingSelfLinks(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanDisabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(planJSONPayloadMissingLink).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.EnablePlan(context.Background(), "teamID", "compositeID", "planID")
	assert.NotNil(t, err, "EnablePlan should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestDisablePlanSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanDisabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.DisablePlan(context.Background(), "teamID", "compositeID", "planID")
	assert.Nil(t, err, "DisablePlan shouldn't return error.")
}

func TestDisablePlanFailureInvalidPlanResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanDisabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.DisablePlan(context.Background(), "teamID", "compositeID", "planID")
	assert.NotNil(t, err, "DisablePlan should return error.")
}

func TestDisablePlanFailureInvalidPlanID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanDisabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.DisablePlan(context.Background(), "teamID", "compositeID", "invalidPlanID")
	assert.NotNil(t, err, "DisablePlan should return error.")
	assert.Equal(t, "Failed to disable plan ID invalidPlanID found under team ID teamID and composite ID compositeID.", err.Error())
}

func TestDisablelanFailureToUpdatePlan(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanEnabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.DisablePlan(context.Background(), "teamID", "compositeID", "invalidPlanID")
	assert.NotNil(t, err, "DisablePlan should return error.")
	assert.Equal(t, "Failed to disable plan ID invalidPlanID found under team ID teamID and composite ID compositeID.", err.Error())
}

func TestDisablePlanFailureToUpdatePlanBadRequest(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanEnabled).WithStatus(http.StatusBadRequest)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.DisablePlan(context.Background(), "teamID", "compositeID", "planID")
	assert.NotNil(t, err, "DisablePlan should return error.")
}

func TestDisablePlanFailureMissingSelfLinks(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanEnabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(planJSONPayloadMissingLink).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.DisablePlan(context.Background(), "teamID", "compositeID", "planID")
	assert.NotNil(t, err, "DisablePlan should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestInitPlanSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID/planconfig").WithMethod("GET").WithBody(planConfigPayload).WithStatus(http.StatusOK)
	ts.Path("/api/plans/planid").WithMethod("PUT").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("POST").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.InitPlan(context.Background(), "branch", "name", "region", "teamID", "cloudID", "compositeID", "revision", 1)

	assert.Nil(t, err, "Plan init failed")
}

func TestCreatePlanSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planid").WithMethod("PUT").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/plans/planID").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/planconfigs/planConfigID").WithMethod("PUT").WithBody(fmt.Sprintf(planConfigPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreatePlan(context.Background(), []byte(fmt.Sprintf(planConfigPayload, ts.URL, ts.URL)))

	assert.Nil(t, err, "Plan creation failed")
	//assert.True(t, false, plan.IsDraft)
}

func TestPanelSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPanelInfo(context.Background(), "teamID", "compositeID", "planID")

	assert.Nil(t, err, "Plan planel failed")
}