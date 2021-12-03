package resolver

import (
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/ghodss/yaml"
)

var (
	refRegex = regexp.MustCompile(`\$ref.*\.yaml.*`)
)

func NewLine(b []byte) *Line {
	l := &Line{
		bytes: b,
		str:   string(b),
	}

	l.setIndentation()

	return l
}

type Ref struct {
	key   string
	fname string
}

type Line struct {
	bytes  []byte
	str    string
	indent string
	ref    *Ref
}

func (s *Line) Parse() error {
	if s.isRef() {
		err := s.processRef()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Line) String() string {
	return s.str + "\n"
}

func (s *Line) extractRef() {
	parts := strings.Split(s.str, "#/")
	if cwd != "" {
		cwd = cwd + "/"
	}
	s.ref = &Ref{
		key:   parts[1],
		fname: cwd + strings.TrimLeft(parts[0], "$ref: "),
	}
}

func (s *Line) isRef() bool {
	return len(refRegex.FindStringSubmatch(s.str)) > 0
}

func (s *Line) processRef() error {
	s.extractRef()
	bytes, err := os.ReadFile(s.ref.fname)
	if err != nil {
		return err
	}
	var refAsMap map[string]interface{}
	yaml.Unmarshal(bytes, &refAsMap)
	bytes, err = yaml.Marshal(refAsMap[s.ref.key])
	if err != nil {
		return err
	}
	s.setString(strings.TrimRight(string(bytes), "\n"))
	return nil
}

func (s *Line) setString(str string) {
	if s.indent == "" || str == "" {
		return
	}
	lines := strings.SplitAfter(str, "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	s.str = strings.Join(append([]string{""}, lines...), s.indent)
}

func (s *Line) setIndentation() {
	for _, c := range s.str {
		if unicode.IsSpace(c) {
			s.indent += " "
		} else {
			break
		}
	}
}
