// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/batch"
)

func ListBatchJobDefinition(client *Client) ([]Resource, error) {
	req := client.Batchconn.DescribeJobDefinitionsRequest(&batch.DescribeJobDefinitionsInput{})

	var result []Resource

	p := batch.NewDescribeJobDefinitionsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.JobDefinitions {

			result = append(result, Resource{
				Type:      "aws_batch_job_definition",
				ID:        *r.JobDefinitionArn,
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
