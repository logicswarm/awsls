// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListDbSubnetGroup(client *Client) ([]Resource, error) {
	req := client.Rdsconn.DescribeDBSubnetGroupsRequest(&rds.DescribeDBSubnetGroupsInput{})

	var result []Resource

	p := rds.NewDescribeDBSubnetGroupsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.DBSubnetGroups {

			result = append(result, Resource{
				Type:      "aws_db_subnet_group",
				ID:        *r.DBSubnetGroupName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
