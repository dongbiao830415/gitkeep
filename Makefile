all:
	GOOS=linux CGO_ENABLED=0 go build -v -ldflags="-s -w"
	GOOS=windows go build -v -ldflags="-s -w"
debug:
	go build -v -gcflags="all=-N -l"
