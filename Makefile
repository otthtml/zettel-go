run:
	cd $(d) && go run main.go && cd ..

test:
	cd $(d) && go test ./... && cd ..