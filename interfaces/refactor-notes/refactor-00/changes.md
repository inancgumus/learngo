## CHANGES

### PROBLEM
+ adding new fields makes the code complex
+ needs to update: `result`, `parser`, `summarizer`
+ needs to add new fields to `parser`: `totalVisits` + `totalUniques`
+ in `parse()`: repeating line errors
  + if we parsing out of it we'd need to have *parser — superfluous

### SOLUTION
+ move all the result related logic to result.go

+ move `parser.go/result` -> `result.go`
  + move `parser.go/parsing` logic -> `result.go`

+ add `addResult` -> `result.go`
  + remove `parser struct`'s: `totalVisits`, `totalUniques`
  + change `update()`'s last line: `p.sum[r.domain] = addResult`

+ remove `(line #d)` errors from `result.go`
  + add: `return r, err` — named params are error prone
  + always check for the error first
    + `if r.visits < 0 || err != nil` -> `if err != nil || r.visits < 0`

+ `parser.go`: check the `parseFields()`:
    ```golang
    r, err := parseFields(line)
    if err != nil {
        p.lerr = fmt.Errorf("line %d: %v", p.lines, err)
    }```

+ - `parser.go` and `summarize.go`
  - remove `total int`
  - let `summarize()` calculate the totals