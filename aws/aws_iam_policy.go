// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamPolicy(client *Client) ([]Resource, error) {
	req := client.Iamconn.ListPoliciesRequest(&iam.ListPoliciesInput{
		Scope: "Local",
	})

	var result []Resource

	p := iam.NewListPoliciesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Policies {

			t := *r.CreateDate
			result = append(result, Resource{
				Type:      "aws_iam_policy",
				ID:        *r.Arn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
