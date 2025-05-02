package drivenports

type ImageCollectionInterface interface {
	SaveImage(img []byte) (string, error)
}
