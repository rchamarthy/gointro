package greeting_test

import (
	"testing"

	"github.com/rchamarthy/gointro/greeting"
	"github.com/stretchr/testify/assert"
)

func allGreetings() map[greeting.Greeter]string {
	return map[greeting.Greeter]string{
		greeting.New("english"):   "hello",
		greeting.New("chinese"):   "你好",
		greeting.New("telugu"):    "నమస్కరం",
		greeting.New("tamil"):     "வணக்கம்",
		greeting.New("sanskrit"):  "नमस्ते",
		greeting.New("hindi"):     "नमस्ते",
		greeting.New("kannada"):   "ನಮಸ್ಕಾರ",
		greeting.New("gibberish"): "hello",
	}
}

func TestGreeting(t *testing.T) { //nolint: paralleltest
	assert := assert.New(t)
	for greeter, greeting := range allGreetings() {
		assert.Equal(greeter.Greet(), greeting)
		assert.NotNil(greeter.Greet())
	}
}
