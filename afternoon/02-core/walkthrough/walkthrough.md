# Core dumps

In this section students will learn how to create and debug core dumps.

## What students will learn

* Detailed information on every command the debugger offers
* How and when to use each command
* Flags and modifiers for each command

## Walkthrough

### Step 1

Explain what a core dump is:

* Snapshot of program
    * Contains memory dump of the process
    * Contains process status
* Usually produced when program exits
* Can be produced on demand
* Will not be able to manipulate execute as program is not actually running

---

### Step 3

Explain how to get system to create core dump when process crashes.


In shell:

```
make build-crash
make run-crash
```

Run the program but note we don't get a core dump.

Note the crash, but no core dump.

Then in shell:

```shell
make run-crash-gotraceback
```

Note how we get more information but still no core dump.

In shell

```shell
ulimit -c unlimited
```

* Explain this sets limit for core dump size
* By default it is 0 which means core dumps will not be produced

In shell:

```shell
make run-crash-gotraceback
```

Make note how we still don't have a core dump.

Then in shell:

```shell
make set-core-pattern
make run-crash-gotraceback
```

And note how we finally get a core dump!

Now let's start debugging it:

Then execute:

```shell
make dlv-core PID=<pid>
```

Try to continue it:

```
continue
```

Note error about not being able to continue core process. Explain why.

Start executing other commands:

```
list
stack
goroutines
threads
```

---

### Step 4

Explain how to create a core dump of running process via signal (pressing ctrl+\).


Run it:

```shell
make build-loop
make run-loop
```

Then type ctrl-\ and show we get a core dump.

Then, run it again:

```shell
make run-loop
```

And in another terminal run:

```shell
make send-sigabrt
```

And note how you can send SIGABRT to Go process to generate core dump
(if GOTRACEBACK=crash is set).

---

### Step 5

Explain how to get core dump of running process with gcore.

Then in one terminal run:

```shell
make run-loop
```

And in another shell:

```shell
make gcore
```

And note how core dump is created without having to crash the program.

---

### Step 6

Explain how to generate core dump within Delve session.

Then in shell:

```shell
make dlv-exec
```

Then in debug session:

```
continue
```

Let the process run for a bit then stop it.

After stopping execute:

```
dump core.dlv
exit
```

From command line:

```shell
dlv core ./bin/loop core.dlv
```

Find Goroutine that is running user code (likely `time.Sleep`), then:

```
goroutine <goroutineID> frame 2 print i
```

And note that we can get process state from core dump.

---

## Summary

In this sections students learned:

* How to ensure their system is setup to produce core dumps
* How to produce core dump when Go program crashes
* How to force Go program to crash and produce a core dump
* How to produce a core dump without Go program crashing
* How to produce core dump from within Delve debug session