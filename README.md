# sthingsBase
module providing basic building blocks

## IMPORT

<details><summary>go.mod</summary>
	
```
module github..
go 1.20

require (
  ..
  github.com/stuttgart-things/sthingsBase v0.1.3
  ..
)
```

</details>

<details><summary>code.go</summary>
	
```
package xy

import (
  ..
  sthingsBase "github.com/stuttgart-things/sthingsBase"
  ..
)	
```

</details>

## FUNCTIONS

### FILESYSTEM

<details><summary>CreateNestedDirectoryStructure</summary>
  tbd!

  EXAMPLE USAGE:
  ```
  sthingsBase.CreateNestedDirectoryStructure("/tmp/terraform", 0777)
  ```
</details>

### FILE

<details><summary>ReadFileToVariable</summary>
  tbd!
</details>

<details><summary>WriteDataToFile</summary>
  tbd!
</details>

### REGEX

<details><summary>GetRegexSubMatch</summary>
  tbd!
</details>

### TEMPLATING

<details><summary>RenderTemplateInline</summary>
 
  INPUT:
 
  | templateData | renderOption | delimStart | delimEnd | templateVariables         |
  |--------------|--------------|------------|----------|------------------------   |
  | string       | string       | string     | string   | map[string]interface{}   |
  | hello {{ .name }}   | "missingkey=zero"    | "{{"        |  "}}"     | vars:= map[string]interface{}{"name": "delicious",}                        
  
  OUTPUT:
 
  | yamlBytes | err   |
  |-----------|-------|
  | []byte    | error |
  |           |       |

  EXAMPLE USAGE:
  ```
  ...
  yamlBytes, err := sthingsBase.RenderTemplateInline(
	metaDataFile.template, 
	"missingkey=zero", 
	"{{", 
	"}}", 
	chartData)
  
  if err != nil {
    log.Fatal(err)
  }
  ```
  
</details>

TASKFILE
-------
* TAG

License
-------

BSD

Author Information
------------------

Patrick Hermann, stuttgart-things 04/2023
