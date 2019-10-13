# go-mibdb
Library that converts SNMP MIB names and OIDs

[![Godoc Reference](https://godoc.org/github.com/twsnmp/go-mibdb?status.svg)](http://godoc.org/github.com/twsnmp/go-mibdb)
[![Build Status](https://travis-ci.org/twsnmp/go-mibdb.svg?branch=master)](https://travis-ci.org/twsnmp/go-mibdb)
[![Coverage Status](https://coveralls.io/repos/github/twsnmp/go-mibdb/badge.svg?branch=master)](https://coveralls.io/github/twsnmp/go-mibdb?branch=master)
[![Go Report Card](https://goreportcard.com/badge/twsnmp/go-mibdb)](https://goreportcard.com/report/twsnmp/go-mibdb)


## Usage (使用方法)

### Import (インポート)

```go
	import mibdb "github.com/twsnmp/go-mibdb"
```

### Load MIBDB / Convert Name to OID / Convert OID to Name

名前とOID間の相互変換

```go
	m, err := mibdb.NewMIBDB("./mib.txt")
	if err != nil {
		fmt.Printf("NewMIBDB failed err=%v", err)
		return
	}
	fmt.Printf("sysDescr=%s", m.NameToOID("sysDescr"))
	fmt.Printf(".1.3.6.1.2.1.1.1.0=%s", m.OIDToName(".1.3.6.1.2.1.1.1.0"))

```
# How to make MIB DB File(mib.txt) 

To create mib.txt, execute the snmptranslate command of Net-SNMP as follows.

mib.txtを作るには、Net-SNMPのsnmptranslateコマンドを次のように実行します。

```
$snmptranslate  -T os  > mib.txt
```

If an extended MIB is installed in the Net-SNMP environment, it can also be extended.

Net-SNMPの環境に拡張MIBをインストールすれば、拡張MIBにも対応できます。


# Copyright

see ./LICENSE

```
Copyright 2019 Masayuki Yamai
```
