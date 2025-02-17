// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package servicecatalog

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
		"aws_servicecatalog_constraint":            DataSourceConstraint,
		"aws_servicecatalog_launch_paths":          DataSourceLaunchPaths,
		"aws_servicecatalog_portfolio":             DataSourcePortfolio,
		"aws_servicecatalog_portfolio_constraints": DataSourcePortfolioConstraints,
		"aws_servicecatalog_product":               DataSourceProduct,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_servicecatalog_budget_resource_association":     ResourceBudgetResourceAssociation,
		"aws_servicecatalog_constraint":                      ResourceConstraint,
		"aws_servicecatalog_organizations_access":            ResourceOrganizationsAccess,
		"aws_servicecatalog_portfolio":                       ResourcePortfolio,
		"aws_servicecatalog_portfolio_share":                 ResourcePortfolioShare,
		"aws_servicecatalog_principal_portfolio_association": ResourcePrincipalPortfolioAssociation,
		"aws_servicecatalog_product":                         ResourceProduct,
		"aws_servicecatalog_product_portfolio_association":   ResourceProductPortfolioAssociation,
		"aws_servicecatalog_provisioned_product":             ResourceProvisionedProduct,
		"aws_servicecatalog_provisioning_artifact":           ResourceProvisioningArtifact,
		"aws_servicecatalog_service_action":                  ResourceServiceAction,
		"aws_servicecatalog_tag_option":                      ResourceTagOption,
		"aws_servicecatalog_tag_option_resource_association": ResourceTagOptionResourceAssociation,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ServiceCatalog
}

var ServicePackage = &servicePackage{}
