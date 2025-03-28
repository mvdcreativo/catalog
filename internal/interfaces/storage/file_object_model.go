package storage

type FileObject struct {
	AltText      string `bson:"alt_text" json:"alt_text"`
	IsPrimary    bool   `bson:"is_primary" json:"is_primary"` // Indica si es la imagen principal.
	ID           string `bson:"id" json:"id"`                 // UUID usado como objectName
	URL          string `bson:"url" json:"url"`
	FileName     string `bson:"file_name" json:"file_name"` // Nombre real del archivo (metadata "name")
	Size         int64  `bson:"size" json:"size"`
	ETag         string `bson:"etag" json:"etag"`
	ContentType  string `bson:"content_type" json:"content_type"`
	DisplayOrder int    `bson:"display_order" json:"display_order"` // Orden de visualización.
	RefId        string `bson:"refId" json:"refId"`
}
