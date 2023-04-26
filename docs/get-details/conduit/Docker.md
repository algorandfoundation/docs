# Docker Image

Algorand's Conduit data pipeline packaged for docker.

This document provides some basic guidance and commands tailored
for the docker container. For additional information refer to
the [full documentation](https://developer.algorand.org/docs/get-details/conduit/GettingStarted/).

# Configuration

Special locations in the container designed to modify behavior:

| File | Description |
| ---- | ----------- |
| /etc/algorand/config.yml | Required. Conduit configuration file. Definition for the pipeline behavior. |
| /data | Optional. Data directory. For persistence you may mount a volume at this location. See volume permissions sections for additional information. |

# Usage

Conduit needs a configuration file to define the data pipeline.
There are built-in utilities to help create the configuration.
Once a configuration is made launch conduit and pass in the configuration
to use.

## Creating a conduit.yml configuration

The init subcommand can be used to create a configuration template.
See the options here:
```
docker run algorand/conduit init -h
```

For a simple default configuration, run with no arguments:
```
docker run algorand/conduit init > conduit.yml
```

Plugins can also be named directly.

See a list of available plugins:
```
docker run algorand/conduit list
```

Provide plugins to the init command:
```
docker run algorand/conduit init --importer algod --processors filter_processor --exporter postgresql > conduit.yml
```

## Run with conduit.yml

With `conduit.yml` in your current working directory, it can be mounted directly to `/data/conduit.yml`. This is good for testing and some deployments which override the starting round. For a more complete deployment see the next section which explains how to mount the entire data directory.

Mount `conduit.yml` with the following command:
```
docker run -it -v $(pwd)/conduit.yml:/etc/algorand/conduit.yml algorand/conduit
```

# Mounting the Data Directory

For production deployments, you should consider mounting the entire data directory. This way you can persist state across images during an upgrade, or for backups. The data directory is located at `/data`. When mounting a data directory, it must contain the `conduit.yml` file.

```
docker run -it -v $(pwd)/local_data_dir:/data algorand/conduit
```

## Volume Permissions

The container executes in the context of the `algorand` user with UID=999 and GID=999 which is handled differently depending on your operating system or deployment platform. During startup the container temporarily runs as root in order to modify the permissions of `/data`. It then changes to the `algorand` user. This can sometimes cause problems, for example if your deployment platform doesn't allow containers to run as the root user.

### Use specific UID and GID

If you do not want the container to start as the root user you can specify a UID and GID. On the host system, ensure the directory being mounted uses UID=999 and GID=999. If the directory already has these permissions you may override the default user with `-u 999:999`.
