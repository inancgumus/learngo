### PROBLEM
+ `summarize()` knows a lot about the internals of the `parser`.
  + coupled to the `parser`.

## SOLUTION
+ remove: `parser.go` `sum` and `domains` fields
  + remove: `parser.go/newParser()`
  + change: `parser.go/parse(p *parser) error` -> `parse() ([]result, error)`
  + initialize: `sum` inside `parse()`
  + remove: `update()`
  + call: `sum` update in the `parse()`
  + collect the grouped results and return them from `parser()`

+ `summarize(p *parser)` -> `summarize([]result)`
+ in `summarize()`
  + `sort.Slice`
  + range over `[]result`

+ `main.go`
  + just: `res, err := parse()`