package upload_test

import (
	"errors"
	"file-management-api/ports/bucketexists"
	"file-management-api/ports/config"
	"file-management-api/ports/createbucket"
	"file-management-api/ports/listfiles"
	"file-management-api/ports/uploadfile"
	"file-management-api/services/upload"
	"file-management-api/utils/errs"
	"mime/multipart"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExecute(t *testing.T) {
	type given struct {
		request upload.Request
	}
	type bucketExists struct {
		response *bucketexists.Response
		err      error
	}
	type listFiles struct {
		response *listfiles.Response
		err      error
	}
	type createBucket struct {
		err error
	}
	type uploadFile struct {
		err error
	}
	type when struct {
		bucketExists bucketExists
		listFiles    listFiles
		createBucket createBucket
		uploadFile   uploadFile
		config       config.AppConfig
	}
	type expect struct {
		response *upload.Response
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
			name: "All Green if bucket is not exists",
			given: &given{
				request: upload.Request{
					UserID: "ba347ea4-3e0c-4215-af5f-c2996169474f",
					File: &multipart.FileHeader{
						Filename: "test.png",
						Header: map[string][]string{
							"Content-Type": {"png"},
						},
						Size: 200,
					},
				},
			},
			when: &when{
				bucketExists: bucketExists{
					response: &bucketexists.Response{
						Exists: false,
					},
				},
				listFiles: listFiles{
					response: &[]listfiles.List{
						{
							FileName:     "cat.png",
							UploadDateAt: time.Time{},
						},
					},
				},
				createBucket: createBucket{
					err: nil,
				},
				uploadFile: uploadFile{
					err: nil,
				},
				config: config.AppConfig{
					AllowFileExtensions: map[string]bool{
						"png": true,
					},
					MaxSizeFile: 5000000,
				},
			},
			expect: &expect{
				response: &upload.Response{},
			},
		},
		{
			name: "All Green if bucket is exists",
			given: &given{
				request: upload.Request{
					UserID: "ba347ea4-3e0c-4215-af5f-c2996169474f",
					File: &multipart.FileHeader{
						Filename: "test.png",
						Header: map[string][]string{
							"Content-Type": {"png"},
						},
						Size: 200,
					},
				},
			},
			when: &when{
				bucketExists: bucketExists{
					response: &bucketexists.Response{
						Exists: true,
					},
				},
				listFiles: listFiles{
					response: &[]listfiles.List{
						{
							FileName:     "cat.png",
							UploadDateAt: time.Time{},
						},
					},
				},
				uploadFile: uploadFile{
					err: nil,
				},
				config: config.AppConfig{
					AllowFileExtensions: map[string]bool{
						"png": true,
					},
					MaxSizeFile: 5000000,
				},
			},
			expect: &expect{
				response: &upload.Response{},
			},
		},
		{
			name: "Failed when create bucket",
			given: &given{
				request: upload.Request{
					UserID: "ba347ea4-3e0c-4215-af5f-c2996169474f",
					File: &multipart.FileHeader{
						Filename: "test.png",
						Header: map[string][]string{
							"Content-Type": {"png"},
						},
						Size: 200,
					},
				},
			},
			when: &when{
				bucketExists: bucketExists{
					response: &bucketexists.Response{
						Exists: false,
					},
				},
				createBucket: createBucket{
					err: errors.New("create bucket error"),
				},
				config: config.AppConfig{
					AllowFileExtensions: map[string]bool{
						"png": true,
					},
					MaxSizeFile: 5000000,
				},
			},
			expect: &expect{
				err: errs.New(http.StatusInternalServerError, errs.T001, ""),
			},
		},
		{
			name: "Failed when upload response error",
			given: &given{
				request: upload.Request{
					UserID: "ba347ea4-3e0c-4215-af5f-c2996169474f",
					File: &multipart.FileHeader{
						Filename: "test.png",
						Header: map[string][]string{
							"Content-Type": {"png"},
						},
						Size: 200,
					},
				},
			},
			when: &when{
				bucketExists: bucketExists{
					response: &bucketexists.Response{
						Exists: true,
					},
				},
				listFiles: listFiles{
					response: &[]listfiles.List{
						{
							FileName:     "cat.png",
							UploadDateAt: time.Time{},
						},
					},
				},
				uploadFile: uploadFile{
					err: errors.New("upload failed"),
				},
				config: config.AppConfig{
					AllowFileExtensions: map[string]bool{
						"png": true,
					},
					MaxSizeFile: 5000000,
				},
			},
			expect: &expect{
				err: errs.New(http.StatusInternalServerError, errs.T001, ""),
			},
		},
		{
			name: "Failed when exists file in bucket",
			given: &given{
				request: upload.Request{
					UserID: "ba347ea4-3e0c-4215-af5f-c2996169474f",
					File: &multipart.FileHeader{
						Filename: "test.png",
						Header: map[string][]string{
							"Content-Type": {"png"},
						},
						Size: 200,
					},
				},
			},
			when: &when{
				bucketExists: bucketExists{
					response: &bucketexists.Response{
						Exists: true,
					},
				},
				listFiles: listFiles{
					response: &[]listfiles.List{
						{
							FileName: "test.png",
						},
					},
				},
				config: config.AppConfig{
					AllowFileExtensions: map[string]bool{
						"png": true,
					},
					MaxSizeFile: 5000000,
				},
			},
			expect: &expect{
				err: errs.New(http.StatusBadRequest, errs.E001, ""),
			},
		},
		{
			name: "Failed when list files response error",
			given: &given{
				request: upload.Request{
					UserID: "ba347ea4-3e0c-4215-af5f-c2996169474f",
					File: &multipart.FileHeader{
						Filename: "test.png",
						Header: map[string][]string{
							"Content-Type": {"png"},
						},
						Size: 200,
					},
				},
			},
			when: &when{
				bucketExists: bucketExists{
					response: &bucketexists.Response{
						Exists: true,
					},
				},
				listFiles: listFiles{
					err: errors.New("upload failed"),
				},
				config: config.AppConfig{
					AllowFileExtensions: map[string]bool{
						"png": true,
					},
					MaxSizeFile: 5000000,
				},
			},
			expect: &expect{
				err: errs.New(http.StatusInternalServerError, errs.T001, ""),
			},
		},
		{
			name: "Failed when bucket exists response error",
			given: &given{
				request: upload.Request{
					UserID: "ba347ea4-3e0c-4215-af5f-c2996169474f",
					File: &multipart.FileHeader{
						Filename: "test.png",
						Header: map[string][]string{
							"Content-Type": {"png"},
						},
						Size: 200,
					},
				},
			},
			when: &when{
				bucketExists: bucketExists{
					err: errors.New("upload failed"),
				},
				config: config.AppConfig{
					AllowFileExtensions: map[string]bool{
						"png": true,
					},
					MaxSizeFile: 5000000,
				},
			},
			expect: &expect{
				err: errs.New(http.StatusInternalServerError, errs.T001, ""),
			},
		},
		{
			name: "Failed when request is empty",
			given: &given{
				request: upload.Request{},
			},
			when: &when{},
			expect: &expect{
				err: errs.New(http.StatusBadRequest, errs.E001, ""),
			},
		},
		{
			name: "Failed when file extension not found",
			given: &given{
				request: upload.Request{
					UserID: "ba347ea4-3e0c-4215-af5f-c2996169474f",
					File: &multipart.FileHeader{
						Filename: "test",
					},
				},
			},
			when: &when{},
			expect: &expect{
				err: errs.New(http.StatusBadRequest, errs.E001, ""),
			},
		},
		{
			name: "Failed when file extension not allow",
			given: &given{
				request: upload.Request{
					UserID: "ba347ea4-3e0c-4215-af5f-c2996169474f",
					File: &multipart.FileHeader{
						Filename: "test.txt",
					},
				},
			},
			when: &when{
				config: config.AppConfig{
					AllowFileExtensions: map[string]bool{
						"png": true,
					},
					MaxSizeFile: 5000000,
				},
			},
			expect: &expect{
				err: errs.New(http.StatusBadRequest, errs.E001, ""),
			},
		},
		{
			name: "Failed when file over max size",
			given: &given{
				request: upload.Request{
					UserID: "ba347ea4-3e0c-4215-af5f-c2996169474f",
					File: &multipart.FileHeader{
						Filename: "test.png",
						Size:     5000000 + 1,
					},
				},
			},
			when: &when{
				config: config.AppConfig{
					AllowFileExtensions: map[string]bool{
						"png": true,
					},
					MaxSizeFile: 5000000,
				},
			},
			expect: &expect{
				err: errs.New(http.StatusBadRequest, errs.E001, ""),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config.Config = tc.when.config

			bucketExists := bucketexists.NewMock()
			bucketExists.On("Execute", mock.Anything).Return(tc.when.bucketExists.response, tc.when.bucketExists.err)

			listFiles := listfiles.NewMock()
			listFiles.On("Execute", mock.Anything).Return(tc.when.listFiles.response, tc.when.listFiles.err)

			createBucket := createbucket.NewMock()
			createBucket.On("Execute", mock.Anything).Return(tc.when.createBucket.err)

			uploadFile := uploadfile.NewMock()
			uploadFile.On("Execute", mock.Anything).Return(tc.when.uploadFile.err)

			service := upload.New(bucketExists, listFiles, createBucket, uploadFile)

			response, err := service.Execute(tc.given.request)
			if tc.expect.err != nil {
				assert.Equal(t, tc.expect.err.GetCode(), err.GetCode())
			}
			assert.Equal(t, tc.expect.response, response)
		})
	}
}
