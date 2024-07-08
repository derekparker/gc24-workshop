.PHONY: build-crash
build-crash:
	go build -o bin/crash crash.go

.PHONY: run-crash
run-crash:
	./bin/crash

.PHONY: run-crash-gotraceback
run-crash-gotraceback:
	GOTRACEBACK=crash ./bin/crash

.PHONY: set-core-pattern
set-core-pattern:
	echo core | sudo tee /proc/sys/kernel/core_pattern

.PHONY: dlv-core
dlv-core:
	dlv core ./bin/crash core.$(PID)

.PHONY: build-loop
build-loop:
	go build -o bin/loop loop.go

.PHONY: run-loop
run-loop:
	GOTRACEBACK=crash ./bin/loop

.PHONY: send-sigabrt
send-sigabrt:
	kill -6 $$(pgrep -f bin/loop)

.PHONY: gcore
gcore:
	gcore $$(pgrep -f bin/loop)

.PHONY: dlv-exec
dlv-exec:
	dlv exec ./bin/loop