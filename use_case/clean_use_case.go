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

func (u *CleanUseCase) Run() {
	for {
		_, err := u.reader.Read()
		if err == io.EOF {
			break
		}
	}
}
