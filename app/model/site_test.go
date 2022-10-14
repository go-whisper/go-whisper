package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSite(t *testing.T) {
	site := GetSiteParameter()
	t.Log("site name:", site.Name)
	assert.Equal(t, "<!–more–>", site.Separator)
}
