package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseAttrModifier(t *testing.T) {
	for _, tc := range []struct {
		s           string
		expected    *modifier
		expectedErr bool
	}{
		{
			s: "empty",
			expected: &modifier{
				empty: boolModifierSet,
			},
		},
		{
			s: "+empty",
			expected: &modifier{
				empty: boolModifierSet,
			},
		},
		{
			s: "-empty",
			expected: &modifier{
				empty: boolModifierClear,
			},
		},
		{
			s: "noempty",
			expected: &modifier{
				empty: boolModifierClear,
			},
		},
		{
			s: "e",
			expected: &modifier{
				empty: boolModifierSet,
			},
		},
		{
			s: "encrypted",
			expected: &modifier{
				encrypted: boolModifierSet,
			},
		},
		{
			s: "executable",
			expected: &modifier{
				executable: boolModifierSet,
			},
		},
		{
			s: "x",
			expected: &modifier{
				executable: boolModifierSet,
			},
		},
		{
			s: "b",
			expected: &modifier{
				order: orderModifierSetBefore,
			},
		},
		{
			s: "-b",
			expected: &modifier{
				order: orderModifierClearBefore,
			},
		},
		{
			s: "after",
			expected: &modifier{
				order: orderModifierSetAfter,
			},
		},
		{
			s: "noafter",
			expected: &modifier{
				order: orderModifierClearAfter,
			},
		},
		{
			s: "once",
			expected: &modifier{
				condition: conditionModifierSetOnce,
			},
		},
		{
			s: "private",
			expected: &modifier{
				private: boolModifierSet,
			},
		},
		{
			s: "p",
			expected: &modifier{
				private: boolModifierSet,
			},
		},
		{
			s: "template",
			expected: &modifier{
				template: boolModifierSet,
			},
		},
		{
			s: "t",
			expected: &modifier{
				template: boolModifierSet,
			},
		},
		{
			s: "create",
			expected: &modifier{
				sourceFileType: sourceFileTypeModifierSetCreate,
			},
		},
		{
			s: "-create",
			expected: &modifier{
				sourceFileType: sourceFileTypeModifierClearCreate,
			},
		},
		{
			s: "modify",
			expected: &modifier{
				sourceFileType: sourceFileTypeModifierSetModify,
			},
		},
		{
			s: "-script",
			expected: &modifier{
				sourceFileType: sourceFileTypeModifierClearScript,
			},
		},
		{
			s: "+symlink",
			expected: &modifier{
				sourceFileType: sourceFileTypeModifierSetSymlink,
			},
		},
		{
			s: "empty,+executable,noprivate,-t",
			expected: &modifier{
				empty:      boolModifierSet,
				executable: boolModifierSet,
				private:    boolModifierClear,
				template:   boolModifierClear,
			},
		},
		{
			s: " empty , -private, notemplate ",
			expected: &modifier{
				empty:    boolModifierSet,
				private:  boolModifierClear,
				template: boolModifierClear,
			},
		},
		{
			s: "p,,-t",
			expected: &modifier{
				private:  boolModifierSet,
				template: boolModifierClear,
			},
		},
		{
			s:           "unknown",
			expectedErr: true,
		},
	} {
		t.Run(tc.s, func(t *testing.T) {
			actual, err := parseModifier(tc.s)
			if tc.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expected, actual)
			}
		})
	}
}
