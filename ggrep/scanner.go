package ggrep

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Scanner struct {
	fileName   string
	pattern    string
	ignoreCase bool
}

type Line struct {
	File   string
	Number int
	Line   string
	Error  error
}

func (l Line) WithFileAndNum() string {
	if l.Error != nil {
		return l.Error.Error()
	}
	return fmt.Sprintf("%s:%d:%s", l.File, l.Number, l.Line)
}

func New(f string, pattern string, ignoreCase bool) Scanner {
	return Scanner{
		fileName:   f,
		pattern:    pattern,
		ignoreCase: ignoreCase,
	}
}

func (s Scanner) Scan(c chan<- Line) {

	fd, err := os.Open(s.fileName)
	if err != nil {
		c <- Line{Error: err}
		return
	}
	defer fd.Close()

	if s.ignoreCase {
		err = s.matchIgnoreCase(fd, c)
	} else {
		err = s.match(fd, c)
	}

	if err != nil {
		c <- Line{Error: err}
	}
}

func (s Scanner) match(r io.Reader, c chan<- Line) error {
	max := 64 * 1024
	scanner := bufio.NewScanner(r)
	buf := make([]byte, max)
	scanner.Buffer(buf, max)
	lineNum := 1
	for scanner.Scan() {
		l := Line{
			File:   s.fileName,
			Number: lineNum,
			Line:   scanner.Text(),
		}
		if strings.Contains(l.Line, s.pattern) {
			// Matched
			c <- l
		}
		lineNum++
	}

	return scanner.Err()
}

func (s Scanner) matchIgnoreCase(r io.Reader, c chan<- Line) error {
	max := 64 * 1024
	scanner := bufio.NewScanner(r)
	buf := make([]byte, max)
	scanner.Buffer(buf, max)
	lineNum := 1
	p := strings.ToLower(s.pattern)
	for scanner.Scan() {
		l := Line{
			File:   s.fileName,
			Number: lineNum,
			Line:   scanner.Text(),
		}

		s := strings.ToLower(l.Line)
		if strings.Contains(s, p) {
			// Matched
			c <- l
		}
		lineNum++
	}

	return scanner.Err()
}
