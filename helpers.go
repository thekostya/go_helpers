package go_helpers

type Closer interface {
	Close() error
}

type Logger interface {
	Log(keyvals ...interface{}) error
}

func DefferedCloser(closer Closer, logger Logger, msg string) {
	err := closer.Close()
	if err != nil {
		_ = logger.Log(
			"err", err,
			"msg", msg,
		)
	}
}
