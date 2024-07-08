# Walkthrough

On Linux machine, exec:

    dlv exec ~/Code/go/bin/go -- build .

Enable follow mode via:

    target follow-exec -on

Set breakpoint on `Cmd.Wait`:

    break Cmd.Wait

Continue and then list targets:

    target list

Switch to another process:

    target switch <pid>