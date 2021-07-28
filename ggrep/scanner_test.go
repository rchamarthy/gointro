package ggrep_test

import (
	"sync"
	"testing"

	"github.com/rchamarthy/gointro/ggrep"
	"github.com/stretchr/testify/assert"
)

func TestScanner(t *testing.T) {
	testScanner(t, "test.txt", "sometimes", false, 2, 1)
	testScanner(t, "test.txt", "Sometimes", false, 2, 1)
	testScanner(t, "test.txt", "verify", false, 4, 1)
	testScanner(t, "test.txt", "VeRify", true, 4, 1)
}

func testScanner(t *testing.T, f string, pattern string, ignoreCase bool, line int, numMatches int) {
	assert := assert.New(t)
	s := ggrep.New("test.txt", pattern, ignoreCase)
	c := make(chan ggrep.Line, 64)
	wgProducer := &sync.WaitGroup{}
	wgProducer.Add(1)

	go func() {
		defer wgProducer.Done()
		s.Scan(c)
	}()

	wgConsumer := &sync.WaitGroup{}
	wgConsumer.Add(1)
	go func() {
		defer wgConsumer.Done()
		n := 0
		for l := range c {
			assert.Equal(line, l.Number)
			assert.Nil(l.Error)
			n++
		}
		assert.Equal(numMatches, n)
	}()

	wgProducer.Wait()
	close(c) // All data produced

	wgConsumer.Wait()
}

func TestBadFile(t *testing.T) {
	assert := assert.New(t)
	s := ggrep.New("its not there", "", false)
	c := make(chan ggrep.Line, 64)
	s.Scan(c)
	l := <-c
	assert.NotNil(l.Error)
	assert.NotNil(l.WithFileAndNum())
}
