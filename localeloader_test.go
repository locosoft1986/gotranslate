package gotranslate

import "testing"

func TestFindFilePathGivenSingleOption(t *testing.T) {

	singleOpt := StaticFileOption{ "testdata/testdir/locale_", ".json"}

	fileLoader := StaticFileLoader{}

	fileLoader.Config(singleOpt)
	retPath := fileLoader.FindFilePath("en-US")

	EXPECT_PATH := "testdata/testdir/locale_en-US.json"
	if retPath != EXPECT_PATH {
		t.Errorf("The paths are not equal given a single option, \n Expect %s tobe %s", retPath, EXPECT_PATH)
	}
}
