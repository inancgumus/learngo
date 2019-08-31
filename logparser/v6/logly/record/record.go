package record

const fieldsLength = 4

// Record stores fields of a log line.
type Record struct {
	Domain  string
	Page    string
	Visits  int
	Uniques int
}

// Sum the numeric fields with another record.
func (r *Record) Sum(other Record) {
	r.Visits += other.Visits
	r.Uniques += other.Uniques
}

// Reset all the fields of this record.
func (r *Record) Reset() {
	r.Domain = ""
	r.Page = ""
	r.Visits = 0
	r.Uniques = 0
}
