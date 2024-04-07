# The file must be run in the root directory of the project

# path to project must not contain spaces
mkdir src
go env -w GOPATH="$(pwd)/src"

cd api_server && go mod download && cd ..
cd auth_service && go mod download && cd ..
cd expression_solver && go mod download && cd ..
cd internal && go mod download && cd ..