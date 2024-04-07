module auth_service

go 1.21

require (
	github.com/golang-jwt/jwt/v5 v5.2.1
	golang.org/x/crypto v0.22.0
	google.golang.org/grpc v1.63.0
	internal v1.0.0
)

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/ilyakaznacheev/cleanenv v1.5.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240401170217-c3f982113cda // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

replace internal v1.0.0 => ../internal
