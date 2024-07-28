package config

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestConfig(t *testing.T) {

	t.Setenv("ENV", "dev")
	t.Setenv("PORT", "80")
	t.Setenv("DATABASE_URL", "todo:password@tcp(dbprod01:3306)/todo")
	t.Setenv("PROJECTID", "xxx")
	t.Setenv("LINE_CHANNEL_SECRET", "0000000000")
	t.Setenv("LINE_ACCESS", "0000000000")
	t.Setenv("LINE_USER_ID", "0000000000")

	expected := &Config{
		Env:              "dev",
		Port:             "80",
		Database_url:     "todo:password@tcp(dbprod01:3306)/todo",
		ProjectID:        "xxx",
		LineChannelscret: "0000000000",
		LineAccesstoken:  "0000000000",
		LineUserId:       "0000000000",
	}
	got, err := New()
	assert.Equal(t, expected, got)
	assert.Equal(t, nil, err)
}
