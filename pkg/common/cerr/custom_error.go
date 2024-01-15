package cerr

// Variables to hold error codes
var (
	InternalServerErrorCode = 500
	InvalidRequestErrorCode = 400
	NotFoundErrorCode       = 404
	DuplicateEntryErrorCode = 409
)

type CustomError struct {
	Origin    error
	ErrorCode int
	Message   string
}

func (ce *CustomError) Error() string {
	return ce.Message
}

func (r *CustomError) SetErrorCode(code int) error {
	r.ErrorCode = code
	return r
}

func (r *CustomError) SetMessage(message string) error {
	r.Message = message
	return r
}

func (r *CustomError) SetOrigin(err error) error {
	r.Origin = err
	return r
}

func (r *CustomError) GetErrorCode() int {
	return r.ErrorCode
}

func (r *CustomError) GetMessage() string {
	return r.Message
}

func (r *CustomError) GetOrigin() error {
	return r.Origin
}

func NewCustomError(s string) *CustomError {
	return &CustomError{
		ErrorCode: InternalServerErrorCode,
		Message:   s,
	}
}

func NewCustomErrorWithOrigin(s string, err error) *CustomError {
	return &CustomError{
		Origin:    err,
		ErrorCode: InternalServerErrorCode,
		Message:   s,
	}
}

func NewCustomErrorWithCodeAndOrigin(s string, code int, err error) *CustomError {
	return &CustomError{
		Origin:    err,
		ErrorCode: code,
		Message:   s,
	}
}

func Wrap(s string, e error) *CustomError {
	return &CustomError{
		Origin:    e,
		ErrorCode: InternalServerErrorCode,
		Message:   s,
	}
}

func SetErrorCode(code int, e error) error {
	if err, ok := e.(*CustomError); ok {
		err.ErrorCode = code
		return err
	}
	return e
}

func GetErrorCode(e error) int {
	if err, ok := e.(*CustomError); ok {
		return err.ErrorCode
	}
	return InternalServerErrorCode
}

func SetErrorMessage(message string, e error) error {
	if err, ok := e.(*CustomError); ok {
		err.Message = message
		return err
	}
	return e
}

func GetErrorMessage(e error) string {
	if err, ok := e.(*CustomError); ok {
		return err.Message
	}
	return e.Error()
}
