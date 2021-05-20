package tests

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"text/template"
)

type TestModel struct {
	ModelName  string
	FieldName  string
	FieldType  string
	ColumnName string
}

func TestParseTemplate(t *testing.T) {
	td := TestModel{"TestModel", "AlertID", "string", "alert_id"}

	content, err := ioutil.ReadFile("../templates/model.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)

	tmpl, err := template.New("test_model").Parse(text)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}
}
