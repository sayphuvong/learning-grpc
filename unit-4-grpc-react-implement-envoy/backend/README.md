Step 1: create 'makefile' and implement "gen_go" script

Step 2: run makefile: $ make gen_go
--> two files are generated:
  + module/chatpb/chat_grpc.pb.go
  + module/chatpb/chat.pb.go

Step 3: run server: $ go run server.go

Step 4: run envoy.yaml: $ func-e run -c envoy.yaml

Step 5: run client: $ go run client.go
--> "Hello, Envoy!"
