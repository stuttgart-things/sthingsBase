/*
Copyright © 2023 Patrick Hermann patrick.hermann@sva.de
*/

package base

import (
	"fmt"
	"testing"
)

func TestGetVariablesAndDefaultsFromTemplate(t *testing.T) {

	curlyTemplate := `This is a {{ or .testKind "simple" }} template
	It can {{ "output" }} something.
	It also
	{{ if .IsAttachment }}
	-- attachement presentation code  --
	{{ end }}
	`

	var wantTemplateVariablesAndDefaults = make(map[string]string)
	var wantPriority = make(map[int]string)
	var wantAllConditions []TemplateCondtition
	var elementArray []string
	var wantAllNoneConditionalVariables []string
	wantAllNoneConditionalVariables = append(wantAllNoneConditionalVariables, "dfd")

	wantTemplateVariablesAndDefaults["output"] = ""
	wantTemplateVariablesAndDefaults["testKind"] = "simple"
	wantPriority[0] = "testKind"
	wantPriority[1] = "output"
	elementArray = append(elementArray, "dfd")
	wantAllConditions = append(wantAllConditions, TemplateCondtition{"hello", elementArray})

	allTemplateVariablesAndDefaults, priority, allConditions, allNoneConditionalVariables := GetVariablesAndDefaultsFromTemplate(curlyTemplate, "curly")
	fmt.Println(allTemplateVariablesAndDefaults, priority, allConditions, allNoneConditionalVariables)

}
