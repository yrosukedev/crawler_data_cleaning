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

func TestOneRow_oneField(t *testing.T) {
	// Given
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	reader := NewMockDataReader(ctrl)
	writer := NewMockDataWriter(ctrl)
	trans := NewMockFieldTransformer(ctrl)
	transGroup := FieldTransformerGroup{
		"field 1": trans,
	}
	useCase := NewCleanUseCase(ctx, reader, writer, transGroup)

	rows := [][]Field{
		{
			{
				Name:  "field 1",
				Value: "value 1",
			},
		},
	}

	givenReaderReadsRows(reader, rows)

	transformedRows := [][]Field{
		{
			{
				Name:  "field 1",
				Value: "transformed value 1",
			},
		},
	}

	trans.EXPECT().Transform(gomock.Eq(rows[0][0].Value)).Return(transformedRows[0][0].Value, nil).Times(1)

	// Then
	writer.EXPECT().Write(gomock.Eq(transformedRows[0])).Times(1)

	// When
	useCase.Run()
}

func givenReaderReadsRows(reader *MockDataReader, rows [][]Field) {
	i := 0
	reader.EXPECT().Read().DoAndReturn(func() ([]Field, error) {
		if i < len(rows) {
			defer func() { i++ }()
			return rows[i], nil
		}
		return nil, io.EOF
	}).AnyTimes()
}
