package gotranslate

import "testing"

func TestFindFilePathGivenNoOption(t *testing.T) {
	fileLoader := StaticFileLoader{}

	_, status := fileLoader.FindFilePath("en")

	if status {
		t.Error("Expect status to be false when there is no config option")
	}
}

func TestFindFilePathGivenSingleOption(t *testing.T) {

	singleOpt := StaticFileOption{ "testdata/testdir/locale_", ".json"}

	fileLoader := StaticFileLoader{}

	fileLoader.Config(singleOpt)
	retPath, _ := fileLoader.FindFilePath("en-US")

	EXPECT_PATH := "testdata/testdir/locale_en-US.json"
	if retPath != EXPECT_PATH {
		t.Errorf("The paths are not equal given a single option, \n Expect %s equal to %s", retPath, EXPECT_PATH)
	}
}

func TestFindFilePathGivenMultipleOptions(t *testing.T) {
	multiOpts := []StaticFileOption{{"testdata/testdir/locale_", ".json"}, {"testdata/otherdir/local-", ".test.json"}}

	fileLoader := StaticFileLoader{}

	fileLoader.Config(multiOpts)
	retPathUS, statusUS := fileLoader.FindFilePath("en-US")
	retPathUK, statusUK := fileLoader.FindFilePath("en-UK")
	_, statusOther := fileLoader.FindFilePath("zh")

	EXPECT_US_PATH := "testdata/testdir/locale_en-US.json"
	EXPECT_UK_PATH := "testdata/otherdir/local-en-UK.test.json"


	if retPathUS != EXPECT_US_PATH {
		t.Errorf("The paths are not equal given a single option, \n Expect %s equal to %s", retPathUS, EXPECT_US_PATH)
	}

	if !statusUS {
		t.Error("Expect loading en-US returning true")
	}

	if retPathUK != EXPECT_UK_PATH {
		t.Errorf("The paths are not equal given a single option, \n Expect %s equal to %s", retPathUK, EXPECT_UK_PATH)
	}

	if !statusUK {
		t.Error("Expect loading en-UK returning true")
	}

	if statusOther {
		t.Error("Expect the localization which does not exist to be not found")
	}
}