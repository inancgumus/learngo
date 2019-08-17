### PROBLEM
+ `parser.go/parse()` also does updating. back to square one.
  + we need to extract the reusable behavior: scanning.

+ inflexible:
  + adding a filter is hard. needs to change the `scan()` code.
  + adding a grouper is also hard. domain grouping is hardcoded.

## SOLUTION
+ 

## IDEAS:

+ make domain filter accept variadic args