# Hints

+ `millions` slice's backing array uses 65 MB of memory.

+ `make` a new slice with 10 elements with a new backing array.

  + `copy` the last 10 elements of the `millions` slice to the new slice.

  + So you will have slice with a new backing array only with 10 elements.

  + Then overwrite the `millions` slice by simply `assigning` `last10` slice to it.

+ **Remember:** slice ~= pointer to a backing array.

  If you overwrite the slice, it will lose that
  pointer. So Go can collect the unused memory.