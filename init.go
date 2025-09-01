package language_wizard

import (
	"sync"
)

// // // // // // // // // // // //

type LanguageWizardObj struct {
	currentLanguage string
	words           map[string]string
	mx              sync.RWMutex

	changedCh chan struct{}
	closed    bool

	log func(string)
}

func New(isoLanguage string, words map[string]string) (*LanguageWizardObj, error) {
	if isoLanguage == "" {
		return nil, ErrNilIsoLang
	}
	if words == nil || len(words) == 0 {
		return nil, ErrNilWords
	}

	obj := new(LanguageWizardObj)
	obj.currentLanguage = isoLanguage

	copyWords := make(map[string]string, len(words))
	for k, v := range words {
		copyWords[k] = v
	}
	obj.words = copyWords

	obj.changedCh = make(chan struct{})
	obj.log = func(s string) {}

	return obj, nil
}

func (obj *LanguageWizardObj) Close() {
	obj.mx.Lock()
	defer obj.mx.Unlock()

	if obj.closed {
		return
	}

	obj.closed = true
	close(obj.changedCh)

	obj.words = make(map[string]string)
}
