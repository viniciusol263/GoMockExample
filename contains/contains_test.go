package contains

import (
	"gomockexample/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestContains_SearchId(t *testing.T) {
	ctr := gomock.NewController(t)

	mockedValue := mock.NewMockValue(ctr)

	mockedValue.EXPECT().Print(true).Return(nil)

	contains := &Contains{
		values: []ValueWrapper{{id: 1, value: mockedValue}},
	}

	_, err := contains.searchId(1)
	if err != nil {
		t.Fatalf("Error occoured %v\n", err)
	}
}
