package gotranslate

type localization struct {
	lang string
	langData map[string]interface{}
}

type transLoaderBase interface {
	Config(options interface{})
	LoadLanguage(langUrl string) (map[string]interface{}, bool)
}

type localLangStore struct {
	defaultLang string
	currentLang string
	loader 		transLoaderBase
	langNames []string
	localMap map[string]localization
}

func (self *localLangStore) Add(langName string, locale localization) {

}

func (self *localLangStore) Get(key string, loader transLoaderBase)(string, bool) {

	return "", false;
}

func UseStaticFileLoader(option interface{}) {

}

func UseLoader(loader transLoaderBase, option interface{}) {

}

func TR(localeId string, params ...interface{}) string {
	var retValue string
	retValue = ""

	return retValue
}