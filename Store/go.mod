module Store

go 1.14

replace Common => ../Common

require (
	Common v0.0.0-00010101000000-000000000000
	github.com/araddon/gou v0.0.0-20190110011759-c797efecbb61 // indirect
	github.com/bitly/go-hostpool v0.1.0 // indirect
	github.com/garyburd/redigo v1.6.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/mattbaird/elastigo v0.0.0-20170123220020-2fe47fd29e4b
)
