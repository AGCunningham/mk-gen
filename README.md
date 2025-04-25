# Mario Kart Generator

A more controlled randomiser for MarioKart 8 tracks (including booster courses)

## Usage

The server can be built and run with the following commands from the root of the repo

```shell
# note that the tag can be changed
docker build -t mk-gen:v0.1.0 .
# note that the hostport can be changed as required
docker run -p 80:8080 mk-gen:v0.1.0
```

And then the following endpoints can be used

* `localhost/random` - will return 4 random tracks that haven't previously been selected
* `localhost/reload` - reset any tracks previously marked as selected

## Configuration

The following options can be configured via environment variables

| Env Var               | Default                | Description                                                            |
|-----------------------|------------------------|------------------------------------------------------------------------|
| `MK_GEN_TRACK_FILE`   | `./static/tracks.yaml` | Path to the YAML file containing the data on the tracks to select from |
| `MK_GEN_TEMPLATE_DIR` | `./templates/`         | Directory containing the webserver HTML templates                      |

## Development

![Static Badge](https://img.shields.io/badge/Go-1.24-blue)

The server can be run locally using the following command

```shell
go run github.com/AGCunnigham/mk-gen
```

> [!NOTE]
> If running locally then the port `8080` must be included in any URLs

## Other Notes

### Copying Image to Network Limited Devices

If wishing to run the server on a RaspberryPi or an equally under powered device it is recommended to build the docker image locally and then copy to the target.

This can be done with the following commands
```shell
# note the output file & image name should be changed as appropriate
docker build --platform linux/arm/v6 -t mk-gen:v0.1.0 .
docker save -o mk-gen-v-0-1-0.tar mk-gen:v0.1.0

rsync -e "ssh -i ${PATH_TO_SSH_KEY}" -P mk-gen-v-0-1-0.tar [USER]@[HOST]:${TARGET_DIRECTORY}/mk-gen/
```

and then on the target
```shell
docker load -i ${TARGET_DIRECTORY}/mk-gen/mk-gen-v-0-1-0.tar
docker run -p 80:8080 --name mk-gen -d mk-gen:v0.1.0
rm -r ${TARGET_DIRECTORY}/mk-gen/mk-gen-v-0-1-0.tar
```
