package util

import (
	"encoding/json"
	"github.com/shawnfeng/sutil/slog"
)

func JsonEscape(raw string) string {
	b, err := json.Marshal(raw)
	if err != nil {
		slog.Errorf("JsonEscape Err: %v", err)
		return ""
	}
	// Trim the beginning and trailing " character
	return string(b[1 : len(b)-1])

}
