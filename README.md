# Build a Shell with Go

[GitHub repo](https://github.com/JamesBenji/GoShell)

## What surprised me

- Building a shell is not hard if one knows the OS APIs exposed by the language and how to use them. The foundation Read-Evaluate-Print-Loop (REPL) idea is the main thing to implement.
- The shell handles `cd`, `>`, `<`, and `|` by itself.

## What I learnt

- The main idea to follow when building a shell (REPL) and how to access delegate to the OS user commands.
- Working with a pipe in Go (Inter-process communication)

## How I used AI

- I used ChatGPT to understand the C examples and structure my Go program.
- ChatGPT helped me fix the Go code I wrote.

## References

- `https://healeycodes.com/building-a-shell`
- `https://www.cs.purdue.edu/homes/grr/SystemsProgrammingBook/Book/Chapter5-WritingYourOwnShell.pdf`
- `https://brennan.io/2015/01/16/write-a-shell-in-c/`
- `https://www.geeksforgeeks.org/c/making-linux-shell-c/`
