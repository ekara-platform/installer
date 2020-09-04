module github.com/ekara-platform/installer

go 1.13

require (
	github.com/GroupePSA/componentizer v0.0.0-20200904074711-6001dbb137ca
	github.com/ekara-platform/engine v1.0.1-0.20200226145759-8a60b2833327
	github.com/ekara-platform/model v1.0.1-0.20200214092618-53f62fb07250
	github.com/stretchr/testify v1.6.1
	gopkg.in/src-d/go-git.v4 v4.13.1 // indirect
)

replace github.com/ekara-platform/engine => ../engine
