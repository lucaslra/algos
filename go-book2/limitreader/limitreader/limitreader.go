package limitreader

import (
	"io"
)

type limitReader struct {
	limit        int64
	currentIndex int
	data         io.Reader
}

func (r *limitReader) Read(p []byte) (int, error) {
	if r.limit <= int64(r.currentIndex) {
		return 0, io.EOF
	}

	arr := make([]byte, r.limit)
	_, err := r.data.Read(arr)
	n := copy(p, arr)

	r.currentIndex += n

	if err == io.EOF {
		return n, io.EOF
	}

	return n, nil
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{limit: n, data: r}
}
