package createbucket

type Port interface {
	Execute(request Request) error
}

type Request struct {
	BucketName string
}
