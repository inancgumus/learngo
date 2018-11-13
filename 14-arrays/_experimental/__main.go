const (
	Mon Weekday = iota
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

var names = [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}


for d := Mon; d <= Sun; d++ {
	fmt.Println(names[d])
}