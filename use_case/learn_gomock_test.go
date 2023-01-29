package use_case

import (
	"github.com/golang/mock/gomock"
	"github.com/yrosukedev/crawler_data_cleaning/mock_use_case"
	"testing"
)

type Foo interface {
	Bar(x int) int
}

func TestLearnMock_functionCalled(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	f := mock_use_case.NewMockFoo(ctrl)

	// Then
	f.
		EXPECT().
		Bar(gomock.Eq(99)).
		Times(1)

	// When
	f.Bar(99)
}

func TestLearnMock_functionNotCalled(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	f := mock_use_case.NewMockFoo(ctrl)

	// Then
	f.
		EXPECT().
		Bar(gomock.Eq(99)).
		Times(0)

	// When
}

func TestLearnMock_functionCalledWithWrongArguments(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	f := mock_use_case.NewMockFoo(ctrl)

	// Then
	f.
		EXPECT().
		Bar(gomock.Eq(99)).
		Times(1)

	// When
	f.Bar(10)
}
