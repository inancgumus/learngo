## When should you use a struct type?
1. For storing the same type of values
2. For adding an additional type of values in runtime
3. For combining different types in a single type to represent a concept *CORRECT*

> **1:** Arrays, slices, and maps are better candidates for that.
>
> **2:** Struct fields are fixed at compile-time, you cannot add additional fields in runtime, neither you can remove them.
>
> **3:** That's right. A struct type combines different types of fields in a single type. You can use a struct type to represent a concept.
>


## What are the properties of struct fields?
1. They all should be of the same type
2. Each one should have a name and possibly a different type *CORRECT*
3. You can add additional fields in runtime
4. You can remove the existing fields in runtime

> **2:** Yes, each field should have a unique name. Also, each field should have a type, but every field can have a different type.
>


## What is wrong with the following code?
```go
type weather struct {
    temperature, humidity float64
    windSpeed             float64
    temperature           float64
}
```
1. Nothing is wrong with it
2. `temperature, humidity float64` field is a syntax error
3. `temperature` field is not unique *CORRECT*

> **2:** That's a parallel definition. It defines two float64 fields: temperature and humidity. It is correct.
>
> **3:** Right! Struct field names should be unique.
>


## What is the zero-value of the following struct value?
```go
var movie struct {
    title, genre string
    rating       float64
    released     bool
}
```
1. `{}`
2. `{title: "", genre: "", rating: 0, released: false}` *CORRECT*
3. `{title: "", genre: "", rating: 0, released: true}`
4. `{"title, genre": "", rating: 0, released: false}`

> **1:** That's an empty struct value with no fields.
>
> **2:** Right! Go initializes a struct's fields to zero-values depending on their type.
>



## What is the type of the following struct?
```go
avengers := struct {
    title, genre string
    rating       float64
    released     bool
}{
    "avengers: end game", "sci-fi", 8.9, true,
}

fmt.Printf("%T\n", avengers)
```
1. `struct{}`
2. `struct{ string; string; float64; bool }`
3. `struct{ title string; genre string; rating float64; released bool }` *CORRECT*
4. `{title: "avengers: end game"; genre: "sci-fi"; rating: 8.9; released: true}`

> **1:** That's an empty struct type with no fields.
>
> **2:** Fields names is also a part of a struct's type.
>
> **3:** Right! Field names and types are part of a struct's type.
>
> **4:** Nope, that's a struct value.
>


## Are the following struct values equal?
```go
type movie struct {
    title, genre string
    rating       float64
    released     bool
}

avengers := movie{"avengers: end game", "sci-fi", 8.9, true}
clone    := movie{
                title: "avengers: end game", genre: "sci-fi",
                rating: 8.9, released: true,
            }
```
1. There is a syntax error
2. Yes *CORRECT*
3. No

> **2:** When creating a struct value, it doesn't matter whether you use the field names or not. So, they are equal.
>



## Are the following struct values equal? If not, why?
```go
type movie struct {
    title, genre string
    rating       float64
    released     bool
}

avengers := movie{
    title: "avengers: end game", genre: "sci-fi",
    rating: 8.9, released: true,
}

clone := movie{title: "avengers: end game", genre: "sci-fi"}

fmt.Println(avengers == clone)
```
1. Yes: They have the same set of fields
2. No : They are not comparable
3. No : Field values are different *CORRECT*

> **1:** That's right, this means they are comparable, but that's not enough.
>
> **2:** Yes, they are. They use the same struct type.
>
> **3:** Yes, when you omit some of the fields, Go assigns zero values to them. Here, "clone" struct value's "rating" and "released" fields are: 0, and false, respectively.
>


## Do the movie and performance struct types have the same types?
```go
type item        struct { title string }
type movie       struct { item }
type performance struct { item }
```
1. Yes: They have the same set of fields
2. No : They have different type names *CORRECT*
3. No : An embedded field cannot be compared

> **2:** Right! Types with different names cannot be compared. However, you can convert one of them to the other because they have the same set of fields. movie{} == movie(performance{}) is ok, or vice versa.
>


## What does the program print?
```go
type item struct{ title string }

type movie struct {
    item
    title string
}

m := movie{
    title: "avengers: end game",
    item:  item{"midnight in paris"},
}

fmt.Println(m.title, "&", m.item.title)
```
1. midnight in paris & midnight in paris
2. avengers: end game & avengers: end game
3. midnight in paris & avengers: end game
4. avengers: end game & midnight in paris *CORRECT*

> **4:** Right! `m.title` returns "avengers: end game" because the outer type always takes priority. However, `m.item.title` returns "midnight in paris" because you explicitly get it from the inner type: item.
>

