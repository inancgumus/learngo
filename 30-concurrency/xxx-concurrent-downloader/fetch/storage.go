package fetch

import (
	"fmt"
	"io"
)

// TODO: let main func hand the "TransferFactory" to "Session"
// instead of "StorageFactory"
//
// Because: "Session" duplicates almost everything for "Transfer"

// StorageFactory func allows to switch storage implementations
// When called it may return a FileStore
// The transfer will call its Write method
//
// Why a func rather than an interface?
// Because: We don't have to store any state
// Storage will return its own io.Writer and the state will be in it
//   However, before invoking the storage, we don't need any state
//   We let us the storage manage its own state, we don't care about it
type StorageFactory func(url string) io.Writer

// FileStorage is the default storage mechanism for the downloader
// It writes to files
type FileStorage struct {
	url   string
	saved int
}

// FileStorageFactory creates and returns a new FileStorage
func FileStorageFactory(url string) io.Writer {
	return &FileStorage{url: url}
}

func (f *FileStorage) Write(p []byte) (int, error) {
	if debug {
		fmt.Println("[FILESTORAGE]", string(p), "for", f.url)
	}
	// TODO:
	// if not exists create it
	// if exists update it
	return 0, nil
}
