# Delve Exec

In this section we cover how to debug a preexisting binary.

## What students will learn

* How to exec a binary
* How to build a binary ideal for debugging

## Walkthrough

### Step 1

*** Make sure we're at root of the project ***

In shell:

```shell
mkdir -p day-1/cmds/exec
cd day-1/cmds/exec
```

---

### Step 2

Run in shell:

```shell
make bin/execme
```

In shell:

```shell
make dlv-exec
```

In debugger:

First start transcript:

```
transcript debugpt1.txt
```

```
break Rectangle.Area
```

Note how there are multiple addresses for the breakpoint.

Next, try and set a breakpoint on `Rectangle.Width`:

```
break Rectangle.Width
```

Note how the debugger cannot set a breakpoint there as there is no information
for the debugger to use.

Next try and see if the function exists anywhere:

```
funcs Rectangle
```

Note how only `Area` and `Height` come back.


---

### Step 5

Explain (briefly) function inlining, Elf files and DWARF debug information.

Execute:

```shell
make build-showing-inlining
```

---

### Step 6

Now we will build without optimizations.

Execute:

```shell
make bin/execmenooptimizations
make dlv-exec-no-optimizations
```

In debugger:

```
break Rectangle.Area
break Rectangle.Width
```

* Note how we only see single address entry.
* Note how we can actually set a breakpoint on `Rectangle.Width`.

---

Split terminal in 2.

In first terminal:

```shell
make dlv-exec
```

In second terminal:

```shell
make dlv-exec-no-optimizations
```

In both debug sessions:

```
disassemble -l main.main
```

Show students how in the optimized there is no `call` instruction.
Show students then how in the unoptimized function there are `call` instructions for the previously inlined functions.

---

## Summary

In this section students learned:

* How to compile Go programs without optimizations
* How to use `dlv exec`
* How to disassemble code in Delve
* How to use the `objdump` tool
* How to get the Go compiler to give information on what it's doing