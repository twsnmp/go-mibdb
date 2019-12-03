package mibdb

import (
	"fmt"
	"testing"
)

// Test_MIBDB is MBDB Test
func Test_MIBDB(t *testing.T) {
	_, err := NewMIBDB("./bad.txt")
	if err == nil {
		t.Fatalf("NewMIBDB open bad file ")
	}
	t.Logf("NewMIBDB open err=%v", err)
	m, err := NewMIBDB("./mib.txt")
	if err != nil {
		t.Fatalf("NewMIBDB failed err=%v", err)
	}
	cases := []struct {
		name string
		oid  string
	}{
		{name: "sysDescr", oid: ".1.3.6.1.2.1.1.1"},
		{name: "sysDescr.0", oid: ".1.3.6.1.2.1.1.1.0"},
	}
	for _, c := range cases {
		if m.NameToOID(c.name) != c.oid {
			t.Errorf("NameToOID  name='%s' '%s' != '%s'", c.name, m.NameToOID(c.name), c.oid)
		}
		if m.OIDToName(c.oid) != c.name {
			t.Errorf("OIDToName  oid='%s' '%s' != '%s'", c.oid, m.NameToOID(c.oid), c.name)
		}
	}
	if m.OIDToName("8.999") != ".8.999" {
		t.Errorf("OIDToName   '.8.999' != '%s'", m.OIDToName(".8.999"))
	}
	if m.NameToOID("badname") != ".0.0" {
		t.Errorf("NameToOID   name='badname' '.0.0' != '%s'", m.NameToOID("badname"))
	}
	nl := m.GetNameList()
	t.Log(nl)
	if len(nl) < 1 {
		t.Errorf("GetNameList return len=0")
	}
	m.Clear()
	nl2 := m.GetNameList()
	t.Log(nl2)
	if len(nl2) != 0 {
		t.Errorf("Clear Error")
	}
	if err := m.Load("mib.txt"); err != nil {
		t.Errorf("Load err=%v", err)
	}
	if err := m.Add("test123", ".1.999"); err != nil {
		t.Errorf("Add err=%v", err)
	}
	if m.NameToOID("test123") != ".1.999" {
		t.Errorf("Add test = %s", m.NameToOID(".1.999"))
	}
	t.Log("Done")
}

func Example() {
	m, err := NewMIBDB("./mib.txt")
	if err != nil {
		fmt.Printf("NewMIBDB failed err=%v", err)
		return
	}
	fmt.Printf("sysDescr=%s", m.NameToOID("sysDescr"))
	fmt.Printf(".1.3.6.1.2.1.1.1.0=%s", m.OIDToName(".1.3.6.1.2.1.1.1.0"))
	fmt.Println(m.GetNameList())
}
