package common

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"
)

// Output takes json output
type Output struct {}

// ParseTmpl parse the template to a file
// vars: the data structure of struct type
// tmplStr: the template const
// tmplName: the template name
// output: the location of output
// fimeMode: e.g. 0644
func ParseTmpl(vars interface{}, tmplStr string, tmplName string, output string, fileMode os.FileMode) {
	buf := new(bytes.Buffer) // the buf used to output the template
	tmpl, err := template.New(tmplName).Parse(tmplStr)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(buf, vars)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(output, []byte(buf.String()), fileMode)
}
