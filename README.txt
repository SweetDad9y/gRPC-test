//Дополняет прото файл нереализованными методами
protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. example.proto

//Evans утилита для тестировки сервера
evans <protofile path> -p <port>

//Генерация прото и grpc файла
protoc --go_out=. example.proto --go-grpc_out=. example.proto
