# Generics function and interface

## Prerequisites
- **[Tutorial: Get started with Go](https://go.dev/doc/tutorial/getting-started)**
- **Some programming experience.** 
    The code here is pretty simple, but it helps to know something about functions.

- **A tool to edit your code.** 
    ny text editor you have will work fine. Most text editors have good support for Go. The most popular are VSCode (free), GoLand (paid), and Vim (free).

- **A command terminal.** 
    Go works well using any terminal on Linux and Mac, and on PowerShell or cmd in Windows.

- **An environment that supports fuzzing.** 
    Go fuzzing with coverage instrumentation is only available on **AMD64** and **ARM64** architectures currently.

## Tutorial

Go tutorials topic [Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz) : 
Introduces the basics of fuzzing in Go. Fuzzing can generate inputs to your tests that can catch edge cases and security issues that you may have missed.

**Fuzz** :
Until this commit (Apr 21, 2023), The fuzzing arguments can only be the following types:
- string, []byte
- int, int8, int16, int32/rune, int64
- uint, uint8/byte, uint16, uint32, uint64
- float32, float64
- bool

For more information, refer below docs:
- [Go Fuzzing docs](https://go.dev/security/fuzz/#requirements)
- [Go Fuzzing glossary](https://go.dev/security/fuzz/#glossary)