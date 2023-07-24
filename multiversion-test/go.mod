module multiversion-test

go 1.20

replace github.com/fnaranjo-vmw/multiversion-main/001 => github.com/fnaranjo-vmw/multiversion-main v0.0.1

replace github.com/fnaranjo-vmw/multiversion-main/002 => github.com/fnaranjo-vmw/multiversion-main v0.0.2

replace github.com/fnaranjo-vmw/multiversion-main/dev => ../

require (
	github.com/fnaranjo-vmw/multiversion-main/001 v0.0.0-00010101000000-000000000000
	github.com/fnaranjo-vmw/multiversion-main/002 v0.0.0-00010101000000-000000000000
	github.com/fnaranjo-vmw/multiversion-main/dev v0.0.0-00010101000000-000000000000
)

require github.com/fnaranjo-vmw/multiversion-companion v0.0.2 // indirect
