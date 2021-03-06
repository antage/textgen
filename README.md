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
Sequence is lazy (implemented via Go's channel).

Examples:

* `for 1 5` - returns [1, 2, 3, 4, 5].
* `for 3 1` - returns [3, 2, 1].

### loop _arg1_ _arg2_ _..._

Convert arguments list (arg1, arg2, ...) to a sequence.

Examples:

* `list "hello" "world"` - returns ["hello", "world"].

### map _key1_ _value1_ _key2_ _value2_ _..._

Make a map. Only string keys are supported.

Example:

* `map "key1" 2 "key2" "s"` - returns {"key1": 2, "key2": "s"}.

### uppercase _str_

Transform `str` string to upper case.

Example:

* `uppercase "hello"` - returns "HELLO".

### lowercase _str_

Transform `str` string to lower case.

Example:

* `lowercase "HELLO"` - returns "hello".

### toInt _val_

Convert `val` to `int` type.

Example:

* `toInt 5.5` - returns 5.
