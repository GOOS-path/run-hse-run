test_token ?= eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDQ3NjIwMjEsImlhdCI6MTY4MTY5MDAyMSwidXNlcl9pZCI6OH0.CHd-HM6JfZWIZC89fm-g5xJKbJyvAy08Vo-bvc4wYrc

generate:
	protoc --proto_path=proto proto/*.proto  --go_out=:genproto --go-grpc_out=require_unimplemented_servers=false:genproto

send-email:
	grpcurl -plaintext -d '{"email": "$(email)"}' localhost:8000 run.hse.run.AuthService/SendVerifyEmail

verify:
	grpcurl -plaintext -d '{"email": "$(email)", "code": $(code)}' localhost:8000 run.hse.run.AuthService/Verify

create-user:
	grpcurl -plaintext -d '{"email": "$(email)", "nickname": "$(nickname)", "image": "$(image)"}' localhost:8000 run.hse.run.AuthService/Registration

get-me:
	grpcurl -plaintext -rpc-header "authorization: $(test_token)" localhost:8000 run.hse.run.UserService/GetMe

get-user-by-id:
	grpcurl -plaintext -rpc-header "authorization: $(test_token)" -d '{"id": $(user_id)}' localhost:8000 run.hse.run.UserService/GetUserByID

get-user-by-nickname:
	grpcurl -plaintext -rpc-header "authorization: $(test_token)" -d '{"nickname": "$(nickname)"}' localhost:8000 run.hse.run.UserService/GetUserByNickname

change-nickname:
	grpcurl -plaintext -rpc-header "authorization: $(test_token)" -d '{"new_nickname": "$(nickname)"}' localhost:8000 run.hse.run.UserService/ChangeNickname

change-image:
	grpcurl -plaintext -rpc-header "authorization: $(test_token)" -d '{"new_image": "$(image)"}' localhost:8000 run.hse.run.UserService/ChangeImage

get-leader-board:
	grpcurl -plaintext -rpc-header "authorization: $(test_token)" localhost:8000 run.hse.run.UserService/GetLeaderBoard
