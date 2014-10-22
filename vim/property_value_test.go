package vim

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func TestValue(t *testing.T) {
	pv := &PropertyValue{
		Type: "ArrayOfManagedObjectReference",
		Xml:  `<ManagedObjectReference type="Datacenter" xsi:type="ManagedObjectReference">ha-datacenter</ManagedObjectReference>`,
	}

	data, err := pv.Value()
	ok(t, err)

	objs := data.(*ArrayOfManagedObjectReference).Things
	assert(t, len(objs) == 1, fmt.Sprintf("%d != %d\n", len(objs), 1))
	assert(t, objs[0].Type == "Datacenter", fmt.Sprintf("%s != %s\n", objs[0].Type, "Datacenter"))
	assert(t, objs[0].Value == "ha-datacenter", fmt.Sprintf("%s != %s\n", objs[0].Value, "ha-datacenter"))
}

func TestArrayOfString(t *testing.T) {
	pv := &PropertyValue{
		Type: "ArrayOfString",
		Xml:  `<string xsi:type="xsd:string">Datacenter</string><string xsi:type="xsd:string">Alarm</string>`,
	}

	data, err := pv.Value()
	ok(t, err)
	objs := data.(*ArrayOfString).Things
	assert(t, len(objs) == 2, fmt.Sprintf("%d != %d\n", len(objs), 2))
	assert(t, objs[0] == "Datacenter", fmt.Sprintf("%s != %s\n", objs[0], "Datacenter"))
	assert(t, objs[1] == "Alarm", fmt.Sprintf("%s != %s\n", objs[1], "Alarm"))
}

func TestUnmarshalPropertyValue(t *testing.T) {
	tests := []struct {
		value    string
		expType  string
		expValue string
	}{
		{`<returnval>
			<obj type="Datacenter">ha-datacenter</obj>
			<propSet>
				<name>vmFolder</name>
				<val type="Folder" xsi:type="ManagedObjectReference">ha-folder-vm</val>
			</propSet>
			</returnval>`,
			"Folder",
			"ha-folder-vm",
		},
		{`<returnval>
			<obj type="Folder">ha-folder-vm</obj>
			<propSet>
				<name>childEntity</name>
				<val xsi:type="ArrayOfManagedObjectReference"><ManagedObjectReference type="VirtualMachine" xsi:type="ManagedObjectReference">1</ManagedObjectReference></val>
			</propSet>
			</returnval>`,
			"ArrayOfManagedObjectReference",
			`<ManagedObjectReference type="VirtualMachine" xsi:type="ManagedObjectReference">1</ManagedObjectReference>`,
		},
	}

	for _, tt := range tests {
		var obj ObjectContent
		err := xml.Unmarshal([]byte(tt.value), &obj)
		ok(t, err)

		motype := obj.PropSet[0].Val.Type
		movalue := obj.PropSet[0].Val.Xml
		assert(t, motype == tt.expType, "%s != %s", motype, tt.expType)
		assert(t, movalue == tt.expValue, "%s != %s", movalue, tt.expValue)
	}
}

func TestFixForIssue14(t *testing.T) {
	data := `<returnval>
			<obj type="Datacenter">ha-datacenter</obj>
			<propSet>
				<name>vmFolder</name>
				<val type="Folder" xsi:type="ManagedObjectReference">ha-folder-vm</val>
			</propSet>
			</returnval>`
	var obj ObjectContent
	err := xml.Unmarshal([]byte(data), &obj)
	ok(t, err)
	value, err := obj.PropSet[0].Val.Value()
	ok(t, err)
	folder := value.(*Folder)
	assert(t, folder.Type == "Folder", "%s != %s", folder.Type, "Folder")
}
