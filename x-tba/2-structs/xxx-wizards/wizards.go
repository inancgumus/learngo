package wizards

// Wizard is one of the greatest of people
type Wizard struct {
	// name won't be marshalled (should be exported)
	Name     string `json:name`
	Lastname string `json:"-"`
	Nick     string `json:"nick"`
}
