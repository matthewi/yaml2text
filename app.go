package yaml2text

import (
	"io"
	"io/ioutil"
	"os"
        "github.com/Masterminds/sprig/v3"
	"text/template"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// App is yaml converter with go template
type App struct {
	t *template.Template
}

// New is App constractor, maybe return error
func New(pattern io.Reader) (*App, error) {
	bs, err := ioutil.ReadAll(pattern)
	if err != nil {
		return nil, errors.Wrap(err, "template read")
	}
	funcMap := sprig.TxtFuncMap() // template.FuncMap{
//		"add": func(a, b int) int { return a + b },
//		"sub": func(a, b int) int { return a - b },
//		"mul": func(a, b int) int { return a * b },
//		"div": func(a, b int) int { return a / b },
//	}
	t, err := template.New("yaml2text").Funcs(funcMap).Parse(string(bs))
	if err != nil {
		return nil, errors.Wrap(err, "template parse")
	}

	app := &App{
		t: t,
	}
	return app, nil
}

// NewWithFile is App constractor, with filename
func NewWithFile(templateFile string) (*App, error) {
	file, err := os.Open(templateFile)
	if err != nil {
		return nil, errors.Wrap(err, "template file")
	}
	defer file.Close()
	return New(file)
}

// Execute do convert yaml data to any text
func (app *App) Execute(input io.Reader, output io.Writer) error {
	decoder := yaml.NewDecoder(input)
	var data map[string]interface{}
	if err := decoder.Decode(&data); err != nil {
		return errors.Wrap(err, "yaml decode")
	}
	return errors.Wrap(app.t.Execute(output, data), "transform execute")
}

// ExecuteWithFile do convert yaml file to any text
func (app *App) ExecuteWithFile(yamlFile string, output io.Writer) error {
	file, err := os.Open(yamlFile)
	if err != nil {
		return errors.Wrap(err, "yaml file")
	}
	defer file.Close()
	return app.Execute(file, output)
}
