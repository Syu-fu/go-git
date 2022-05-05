package command

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
)

func getObjectId(store string) string {
	h := sha1.New()
	io.WriteString(h, store)
	return hex.EncodeToString(h.Sum(nil))
}
