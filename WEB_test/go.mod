module go_WEB/WEB_test

go 1.20

replace go_WEB/WEB_UUID => ../WEB_UUID

require (
	go_WEB/WEB_UUID v0.0.0-00010101000000-000000000000 // indirect
	go_WEB/WEB_test/test v0.0.0-00010101000000-000000000000
)

require github.com/google/uuid v1.3.0 // indirect

replace go_WEB/WEB_test/test => ./test
