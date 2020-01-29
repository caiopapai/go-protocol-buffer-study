protoc -I src/ --go_out=src/ src/simple.proto
protoc -I src/ --go_out=src/ --plugin=rpc src/search-request.proto
protoc -I enum/ --go_out=enum/ enum/day_of_week.proto