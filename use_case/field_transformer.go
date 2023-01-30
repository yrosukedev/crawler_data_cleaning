package use_case

type FieldTransformer interface {
	Transform(fieldValue FieldValue) (v FieldValue, err error)
}

type FieldTransformerGroup = map[FieldName]FieldTransformer
