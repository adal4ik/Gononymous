package structures

// Structure for storing data about the object and retriving it in the xml format
type Object struct {
	ObjectKey    string `xml:"Objectkey"`
	Size         string `xml:"Size"`
	ContentType  string `xml:"ContentType"`
	LastModified string `xml:"LastModified"`
}
