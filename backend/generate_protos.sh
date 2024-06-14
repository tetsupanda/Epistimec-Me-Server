#!/bin/bash

rm -rf pb
find ./proto -name "*.proto" -print0 | xargs -0 protoc --proto_path=./ --go_out=. --go_opt=module=epistemic-me-backend --go-grpc_out=. --go-grpc_opt=module=epistemic-me-backend 