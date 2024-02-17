package bucketexists

type Port interface {
	Execute(request Request) (*Response, error)
}
type Request struct {
	BucketName string
}

type Response struct {
	Exists bool
}
