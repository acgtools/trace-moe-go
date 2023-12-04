package util

import "io"

func Close(w io.Closer, err *error) {
	// respect the existed err
	if e := w.Close(); *err == nil {
		*err = e
	}
}
