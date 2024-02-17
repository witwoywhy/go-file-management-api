package integrationtest

import (
	"context"
	"file-management-api/httpserv"
	"testing"
)

func TestIntegrationPipeline(t *testing.T) {
	type testCase struct {
		name     string
		testFunc func(t *testing.T)
	}

	ctx := context.Background()
	container := runContainer(t, ctx)
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			t.Error(err)
		}
	}()

	initApp(t, container, ctx)
	go httpserv.Run()

	testCases := []testCase{
		{
			name:     "upload",
			testFunc: upload,
		},
		{
			name:     "list",
			testFunc: list,
		},
		{
			name:     "remove",
			testFunc: remove,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testFunc(t)
		})
	}
}
