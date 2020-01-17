module github.com/Smilefish0/sailing/user

go 1.13

require (
	github.com/Smilefish0/gener v0.0.0-20200116090144-b41e300253c4 // indirect
	github.com/Smilefish0/sailing/configer v0.0.0
	github.com/Smilefish0/sailing/starter v0.0.0
	github.com/caarlos0/env/v6 v6.1.0
	github.com/davecgh/go-spew v1.1.1
	github.com/fatih/color v1.9.0
	github.com/go-redis/redis/v7 v7.0.0-beta.5
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/configor v1.1.1
	github.com/jinzhu/gorm v1.9.12
	github.com/joho/godotenv v1.3.0
	github.com/micro/go-micro v1.18.0
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sys v0.0.0-20200116001909-b77594299b42 // indirect
)

replace github.com/Smilefish0/sailing/starter => ../starter //本地包相对路径或绝对路径

replace github.com/Smilefish0/sailing/configer => ../configer //本地包相对路径或绝对路径
