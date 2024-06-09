package proto

func NewUploadImageURLRequest(opts UploadImageURLOpts) *UploadImageURLRequest {
	return &UploadImageURLRequest{
		Meta: &ObjectParams{
			ObjectID:   int32(opts.ObjectID),
			BucketName: opts.Bucket,
		},
		Url: opts.URL,
	}
}

func NewUploadImageFileRequest(opts UploadImageFileOpts) *UploadImageFileRequest {
	return &UploadImageFileRequest{
		Meta: &ObjectParams{
			ObjectID:   int32(opts.ObjectID),
			BucketName: opts.Bucket,
			Format:     opts.Format,
		},
		Data: opts.Data,
	}
}
