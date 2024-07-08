# Delve Scripting

In this section we cover how to script Delve.

## What students will learn

* How to add new commands to Delve via the scripting interface
* How to enter an interactive scripting session

## Walkthrough

### Step 1

In shell:

```shell
make debug
```

And within debugger:

```
source find_array_elem.star
break m65 main.go:65
cond m65 len(buf) == 100
continue
find_array "buf", lambda x: x.id == <id>
```

---

### Step 2

Within debugger:

```
source switch_to_goroutine_running.star
switch_to_g_running "main.produceValues"
```

---

### Step 3

Make note of `main` function.

In debugger:

```
source goroutine_start_line.star
gsl
```

---

### Step 3

Show off interactive session.

In debugger:

```
source -
goroutines()
goroutines().Goroutines
goroutines().Goroutines[0]
```

---

## Summary

In this sections students learned:

* How to use the starlark scripting interface.