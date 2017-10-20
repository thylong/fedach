package fedach

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

// RoutingDirectoryRecord is a record of a Routing Directory file.
//
// This file is provided by the Federal Bank Reserve Service and can be found
// here: https://www.frbservices.org/EPaymentsDirectory/download.html
type RoutingDirectoryRecord struct {
	RoutingNumber         string `length:"9"  pos:"1-9"`
	OfficeCode            string `length:"1"  pos:"10"`
	ServicingFRBNumber    string `length:"9"  pos:"11-19"`
	RecordTypeCode        string `length:"1"  pos:"20"`
	ChangeDate            string `length:"6"  pos:"21-26"`
	NewRoutingNumber      string `length:"9"  pos:"27-35"`
	CustomerName          string `length:"36" pos:"36-71"`
	Address               string `length:"36" pos:"72-107"`
	City                  string `length:"20" pos:"108-127"`
	StateCode             string `length:"2"  pos:"128-129"`
	Zipcode               string `length:"5"  pos:"130-134"`
	ZipcodeExtension      string `length:"4"  pos:"135-138"`
	TelephoneAreaCode     string `length:"3"  pos:"139-141"`
	TelephonePrefixNumber string `length:"3"  pos:"142-144"`
	TelephoneSuffixNumber string `length:"4"  pos:"145-148"`
	InstitutionStatusCode string `length:"1"  pos:"149"`
	DataViewCode          string `length:"1"  pos:"150"`
	Filler                string `length:"5"  pos:"151-155"`
}

// Bytes returns RoutingDirectoryRecord bytes representation.
func (r *RoutingDirectoryRecord) Bytes() []byte {
	var routingDirectoryLine bytes.Buffer

	t := reflect.TypeOf(r).Elem()
	for i := 0; i < t.NumField(); i++ {
		fieldBytes := FieldBytes(r, i)

		routingDirectoryLine.Write(fieldBytes)
		routingDirectoryLine.Bytes()
	}
	return routingDirectoryLine.Bytes()
}

// FieldBytes encode field and return bytes.
func FieldBytes(r *RoutingDirectoryRecord, i int) []byte {
	t := reflect.TypeOf(r).Elem()
	field := t.Field(i)

	fieldValue := reflect.ValueOf(*r).FieldByName(field.Name).String()

	length, err := strconv.ParseInt(field.Tag.Get("length"), 10, 64)
	if err != nil {
		panic("wrong length struct tag format")
	}

	b := bytes.NewBufferString(fieldValue)

	whiteSpaceBytes := int(length - int64(len(fieldValue)))
	if whiteSpaceBytes < 0 {
		panic(
			fmt.Sprintf("Error while encoding %s: wrong field length.", field.Name),
		)
	}
	b.Write(bytes.Repeat([]byte(" "), whiteSpaceBytes))
	return b.Bytes()
}
