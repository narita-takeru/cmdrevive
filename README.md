
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

if you want monitoring recursion directories,

Monitoring directories separate space.

```
cmdrevive "./htmls/ ./htmls/users ./htmls/items ./htmls/products" ".html$" (start your application command)
```

It's Cool schell script code.
```
DIRS=`find ./htmls -type d | tr "\n" " "`
cmdrevive "$DIRS" ".html$" (your application) (arguments)
```

