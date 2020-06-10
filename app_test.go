package yaml2text_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/andreyvit/diff"
	"github.com/mashiike/yaml2text"
	"github.com/pkg/errors"
)

func TestExecute(t *testing.T) {

	cases := []struct {
		yamlFile     string
		templateFile string
		expectedFile string
		isSuccess    bool
	}{
		{"testdata/basic.yaml", "testdata/basic.tpl", "testdata/basic.sql", true},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			t.Logf("%s => %s ,(with %s)", c.yamlFile, c.expectedFile, c.templateFile)
			app, err := yaml2text.NewWithFile(c.templateFile)
			if check(err, c.isSuccess) != nil {
				t.Errorf("unexpected constract error: %s", err)
				return
			}
			var buf bytes.Buffer
			err = app.ExecuteWithFile(c.yamlFile, &buf)
			if check(err, c.isSuccess) != nil {
				t.Errorf("unexpected execute error: %s", err)
				return
			}

			if c.expectedFile == "" {
				return
			}

			expectedBs, err := ioutil.ReadFile(c.expectedFile)
			if err != nil {
				t.Fatalf("expected data read failed: %s", err)
			}

			if a, e := strings.TrimSpace(buf.String()), strings.TrimSpace(string(expectedBs)); a != e {
				t.Errorf("result not as expected:\n%v", diff.LineDiff(e, a))
			}
		})
	}

}

func check(err error, mustSuccess bool) error {
	if !mustSuccess {
		return err
	}
	if err != nil {
		return errors.New("must success")
	}
	return err
}
