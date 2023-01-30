package use_case

type FieldName = string

type FieldValue = string

type Field struct {
	Name  FieldName
	Value FieldValue
}

func (f *Field) transformValue(trans FieldTransformer) (newField *Field, err error) {
	if trans == nil {
		return f, nil
	}

	var transformedValue FieldValue

	transformedValue, err = trans.Transform(f.Value)
	if err != nil {
		return nil, err
	}

	newField = &Field{
		Name:  f.Name,
		Value: transformedValue,
	}

	return newField, nil
}
