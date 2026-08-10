package common

import (
	"context"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/require"
)

func awsRegion() string {
	for _, v := range []string{os.Getenv("AWS_DEFAULT_REGION"), os.Getenv("AWS_REGION")} {
		if v != "" {
			return v
		}
	}
	return "us-east-2"
}

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	appmeshClient := appmesh.NewFromConfig(GetAWSConfig(t))
	meshName := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "name")
	meshArn := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "arn")

	t.Run("TestDoesMeshExist", func(t *testing.T) {
		output, err := appmeshClient.DescribeMesh(context.Background(), &appmesh.DescribeMeshInput{MeshName: aws.String(meshName)})
		require.NoError(t, err, "describe mesh")

		require.Equal(t, meshName, *output.Mesh.MeshName, "Expected mesh name to be %s, but got %s", meshName, *output.Mesh.MeshName)
		require.Equal(t, meshArn, *output.Mesh.Metadata.Arn, "Expected mesh ARN to be %s, but got %s", meshArn, *output.Mesh.Metadata.Arn)
	})
}

func GetAWSConfig(t *testing.T) aws.Config {
	t.Helper()

	loadOpts := []func(*config.LoadOptions) error{
		config.WithRegion(awsRegion()),
	}
	if profile := os.Getenv("AWS_PROFILE"); profile != "" {
		loadOpts = append(loadOpts, config.WithSharedConfigProfile(profile))
	}

	cfg, err := config.LoadDefaultConfig(context.Background(), loadOpts...)
	require.NoErrorf(t, err, "unable to load SDK config, %v", err)
	return cfg
}
