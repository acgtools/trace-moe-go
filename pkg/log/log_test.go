package log_test

import (
	"log/slog"
	"testing"

	"github.com/acgtools/trace-moe-go/pkg/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseLevel(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name        string
		lvl         string
		parsedLevel slog.Level
		wantErr     bool
	}{
		{
			name:        "Empty string",
			lvl:         "",
			parsedLevel: 0,
			wantErr:     true,
		},
		{
			name:        "Uppercase level",
			lvl:         "DEBUG",
			parsedLevel: log.DebugLevel,
			wantErr:     false,
		},
		{
			name:        "Debug",
			lvl:         "debug",
			parsedLevel: log.DebugLevel,
			wantErr:     false,
		},
		{
			name:        "Info",
			lvl:         "info",
			parsedLevel: log.InfoLevel,
			wantErr:     false,
		},
		{
			name:        "Warn",
			lvl:         "warn",
			parsedLevel: log.WarnLevel,
			wantErr:     false,
		},
		{
			name:        "Error",
			lvl:         "error",
			parsedLevel: log.ErrorLevel,
			wantErr:     false,
		},
		{
			name:        "Unsupported level",
			lvl:         "XXX",
			parsedLevel: 0,
			wantErr:     true,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l, err := log.ParseLevel(tc.lvl)

			assert.Equal(t, tc.parsedLevel, l)

			if tc.wantErr {
				require.Error(t, err)
				require.ErrorContains(t, err, "unrecognized level: ")
			} else {
				require.NoError(t, err)
			}
		})
	}
}
