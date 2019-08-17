# Func Values

+ Please feel free to put any function that match this signature here and it will be called each time my function is called.

+ They are extremely useful when wanting to declare a block of code that you want to pass around.

+ They allow functions to be passed as parameters.

+ They can be chained together

# Interfaces vs Func Values

+ Use interfaces when a method needs to belong to an object (that needs to operate on the value of a type: io.Reader, sort.Interface, etc).

+ Use func values when a function doesn't need to belong to an object. For simple, not-data oriented jobs: Filtering, splitting, mapping etc.

+ An implementation of the interface means that the behaviour is intrinsic to the implementing object. The behaviour does not change based on the caller or the circumstances of the call.

  + The func value is saying that the operation is not intrinsic to the object but based on context or the caller to define.

+ An interface is a language feature representing some contract: "I hereby guarantee I make the following methods available".

+ The benefit of using interfaces is that you give the type a name (you would get that with records too) and you are more clearly expressing your intention (other components can implement the interface).

+ Use interface values when:
  + Well known interfaces already exist, e.g. io.Reader
  + More than one behavior required
  + Typically stateful
  + Implementations non-trivial

  + Use function values when:
    + Only one behavior
    + Typically stateless
    + In-line implementations typical

  + http://go-talks.appspot.com/github.com/ChrisHines/talks/non-orthogonal-choices-in-go/non-orthogonal-choices-in-go.slide#1