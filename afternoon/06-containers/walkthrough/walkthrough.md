# Debug from container part 1

This section starts to introduce students to the difficult aspects of debugging within a container.

### Things to talk about in intro:

* Namespace isolation: pid, mount, network, etc...
* Resource limits
* Limits of what may be present in container filesystem

## Talk Walkthrough

*** BEFORE CONTINUING ***

Ensure the host yama/ptrace_scope is set to 1:

```shell
echo "1" | sudo tee /proc/sys/kernel/yama/ptrace_scope
```

### Step 1

Make `build` directory:

```shell
mkdir build
```

### Step 2

Copy `Dockerfile-basic` into the `build` directory:

```Dockerfile
# This Dockerfile shows using a builder container and then
# configuring the container to run the debugger from a scratch container.
FROM golang:1.17 as builder

WORKDIR /project

COPY . .

RUN go build -o /project/bin/app

FROM ubuntu:22.04

COPY --from=builder /project/bin/app .

CMD [ "./app" ]
```

### Step 3

Explain Dockerfile:

* We have a Go builder image
* We have a stage 2 which is Ubuntu
* Stage 2 could be scratch, but we'll go into that later

### Step 4

Copy Makefile command to build image:

```Makefile
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

```

Then execute the following in a shell:

```shell
make build-image
```

### Step 5

Copy the following to the Makefile:

```Makefile
.PHONY: run-basic-image
run-basic-image:
	docker run -it --detach -p 8080:8080 --rm $(BASIC_IMG)
```

Then execute the following to run the container:

```shell
make run-basic-image
```

Then in another shell execute the following to show the app is running in the container:

```shell
make curl-app
```

At this point take a second to explain:

* The process within the container is not visible to Delve from the host
* From within the container we cannot see host files (source code, etc...)

### Step 6

Copy the following command to the Makefile:

```Makefile
# Copy dlv binary into basic container.
.PHONY: copy-dlv-to-container
copy-dlv-to-container:
	docker cp $$(which dlv) $$(docker ps -aqf "ancestor=$(BASIC_IMG)"):/dlv
```

And then execute it:

```shell
make copy-dlv-to-container
```

At this point explain:

* The debugger is in the container, but we're outside the container mount namespace
* We must somehow exec into the container to be able to execute Delve

### Step 7

Copy the following to the Makefile:

```Makefile
# Exec dlv within basic container.
.PHONY: exec-dlv-basic-container
exec-dlv-basic-container:
	docker exec -it $$(docker ps -aqf "ancestor=$(BASIC_IMG)") /dlv attach 1
```

Execute the following:

```shell
make exec-dlv-basic-container
```

Note error about yama/ptrace_scope:

On systems with the Yama Linux Security Module (LSM) installed
the /proc/sys/kernel/yama/ptrace_scope file (available since Linux 3.4)
can be used to restrict the ability to trace a process with ptrace().

* Delve (and other debuggers) use ptrace under the hood to attach to process
* Blocking ptrace prevents Delve or other tools from working
* Explain we can change this on the host

### Step 8

Copy the following to the Makefile:

```Makefile
.PHONY: change-ptrace-yama
change-ptrace-yama:
	echo "0" | sudo tee /proc/sys/kernel/yama/ptrace_scope

```

Then execute:

```shell
make change-ptrace-yama
```

Then re-rerun the following command:

```shell
make exec-dlv-basic-container
```

Now we have a debug session actually started.

### Step 9

Explain we can also change this more surgically by adding certain linux capabilities
to the container when we start it.

Copy the following to the Makefile:

```Makefile
# Run basic image with ptrace SYS_CAP.
.PHONY: run-basic-image-with-ptrace
run-basic-image-with-ptrace:
	docker run -it --detach --rm -p 8080:8080 --cap-add=SYS_PTRACE $(BASIC_IMG)
```

Then execute the following:

```shell
make run-basic-image-with-ptrace && make copy-dlv-to-container && make exec-dlv-basic-container
```

At this point:

* Make note of how the debugger has now started and attached to the process.
* Execute `list` command within debugger and note how source code isn't found
* Execute `break helloServer` within debugger
* Execute `continue` within debugger

In split terminal:

```shell
make curl-app
```

Show that users source code is not there.

Stop debug session.

### Step 10

* Explain that since the container is in a different mount namespace we cannot
  see the source code, so even though the debugger is working we cannot see
  the source code we are stepping through.
* Explain we can use Delves config and some docker magic to get source code.

Execute the following in the shell:

```shell
mkdir hack
touch hack/delve-container-initfile
```

Copy the following to `hack/delve-container-initfile`:

```
# Configure substitute path rule for GOROOT.
config substitute-path /usr/local/go /goroot

# Configure substitute path rule for user code.
config substitute-path /project /src
```


Copy the following to the Makefile:

```Makefile
# Exec dlv within basic container, using substitute path config.
.PHONY: exec-dlv-basic-container-with-src
exec-dlv-basic-container-with-src:
	docker cp $$(pwd) $$(docker ps -aqf "ancestor=$(BASIC_IMG)"):/src
	docker cp /usr/local/go $$(docker ps -aqf "ancestor=$(BASIC_IMG)"):/goroot
	docker exec -it $$(docker ps -aqf "ancestor=$(BASIC_IMG)") /dlv --init=/src/hack/delve-container-initfile attach 1
```

Then execute the following in the shell:

```shell
make exec-dlv-basic-container-with-src
```

Then:

* Execute `list` command within debugger, note how we see source code
* Execute `break helloServer` within debugger
* Execute `continue` within debugger
* In split terminal execture `make curl-app`
* Note how we have user source code too

## Summary

* We learned how to configure container to be able to run debugger
* We learned how to copy Delve into a container
* We learned how to get source code into container
* We learned how to configure Delve to find source code

---

# Part 2

# Debug from container part 2

This section introduces how to debug a local scratch container.

### Step 1

Within terminal execute:

```shell
touch build/Dockerfile-scratch
```

And then paste the following into the file:

```Dockerfile
FROM golang:1.17 as builder

WORKDIR /project

COPY . .

ENV CGO_ENABLED=0
RUN go build -o /project/bin/app

FROM scratch

COPY --from=builder /project/bin/app /app

CMD [ "/app" ]
```

Then copy the following into the Makefile:

```Makefile
.PHONY: build-scratch-image
build-scratch-image:
	docker build --pull --rm -f build/Dockerfile-scratch -t $(SCRATCH_IMG):latest .
```

And then execute it:

```shell
make build-scratch-image
```

### Step 2

Create a new Dockerfile for our debug container:

```shell
touch build/Dockerfile-debug
```

Then copy the following into it:

```Dockerfile
FROM golang:1.17

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN mv /go/bin/dlv /dlv
COPY . /project
```

Then copy the following to the Makefile:

```Makefile
.PHONY: build-debug-image
build-debug-image:
	docker build --pull --rm -f build/Dockerfile-debug -t buggy-debug:latest .
```

Then run that command:

```shell
make build-debug-image
```

### Step 3

Copy the following to the Makefile:

```Makefile
# Run scratch image.
.PHONY: run-scratch-image
run-scratch-image:
	docker run -it --detach -p 8080:8080 --name=buggy-scratch --rm $(SCRATCH_IMG)
```

Start scratch container:

```shell
make run-scratch-image
```

Copy the following to the Makefile:

```Makefile
# Debug scratch image.
.PHONY: debug-scratch-image
debug-scratch-image:
	docker run -it --rm --pid="container:buggy-scratch" buggy-debug /bin/bash
```

Then attach our debug image to the process namespace:

```shell
make debug-scratch-image
```

### Step 4

Make note of how source code, everything is where it is supposed to be.

Removes steps to start debugging.

## Summary

In this section we learned:

* How to make a scratch image
* How to make a debug image
* How to use the debug image to attach to the scratch and debug it

---

# Part 3

# Debug from container part 3

This section shows how we can start the container with the debugger
already running and then connect from outside.

### Step 1

Execute the following in a shell:

```shell
touch build/Dockerfile-with-delve
```

Then copy the following to it:

```Dockerfile
# This Dockerfile shows how to build a container that contains the
# debugger already and then shows how to run your app via the debugger
# so that you can easily connect to it from the host.
FROM golang:1.17 as builder

WORKDIR /project

COPY . .

ENV CGO_ENABLED=0
RUN go build -o /project/bin/app
RUN go install github.com/go-delve/delve/cmd/dlv@latest

FROM ubuntu:22.04

COPY --from=builder /project/bin/app .
COPY --from=builder /go/bin/dlv /dlv

EXPOSE 9090

CMD [ "/dlv", "exec", "--continue", "--headless", "--accept-multiclient", "--api-version", "2", "--listen", "0.0.0.0:9090", "./app" ]
```

Copy the following to the Makefile:

```Makefile
# Build docker image containing Delve binary already.
.PHONY: build-image-with-delve
build-image-with-delve:
	docker build --pull --rm -f build/Dockerfile-with-delve -t $(DELVE_IMG):latest .
```

Then execute the following in a shell:

```shell
make build-image-with-delve
```

### Step 2

Copy the following to the Makefile:

```Makefile
# Run docker image containing delve binary.
.PHONY: run-dlv-container
run-dlv-container:
	docker run --cap-add=SYS_PTRACE --rm -it --detach -p 8080:8080 -p 9090:9090 $(DELVE_IMG)
```

Execute the following in a shell:

```shell
make run-dlv-container
```

Execute the following to prove it's running:

```shell
make curl-app
```

### Step 3

Copy the following to the Makefile:

```Makefile
# Connect to headless dlv server within container.
.PHONY: connect-to-remote-dlv
connect-to-remote-dlv:
	dlv connect localhost:9090
```

### Step 4

* Explain we have connected to the remote instance but we don't have source code.
* Explain in these images maybe include the source code as well to aid debugging.

## Summary

In this final container section we learned:

* How to create a container with Delve already running our app
* How to port forward from container
* How to connect to remote Delve instance in container
