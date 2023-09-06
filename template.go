/*
Copyright © 2023 Patrick Hermann patrick.hermann@sva.de
*/

package base

import (
	"bytes"
	"strings"
	"text/template"
	"unicode/utf8"
)

type TemplateCondtition struct {
	VariableName         string   `mapstructure:"VariableName"`
	ConditionalVariables []string `mapstructure:"ConditionalVariables"`
}

type VariableDelimiter struct {
	begin        string `mapstructure:"begin"`
	end          string `mapstructure:"end"`
	regexPattern string `mapstructure:"regex-pattern"`
}

var Patterns = map[string]VariableDelimiter{
	"curly":  VariableDelimiter{"{{", "}}", `\{\{(.*?)\}\}`},
	"square": VariableDelimiter{"[[", "]]", `\[\[(.*?)\]\]`},
}

func RenderTemplateInline(templateData, renderOption, delimStart, delimEnd string, templateVariables map[string]interface{}) ([]byte, error) {

	tmpl, err := template.New("tpl").Delims(delimStart, delimEnd).Option(renderOption).Parse(templateData)

	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, templateVariables); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func GetVariablesAndDefaultsFromTemplate(templateFileContent string, variableDelimiter string) (map[string]interface{}, map[int]string, []TemplateCondtition, []string) {

	var priority = make(map[int]string)
	var allVariablesAndDefaults = make(map[string]interface{})
	var alllConditionals []TemplateCondtition
	var conditionalVariables []string
	var allNoneConditionalVariables []string

	allMatches := GetAllRegexMatches(templateFileContent, Patterns[variableDelimiter].regexPattern)

	conditionName := ""
	priorityIndex := 0
	openIfCondition := false

	for _, element := range allMatches {

		if strings.Contains(element, "if .") {
			element = strings.ReplaceAll(element, "if .", "")
			element, _ := GetRegexSubMatch(element, Patterns[variableDelimiter].regexPattern)
			element = strings.ReplaceAll(element, ".", "")
			element = strings.TrimSpace(element)

			conditionName = element
			openIfCondition = true

		} else if strings.Contains(element, " end ") || strings.Contains(element, " else ") {

			alllConditionals = append(alllConditionals, TemplateCondtition{conditionName, conditionalVariables})
			conditionalVariables = nil
			openIfCondition = false

		} else {

			element, _ := GetRegexSubMatch(element, Patterns[variableDelimiter].regexPattern)
			element = strings.TrimSpace(element)
			element = trimFirstRune(element)

			elementArray := strings.Fields(element)

			// IF NO DEFAULT IS FOUND
			if len(elementArray) == 1 {
				allVariablesAndDefaults[elementArray[0]] = ""
				priority[priorityIndex] = elementArray[0]
				priorityIndex = priorityIndex + 1

				if openIfCondition {
					conditionalVariables = append(conditionalVariables, elementArray[0])
				} else {
					allNoneConditionalVariables = append(allNoneConditionalVariables, elementArray[0])
				}

				// IF DEFAULT IS FOUND
			} else if len(elementArray) == 3 {

				element = strings.ReplaceAll(elementArray[1], ".", "")
				allVariablesAndDefaults[element] = strings.ReplaceAll(elementArray[2], "\"", "")
				priority[priorityIndex] = element
				priorityIndex = priorityIndex + 1

				if openIfCondition {
					conditionalVariables = append(conditionalVariables, element)
				} else {
					allNoneConditionalVariables = append(allNoneConditionalVariables, element)
				}

			}

		}

	}

	return allVariablesAndDefaults, priority, alllConditionals, allNoneConditionalVariables
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
