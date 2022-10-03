## 哪一个是正确的类型转换表达式?

* convert(40)
* var("hi")
* int(4.) *正确*
* int[4]

## 这段代码打印什么?

```go
age := 6.5
fmt.Print(int(age))
```

* 6.5
* 65
* 6 *正确*
* .5

> 当你把一个 float 转换为 int 时
> 它失去了它的小数部分

## 这段代码打印什么?

```go
fmt.Print(int(6.5))
```

* 6.5
* 65
* 6
* Compile-Time Error *正确*

> Go可以在编译时检测转换错误

## 这段代码打印什么?

```go
area := 10.5
fmt.Print(area/2)
```

* 5.25 *正确*
* 5
* 0
* Error

## 这段代码打印什么?

```go
area := 10.5
div := 2
fmt.Print(area/div)
```

* 5.25
* 5
* ERROR *正确*

> 你不能除以不同类型的值
> 你需要转换:`area / float64(div)`

## 下面哪段代码能修复这段代码?

```go
area := 10.5
div := 2
fmt.Print(area/div)
```

* `fmt.Print(int(area)/div)`      // 5
* `fmt.Print(area/int(div))`      // 类型不匹配
* `fmt.Print(int(area)/int(div))` // 5
* `fmt.Print(area/float64(div))`  *正确*

