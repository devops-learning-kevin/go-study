module gomodtest

go 1.18

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/jinzhu/configor v1.2.1 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
	sum v0.0.0
)

replace (
	sum => ./sum
)