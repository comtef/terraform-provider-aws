// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func testAccTransitGatewayRouteTableRoutesDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_ec2_transit_gateway_route_table_routes.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTransitGatewayRouteTableRoutesDataSourceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckResourceAttrGreaterThanValue(dataSourceName, "routes.#", 0),
				),
			},
		},
	})
}

func testAccTransitGatewayRouteTableRoutesDataSource_filter(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_ec2_transit_gateway_route_table_routes.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTransitGatewayRouteTableRoutesDataSourceConfig_filter(rName),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckResourceAttrGreaterThanValue(dataSourceName, "routes.#", 1),
				),
			},
		},
	})
}

func testAccTransitGatewayRouteTableRoutesDataSourceConfig_basic(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
  tags = {
    Name = %[1]q
  }
}

resource "aws_ec2_transit_gateway_vpc_attachment" "test" {
  subnet_ids         = aws_subnet.test[*].id
  transit_gateway_id = aws_ec2_transit_gateway.test.id
  vpc_id             = aws_vpc.test.id

  transit_gateway_default_route_table_association = false

  tags = {
    Name = %[1]q
  }
}

resource "aws_ec2_transit_gateway_route_table" "test" {
  transit_gateway_id = aws_ec2_transit_gateway.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_ec2_transit_gateway_route_table_association" "test" {
  transit_gateway_attachment_id  = aws_ec2_transit_gateway_vpc_attachment.test.id
  transit_gateway_route_table_id = aws_ec2_transit_gateway_route_table.test.id
}

resource "aws_ec2_transit_gateway_route_table_propagation" "test" {
  transit_gateway_attachment_id  = aws_ec2_transit_gateway_vpc_attachment.test.id
  transit_gateway_route_table_id = aws_ec2_transit_gateway_route_table.test.id
}

data "aws_ec2_transit_gateway_route_table_routes" "test" {
  transit_gateway_route_table_id = aws_ec2_transit_gateway_route_table.test.id
  depends_on                     = [aws_ec2_transit_gateway_route_table_association.test, aws_ec2_transit_gateway_route_table_propagation.test]
}
`, rName))
}

func testAccTransitGatewayRouteTableRoutesDataSourceConfig_filter(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
  tags = {
    Name = %[1]q
  }
}

resource "aws_ec2_transit_gateway_vpc_attachment" "test" {
  subnet_ids         = aws_subnet.test[*].id
  transit_gateway_id = aws_ec2_transit_gateway.test.id
  vpc_id             = aws_vpc.test.id

  transit_gateway_default_route_table_association = false

  tags = {
    Name = %[1]q
  }
}

resource "aws_ec2_transit_gateway_route_table" "test" {
  transit_gateway_id = aws_ec2_transit_gateway.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_ec2_transit_gateway_route_table_association" "test" {
  transit_gateway_attachment_id  = aws_ec2_transit_gateway_vpc_attachment.test.id
  transit_gateway_route_table_id = aws_ec2_transit_gateway_route_table.test.id
}

resource "aws_ec2_transit_gateway_route_table_propagation" "test" {
  transit_gateway_attachment_id  = aws_ec2_transit_gateway_vpc_attachment.test.id
  transit_gateway_route_table_id = aws_ec2_transit_gateway_route_table.test.id
}

data "aws_ec2_transit_gateway_route_table_routes" "test" {
  transit_gateway_route_table_id = aws_ec2_transit_gateway_route_table.test.id

  filter {
    name   = "type"
    values = ["propagated"]
  }

  depends_on = [aws_ec2_transit_gateway_route_table_association.test, aws_ec2_transit_gateway_route_table_propagation.test]
}
`, rName))
}
