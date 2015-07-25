package gotranslate

type localization struct {
	lang string
	langData map[string]interface{}
}

type transLoaderBase interface {
	Config(options interface{})
	LoadLanguage(langUrl string) (map[string]interface{}, bool)
}

type translateStore struct {
	currentLang 		string
	fullbackLang 		string
	loader 				*transLoaderBase
	langNames 			[]string
	locales 			map[string]*localization
}

var tranlates = &translateStore{locales: make(map[string]*localization))}

func (self *translateStore) Add(langName string, locale localization) {

}

func (self *translateStore) Get(key string, loader transLoaderBase)(string, bool) {

	return "", false;
}

func UseStaticFileLoader(option interface{}) {
	UseLoader(new(StaticFileLoader), option)
}

func UseLoader(loader transLoaderBase, option interface{}) {
	tranlates.loader = &loader
	(*tranlates.loader).Config(option)
}

func Use(langName string) {

}

func PreferredLanguage(langName string) {

}

func FallbackLanguage(langName string) {

}

func Reload(langName string) {

}

func TR(localeId string, params ...interface{}) string {
	var retValue string
	retValue = ""

	return retValue
}