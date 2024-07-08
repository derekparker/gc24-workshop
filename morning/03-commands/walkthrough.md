# Debugger Commands

In this section we cover the commands to be used within a debug session.

## What students will learn

* Detailed information on every command the debugger offers
* How and when to use each command
* Flags and modifiers for each command

## Walkthrough

### Step 1

Start debug session:

```shell
make debug
```

Within debug session:

```
continue main.main
```

Explain how this sets a temporary breakpoint and continues to it, removing afterwards.

Then, use `next` to move forward a few lines:

```
next
next
```

Then demonstrate how we can jump into a function call:

```
step
```

Then explain how we can step out of a function:

```
stepout
```

Then explain how we can step single instructions.

First, disassemble to show which instruction we are on:

```
disassemble
```

Then once we've noted which instruction we are on, step to the next one:

```
step-instruction
```

Note how we've moved forward one CPU instruction.

Explain next how we can start the program and change the redirections:

In separate terminal:

```
echo "and now for something completely different" > stdin2.txt
```

In debug session:

```
restart <stdin2.txt
```

Then continue until it gets printed:

```
continue main.go:89
```

---

### Step 5

This section will focus on breakpoints.

Note we can assign name to breakpoint:

```
break mainfunc main.main
```

Then we can clear it by name:

```
clear mainfunc
```

Set another breakpoint on `main.main`:

```
break main.main
```

Then clean by id:

```
clear <id>
```

Show how to set breakpoint by regex:

```
break /^main.*/
```

Then how we can clear all of them:

```
clearall
```

Set breakpoint:

```
break m33 main.go:33
```

Set condition to only stop when breakpoint has been hit 5 times:

```
cond -hitcount m33 == 5
```

Then continue program:

```
continue
```

Note how the total hit count for the breakpoint is the number we set on the condition.

Clear the condition and the breakpoint:

```
cond -clear m33
clear m33
```

Now, set another breakpoint:

```
break m33 main.go:33
```

Set a command to execute when breakpoint is hit:

```
on m33 print id
```

Note how this is useful for when you always execute a series of commands
when a breakpoint is hit. This way you can automate it.

Next toggle breakpoint off and continue:

```
toggle m33
continue
```

Toggle it back on:

```
toggle m33
continue
```

Note how everything about breakpoint is preserved.

Note how we can change breakpoint to a tracepoint:

```
on m33 trace
continue
```

Note how we still get the same output from the other "on" command invocations.

Clear all breakpoints:

```
clearall
```

Continue to location:

```
continue main.go:62
```

Now note how we can set a watchpoint.

Note how we can set on read and write or either seperately.

```
watch -rw id
```

Explain error in output:

* Cannot watch for reads even though it will technically work.
* The runtime walks the stack so we would get stops on this erroneously.

Now try just with writes:

```
watch -w id
```

Now continue program:

```
continue
```

And note how we've stopped where the id is changed.

---

### Step 6

This section will cover viewing / writing to program variables and memory.

```
continue main.go:33
```

See all args to function:

```
args
```

See all local vars:

```
locals
```

Display variable anytime program stops:

```
display -a id
continue main.go:33
```

Show how the value of `id` is printed at the bottom once the program stops.

Show how we can examine raw memory (also note how print takes expressions):

```
continue main.go:38
x -fmt hex -count 20 -x &p
```

That will print 20 bytes starting at the address of the `p` variable.

Show how we can print type of variable:

```
whatis p
```

Now show we can change value:

```
print p.id
set p.id = 500
print p.id
```

Finally show how we can see register contents:

```
regs
```

Then show how we can see even more registers with:

```
regs -a
```

---

### Step 7

This section will cover threads and goroutines.

List all goroutines:

```
goroutines
```

Explain the output and that `*` means this is the goroutine we are on.

Now execute:

```
help goroutines
```

And explain help output including grouping, filtering, etc...

Explain how we can execute command on another goroutine:

```
goroutine <id> stack
```

Also explain that without args it prints info on current goroutine:

```
goroutine
```

Now show how we can list threads:

```
threads
```

Also explain how `*` shows the thread we are currently on.

Explain we can switch to another thread:

```
thread <id>
```

But we cannot execute command there.

Switch to goroutine running `wg.Wait`

Execute `print ch` to see list of goroutines waiting on channel

---

### Step 8

This section will focus on the program stack.

Type:

```
help stack
```

Then go through output and execute commands to help explain what the flags do.

Type:

```
up
```

Explain out context is now the parent (caller) frame.

Type:

```
down
```

Then explain we are back to the frame we were at before.

---

### Step 9

Type `help` and go through the "other" commands, showing expamples where it makes sense.

---

## Summary

In this section students learned all debugger commands in detail.