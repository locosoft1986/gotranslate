package gotranslate

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
)

type StaticFileOption struct {
	prefix string
	suffix string
}

type StaticFileLoader struct {
	options interface{}
}

func (self *StaticFileLoader) Config(options interface{}) {
	self.options = options;
}

func (self *StaticFileLoader) FindFilePath(langName string) string {

	retFileName := langName

	switch self.options.(type) {
	case StaticFileOption://single option
		option := self.options.(StaticFileOption)
		retFileName = option.prefix + langName + option.suffix


	case []StaticFileOption://multiple options
		for _, option := range self.options.([]StaticFileOption)  {
			tmpFullPath := option.prefix + langName + option.suffix

			if _,err := os.Stat(tmpFullPath); err == nil {
				retFileName = tmpFullPath
				break
			}
		}

	}



	return retFileName
}

func (self *StaticFileLoader) LoadLanguage(langName string) (map[string]interface{}, bool) {
	filename := self.FindFilePath(langName)

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("ERROR: Load json file failed, err = ", err)
		return nil, false
	}

	var tmpResult interface{}

	err = json.Unmarshal(content, &tmpResult)

	if err != nil {
		fmt.Println("ERROR: Parsing json failed, err = ", err)
		return nil, false
	}

	retMap := tmpResult.(map[string]interface{})

	return retMap, true
}
