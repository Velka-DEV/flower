package engine

import (
	"bytes"
	"text/template"
)

func renderTemplate(tmpl string, data interface{}) (string, error) {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func resolveInputs(ctx *Context, inputs map[string]interface{}) (map[string]interface{}, error) {
	resolvedInputs := make(map[string]interface{})

	for key, value := range inputs {
		switch v := value.(type) {
		case string:
			rendered, err := renderTemplate(v, ctx.stepOutputs)
			if err != nil {
				return nil, err
			}
			resolvedInputs[key] = rendered
		default:
			resolvedInputs[key] = v
		}
	}

	return resolvedInputs, nil
}
