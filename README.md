# textgen

textgen is CLI tool to generate text file from template file(s).

## Usage

`textgen [flags] [arguments]`

Flags:

* `-o` - output filename (default: stdout)

Arguments: one or more template filename(s).

## Template syntax

See [documention](http://golang.org/pkg/text/template).

## Example

`sample.t`:
```
{{ range $index, $value := list "hello" "world" }}
  {{ $value }}
{{ end }}
```

```
$ textgen sample.t
  hello

  world


```

## Additional template functions

### for _start_ _end_

Generates a sequence of integer numbers from `start` to `end`.

Examples:

* `for 1 5` - generates [1, 2, 3, 4, 5].
* `for 3 1` - generates [3, 2, 1].

### loop _arg1_ _arg2_ _..._

Convert arguments list (arg1, arg2, ...) to a sequence.

Examples:

* `list "hello" "world"` - returns ["hello", "world"].
