SRC = ./src/...
COVER = cover/cover

test:
	go test -p 1 $(SRC)

test-cover:
	mkdir -p cover && go test -p 1 $(SRC) -v -coverprofile $(COVER).out && go tool cover -html $(COVER).out -o $(COVER).html && open $(COVER).html