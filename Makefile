run:
	cd $(d) && go run . && cd ..

test:
	cd $(d) && go test ./... && cd ..