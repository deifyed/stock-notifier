package config

import (
	"net/url"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
)

func generateValidEnvironment() map[string]string {
	return map[string]string{
		"PUSH_SERVER_URL": "https://push.domain.io/message?token=abcdefg",
		"SYMBOLS":         "ABC,DEF",
		"ABC_TARGETS":     "10-,20+",
		"DEF_TARGETS":     "1-,2+",
	}
}

func TestLoadConfig(t *testing.T) {
	testCases := []struct {
		name string

		withEnvironment map[string]string

		expectConfig Config
		expectErr    string
	}{
		{
			name: "sanity check",

			withEnvironment: generateValidEnvironment(),

			expectConfig: Config{
				Symbols: []string{"ABC", "DEF"},
				PushServerURL: url.URL{
					Scheme:   "https",
					Host:     "push.domain.io",
					Path:     "/message",
					RawQuery: "token=abcdefg",
				},
				PriceTargets: map[string][]string{
					"ABC": {"10-", "20+"},
					"DEF": {"1-", "2+"},
				},
			},
		},
		{
			name: "should err on missing url",

			withEnvironment: func() map[string]string {
				env := generateValidEnvironment()

				env["PUSH_SERVER_URL"] = ""

				return env
			}(),

			expectErr: "cannot be blank",
		},
		{
			name: "should err on missing symbols",

			withEnvironment: func() map[string]string {
				env := generateValidEnvironment()

				env["SYMBOLS"] = ""

				return env
			}(),

			expectErr: "cannot be blank",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			getter := func(key string) string {
				value, ok := tc.withEnvironment[key]
				if !ok {
					return ""
				}

				return value
			}

			cfg := LoadConfig(getter)
			err := cfg.Validate()

			if tc.expectErr == "" {
				assert.NoError(t, err)

				assert.Equal(t, tc.expectConfig, cfg)
			} else {
				assert.Error(t, err)

				assert.Equal(t, tc.expectErr, err.Error())
			}
		})
	}
}

func TestInRule(t *testing.T) {
	a := []string{"a", "b"}
	b := []string{"a", "b"}

	err := validation.Validate(a, validation.In(b))
	assert.Nil(t, err)
}
