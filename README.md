# go-mibdb
Library that converts SNMP MIB names and OIDs

[![Godoc Reference](https://godoc.org/github.com/twsnmp/go-mibdb?status.svg)](http://godoc.org/github.com/twsnmp/go-mibdb)
[![Build Status](https://travis-ci.org/twsnmp/go-mibdb.svg?branch=master)](https://travis-ci.org/twsnmp/go-mibdb)
[![Coverage Status](https://coveralls.io/repos/github/twsnmp/go-mibdb/badge.svg?branch=master)](https://coveralls.io/github/twsnmp/go-mibdb?branch=master)
[![Go Report Card](https://goreportcard.com/badge/twsnmp/go-mibdb)](https://goreportcard.com/report/twsnmp/go-mibdb)


## Usage

### Import

```go
	import mibdb "github.com/twsnmp/go-mibdb"
```

### Load MIBDB / Convert Name to OID / Convert OID to Name

```go
	m, err := NewMIBDB("./mib.txt")
	if err != nil {
		fmt.Printf("NewMIBDB failed err=%v", err)
		return
	}
	fmt.Printf("sysDescr=%s", m.NameToOID("sysDescr"))
	fmt.Printf(".1.3.6.1.2.1.1.1.0=%s", m.OIDToName(".1.3.6.1.2.1.1.1.0"))

```

# Copyright

see ./LICENSE

```
Copyright 2019 Masayuki Yamai
```
