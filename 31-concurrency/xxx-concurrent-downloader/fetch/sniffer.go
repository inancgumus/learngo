package fetch

// sniffer converts a function to a sniffer that can sniff from the
// on-going io.Writer call such as request -> file.
//
// This is here just for simplifying the logic of HTTPTransfer.
type sniffer func(p []byte) bool

// Write satistifes io.Writer interface to sniff from it.
// It can be used through a io.MultiWriter.
func (f sniffer) Write(p []byte) (n int, err error) {
	n = len(p)

	// if the sniffer returns false, terminate with a non-nil error.
	// it used to abrupt the sniffing process, such as abrupting the
	// io.MultiWriter.
	if !f(p) {
		err = f
	}
	return
}

// Error satisfies the Error interface. So the returned error from the
// sniffer.Write func can return itself as an error. This is only used when
// the sniff.Write wants to terminate. So that we can distinguish it from a
// real error.
func (f sniffer) Error() string { return "" }
