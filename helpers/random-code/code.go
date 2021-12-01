package randomcode

import (
	"github.com/mazen160/go-random"
)

func GenerateCode(banyak int) (string, error) {
	data, err := random.String(banyak)

	return data, err
}
