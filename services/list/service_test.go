package list_test

import (
	"errors"
	"file-management-api/ports/listfiles"
	"file-management-api/services/list"
	"file-management-api/utils/errs"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var now = time.Now()

func TestExecute(t *testing.T) {
	type given struct {
		request list.Request
	}
	type listFilesPort struct {
		response *listfiles.Response
		err      error
	}
	type when struct {
		listFiles listFilesPort
	}
	type expect struct {
		response *list.Response
		err      errs.AppError
	}
	type testCase struct {
		name   string
		given  *given
		when   *when
		expect *expect
	}

	testCases := []testCase{
		{
			name: "All Green",
			given: &given{
				request: list.Request{
					UserID: "ccb41151-1a5f-4be6-9d72-423f7f0a97a0",
				},
			},
			when: &when{
				listFiles: listFilesPort{
					response: &[]listfiles.List{
						{
							FileName:     "test.png",
							UploadDateAt: now,
						},
					},
				},
			},
			expect: &expect{
				response: &[]list.List{
					{
						FileName:     "test.png",
						UploadDateAt: now.Local().Format(time.RFC3339),
					},
				},
			},
		},
		{
			name: "Failed when get list response error",
			given: &given{
				request: list.Request{
					UserID: "ccb41151-1a5f-4be6-9d72-423f7f0a97a0",
				},
			},
			when: &when{
				listFiles: listFilesPort{
					err: errors.New("get list error"),
				},
			},
			expect: &expect{
				err: errs.New(http.StatusInternalServerError, errs.T001, ""),
			},
		},
		{
			name: "Failed when Invalid request",
			given: &given{
				request: list.Request{},
			},
			when: &when{},
			expect: &expect{
				err: errs.New(http.StatusBadRequest, errs.E001, ""),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			listFiles := listfiles.NewMock()
			listFiles.On("Execute", mock.Anything).Return(tc.when.listFiles.response, tc.when.listFiles.err)

			service := list.New(listFiles)

			response, err := service.Execute(tc.given.request)
			if tc.expect.err != nil {
				assert.Equal(t, tc.expect.err.GetCode(), err.GetCode())
			}
			assert.Equal(t, tc.expect.response, response)
		})
	}
}
