module es2mysql

go 1.14

replace (
	Common => ./Common
	Store => ./Store
)

require (
	Store v0.0.0-00010101000000-000000000000 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/elastic/go-elasticsearch/v7 v7.6.0 // indirect
)
