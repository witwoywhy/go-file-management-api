package remove_test

import (
	"errors"
	"file-management-api/ports/removefile"
	"file-management-api/services/remove"
	"file-management-api/utils/errs"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExecute(t *testing.T) {
	type given struct {
		request remove.Request
	}
	type removeFile struct {
		err error
	}
	type when struct {
		removeFile removeFile
	}
	type expect struct {
		response *remove.Response
		err      errs.AppError
	}
	type testCase struct {
		name string

		given  *given
		when   *when
		expect *expect
	}

	testCases := []testCase{
		{
			name: "All Green",
			given: &given{
				request: remove.Request{
					UserID:   "0f5384b4-cfe2-4e3e-95d4-3ba6b6d639ec",
					FileName: "test.png",
				},
			},
			when: &when{
				removeFile: removeFile{
					err: nil,
				},
			},
			expect: &expect{
				response: &remove.Response{},
			},
		},
		{
			name: "Failed when remove file",
			given: &given{
				request: remove.Request{
					UserID:   "0f5384b4-cfe2-4e3e-95d4-3ba6b6d639ec",
					FileName: "test.png",
				},
			},
			when: &when{
				removeFile: removeFile{
					err: errors.New("remove failed"),
				},
			},
			expect: &expect{
				err: errs.New(http.StatusInternalServerError, errs.T001, ""),
			},
		},
		{
			name: "Failed when invalid request",
			given: &given{
				request: remove.Request{},
			},
			when: &when{},
			expect: &expect{
				err: errs.New(http.StatusBadRequest, errs.E001, ""),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			removeFile := removefile.NewMock()
			removeFile.On("Execute", mock.Anything).Return(tc.when.removeFile.err)

			service := remove.New(removeFile)

			response, err := service.Execute(tc.given.request)
			if tc.expect.err != nil {
				assert.Equal(t, tc.expect.err.GetCode(), err.GetCode())
			}
			assert.Equal(t, tc.expect.response, response)
		})
	}
}
