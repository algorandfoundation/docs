title: API V2 Migration Guide

In June 2020, Algorand introduced the V2 API for `algod` and deprecated the V1 API. Both APIs remained functional until December 2022 to allow developers time to transition their application code to the fully supported V2 endpoints. Simultaneously, Algorand introduced `algorand-indexer` with only a V2 API. Use this guide to update your preferred SDK for V2 client support and then transition your application to use V2 clients for `algod` and `algorand-indexer` where applicable.

!!! warning
    The V1 API for `algod` was removed from source in December 2022. Please use V2 going forward.

!!! information 
    The `kmd` API remains unchanged at V1. There are no changes required for application code using the V1 `kmd` client. Reference [kmd client instantiations](#kmd-instantiations) below.

