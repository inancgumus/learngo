package main

type logger struct {
	print   func(...interface{}) (int, error)
	printf  func(fmt string, args ...interface{}) (int, error)
	println func(...interface{}) (int, error)
}

func (l logger) Print(args ...interface{}) {
	l.print(args...)
}

func (l logger) Printf(fmt string, args ...interface{}) {
	l.printf(fmt, args...)
}

func (l logger) Println(args ...interface{}) {
	l.println(args...)
}
