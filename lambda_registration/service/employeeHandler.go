package service

import (
	"encoding/xml"
	"log"

	"github.com/ricardo-comar/identity-provider/lib_common/model"
	"github.com/ricardo-comar/identity-provider/registration/core"
	"github.com/ricardo-comar/identity-provider/registration/gateway"

	"github.com/go-resty/resty/v2"
)

func EmployeeService(ctx *model.ExecutionContext) (bool, error) {

	c, err := core.New()
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	log.Println("*** Solicitando dados de funcionários...")
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		SetHeader("X-API-Key", "4da852a0").
		Get("https://my.api.mockaroo.com/employees.xml")
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	if resp.IsError() {
		log.Fatal(resp.Error())
		return false, err
	}

	log.Println("*** Realizando Marshall dos dados...")
	registries := model.EmployeeRegistries{}
	err = xml.Unmarshal(resp.Body(), &registries)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	log.Printf("Registros recebidos: %d", len(registries.Registries))

	for _, registry := range registries.Registries {
		_, err = gateway.SendMessage(c, ctx, registry)
	}

	return true, err
}
