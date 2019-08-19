# Problem

## All these functions operate on a *parser value:
+ `main.go`:
  + `summarize(*p parser)`
+ `parser.go`:
  + `func parse(p *parser, line string) (r result)`
  + `func update(p *parser, r result)`
  + `func err(p *parser) error`

## What is an object in Go?
+ Object = A value of a type : `&parser{}` is an object.
  + `type parser struct` is the definition (class/blueprint) of the object.

### Methods are behaviors of objects:
+ Parser needs to parse, update (analyse) the parsings, and report an error if any.
+ Parser also needs to summarize the results.
+ So, all these functions belong to a parser object.

# Solution

+ Attach the functions to the parser using methods.
+ Methods are the behavior of the parser.
+ Keep the methods that belong to a single type in the same file.

1. Move the `summarize` into `parser.go`.
2. `parser.go`: Use methods.
3. `main()`   : Use methods of the `parser`.

# Convince

+ `p.` shows the fields and methods belong to the `parser`.
+ `p.` selects a method from the method set of the `parser`.
+ `p.method` sends the `p` as the first parameter to the method.
+ `p` is the receiver: `*pointer`. It's a pointer receiver.
+ Receiver is copied to the method.