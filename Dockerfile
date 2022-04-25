# Build
FROM golang as build

RUN rm -f /etc/shadow /etc/group /etc/passwd

RUN useradd -u 10001 -m nonroot

USER nonroot

WORKDIR $GOPATH/src/$SERVICE_NAME

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/$SERVICE_NAME

# production
FROM registry.access.redhat.com/ubi8/ubi

EXPOSE 9000

USER root

COPY --from=build /scripts /scripts

# Install node and secure-spreadsheet
RUN /scripts/set-up-secure-spreadsheet

RUN npm -v
RUN node -v

COPY --from=build /etc/passwd /etc/passwd

USER nonroot

COPY --from=build /go/bin/$SERVICE_NAME /go/bin/$SERVICE_NAME

ENTRYPOINT ["/bin/sh", "-c", "/go/bin/${SERVICE_NAME}"]