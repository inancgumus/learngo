# Hints

You can use the following slice operations to solve the exercise:

+ Prepends "value" to the slice:
  ```go
  slice = append([]string{"value"}, slice...)
  ```

+ Appends some part (N to M) of the same slice to itself:
  ```go
  slice = append(slice, slice[N:M]...)
  ```
+ Copies the last part of the slice starting from M to the first part of the slice until N:
  ```go
  slice = append(slice[:N], slice[M:]...)
  ```
