// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

// compress in bash:
// gzip -k database.json

func main() {
	// removed error handling and x.Close() calls for brevity (normally you shouldn't do this).
	// normally: fr.Close(), gr.Close() should also be called to release the resources.

	// data, _ := ioutil.ReadFile("database.json")
	// l, _ := decode(data)

	// fmt.Print(string(data))
	// fmt.Print(l)

	// ---

	// fr, _ := os.Open("database.json")
	// tr := TeeReader(fr, os.Stdout)
	// l, _ := decodeReader(tr)
	// fmt.Print(l)

	// ---

	// fr, _ := os.Open("database.json.gz")
	// gr, _ := gzip.NewReader(fr)
	// tr := TeeReader(gr, os.Stdout)
	// l, _ := decodeReader(tr)
	// fmt.Print(l)

	// ---

	// TODO: FIX
	// l, err := decodeGzip("database.json.gz", os.Stdout)
	// l, err := decodeGzip("database.json.gz", ioutil.Discard)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Print("Products:\n", l)
}

// func decodeGzipFile(path string) (list, error) {
// 	fr, err := os.Open("database.json.gz")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer fr.Close()

// 	return decodeGzip(fr)
// }

// func decodeGzip(r io.Reader, w io.Writer) (list, error) {
// 	fr, err := os.Open("database.json.gz")
// 	if err != nil {
// 		return nil, err
// 	}

// 	gr, err := gzip.NewReader(fr)
// 	if err != nil {
// 		fr.Close()
// 		return nil, err
// 	}

// 	tr := TeeReader(gr, w)
// 	l, err := decodeReader(tr)
// 	if err != nil {
// 		fr.Close()
// 		gr.Close()
// 		return nil, err
// 	}

// 	fr.Close()
// 	gr.Close()
// 	return l, nil
// }
