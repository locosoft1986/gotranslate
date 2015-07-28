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

When use the default static file loader, first call the UseStaticFileLoader() function and pass a StaticFileOption structure with prefix and suffix attributes and call Use() function to load the actual json file::
 
 ```go
 UseStaticFileLoader(StaticFileOption{"testdata/testdir/locale_", ".json"})
 Use("en-US")
 ```
 
The StaticFileOption structure contains a prefix attribute and a suffix attribute.

The StaticFileLoader will search the file using {{prefix}} + {{language name that was passed into Use() function}} + {{suffix}} pattern.
 
 



