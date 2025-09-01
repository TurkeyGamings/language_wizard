package language_wizard

// // // // // // // // // // // //

func (obj *LanguageWizardObj) SetLog(f func(string)) {
	if f == nil {
		return
	}

	obj.mx.Lock()
	defer obj.mx.Unlock()

	obj.log = f
}

func (obj *LanguageWizardObj) SetLanguage(isoLanguage string, words map[string]string) error {
	if isoLanguage == "" {
		return ErrNilIsoLang
	}
	if words == nil || len(words) == 0 {
		return ErrNilWords
	}

	obj.mx.Lock()
	defer obj.mx.Unlock()

	if obj.closed {
		return ErrClosed
	}

	if isoLanguage == obj.currentLanguage {
		return ErrLangAlreadySet
	}

	obj.currentLanguage = isoLanguage

	copyWords := make(map[string]string, len(words))
	for k, v := range words {
		copyWords[k] = v
	}
	obj.words = copyWords

	close(obj.changedCh)
	obj.changedCh = make(chan struct{})

	return nil
}
