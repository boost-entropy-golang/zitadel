ARG GO_VERSION=1.17

#######################
## These steps set platform / arch type specific variables
#######################
FROM alpine:3 AS arm64-base
ENV PROTOC_ARCH aarch_64

FROM alpine:3 AS amd64-base
ENV PROTOC_ARCH x86_64

#######################
## This step sets up the folder structure,
## initalices go mods,
## downloads the protofiles and protoc for later use
#######################
FROM ${BUILDARCH}-base AS base
ARG PROTOC_VERSION=3.18.0
ARG PROTOC_ZIP=protoc-${PROTOC_VERSION}-linux-${PROTOC_ARCH}.zip
ARG GATEWAY_VERSION=2.10.0
#0.6.3+ has an issue with enums and imports:
#https://github.com/envoyproxy/protoc-gen-validate/issues/565
#https://github.com/envoyproxy/protoc-gen-validate/issues/552
ARG VALIDATOR_VERSION=0.6.2

RUN apk add tar curl rsync
WORKDIR /proto

#protoc
RUN apk add tar curl \
    && curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/$PROTOC_ZIP \
    && unzip -o $PROTOC_ZIP -d /usr/local bin/protoc \
    && unzip -o $PROTOC_ZIP -d /proto include/* \
    && rm -f $PROTOC_ZIP \
    && curl https://raw.githubusercontent.com/envoyproxy/protoc-gen-validate/v${VALIDATOR_VERSION}/validate/validate.proto --create-dirs -o include/validate/validate.proto  \
    && curl https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/v${GATEWAY_VERSION}/protoc-gen-openapiv2/options/annotations.proto --create-dirs -o include/protoc-gen-openapiv2/options/annotations.proto \
    && curl https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/v${GATEWAY_VERSION}/protoc-gen-openapiv2/options/openapiv2.proto --create-dirs -o include/protoc-gen-openapiv2/options/openapiv2.proto \
    && curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto --create-dirs -o include/google/api/annotations.proto \
    && curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto --create-dirs -o include/google/api/http.proto \
    && curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/field_behavior.proto --create-dirs -o include/google/api/field_behavior.proto

#zitadel protos
COPY proto/ include/.

#######################
## Go dependencies
## Speed up this step by mounting your local go mod pkg directory
#######################
FROM golang:${GO_VERSION} as go-dep
RUN mkdir -p src/github.com/zitadel/zitadel
WORKDIR /go/src/github.com/zitadel/zitadel

#download modules
COPY ./go.* .
RUN go mod download

# install tools
COPY tools ./tools
RUN ./tools/install.sh

ARG COCKROACH_BINARY=cockroach
RUN apt install openssl tzdata tar

# cockroach binary used to backup database
RUN mkdir /usr/local/lib/cockroach
RUN wget -qO- https://binaries.cockroachdb.com/cockroach-v21.2.7.linux-amd64.tgz \
    | tar  xvz && cp -i cockroach-v21.2.7.linux-amd64/cockroach /usr/local/bin/
RUN rm -r cockroach-v21.2.7.linux-amd64

#######################
## generates static files
#######################
FROM go-dep AS go-static
COPY internal/ui/login/static internal/ui/login/static
COPY internal/ui/login/statik internal/ui/login/statik
COPY internal/notification/static internal/notification/static
COPY internal/notification/statik internal/notification/statik
COPY internal/static internal/static
COPY internal/statik internal/statik

RUN go generate internal/ui/login/statik/generate.go \
    && go generate internal/ui/login/static/generate.go \
    && go generate internal/notification/statik/generate.go \
    && go generate internal/statik/generate.go

#######################
## generates grpc stub
#######################
FROM go-static AS go-stub
COPY --from=base /proto /proto
COPY --from=base /usr/local/bin /usr/local/bin/.

COPY build/zitadel/generate-grpc.sh build/zitadel/generate-grpc.sh
COPY internal/protoc internal/protoc
COPY openapi/statik openapi/statik
COPY internal/api/assets/generator internal/api/assets/generator
COPY internal/config internal/config
COPY internal/errors internal/errors

RUN build/zitadel/generate-grpc.sh \
    && go generate openapi/statik/generate.go \
    && mkdir -p docs/apis/assets/ \
    && go run internal/api/assets/generator/asset_generator.go -directory=internal/api/assets/generator/ -assets=docs/apis/assets/assets.md

#######################
## Go base build
#######################
FROM go-stub as go-base
# copy remaining zitadel files
COPY . .


#######################
## copy for local dev
#######################
FROM scratch as go-copy
COPY --from=go-stub /go/src/github.com/zitadel/zitadel/internal/ui/login/statik/statik.go internal/ui/login/statik/statik.go
COPY --from=go-stub /go/src/github.com/zitadel/zitadel/internal/notification/statik/statik.go internal/notification/statik/statik.go
COPY --from=go-stub /go/src/github.com/zitadel/zitadel/internal/statik/statik.go internal/statik/statik.go
COPY --from=go-stub /go/src/github.com/zitadel/zitadel/openapi/statik/statik.go openapi/statik/statik.go

COPY --from=go-stub /go/src/github.com/zitadel/zitadel/pkg/grpc pkg/grpc
COPY --from=go-stub /go/src/github.com/zitadel/zitadel/openapi/v2/zitadel openapi/v2/zitadel
COPY --from=go-stub /go/src/github.com/zitadel/zitadel/openapi/statik/statik.go openapi/statik/statik.go
COPY --from=go-stub /go/src/github.com/zitadel/zitadel/internal/protoc/protoc-gen-authoption/templates.gen.go internal/protoc/protoc-gen-authoption/templates.gen.go
COPY --from=go-stub /go/src/github.com/zitadel/zitadel/internal/protoc/protoc-gen-authoption/authoption/options.pb.go internal/protoc/protoc-gen-authoption/authoption/options.pb.go
COPY --from=go-stub /go/src/github.com/zitadel/zitadel/docs/apis/proto docs/docs/apis/proto
COPY --from=go-stub /go/src/github.com/zitadel/zitadel/docs/apis/assets docs/docs/apis/assets

COPY --from=go-stub /go/src/github.com/zitadel/zitadel/internal/api/assets/authz.go ./internal/api/assets/authz.go
COPY --from=go-stub /go/src/github.com/zitadel/zitadel/internal/api/assets/router.go ./internal/api/assets/router.go

######################
## Output generated files by running a container from this target mounting the target dir to /dest
######################
FROM base as merge-go-copy

RUN apk add rsync

COPY --from=go-copy / /src

ENV UID=1000
ENV GID=1000

ENTRYPOINT [ "sh", "-c" ]
CMD [ "rsync -a --chown $UID:$GID /src/ /dest" ]

#######################
## Go test
#######################
FROM go-base as go-test

# Migrations for cockroach-secure
RUN go install github.com/rakyll/statik \
    && go test -race -v -coverprofile=profile.cov $(go list ./... | grep -v /operator/)

#######################
## Go test results
#######################
FROM scratch as go-codecov
COPY --from=go-test /go/src/github.com/zitadel/zitadel/profile.cov profile.cov

#######################
## Go prod build
#######################
FROM go-test as prod-go-build
ARG BUILDARCH
ARG VERSION=""

RUN CGO_ENABLED=0 GOOS=linux GOARCH=${BUILDARCH} go build -a -installsuffix cgo -ldflags "-X main.version=${VERSION:-'dev'} -extldflags \"-static\"" -o zitadel-linux-${BUILDARCH} cmd/zitadel/main.go

#######################
## Go dev build
#######################
FROM go-base as dev-go-build
ENTRYPOINT [ "go", "run", "cmd/zitadel/main.go" ]

#######################
## Only Copy Assets
#######################
FROM scratch as prod-go-export
ARG BUILDARCH
COPY --from=prod-go-build /go/src/github.com/zitadel/zitadel/zitadel-linux-${BUILDARCH} .
