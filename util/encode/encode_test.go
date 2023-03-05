package logic

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeWithSalt(t *testing.T) {
	t.Run("normal: salt = 0", func(t *testing.T) {
		cases := []struct {
			url  string
			salt int
		}{
			{"http://confluence.pri.ibanyu.com/display/server/tinyURL", 0},
			{"https://github.com/jinzhu/gorm", 0},
			{"https://gorm.io/docs/transactions.html", 0},
			{"https://medium.com/@hussachai/error-handling-in-go-a-quick-opinionated-guide-9199dd7c7f76", 0},
			{"https://en.wikipedia.org/wiki/Base64", 0},
		}

		set := map[string]struct{}{}
		for _, c := range cases {
			tiny := EncodeWithSalt(c.url, c.salt)
			assert.Equal(t, NumChars, len(tiny))
			set[tiny] = struct{}{}
		}
		assert.Equal(t, len(cases), len(set))
	})

	t.Run("exhausted: salt >= 0", func(t *testing.T) {
		url := "https://en.wikipedia.org/wiki/Base64"

		set := map[string]struct{}{}
		total := 100000
		for i := 0; i < total; i++ {
			tiny := EncodeWithSalt(url, i)
			assert.Equal(t, NumChars, len(tiny))
			set[tiny] = struct{}{}
		}
		log.Printf("%s total:%d distinct:%d", t.Name(), total, len(set))
	})
}
