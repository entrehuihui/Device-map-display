
echo Set up cross-compilation environment
set GOOS=linux
set GOARCH=amd64

go build -o mymapp main.go

echo Clear cross-compilation environment 
set GOARCH=
set GOOS=

echo swag init
swag init

echo build front end
cd forntmap
yarn install
yarn build

cd ../.