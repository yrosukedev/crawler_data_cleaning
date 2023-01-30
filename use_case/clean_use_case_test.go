package use_case

import (
	"context"
	"github.com/golang/mock/gomock"
	"io"
	"testing"
)

func TestEmptyRow(t *testing.T) {
	// Given
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	reader := NewMockDataReader(ctrl)
	writer := NewMockDataWriter(ctrl)
	trans := NewMockFieldTransformer(ctrl)
	transGroup := FieldTransformerGroup{
		"::any field name::": trans,
	}
	useCase := NewCleanUseCase(ctx, reader, writer, transGroup)

	reader.EXPECT().Read().Return(nil, io.EOF)

	// Then
	writer.EXPECT().Write(gomock.Any()).Times(0)
	trans.EXPECT().Transform(gomock.Any()).Times(0)

	// When
	useCase.Run()
}
