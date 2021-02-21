package mock

import "fmt"

type Retriever struct {
	Contents string
}
// interface Stringer: fmt the value of interface
func (r *Retriever) String() string {
	return fmt.Sprintf( // Sprintf means format the string and return it, but not output it
		"Retriever: {Contents = %s}", r.Contents)
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}
