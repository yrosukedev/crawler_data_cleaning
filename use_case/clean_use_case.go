package use_case

import (
	"context"
	"io"
)

type CleanUseCase struct {
	ctx              context.Context
	reader           DataReader
	writer           DataWriter
	transformerGroup FieldTransformerGroup
}

func NewCleanUseCase(ctx context.Context, reader DataReader, writer DataWriter, transformerGroup FieldTransformerGroup) *CleanUseCase {
	return &CleanUseCase{
		ctx:              ctx,
		reader:           reader,
		writer:           writer,
		transformerGroup: transformerGroup,
	}
}

func (u *CleanUseCase) Run() error {
	for {
		record, err := u.reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		var transformedRecord []Field
		transformedRecord, err = u.transformRecord(record)
		if err != nil {
			return err
		}

		err = u.writer.Write(transformedRecord)
		if err != nil {
			return err
		}
	}

	return nil
}

// !!! SMELL: Primitive Obsession
// []Field should be encapsulated by a type, and transformRecord() should be moved to that type.
func (u *CleanUseCase) transformRecord(record []Field) ([]Field, error) {
	var transformedRecord []Field
	for _, field := range record {
		transformedField, err := field.transformValue(u.transformerGroup[field.Name])
		if err != nil {
			return nil, err
		}

		transformedRecord = append(transformedRecord, *transformedField)
	}
	return transformedRecord, nil
}
