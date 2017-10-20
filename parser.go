package fedach

import (
	"bytes"
	"errors"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

// Unmarshal parses the Fed-format-encoded data and stores the result in the value
// pointed to by v.
// If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.
func Unmarshal(data []byte, v interface{}) error {
	records := []RoutingDirectoryRecord{}
	reader := bytes.NewBuffer(data)
	for {
		line, err := reader.ReadString('\n')

		if len(line) == 0 {
			reflect.ValueOf(v).Elem().Set(reflect.ValueOf(records))
			return nil
		}

		record := RoutingDirectoryRecord{}
		t := reflect.TypeOf(&record).Elem()
		for i := 0; i < t.NumField(); i++ {
			pos := t.Field(i).Tag.Get("pos")
			indexes := strings.Split(pos, "-")

			start, _ := strconv.ParseInt(indexes[0], 10, 64)

			var end int64
			if len(indexes) == 2 {
				end, _ = strconv.ParseInt(indexes[1], 10, 64)
			} else {
				end = start
			}

			value := reflect.ValueOf(line[start-1 : end])
			reflect.ValueOf(&record).Elem().Field(i).Set(value)
		}
		records = append(records, record)
		if err != nil {
			reflect.ValueOf(v).Elem().Set(reflect.ValueOf(records))
			return nil
		}
	}
}

// Marshal returns the Fed file encoding of v.
func Marshal(v interface{}) (bs []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			} else if s, ok := r.(string); ok {
				err = errors.New(s)
			} else {
				err = r.(error)
			}
		}
	}()

	var records Records
	bs = records.bytes(v, records)
	return bs, err
}

// Records is an array of records from a Fed file.
type Records []byte

func (r *Records) bytes(input interface{}, records Records) []byte {
	v := reflect.ValueOf(input)
	var routingDirectoryFileBuf bytes.Buffer

	switch v.Type().String() {
	case "[]fedach.RoutingDirectoryRecord":
		for _, record := range v.Interface().([]RoutingDirectoryRecord) {
			// fmt.Printf("%s\n", record.Bytes())
			routingDirectoryFileBuf.Write(record.Bytes())
		}
	case "[][]string":
		records := make([]RoutingDirectoryRecord, 0, len(input.([][]string)))

		for _, recordLine := range input.([][]string) {
			record := RoutingDirectoryRecord{}
			for i, field := range recordLine {
				value := reflect.ValueOf(field)
				reflect.ValueOf(&record).Elem().Field(i).Set(value)
			}
			records = append(records, record)
		}
		for i, record := range records {
			// fmt.Printf("%s\n", record.Bytes())
			routingDirectoryFileBuf.Write(record.Bytes())
			if i < len(input.([][]string)) {
				routingDirectoryFileBuf.WriteByte(byte('\n'))
			}
		}
	default:
		panic("Cannot marshal provided struct.")
	}

	return routingDirectoryFileBuf.Bytes()
}
