package gotranslate
import "testing"

func TestUseLoaderGivenStaticLoader(t *testing.T) {
	staticLoader := new(StaticFileLoader)

	UseLoader(staticLoader, StaticFileOption{"testprefix", "testsuffix"})

	resultLoader := (*translates.loader).(*StaticFileLoader)

	resultOpt := resultLoader.options.(StaticFileOption)

	if resultOpt.prefix != "testprefix" || resultOpt.suffix != "testsuffix" {
		t.Error("Expect the config option properties to be equal")
	}
}

func TestUseStaticFileLoaderGivenOption(t *testing.T) {
	UseStaticFileLoader(StaticFileOption{"testPrefix", "testSuffix"})

	resultLoader := (*translates.loader).(*StaticFileLoader)

	resultOpt := resultLoader.options.(StaticFileOption)

	if resultOpt.prefix != "testPrefix" || resultOpt.suffix != "testSuffix" {
		t.Error("Expect the config option properties to be equal")
	}
}

func TestDiveGivenKeysAndJsonMapData(t *testing.T) {
	jsonMapData := make(map[string]interface{})
	jsonMapData["test1"] = make(map[string]interface{})
	(jsonMapData["test1"].(map[string]interface{}))["insideTest1"] = "This string is deep inside test1"
	jsonMapData["test2"] = "String at test2"

	keyTest1 := []string{"test1", "insideTest1"}
	keyTest2 := []string{"test2"}
	keyNotExistInTest1 := []string{"test1", "notExistsKey"}

	resultTest1, okTest1 := translates.Dive(keyTest1, 0, jsonMapData)
	resultTest2, okTest2 := translates.Dive(keyTest2, 0, jsonMapData)
	resultTest3, okTest3 := translates.Dive(keyNotExistInTest1, 0, jsonMapData)

	if !okTest1 {
		t.Error("Expect okTest1 to be true")
	}

	if resultTest1 != "This string is deep inside test1" {
		t.Error("resultTest1 is equal to 'deep inside test1' string")
	}

	if !okTest2 {
		t.Error("Expect okTest1 to be true")
	}

	if resultTest2 != "String at test2" {
		t.Error("resultTest1 is equal to 'String at test2' string")
	}

	if okTest3 {
		t.Error("Expect okTest1 to be false")
	}

	if resultTest3 != "" {
		t.Error("resultTest1 is an empty string")
	}

}