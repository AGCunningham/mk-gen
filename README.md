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
