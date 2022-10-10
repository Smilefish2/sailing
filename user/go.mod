module github.com/Smilefish2/sailing/user

go 1.13

require (
	github.com/Smilefish2/sailing/configer v0.0.0
	github.com/Smilefish2/sailing/starter v0.0.0
	github.com/gf-third/mysql v1.4.2
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	golang.org/x/sys v0.0.0-20221006211917-84dc82d7e875 // indirect
)

replace github.com/Smilefish2/sailing/starter => ../starter //本地包相对路径或绝对路径

replace github.com/Smilefish2/sailing/configer => ../configer //本地包相对路径或绝对路径
