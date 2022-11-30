dev:
	cd src && go run -tags git2git ./main.go

test:
	cd src && go test -v -cover ./...

test-report:
	cd src \
	&& go test -v -coverprofile=cover.out ./... \
	&& go tool cover -html=cover.out -o cover.html
