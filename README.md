# sthingsBase
module providing basic building blocks

## FUNCTIONS

### FILE

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
