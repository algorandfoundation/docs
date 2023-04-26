# File Import Plugin

Read files from a directory and import them as blocks. This plugin works with the file exporter plugin to create a simple file-based pipeline.

## Configuration
```yml @sample.yaml
  name: file_reader
  config:
    # BlocksDir is the path to a directory where block data should be stored.
    # The directory is created if it doesn't exist. If no directory is provided
    # blocks are written to the Conduit data directory.
    #block-dir: "/path/to/directory"

    # RetryDuration controls the delay between checks when the importer has
    # caught up and is waiting for new blocks to appear.
    retry-duration: "5s"

    # RetryCount controls the number of times to check for a missing block
    # before generating an error. The retry count and retry duration should
    # be configured according the expected round time.
    retry-count: 5

    # FilenamePattern is the format used to find block files. It uses go string
    # formatting and should accept one number for the round.
    filename-pattern: "%[1]d_block.json"
```
