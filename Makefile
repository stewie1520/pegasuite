OUTPUT_DIR := ./build
DEBUG_PORT_API := :8181

build:
	go build -o ${OUTPUT_DIR}/main main.go
run: clean generate.swag build
	GIN_MODE=release ${OUTPUT_DIR}/main
run.dev:
	reflex -c ./reflex.conf
clean:
	rm -rf ${OUTPUT_DIR}
debug: clean generate.swag build dlv
generate.swag:
	swag init
dlv:
	dlv exec ${OUTPUT_DIR}/main --listen=${DEBUG_PORT_API} --headless=true --api-version=2 --accept-multiclient
