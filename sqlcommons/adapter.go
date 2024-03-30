package sqlcommons

type SQLAdapter interface {
	Translate(query string) string
	ErrorHandler(err error) error
}
