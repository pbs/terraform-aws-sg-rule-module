package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func testSGRule(t *testing.T, variant string) {
	t.Parallel()

	terraformDir := fmt.Sprintf("../examples/%s", variant)

	terraformOptions := &terraform.Options{
		TerraformDir: terraformDir,
		LockTimeout:  "5m",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	expectedName := fmt.Sprintf("example-tf-sg-rule-%s", variant)

	expectedLambdaName := expectedName
	lambdaName := terraform.Output(t, terraformOptions, "lambda_name")
	assert.Equal(t, expectedLambdaName, lambdaName)

	session, err := session.NewSession()
	if err != nil {
		t.Fatalf("Failed to create AWS session: %v", err)
	}

	lambdaSvc := lambda.New(session)

	invokeOutput, err := lambdaSvc.Invoke(&lambda.InvokeInput{
		FunctionName: &lambdaName,
	})

	if err != nil {
		t.Fatal(err)
	}

	var lambdaResp struct {
		Result string `json:"result"`
		Error  string `json:"error"`
	}

	err = json.Unmarshal(invokeOutput.Payload, &lambdaResp)

	if err != nil {
		t.Fatal(err)
	}

	if *invokeOutput.StatusCode != 200 {
		t.Logf("lambda response: %v", lambdaResp)
		t.Fatalf("Expected status code 200, got %d", *invokeOutput.StatusCode)
	}

	if lambdaResp.Result == "failure" {
		t.Logf("lambda response: %v", lambdaResp)
		t.Fatalf("Expected result 'success', got 'failure'")
	}
}
