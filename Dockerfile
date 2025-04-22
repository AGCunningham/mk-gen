# this needs to be build on an image that support linux/arm/v6 which is what the RaspberryPi uses
FROM golang:1.24-alpine AS build

WORKDIR /mk-gen
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o server github.com/AGCunningham/mk-gen

FROM scratch

COPY --from=build /mk-gen/server /bin/mk-gen
COPY ./static/tracks.yaml ./tracks.yaml
COPY ./templates/* ./templates/

ENV MK_GEN_TRACK_FILE="./tracks.yaml"
ENV MK_GEN_TEMPLATE_DIR="./templates/"

CMD ["/bin/mk-gen"]
