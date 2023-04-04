/*
Copyright © 2023 Patrick Hermann patrick.hermann@sva.de
*/

package base

import (
	"bytes"
	"html/template"
)

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
