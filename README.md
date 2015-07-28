# gotranslate

Package gotranslate is for application localization. 

## Introduction

This provides localization support which inspired by angular-translate which supports json static files and custom loader.
You can use following command to install the package:

    go get https://github.com/locosoft1986/gotranslate

## Usage

First, you have to import this package:

```go
import "github.com/locosoft1986/gotranslate"
```

The default format of localization files is JSON when using default static file loader.

### Using Static File Loader

When use the default static file loader, first call the UseStaticFileLoader() function and pass a StaticFileOption structure with prefix and suffix attributes and call Use() function to load the actual json file:
 
 ```go
 UseStaticFileLoader(StaticFileOption{"testdata/testdir/locale_", ".json"})
 Use("en-US")
 ```
 
The StaticFileOption structure contains a prefix attribute and a suffix attribute.

The StaticFileLoader will search the file using {{prefix}} + {{language name that was passed into Use() function}} + {{suffix}} pattern automatically.

The Use function will check the language was loaded and prevent it from loading again when using the same language name next time.

### Multiple Search Location

The static file loader accepts an array of StaticFileOption structures which contain multiple configurations of the prefix and suffix of the search pattern.

 ```go
 staticOptions := []StaticFileOption{{"testdata/testdir/locale_", ".json"}, {"testdata/other/locale-", ".test.json"}}
 UseStaticFileLoader(StaticFileOption{"testdata/testdir/locale_", ".json"})
 Use("en-UK")
 ```
The above code will load the local-en-UK.test.json file in the "testdata/other" directory.
 
### Translate
 
 After the json file is loaded, you can use function TR to translate the localization ID into actual language. The localization ID is a dotted JSON query string:
 
 ```go
 UseStaticFileLoader(StaticFileOption{"testdata/testdir/locale_", ".json"})
 Use("en-US")
 firstName := TR("username.FIRSTNAME") //Expect return value to be "First Name"
 lastName := TR("username.LASTNAME") //Expect return value to be "Last Name" 
 ```
 
 The value corresponding to the localization ID can have format strings. For example, if there is a key "ParamTest" in the "username" field in the JSON file "testdir/locale_en-US.json" from above example:
 
 ```go
 //.... following above example
 paramTest := TR("username.ParamTest", 20) //Expect return value to be "Param value is 20"
 ```

### Preferred Language
 You can use PreferredLanguage() function to load a language JSON file. It will overwrite the previous loaded language with the same name.
 
### Fallback Language
 You can use FallbackLanguage() function to set a backup language when given localization key is not found in the current using language set by Use() or PreferredLanguage() function:
 
### Reload
 You can use Reload() function to reload a given language name.
 
 
## Custom Loader
 There is a base loader interface: 
 ```go
 type transLoaderBase interface {
 	Config(options interface{})
 	LoadLanguage(langUrl string) (map[string]interface{}, bool)
 }
 ```
 
 The Config function accepts an option object which will be used in the loading phase of the language as a parameter.
 The LoadLanguage function will be called when user call Use(), PreferredLanguage() or Reload() function.


#TODO
 1.Default Url Loader: It can fetch the json file from an URL addresses
 
 
 2.Ini file Loader: It can load the ini file and support the ini section as a dotted string key. 
 
 For example, if user want to fetch the key "name" in the section "test", user can use "test.name" as a localization key.
 
 
 3.Partial Loader: If the localization file is too large, user can split it into different parts and load the different parts in different situations
