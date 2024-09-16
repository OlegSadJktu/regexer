# Regexer

Simple Go project for counting possible RegEx variations

## Main regex literals

- `(abc|sad)` - group. Allows any sequence from brackets divided by `|`. In this case satisfies only "abc" or "sad" strings
- `[0123]` - charset. Allows any symbol from square brackets 1 time. In this case satisfies only "0", "1", "2" or "3". Can contains ranges. **Currently not implemented**
- `a{2,4}` - quantifier. Allows previous symbol (or other entity) from `x` to `y` times. In this case satisfies only "aa", "aaa" or "aaaa"

## Examples

### 1. Regex `12(3|4){0,3}` read as

> Start with "12" and can contain a "3" or "4" from 0 to 3 times

#### Possible satisfying strings:

- `12` 0 repetitions
- `123` 1 repetition
- `124`
- `1233` 2 repetitions
- `1234`
- `1243`
- `1244`
- `12333` 3 repetitions
- `12334`
- `12343`
- `12344`
- `12433`
- `12434`
- `12443`
- `12444`

Totally 15

### 2. Regex `[a-z][0-9]` read as

> First symbol in range from "a" to "z" and second symbol in range from "0" to "9"

Simple combinatorics: `26 * 10 = 260`
Totally 260

### 3. Regex `[A-Za-z0-9а-яА-Я ]` read as

> Single symbol in range from "A" to "Z", or from "a" to "z", or from "0" to "9", or from "а" to "я", or from "А" to "Я" or space character

Simple combinatorics: `26 * 2 + 10 + 33 * 2 + 1 = 129`
Totally 129

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

- Groups
- Quantifiers
- Charsets
