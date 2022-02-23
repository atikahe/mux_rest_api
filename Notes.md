# Collections

## Context
Used to handle connection to remote servers, databases, external APIs, etc.
Used as per-call basis, passed to function before usage, not recommended to be declared in struct.


## What is defer?
- Only executes after all the function surrounding it returns
- When multiple defers -> LIFO
