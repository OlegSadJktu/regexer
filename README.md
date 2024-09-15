# Regexer

Simple Go project for counting possible RegEx variations

## How to use

1. Regex in command line

> [!WARNING]
> Don't forget about escape characters

```
$ go run . --regex 123\(1\|2\)
2
$ go run . --regex 123\(1\|2\)\{1,2\}
6
```
2. From file (simple way)
```
$ go run . --file input.txt
123(3|4){0,3}
15
```

## Available possibilities
- Groups: `(asd|basd|ass)`. Has three options
- Quantifiers: `a{1,4}`. May be `a`, `aa`, `aaa` or `aaaa`.