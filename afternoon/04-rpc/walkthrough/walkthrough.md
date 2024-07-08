# Delve JSON RPC API

In this section we cover how to use Delves JSON-RPC API.

## What students will learn

* How to build a client for the JSON-RPC API
* How to connect that client to remote Delve instance
* How to create a small service around it

## Walkthrough

### Step 1

Then in a shell:

```shell
make run-testprog-with-dlv
```

And in another shell:

```shell
make build
make run
```

---

### Step 2

Extend client to set tracepoints in remote session.

Update Go program to look like this:

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-delve/delve/service/api"
	"github.com/go-delve/delve/service/rpc2"
)

func main() {
	if len(os.Args) < 3 {
		bail("Not enough arguments passed in, please provide address to connect to and function to trace")
	}

	serverAddr := os.Args[1]
	funcToTrace := os.Args[2]

	// Create a new connection to the Delve debug server.
	// rpc2.NewClient will log.Fatal if connection fails so there
	// won't be an error to handle here.
	client := rpc2.NewClient(serverAddr)

	defer client.Disconnect(true)

	// Stop the program we are debugging.
	// The act of halting the program will return it's current state.
	state, err := client.Halt()
	if err != nil {
		bail(err)
	}

	bp := &api.Breakpoint{
		FunctionName: funcToTrace,
		Tracepoint:   true,
		LoadLocals: &api.LoadConfig{
			FollowPointers:     false,
			MaxVariableRecurse: 5,
			MaxStringLen:       100,
			MaxArrayValues:     50,
			MaxStructFields:    50,
		},
	}
	tracepoint, err := client.CreateBreakpoint(bp)
	if err != nil {
		bail(err)
	}
	defer client.ClearBreakpoint(tracepoint.ID)

	// Continue the program.
	stateChan := client.Continue()

	// Create JSON encoder to write to stdout.
	enc := json.NewEncoder(os.Stdout)

	for state = range stateChan {
		// Write state to stdout.
		enc.Encode(state)
	}
}

func bail(s interface{}) {
	fmt.Println(s)
	os.Exit(1)
}
```

Then in one shell:

```shell
make run-testprog-with-dlv
```

In another shell:

```shell
./bin/myclient localhost:9090 main.echoHandler
```

And in yet another shell:

```shell
curl localhost:8081/foo/bar
```

---

## Summary

In this sections students learned:

* How to use the JSON-RPC API.

[Next Section](scripting.md)