package fileproc

type Result struct {
	Path     string
	Lines    int
	Bytes    int
	Err      error // non-nil if the file failed
	Panicked bool  // true if a panic was recovered for this file
}
