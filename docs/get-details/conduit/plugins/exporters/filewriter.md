# File Export Plugin

Write block data to files. This plugin works with the file rerader plugin to create a simple file-based pipeine.

## Configuration
```yml @sample.yaml
  name: "file_writer"
  config:
    # BlocksDir is the path to a directory where block data should be stored.
    # The directory is created if it doesn't exist. If no directory is provided
    # blocks are written to the Conduit data directory.
    #block-dir: "/path/to/block/files"

    # FilenamePattern is the format used to write block files. It uses go
    # string formatting and should accept one number for the round.
    # If the file has a '.gz' extension, blocks will be gzipped.
    # Default: "%[1]d_block.json"
    filename-pattern: "%[1]d_block.json"

    # DropCertificate is used to remove the vote certificate from the block data before writing files.
    drop-certificate: true
```
