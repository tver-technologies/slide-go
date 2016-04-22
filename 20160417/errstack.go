import "github.com/go-errors/errors"

if err != nil {
	if e, ok := err.(*errors.Error); ok {
		fmt.Println(e.ErrorStack())
	}
}
