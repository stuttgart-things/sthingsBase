# sthingsBase
module providing basic building blocks

## Functions

### Regex

<details><summary>GetRegexSubMatch</summary>
  tbd!
</details>

### Templating

<details><summary>RenderTemplateInline</summary>
  
  INPUT:
  | templateData | renderOption | delimStart | delimEnd | templateVariables         |
  |--------------|--------------|------------|----------|------------------------   |
  | string       | string       | string     | string   | map[string]interface{}   |
  |              |              |            |          |                           |
  
  OUTPUT:
 
  | yamlBytes | err   |
  |-----------|-------|
  | []byte    | error |
  |           |       |

  EXAMPLE USAGE:
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
