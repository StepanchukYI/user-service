# image for compiling binary
ARG BUILDER_IMAGE="golang:1.20"
# here we'll run binary app
ARG RUNNER_IMAGE="alpine:latest"

FROM ${BUILDER_IMAGE} as build

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN make build

FROM ${RUNNER_IMAGE}

COPY --from=build /app/bin/. .

CMD ["./main"]
