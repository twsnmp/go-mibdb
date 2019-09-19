package mibdb

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// MIBDB is MIB  Name OID Database
type MIBDB struct {
	path      string
	nameToOid map[string]string
	oidToName map[string]string
	Errors    []string
}

// NewMIBDB create new MIBDB struct
func NewMIBDB(path string) (*MIBDB, error) {
	m := MIBDB{
		path:      path,
		nameToOid: make(map[string]string),
		oidToName: make(map[string]string),
		Errors:    []string{},
	}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	a := strings.Split(string(b), "\n")
	for i := 0; i < len(a)-1; i += 2 {
		oid := strings.TrimSpace(a[i])
		name := strings.TrimSpace(a[i+1])
		na := strings.Split(name, ".")
		if len(na) < 1 {
			m.Errors = append(m.Errors, fmt.Sprintf("Invalid Line %#v = %#v", oid, name))
			continue
		}
		sname := na[len(na)-1]
		if val, ok := m.oidToName[oid]; ok {
			m.Errors = append(m.Errors, fmt.Sprintf("Dup OID %#v=%#v : %#v", oid, name, val))
			continue
		}
		if val, ok := m.nameToOid[sname]; ok {
			m.Errors = append(m.Errors, fmt.Sprintf("Dup name %#v=%#v : %#v", oid, sname, val))
			continue
		}
		m.oidToName[oid] = sname
		m.nameToOid[sname] = oid
	}
	if len(m.oidToName) < 1 || len(m.nameToOid) < 1 {
		return nil, fmt.Errorf("Invalid MIBDB file format")
	}
	return &m, nil
}

// OIDToName convert OID to Name function
func (m *MIBDB) OIDToName(oid string) string {
	if len(oid) > 0 && oid[0] != '.' {
		oid = "." + oid
	}
	if n, ok := m.oidToName[oid]; ok {
		return n
	}
	a := strings.Split(oid, ".")
	for i := len(a) - 1; i > 0; i-- {
		o := strings.Join(a[:i], ".")
		if n, ok := m.oidToName[o]; ok {
			return n + "." + strings.Join(a[i:], ".")
		}
	}
	return oid
}

// NameToOID convert Name to OID function
func (m *MIBDB) NameToOID(name string) string {
	a := strings.Split(name, ".")
	if o, ok := m.nameToOid[a[0]]; ok {
		if len(a) == 1 {
			return o
		}
		return o + "." + strings.Join(a[1:], ".")
	}
	return ".0.0"
}