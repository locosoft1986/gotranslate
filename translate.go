package gotranslate
import (
	"fmt"
	"strings"
	"reflect"
)

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
	locales 			map[string]*localization
}

var translates = &translateStore{locales: make(map[string]*localization)}


func (self *translateStore) Load(langName string) {
	if self.loader == nil {
		fmt.Println("ERROR: Loader is not set")
		return
	}

	langData, ok := (*self.loader).LoadLanguage(langName)

	if ok {
		locale := new(localization)
		locale.lang = langName
		locale.langData = langData

		self.locales[langName] = locale
	}
}

func (self *translateStore) Dive(keyArray []string, index int, jsonData map[string]interface{})(string, bool) {
	var retValue string
	var success bool

	if index >= len(keyArray) {
		return "", false
	}

	tmpVar, ok := jsonData[keyArray[index]]

	if !ok {
		return "", false
	}

	switch tmpVar.(type) {
	case map[string]interface{}:
		retValue, success = self.Dive(keyArray, index + 1, tmpVar.(map[string]interface{}))

	case string:
		retValue = tmpVar.(string)
		success = true

	default:
		retValue = ""
		success = false
	}

	return retValue, success;
}

func (self * translateStore) Get(langName string, dottedKey string) (string, bool) {
	jsonKeys := strings.Split(dottedKey, ".")

	locale, langExists := self.locales[langName]

	if !langExists {
		return "", false
	}

	retLangString, ok := self.Dive(jsonKeys, 0, locale.langData)

	return retLangString, ok
}

func UseStaticFileLoader(option interface{}) {
	UseLoader(new(StaticFileLoader), option)
}

func UseLoader(loader transLoaderBase, option interface{}) {
	translates.loader = &loader
	(*translates.loader).Config(option)
}

func Use(langName string) {
	translates.currentLang = langName

	if locale, ok := translates.locales[langName]; !ok || locale == nil {
		translates.Load(langName)
	}
}

func PreferredLanguage(langName string) {
	translates.currentLang = langName
	translates.Load(langName)
}

func FallbackLanguage(langName string) {
	translates.fullbackLang = langName
	translates.Load(langName)
}

func Reload(langName string) {
	translates.Load(langName)
}

func TR(localeId string, args ...interface{}) string {
	var retString string
	var ok bool

	retString, ok = translates.Get(translates.currentLang, localeId)

	if !ok {
		retString, ok = translates.Get(translates.fullbackLang, localeId)
	}

	if !ok {
		return ""
	}

	if len(args) > 0 {
		params := make([]interface{}, 0, len(args))
		for _, arg := range args {
			if arg != nil {
				val := reflect.ValueOf(arg)
				if val.Kind() == reflect.Slice {
					for i := 0; i < val.Len(); i++ {
						params = append(params, val.Index(i).Interface())
					}
				} else {
					params = append(params, arg)
				}
			}
		}
		return fmt.Sprintf(retString, params...)
	}

	return retString
}