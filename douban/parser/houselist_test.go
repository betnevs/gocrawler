package parser

import (
	"io/ioutil"
	"testing"
)

const resultSize = 25

func TestHouseList(t *testing.T) {
	content, err := ioutil.ReadFile("houselist_test_data.html")
	if err != nil {
		t.Error(err)
	}
	result := HouseList(content)
	if len(result.Items) != resultSize {
		t.Errorf("result size should be %d, but had %d", resultSize, len(result.Items))
	}

}
