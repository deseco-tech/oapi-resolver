// resolver manage the process
package resolver

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

var (
	cwd = ""
)

func New(inputFile, outputFile string) (*Manager, error) {
	extractCwdFromInputFile(inputFile)
	in, err := openInputFile(inputFile)
	if err != nil {
		return nil, err
	}
	out, err := os.Create(outputFile)
	if err != nil {
		return nil, err
	}
	return &Manager{
		in:   in,
		out:  out,
		done: false,
	}, nil
}

type Manager struct {
	in   io.Reader
	out  *os.File
	done bool
}

func (s *Manager) Combaine() error {
	buf := bufio.NewReader(s.in)
	for {
		bytes, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		line := NewLine(bytes)
		if err := line.Parse(); err != nil {
			return err
		}
		s.out.WriteString(line.String())
	}
	s.done = true
	return nil
}

func extractCwdFromInputFile(inputFile string) {
	if inputFile != "" {
		pathSlice := strings.Split(inputFile, "/")
		cwd = strings.Join(pathSlice[:len(pathSlice)-1], "/")
	}
}

func openInputFile(inputFile string) (io.Reader, error) {
	if inputFile == "" {
		return nil, errors.New("missing input file")
	}
	return os.Open(inputFile)
}
