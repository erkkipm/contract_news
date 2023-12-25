APP_BIN =  gen/go/news/
APP_PROTO = proto/news/

.PHONY: gen
gen:
	rm -f $(APP_BIN)*
	protoc -I proto $(APP_PROTO)*.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
