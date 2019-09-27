package asession

import (
	"encoding/gob"
	"fmt"

	"github.com/gorilla/sessions"
)

//Configure session storage to use FilesystemStore.

var (
	Store *sessions.FilesystemStore
)

func init() {
	fmt.Println("session store initializing......")
	Store = sessions.NewFilesystemStore("", []byte("something-very-secret"))
	gob.Register(map[string]interface{}{})
}
