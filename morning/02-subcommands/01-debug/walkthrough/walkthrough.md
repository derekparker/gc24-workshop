# Debug Command Walkthrough

## Start Basic Debug Session

1. Execute `dlv debug` at the command line
2. Set a breakpoint at main via `break main.main`
3. Continue to breakpoint
4. Type `next` twice
5. PANIC!
6. Restart process via `restart <Insert Name Here>`
7. Continue again, and next twice
8. Print value of `name` via `print name`
9. Show how you can convert types via `print []byte(name)`
10. Step into function via `step`
11. Print arguments with `args`
12. Execute `next` command
13. Update string via `call name = "<Insert Full Name>"`
14. Step out via `stepout`
15. Execute `next` 
16. Execute `locals`
17. Change value of `x` with `set x = 2`
18. Continue
19. PANIC!
20. NOTE: Stack variable moved into register for comparison
21. Set a breakpoint before panic via `frame 2 break -2`
22. Restart and Continue
23. Execute `stepi`
24. Note where R2 variable is updated and grab address
25. Set breakpoint there via `break *<address>`
26. Continue
27. Step instruction via `stepi` so register is updated
28. Verify address with `examinemem -x int(X2+24)`
29. Update value of stack address with `set *(*uintptr)(<addr from last step>) = 2`
30. Continue and note no more panic