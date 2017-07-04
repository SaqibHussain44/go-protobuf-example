package main

import (
	"log"

	"github.com/alextanhongpin/go-proto/test"
	"github.com/golang/protobuf/proto"
)

func main() {
	test := &example.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Reps:  []int64{1, 2, 3},
		Optionalgroup: &example.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}

	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshalling error: ", err)
	}

	newTest := &example.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshalling error: ", err)
	}
	log.Println(newTest)
	// Now test and newTest contains the same data
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}
}
