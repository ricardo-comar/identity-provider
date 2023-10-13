package gateway

import (
	"context"
	"log"
	"math"
	"os"

	"github.com/ricardo-comar/identity-provider/lib_common/entity"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	TABLE_NAME = "employee"
)

func QueryEmployee(cfg aws.Config, previous string, size int) ([]entity.EmployeeRecordEntity, error) {

	svc := dynamodb.NewFromConfig(cfg)

	localendpoint, found := os.LookupEnv("LOCALSTACK_HOSTNAME")
	if found {
		svc = dynamodb.NewFromConfig(cfg, dynamodb.WithEndpointResolver(dynamodb.EndpointResolverFromURL("http://"+localendpoint+":4566")))
	}

	var employees []entity.EmployeeRecordEntity

	var prevMap map[string]types.AttributeValue
	if len(previous) > 0 {
		prevMap = map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: previous},
		}
	}

	log.Printf("Quering with previous %s and size %d", previous, size)

	input := dynamodb.NewScanPaginator(svc, &dynamodb.ScanInput{
		TableName:         aws.String(TABLE_NAME),
		Limit:             aws.Int32(int32(size)),
		ExclusiveStartKey: prevMap,
	})

	for input.HasMorePages() {
		out, err := input.NextPage(context.TODO())
		if err != nil {
			panic(err)
		}

		var emp []entity.EmployeeRecordEntity
		err = attributevalue.UnmarshalListOfMaps(out.Items, &emp)
		if err != nil {
			panic(err)
		}

		employees = append(employees, emp[:int(math.Min(float64(int(size)-len(employees)), float64(len(emp))))]...)

		if len(employees) == size {
			return employees, nil
		}

	}

	return employees, nil

}
