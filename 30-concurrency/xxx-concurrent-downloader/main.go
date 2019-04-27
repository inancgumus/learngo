package main

import (
	"time"

	dl "github.com/inancgumus/course-experimentations/xxx-concurrent-downloader/0419-experimental/fetch"
)

//
//       +--- transfer #1 (on done/error signals and quits)
//       |
// UI <--+--- transfer #2
//       |
//       +--- transfer #N
//               ^
//               |
// SESSION ------+ launches goroutines
//

//
//       +--- transfer #1
//       |
// UI <--+--- transfer #2
//       |
//       +--- transfer #N
//               ^
//               |
// SESSION ------+ launches goroutines
//

//
//       +--- transfer #1
//       |
// UI <--+--- transfer #2
// ^     |
// |     +--- transfer #N
// |             ^
// |             |
// +--------- SESSION
//
// Session can close the transfers with session.Shutdown()
// (through session.done chan)
//
// Session closes the watcher chan when the transfers end
// Possible problem: If any of the transfers never quit, session will hang.
// I think, I'm going to manage that using context.Context in the transfers.
//

func main() {
	// ========================================================================

	// to := time.Second * 5
	// res, err := dl.HTTPGet("https://jsonplaceholder.typicode.com/todos/1", to)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer res.Body.Close()

	// io.Copy(os.Stdout, res.Body)

	// return

	// ========================================================================

	sess := dl.NewSession()

	// simulate a manual shutdown
	// time.AfterFunc(time.Second, func() {
	// 	sess.Shutdown()
	// })

	to := time.Second * 5
	transfers := []dl.Transfer{
		dl.NewHTTPTransfer("https://inancgumus.github.io/samples/jpg-beach-1020x670-160kb.jpg", to),
		dl.NewHTTPTransfer("https://inancgumus.github.io/samples/jpg-beach-1020x670-610kb.jpg", to),
		dl.NewHTTPTransfer("https://inancgumus.github.io/samples/jpg-beach-1020x670-80kb.jpg", to),
		dl.NewHTTPTransfer("https://inancgumus.github.io/samples/jpg-beach-1900x1250-1700kb.jpg", to),
	}

	// transfers := []dl.Transfer{
	// 	dl.NewHTTPTransfer(ctx, "http://inanc.io/1"),
	// 	dl.NewHTTPTransfer(ctx, "http://inanc.io/2"),
	// 	dl.NewHTTPTransfer(ctx, "http://inanc.io/3"),
	// }

	dl.UI(sess.Start(transfers...))

	// results := dl.Drain(sess.Start(urls))
	// for _, r := range results {
	// 	fmt.Printf("%s [err: %v — %d/%d]\n",
	// 		r.URL, r.Error, r.Downloaded, r.Total)
	// }

	// how to handle ctrl+c signals?
	//   register a signal
	//   let the downloader now that the operation halts
	//   with Downloader.onCancel or "context.Context"

	// run with: GOTRACEBACK=all go run -race main.go
	// time.Sleep(time.Second * 2)
	// panic("give me the stack trace")
}
