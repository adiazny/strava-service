# strava-service
new strava uploader


# Project Layout

## Foundation
The foudation packages are used as a library packages to be used in the business packages. Think of them as our home grown standard library. No logging is done in any of the foundation packages.

### Web
The web package is a mini web framwork built to customize the handler functions for overwritting / extending the current mux handle implemenation to make it our own. 

### Middleware