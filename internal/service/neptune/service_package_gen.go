// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []func(context.Context) (datasource.DataSourceWithConfigure, error) {
	return []func(context.Context) (datasource.DataSourceWithConfigure, error){}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []func(context.Context) (resource.ResourceWithConfigure, error) {
	return []func(context.Context) (resource.ResourceWithConfigure, error){}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_neptune_engine_version":        DataSourceEngineVersion,
		"aws_neptune_orderable_db_instance": DataSourceOrderableDBInstance,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_neptune_cluster":                 ResourceCluster,
		"aws_neptune_cluster_endpoint":        ResourceClusterEndpoint,
		"aws_neptune_cluster_instance":        ResourceClusterInstance,
		"aws_neptune_cluster_parameter_group": ResourceClusterParameterGroup,
		"aws_neptune_cluster_snapshot":        ResourceClusterSnapshot,
		"aws_neptune_event_subscription":      ResourceEventSubscription,
		"aws_neptune_global_cluster":          ResourceGlobalCluster,
		"aws_neptune_parameter_group":         ResourceParameterGroup,
		"aws_neptune_subnet_group":            ResourceSubnetGroup,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Neptune
}

var ServicePackage = &servicePackage{}
