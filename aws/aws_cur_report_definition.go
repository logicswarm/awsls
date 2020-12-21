// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
)

func ListCurReportDefinition(client *Client) ([]Resource, error) {
	req := client.Costandusagereportserviceconn.DescribeReportDefinitionsRequest(&costandusagereportservice.DescribeReportDefinitionsInput{})

	var result []Resource

	p := costandusagereportservice.NewDescribeReportDefinitionsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.ReportDefinitions {

			result = append(result, Resource{
				Type:      "aws_cur_report_definition",
				ID:        *r.ReportName,
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
