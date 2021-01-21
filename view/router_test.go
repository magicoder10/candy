package view

import (
	"testing"

	"candy/observability"

	"github.com/stretchr/testify/assert"
)

func TestRouter_AddRoute(t *testing.T) {
	testCases := []struct {
		name      string
		newRoutes []Route
		hasErr    bool
	}{
		{
			name:      "path is empty",
			newRoutes: []Route{{Path: ""}},
			hasErr:    true,
		},
		{
			name:      "path must start with /",
			newRoutes: []Route{{Path: "path"}},
			hasErr:    true,
		},
		{
			name:      "path is valid",
			newRoutes: []Route{{Path: "/first/second"}},
			hasErr:    false,
		},
		{
			name:      "path already exists",
			newRoutes: []Route{{Path: "/first/second"}, {Path: "/first/second/"}},
			hasErr:    true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			logger := observability.NewLogger(observability.Info)
			router := NewRouter(&logger)

			var err error

			for _, newRoute := range testCase.newRoutes {
				err = router.AddRoute(newRoute)
				if err != nil {
					break
				}
			}

			if testCase.hasErr {
				assert.NotNil(t, err)
				return
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
