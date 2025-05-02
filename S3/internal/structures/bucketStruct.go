package structures

// Structure for storing data about the bucket and retriving it in the xml format
type Bucket struct {
	Name             string `xml:"Name"`
	CreatedTime      string `xml:"CreatedTime"`
	LastModifiedTime string `xml:"LastModifiedTime"`
	Status           string `xml:"Status"`
}

// Structure for storing the buckets in order to retrive the data of all bucket in xml format at once
type ListAllMyBucketsResult struct {
	AllBuckets []Bucket `xml:"Buckets>Bucket"`
}
