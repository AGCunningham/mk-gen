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
COPY ./static/*.csv .
COPY ./templates/* ./templates/

ENV MK_GEN_TRACK_FILE="./tracks.yaml"
ENV MK_GEN_TEMPLATE_DIR="./templates/"
ENV MK_GEN_CHARACTER_FILE="./characters.csv"
ENV MK_GEN_GLIDER_FILE="./gliders.csv"
ENV MK_GEN_KART_FILE="./karts.csv"
ENV MK_GEN_TYRE_FILE="./tyres.csv"

CMD ["/bin/mk-gen"]
