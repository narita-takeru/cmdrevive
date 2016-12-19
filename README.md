
# cmdrevive

cmdrevive is auto restart command.
restart trigger is file changed.

## Example 

```
cmdrevive ./htmls/ ".html$" (application) (arguments)
```

`./htmls/` is want monitoring file changed directory.
` ".html$"` is regex for target file name.

## Requirements

- go 

## Installation

```bash
$ go get github.com/narita-takeru/cmdrevive/cmd/cmdrevive
```

## Usage

```
cmdrevive ./htmls/ ".html$" (start your application command)
```

