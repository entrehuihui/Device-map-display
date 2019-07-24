echo Set up cross-compilation environment
set GOOS=linux
set GOARCH=amd64

go build -o mymapp main.go

echo Clear cross-compilation environment 
set GOARCH=
set GOOS=