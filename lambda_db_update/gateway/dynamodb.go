package gateway

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/ricardo-comar/identity-provider/lib_common/entity"
	"github.com/ricardo-comar/identity-provider/lib_common/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"encoding/json"
	"log"
)

const (
	TABLE_NAME = "employee"
)

func SaveEmployee(cfg aws.Config, model *model.EmployeeMessage) {

	svc := dynamodb.NewFromConfig(cfg)

	localendpoint, found := os.LookupEnv("LOCALSTACK_HOSTNAME")
	if found {
		svc = dynamodb.NewFromConfig(cfg, dynamodb.WithEndpointResolver(dynamodb.EndpointResolverFromURL("http://"+localendpoint+":"+os.Getenv("EDGE_PORT"))))
	}

	emp, err := getEmployee(svc, model)
	if err != nil {
		log.Fatal("Got error retrieving employee: ", err)
		return
	}

	emp.CadastroFuncionario = entity.EmployeeData{}
	j, _ := json.Marshal(model)
	json.Unmarshal(j, &emp.CadastroFuncionario)

	item, err := attributevalue.MarshalMap(emp)
	if err != nil {
		log.Fatal("Got error marshalling entity into attributeMap: ", err)
		return
	}

	_, err = svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item:      item,
	})
	if err != nil {
		log.Panicf("Got error calling PutItem: %s", err)
	}

}

func getEmployee(svc *dynamodb.Client, model *model.EmployeeMessage) (*entity.EmployeeRecordEntity, error) {

	result, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: model.ID},
		},
	})

	if err != nil {
		log.Panicf("Got error calling GetItem: %s", err)
		return nil, err
	}

	employee := entity.EmployeeRecordEntity{}
	employee.CadastroFuncionario = entity.EmployeeData{}
	employee.ID = model.ID
	employee.DataHoraCriacao = time.Now().Format(time.RFC3339)
	employee.TTL = strconv.FormatInt(time.Now().Add(time.Duration(300)*time.Second).UnixNano(), 10) // 5 minutos

	if result.Item != nil {
		attributevalue.UnmarshalMap(result.Item, &employee)
		employee.DataHoraAlteracao = time.Now().Format(time.RFC3339)
	}

	return &employee, nil
}
