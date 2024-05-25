/home/mint/development/code-repo/projects/nest-projects/spotify-nestjs-apis2/ms-auth/node_modules/.bin/grpc_tools_node_protoc --js_out=/home/mint/development/code-repo/projects/nest-projects/spotify-nestjs-apis2/ms-auth/src/gen --grpc_out=/home/mint/development/code-repo/projects/nest-projects/spotify-nestjs-apis2/ms-auth/src/gen/ --plugin=protoc-gen-grpc=grpc_tools_node_protoc_plugin /home/mint/development/code-repo/projects/nest-projects/spotify-nestjs-apis2/ms-auth/src/proto/hello.proto --proto_path=/home/mint/development/code-repo/projects/nest-projects/spotify-nestjs-apis2/ms-auth/src/proto

npx grpc_tools_node_protoc --proto_path=src/proto --js_out=src/gen --grpc_out=src/gen --plugin=protoc-gen-grpc=grpc_tools_node_protoc_plugin src/proto/hello.proto

npx grpc_tools_node_protoc --proto_path=src/proto --js_out=src/gen --grpc_out=src/gen --plugin=protoc-gen-grpc=grpc_tools_node_protoc_plugin src/proto/hello.proto

working below:
pnpm install @nestjs/microservices grpc
pnpm install grpc-tools

➜ ms-auth git:(main) ✗ /home/mint/development/code-repo/projects/nest-projects/spotify-nestjs-apis2/ms-auth/node_modules/.bin/grpc_tools_node_protoc --js_out=src/gen --grpc_out=src/gen --plugin=protoc-gen-grpc=/home/mint/development/code-repo/projects/nest-projects/spotify-nestjs-apis2/ms-auth/node_modules/.bin/grpc_tools_node_protoc_plugin src/proto/hello.proto --proto_path=src/proto

./node_modules/.bin/grpc_tools_node_protoc --js_out=./src/gen --grpc_out=./src/gen --plugin=protoc-gen-grpc=./node_modules/.bin/grpc_tools_node_protoc_plugin ./src/proto/hello.proto --proto_path=./src/proto

rotoc --plugin=protoc-gen-ts_proto=./node_modules/.bin/protoc-gen-ts_proto --ts_proto_out=./src/gen --ts_proto_opt=nestpJs=true,outputServices=grpc-js ./src/proto/\*.proto --proto_path=./src/proto

./node_modules/.bin/grpc_tools_node_protoc --js_out=./src/gen --grpc_out=./src/gen --plugin=protoc-gen-grpc=./node_modules/.bin/grpc_tools_node_protoc_plugin ./src/proto/hello.proto --proto_path=./src/proto --ts_proto_opt=nestpJs=true,outputServices=grpc-js

./node_modules/.bin/grpc_tools_node_protoc --plugin=protoc-gen-ts_proto=./node_modules/.bin/protoc-gen-ts_proto --ts_proto_out=./src/gen --ts_proto_opt=nestJs=true,outputServices=grpc-js ./src/proto/\*.proto --proto_path=./src/proto

npx protoc --plugin=./node_modules/.bin/protoc-gen-tc_proto --ts_proto_out=./ --ts_proto_opt=nestJs=true ./proto/todo.proto
