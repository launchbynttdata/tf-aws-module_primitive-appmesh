package common

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/require"
)

func TestComplete(t *testing.T, ctx types.TestContext) {
	appmeshClient := appmesh.NewFromConfig(GetAWSConfig(t))
	meshName := terraform.Output(t, ctx.TerratestTerraformOptions(), "name")
	meshArn := terraform.Output(t, ctx.TerratestTerraformOptions(), "arn")

	t.Run("TestDoesMeshExist", func(t *testing.T) {
		output, err := appmeshClient.DescribeMesh(context.TODO(), &appmesh.DescribeMeshInput{MeshName: &meshName})
		if err != nil {
			t.Errorf("Error describing mesh: %v", err)
		}

		require.Equal(t, meshName, *output.Mesh.MeshName, "Expected mesh name to be %s, but got %s", meshName, *output.Mesh.MeshName)
		require.Equal(t, meshArn, *output.Mesh.Metadata.Arn, "Expected mesh ARN to be %s, but got %s", meshArn, *output.Mesh.Metadata.Arn)
	})
}

func GetAWSConfig(t *testing.T) (cfg aws.Config) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	require.NoErrorf(t, err, "unable to load SDK config, %v", err)
	return cfg
}
