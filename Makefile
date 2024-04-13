build:
	go build -o build/aws-cdk
deploy:
	cdk deploy --profile instaAdmin
clean:
	rm -rf build
