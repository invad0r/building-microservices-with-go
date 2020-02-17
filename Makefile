check_install:
	which swagger || GO111MODULE=off go get -v -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	GO111MODULE=off swagger generate spec -o ./swagger.yml --scan-models