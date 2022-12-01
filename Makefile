build: compile

clean:
		@echo "Cleaning bin folder"
		@rm -rf bin

prepare:
		@echo "Creating bin folder"
		@mkdir -p bin

test: clean prepare		
		@echo "Testing lambdas"
		@find . -maxdepth 1 -mindepth 1 -type d  -name 'lambda_*'| while read dir; do\
			echo "Testing $$dir"; cd $$dir; go test -coverprofile=../bin/$$dir-coverage.out ./...; cd ..; \
		done

compile: clean prepare		
		@echo "Compiling lambdas"
		@find . -maxdepth 1 -mindepth 1 -type d -name 'lambda_*'| while read dir; do\
			echo "Compiling $$dir"; cd $$dir; GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../bin/$$dir handler/handler.go; cd ..; \
		done

package: build 
		@find . -maxdepth 1 -mindepth 1 -type d  -name 'lambda_*'| while read dir; do \
			zip -j bin/$$dir.zip bin/$$dir; \
		done

localstack:
	aws --endpoint http://localhost:4566 iam create-user --user-name test
	

event:
	aws --endpoint http://localhost:4566 events put-events --entries \
		--entries '[{"Time": "2016-01-14T01:02:03Z", "Source": "com.mycompany.myapp", "Resources": ["resource1", "resource2"], "DetailType": "myDetailType", "Detail": "{ \"key1\": \"value1\", \"key2\": \"value2\" }"}]'

msg:
	aws --endpoint http://localhost:4566 sqs send-message --queue-url http://localhost:4566/000000000000/identity-provider-sqs-employees --message-body "IOT-1 Temp: 51C"

gateway:
	curl http://localhost:4566/restapis/$(shell aws --endpoint-url=http://localhost:4566 apigateway get-rest-apis | jq -r '.items[0].id')/v1/\_user_request_/employees | jq
