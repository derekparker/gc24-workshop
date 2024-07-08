# Record & Replay Debugging

In this section students will learn how to use record & replay debugging.

## What students will learn

* What RR is
* How RR works
* Why record & replay debugging is so useful
* How to setup their system to use RR

## Walkthrough

### Step 1

Explain about RR:

* Only works on Linux / amd64 (maybe a little bit on ARM, experimentally)
* Has trouble running in VMs (needs hardware performance counters)
* Records memory layout
* Records all nondeterminism
* Runs programs single threaded
* Eliminates VDSO optimizations
* Explain how Delve is able to use RR as backend

---

### Step 2

Explain getting host setup for running RR.

Then run:

```shell
make set-perf-event-paranoid
```

Then in shell:

```shell
make check-perf-list
```

Tell students to look for output of [Hardware event].

---

### Step 3

Record using RR, replay with Delve.

Then execute:

```shell
make build
```

Then record the program:

```shell
make rr-record
```

Then replay with Delve:

```shell
make dlv-replay TRACE=<path to trace>
```

---

### Step 4

Use Delve to record and replay.

Then execute:

```shell
make dlv-rr
```

Explain that Delve has taken care of all steps to record and replay binary.

Continue to end then restart and continue to show result is always the same:

```
continue
restart
continue
```

Then prove that memory layout is the same:

```
restart
continue main.main
print runtime.curg.stack
restart
continue main.main
print runtime.curg.stack
```

Now start explaining some RR specific features.

Explore checkpoints:

```
restart
continue main.main
check
continue
restart c1
```

Reverse execution:

```
break main.go:21
continue
continue
rev continue
```

---

## Summary

In this sections students learned:

* How to record programs
* How to replay programs
* How to debug recorded programs