# sthingsBase
module providing basic building blocks

## Functions

### Regex

<details><summary>GetRegexSubMatch</summary>
  tbd!
</details>

### Templating

<details><summary>RenderTemplateInline</summary>
  
  input:
  | templateData | renderOption | delimStart | delimEnd | templateVariables         |
  |--------------|--------------|------------|----------|------------------------   |
  | string       | string       | string     | string   | map[string]interface{}   |
  |              |              |            |          |                           |
  
  output:
 
	
  ```
  func RenderTemplateInline(
	templateData,
	renderOption, 
	delimStart, 
	delimEnd string, 
	templateVariables map[string]interface{}) 
  ([]byte, error) {
  ```
 
  example usage:
  ```
  Patterns := map[string]VariableDelimiter {
    "curly":  VariableDelimiter{"{{", "}}", `\{\{(.*?)\}\}`},
    "square": VariableDelimiter{"[[", "]]", `\[\[(.*?)\]\]`},
  }
  ...
  yamlBytes, err := sthingsBase.RenderTemplateInline(
	metaDataFile.template, 
	"missingkey=zero", 
	Patterns["curly"].begin, 
	Patterns["curly"].end, 
	chartData)
  
  if err != nil {
    log.Fatal(err)
  }
  ```
  
</details>

License
-------

BSD

Author Information
------------------

Patrick Hermann, stuttgart-things 04/2023
