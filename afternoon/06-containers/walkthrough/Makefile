###################
## HOST commands ##
###################

.PHONY: run
run:
	go run main.go

.PHONY: curl-app
curl-app:
	curl localhost:8080/foo/bar

.PHONY: install-delve
install-delve:
	go install github.com/go-delve/delve/cmd/dlv@latest

.PHONY: debug
debug:
	dlv debug

.PHONY: build
build:
	go build -gcflags="all=-N -l" -o bin/app

.PHONY: clean
clean:
	rm -rf bin

.PHONY: dlv-exec
dlv-exec:
	dlv exec bin/app

##############################
## Container Image commands ##
##############################

BASIC_IMG := buggy-basic
SCRATCH_IMG := buggy-scratch
DELVE_IMG := buggy-with-delve

# Build basic container image.
.PHONY: build-image
build-image:
	docker build --pull --rm -f build/Dockerfile-basic -t $(BASIC_IMG):latest .

.PHONY: build-scratch-image
build-scratch-image:
	docker build --pull --rm -f build/Dockerfile-scratch -t $(SCRATCH_IMG):latest .

.PHONY: build-scratch-image
build-debug-image:
	docker build --pull --rm -f build/Dockerfile-debug -t buggy-debug:latest .

# Run basic image.
.PHONY: run-basic-image
run-basic-image:
	docker run --detach -p 8080:8080 --rm $(BASIC_IMG)

.PHONY: build-debug-image
build-debug-image:
	docker build --pull --rm -f build/Dockerfile-debug -t buggy-debug:latest .

# Debug scratch image.
.PHONY: debug-scratch-image
debug-scratch-image:
	docker run -it --rm --pid="container:buggy-scratch" buggy-debug /bin/bash

# Run scratch image.
.PHONY: run-scratch-image
run-scratch-image:
	docker run --detach -p 8080:8080 --name=buggy-scratch --rm $(SCRATCH_IMG)

.PHONY: change-ptrace-yama
change-ptrace-yama:
	echo "0" | sudo tee /proc/sys/kernel/yama/ptrace_scope

# Copy dlv binary into basic container.
.PHONY: copy-dlv-to-container
copy-dlv-to-container:
	docker cp $$(which dlv) $$(docker ps -aqf "ancestor=$(BASIC_IMG)"):/dlv

# Run basic image with ptrace SYS_CAP.
.PHONY: run-basic-image-with-ptrace
run-basic-image-with-ptrace:
	docker run --detach --rm -p 8080:8080 --cap-add=SYS_PTRACE $(BASIC_IMG)

# Exec dlv within basic container.
.PHONY: exec-dlv-basic-container
exec-dlv-basic-container:
	docker exec -it $$(docker ps -aqf "ancestor=$(BASIC_IMG)") /dlv attach 1

# Exec dlv within basic container, using substitute path config.
.PHONY: exec-dlv-basic-container-with-src
exec-dlv-basic-container-with-src:
	docker cp $$(pwd) $$(docker ps -aqf "ancestor=$(BASIC_IMG)"):/src
	docker cp /usr/local/go $$(docker ps -aqf "ancestor=$(BASIC_IMG)"):/goroot
	docker cp hack/delve-container-initfile $$(docker ps -aqf "ancestor=$(BASIC_IMG)"):/delve-container-initfile
	docker exec -it $$(docker ps -aqf "ancestor=$(BASIC_IMG)") /dlv --init=/delve-container-initfile attach 1

# Stop running basic container.
.PHONY: stop-basic-container
stop-basic-container:
	docker stop $$(docker ps -aqf "ancestor=$(BASIC_IMG)")

# Build docker image containing Delve binary already.
.PHONY: build-image-with-delve
build-image-with-delve:
	docker build --pull --rm -f build/Dockerfile-with-delve -t $(DELVE_IMG):latest .

# Run docker image containing delve binary.
.PHONY: run-dlv-container
run-dlv-container:
	docker run --cap-add=SYS_PTRACE --rm --detach -p 8080:8080 -p 9090:9090 $(DELVE_IMG)

# Connect to headless dlv server within container.
.PHONY: connect-to-remote-dlv
connect-to-remote-dlv:
	dlv connect localhost:9090

# Connect to headless dlv server within container.
.PHONY: connect-to-remote-dlv-with-src
connect-to-remote-dlv-with-src:
	dlv --init=hack/delve-remote-initfile connect localhost:9090

# Stop container running headless dlv server.
.PHONY: stop-dlv-container
stop-dlv-container:
	docker stop $$(docker ps -aqf "ancestor=$(DELVE_IMG)")