package yuque

import "errors"

var (
	ErrTooManyOptions = errors.New("options length more than one")
	ErrGetDocumentID  = errors.New("Error get documentID")
)
