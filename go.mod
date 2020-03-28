module es2mysql

go 1.14

replace (
	Common => ./Common
	Store => ./Store
)

require Store v0.0.0-00010101000000-000000000000 // indirect
