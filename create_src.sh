# The file must be run in the root directory of the project

# path to project must not contain spaces

if [ ! -d src ]; then
  mkdir src
fi

OLD_GOPATH=$(go env GOPATH)

export GOPATH="$(pwd)/src"

cd api_server && go mod download && cd ..
cd auth_service && go mod download && cd ..
cd expression_solver && go mod download && cd ..
cd internal && go mod download && cd ..

export GOPATH="$OLD_GOPATH"