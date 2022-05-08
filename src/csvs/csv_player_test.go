package csvs

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	for i, v := range ConfigPlayerLevelSlice {
		fmt.Println(i, v)
	}

}
