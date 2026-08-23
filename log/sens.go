package log

import "log/slog"

var (
	sensGroupName    = "sens"
	sensGroupEnabled bool
)

// SetSensGroupName sets the group name for sensitive information in logs.
func SetSensGroupName(name string) {
	sensGroupName = name
}

// EnableSensGroup enables the grouping of sensitive information in logs under the sensGroupName key.
func EnableSensGroup() {
	sensGroupEnabled = true
}

// Sens returns a group attr with the given attrs under the sensGroupName key.
// This is useful for grouping sensitive information in logs.
func Sens(attrs ...slog.Attr) slog.Attr {
	if !sensGroupEnabled {
		return slog.Group(sensGroupName)
	}

	return slog.GroupAttrs(sensGroupName, attrs...)
}
