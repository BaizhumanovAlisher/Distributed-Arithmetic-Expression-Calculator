module api_server

go 1.21.6

require (
	github.com/go-chi/chi/v5 v5.0.12
	github.com/go-chi/render v1.0.3
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/justinas/alice v1.2.0
	google.golang.org/grpc v1.63.0
	internal v1.0.0
)

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/ilyakaznacheev/cleanenv v1.5.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/nxadm/tail v1.4.11 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240401170217-c3f982113cda // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

replace internal v1.0.0 => ../internal
