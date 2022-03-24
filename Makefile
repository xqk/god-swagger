

mac_amd64:
	GOARCH=amd64 CGO_ENABLED=0 GOOS=darwin go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o god-swagger main.go
	#$(if $(shell command -v upx), upx god)
	mv god-swagger ~/go/bin/
