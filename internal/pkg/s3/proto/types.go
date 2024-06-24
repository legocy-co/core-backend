package proto

type UploadImageURLOpts struct {
	URL      string
	ObjectID int
	Bucket   string
}

type UploadImageFileOpts struct {
	Data     []byte
	ObjectID int
	Bucket   string
	Format   string
}
