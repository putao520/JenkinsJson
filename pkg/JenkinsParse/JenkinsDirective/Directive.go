package JenkinsDirective

import (
	"github.com/pkg/errors"
	"strings"
)

type DirectiveProps struct {
	Name string
	Value string
}

func New(str string) (*DirectiveProps, error) {
	dProp := new(DirectiveProps)
	vSeg := strings.Split(str, ":")
	if len(vSeg) < 2 {
		return nil, errors.Errorf("directive struct error")
	}
	dProp.Name = vSeg[0]
	dProp.Value = vSeg[1]
	return dProp, nil
}