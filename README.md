# stuttgart-things/sthingsBase
module which provides basic golang functions & building blocks

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

### CONVERTING

ConvertIntegerToString

### LOGGING

<details><summary>StdOutFileLogger</summary>
  tbd!

  EXAMPLE USAGE:
  ```
  log := sthingsBase.StdOutFileLogger(logfilePath, "2006-01-02 15:04:05", 50, 3, 28)
  ..
  log.Info("gRPC server running on port " + serverPort)  
  ```
</details>

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

<details><summary>StoreVariableInFile</summary>
  tbd!
</details>

### REGEX

<details><summary>GetRegexSubMatch</summary>
  tbd!
</details>

<details><summary>GetAllRegexMatches</summary>
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

## LICENSE

<details><summary><b>APACHE 2.0</b></summary>

Copyright 2023 patrick hermann.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

</details>

Author Information
------------------
Patrick Hermann, stuttgart-things 04/2023

