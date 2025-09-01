package language_wizard

// // // // // // // // // // // //

func (obj *LanguageWizardObj) CurrentLanguage() string {
	obj.mx.RLock()
	defer obj.mx.RUnlock()
	return obj.currentLanguage
}

func (obj *LanguageWizardObj) Words() map[string]string {
	obj.mx.RLock()
	defer obj.mx.RUnlock()

	words := make(map[string]string, len(obj.words))
	for k, v := range obj.words {
		words[k] = v
	}

	return words
}

func (obj *LanguageWizardObj) Get(id, def string) string {
	if id == "" {
		return def
	}

	obj.mx.RLock()
	val, ok := obj.words[id]
	obj.mx.RUnlock()

	if ok {
		return val
	}

	obj.log("undef: " + id)
	return def
}
