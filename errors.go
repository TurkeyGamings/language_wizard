package language_wizard

import "errors"

// // // // // // // // // // // //

var (
	ErrNilIsoLang = errors.New("iso language is required")
	ErrNilWords   = errors.New("words is required")

	ErrLangAlreadySet = errors.New("iso language is already set")
	ErrClosed         = errors.New("language-wizard is closed")
)
