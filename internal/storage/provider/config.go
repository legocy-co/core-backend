package provider

const (
	UserObjectsBucketName      = "users"
	MarketItemImagesBucketName = "marketItems"
)

func isValidBucketName(bucketName string) bool {
	return bucketName == UserObjectsBucketName ||
		bucketName == MarketItemImagesBucketName
}
