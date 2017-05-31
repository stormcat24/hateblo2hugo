package hugo

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/require"
	"time"
)

func TestHugoPage_Render(t *testing.T) {

	p := HugoPage{
		Date: time.Now().String(),
		Draft: false,
		Title: "title",
		Tags: []string{"tag1", "tag2"},
	}
	actual, err := p.Render()
	require.NoError(t, err)

	fmt.Println("---")
	fmt.Println(string(actual))
}