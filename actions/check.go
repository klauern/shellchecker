package actions

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

// ShellCheckErrors represents the output from running shellcheck and outputting
// the errors in JSON format.  This is tied to the 'shellcheck' command's output.
type ShellCheckErrors struct {
	File      string `json:"file"`
	Line      int    `json:"line"`
	EndLine   int    `json:"endLine"`
	Column    int    `json:"column"`
	EndColumn int    `json:"endColumn"`
	Level     string `json:"level"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

func writeTempFile(file []byte) (*os.File, error) {
	f, err := ioutil.TempFile("", "shellcheck")
	if err != nil {
		return nil, errors.WithMessage(err, "unable to create TempFile")
	}
	num, err := f.Write(file)
	if err != nil {
		return nil, errors.WithMessage(err, "unable to write contents")
	}
	if num == 0 || num < len(file) {
		return nil, errors.Errorf("bytes written: %v, bytes expected: %v", num, len(file))
	}

	return f, nil
}

func runShellCheck(input []byte) ([]ShellCheckErrors, error) {
	log := buffalo.NewLogger("DEBUG")
	path, err := exec.LookPath("shellcheck")
	if err != nil {
		return nil, errors.WithMessage(err, "cannot find shellcheck")
	}
	log.Debugf("Found shellcheck executable")

	file, err := writeTempFile(input)
	if err != nil {
		return nil, errors.WithMessage(err, "cannot write tempfile to run shellcheck")
	}
	log.Debugf("wrote to temp file %v", file.Name())
	defer removeThing(file.Name())

	scCmd := exec.Command(path, "-f", "json", file.Name())
	output, err := scCmd.Output()
	log.Debugf("output is %s", output)
	log.Debugf("err is %v", err)
	if err != nil {
		if !strings.HasPrefix(err.Error(), "exit status 1") {
			return nil, errors.WithMessage(err, "error executing shellcheck command")
		}

	}
	var errs []ShellCheckErrors
	if err = json.Unmarshal(output, &errs); err != nil {
		return nil, errors.WithMessage(err, "unable to unmarshal output for shellcheck")
	}
	return errs, nil
}

// CheckShellCodeHandler handles calls for /check paths.
func CheckShellCodeHandler(c buffalo.Context) error {
	script, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return errors.WithMessage(err, "not able to read POST body")
	}
	errs, err := runShellCheck(script)
	if err != nil {
		return err
	}
	for i, v := range errs {
		v.File = ""
		errs[i] = v
	}

	c.Render(200, r.JSON(errs))
	return nil
}

func removeThing(name string) {
	err := os.Remove(name)
	if err != nil {
		errLog := buffalo.NewLogger("ERROR")
		errLog.Errorf("error removing %v: %v", name, err)
	}
}
