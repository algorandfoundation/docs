# Algod Importer

Fetch blocks one by one from the [algod REST API](https://developer.algorand.org/docs/rest-apis/algod/v2/). The node must be configured as an archival node in order to
provide old blocks.

Block data from the Algod REST API contains the block header, transactions, and a vote certificate.

# Config
```yaml
importer:
    name: algod
    config:
      - netaddr: "algod URL"
        token: "algod REST API token"
        # This is an optional section which will enable Conduit to run fast catchup on your behalf using the provided catchpoint
        catchup-config:
          catchpoint: "the catchpoint to use for fast catchup"
          admin-token: "the algod admin API token (used to run catchpoint catchup)"
```

