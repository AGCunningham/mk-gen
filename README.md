# Mario Kart Generator

## Usage

* `/random` - will return 4 random tracks that haven't previously been selected
* `/reload` - reset any tracks previously marked as selected

The following options can be configured via environment variables

| Env Var             | Default                | Description                                                            |
|---------------------|------------------------|------------------------------------------------------------------------|
| `MK_GEN_TRACK_FILE` | `./static/tracks.yaml` | Path to the YAML file containing the data on the tracks to select from |

## Development

![Static Badge](https://img.shields.io/badge/Go-1.24-blue)

The server can be run locally using the following command

```shell
go run github.com/AGCunnigham/mk-gen
```
