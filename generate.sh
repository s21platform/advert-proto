protoc  --go_out=./  \
        --go-grpc_out=./ \
        advert.proto

protoc --doc_out=. --doc_opt=markdown,README.md ./advert.proto

#protoc  --go_out=./  \
#        --go-grpc_out=./ \
        #new_user.proto