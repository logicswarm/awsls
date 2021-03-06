// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListSpotFleetRequest(client *Client) ([]Resource, error) {
	req := client.Ec2conn.DescribeSpotFleetRequestsRequest(&ec2.DescribeSpotFleetRequestsInput{})

	var result []Resource

	p := ec2.NewDescribeSpotFleetRequestsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.SpotFleetRequestConfigs {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t := *r.CreateTime
			result = append(result, Resource{
				Type:      "aws_spot_fleet_request",
				ID:        *r.SpotFleetRequestId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				Tags:      tags,
				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
