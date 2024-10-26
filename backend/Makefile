portal:
	go run server/cmd/portal/main.go

openapi:
	go run server/cmd/openapi/main.go

exporter:
	go run server/cmd/export/main.go

work:
	go work init

mod:
	mkdir $(name)
	cd $(name) && go mod init $(name)
	go work use $(name)