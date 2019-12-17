//Usage: go generate <path_to_this_directory>

//go:generate protoc -I=. --go_out=plugins=grpc:. ./add.proto

package addpb
