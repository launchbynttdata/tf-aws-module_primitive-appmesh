package test

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

const (
	base            = "../../examples/"
	testVarFileName = "/test.tfvars"
)

var standardTags = map[string]string{
	"provisioner": "Terraform",
}

func TestAppMesh(t *testing.T) {
	t.Parallel()
	stage := test_structure.RunTestStage

	files, err := os.ReadDir(base)
	if err != nil {
		assert.Error(t, err)
	}
	for _, file := range files {
		dir := base + file.Name()
		if file.IsDir() {
			defer stage(t, "teardown_appmesh", func() { tearDownAppMesh(t, dir) })
			stage(t, "setup_and_test_appmesh", func() { setupAndTestAppMesh(t, dir) })
		}
	}
}

func setupAndTestAppMesh(t *testing.T, dir string) {

	terraformOptions := &terraform.Options{
		TerraformDir: dir,
		VarFiles:     []string{dir + testVarFileName},
		NoColor:      true,
		Logger:       logger.Discard,
	}
	expectedPatternARN := "^arn:aws:appmesh:[a-z0-9-]+:[0-9]{12}:mesh/.+$"

	test_structure.SaveTerraformOptions(t, dir, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	actualARN := terraform.Output(t, terraformOptions, "arn")

	expectedName, err := terraform.GetVariableAsStringFromVarFileE(t, dir+testVarFileName, "name")
	expectedSpecEgressFilterType := terraform.GetVariableAsStringFromVarFile(t, dir+testVarFileName, "spec_egress_filter_type")

	if err == nil {
		actualName := terraform.Output(t, terraformOptions, "id")
		assert.True(t, strings.HasPrefix(actualName, expectedName), actualName, "Name did not match expected")
	}

	assert.Regexp(t, expectedPatternARN, actualARN, "ARN does not match expected pattern")

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile(os.Getenv("AWS_PROFILE")),
	)
	if err != nil {
		assert.Error(t, err, "can't connect to aws")
	}

	client := appmesh.NewFromConfig(cfg)

	input := &appmesh.DescribeMeshInput{
		MeshName: aws.String(expectedName),
	}

	result, err := client.DescribeMesh(context.TODO(), input)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("The expected mesh was not found %s", err.Error()))

	}

	mesh := result.Mesh
	actualName := *mesh.MeshName
	actualEgressStatus := string(mesh.Spec.EgressFilter.Type)

	if err == nil {
		assert.Equal(t, expectedName, actualName, "Mesh name does not match")
		assert.Equal(t, expectedSpecEgressFilterType, actualEgressStatus, "Mesh Egress does not match")

		expectedTags, err := terraform.GetVariableAsMapFromVarFileE(t, dir+testVarFileName, "tags")

		if err == nil {
			result2, err2 := client.ListTagsForResource(context.TODO(), &appmesh.ListTagsForResourceInput{ResourceArn: aws.String(actualARN)})
			if err2 != nil {
				assert.Error(t, err2, "Failed to retrieve tags from AWS")
			}
			// convert AWS Tag[] to map so we can compare
			actualTags := map[string]string{}
			for _, tag := range result2.Tags {
				actualTags[*tag.Key] = *tag.Value
			}

			// add the standard tags to the expected tags
			for k, v := range standardTags {
				expectedTags[k] = v
			}
			expectedTags["env"] = actualTags["env"]
			assert.True(t, reflect.DeepEqual(actualTags, expectedTags), fmt.Sprintf("tags did not match, expected: %v\nactual: %v", expectedTags, actualTags))
		}
	}
}

func tearDownAppMesh(t *testing.T, dir string) {
	terraformOptions := test_structure.LoadTerraformOptions(t, dir)
	terraformOptions.Logger = logger.Discard
	terraform.Destroy(t, terraformOptions)
}
