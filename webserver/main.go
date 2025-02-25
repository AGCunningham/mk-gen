package webserver

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func LoadRenderAndWrite(templateName, templatePath string, w http.ResponseWriter, data any) error {
	tpl, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to open template '%s': %v", templateName, err)
	}

	t, err := template.New(templateName).Parse(string(tpl))
	if err != nil {
		return fmt.Errorf("failed to parse template '%s': %v", templateName, err)
	}

	err = t.Execute(w, data)
	if err != nil {
		return fmt.Errorf("failed to write template '%s': %v", templateName, err)
	}

	return nil
}

func PrintAndReturnError(err error, w http.ResponseWriter) {
	// no benefit to catching an error that failed to be written
	_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
