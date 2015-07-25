package gotranslate
import "testing"

func TestUseLoaderGivenStaticLoader(t *testing.T) {
	staticLoader := new(StaticFileLoader)

	UseLoader(staticLoader, StaticFileOption{"testprefix", "testsuffix"})

	resultLoader := (*tranlates.loader).(*StaticFileLoader)

	resultOpt := resultLoader.options.(StaticFileOption)

	if resultOpt.prefix != "testprefix" || resultOpt.suffix != "testsuffix" {
		t.Error("Expect the config option properties to be equal")
	}
}

func TestUseStaticFileLoaderGivenOption(t *testing.T) {
	UseStaticFileLoader(StaticFileOption{"testPrefix", "testSuffix"})

	resultLoader := (*tranlates.loader).(*StaticFileLoader)

	resultOpt := resultLoader.options.(StaticFileOption)

	if resultOpt.prefix != "testPrefix" || resultOpt.suffix != "testSuffix" {
		t.Error("Expect the config option properties to be equal")
	}
}