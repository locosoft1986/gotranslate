package gotranslate

type localization struct {
	lang string
	langData map[string]interface{}
}

type localLangMap struct {
	langNames []string
	localMap map[string]localization
}

type transLoaderBase interface {
	LoadLanguage(langUrl string) map[string]interface{}
}


func UseLoader(loader transLoaderBase) {

}

func Translate(localeId string, params ...interface{}) string {
	var retValue string
	retValue = ""

	return retValue
}