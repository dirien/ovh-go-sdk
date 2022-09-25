# ovh-go-sdk

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/github.com/dirien/ovh-go-sdk?tab=doc)
![OVH](https://img.shields.io/badge/ovh-123F6D?style=for-the-badge&logo=ovh&logoColor=white)
![Go](https://img.shields.io/badge/go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

### Motivation

An opinionated OVHCloud Go SDK based on github.com/ovh/go-ovh

### Why writing an own SDK? 🤔

The official go-ovh is a lightweight Go wrapper around the whole OVH APIs. I wanted to create an opinionated view on a
subset of features.

It is build, that it resembles the SDKs from other cloud-providers like DigitalOcean.

### Authentication

When using the `NewOVHDefaultClient(region, serviceName string)` the client will look for following environment
variables:

- ``OVH_ENDPOINT``,
- ``OVH_APPLICATION_KEY``,
- ``OVH_APPLICATION_SECRET``
- ``OVH_CONSUMER_KEY``

If either of these parameter is not provided, it will look for a configuration file of the form:

```ini
[default]
; general configuration: default endpoint
endpoint=ovh-eu

[ovh-eu]
; configuration specific to 'ovh-eu' endpoint
application_key=my_app_key
application_secret=my_application_secret
consumer_key=my_consumer_key
```

Depending on the API you want to use, you may set the ``endpoint`` to:

* ``ovh-eu`` for OVH Europe API
* ``ovh-us`` for OVH US API
* ``ovh-ca`` for OVH Canada API
* ``soyoustart-eu`` for So you Start Europe API
* ``soyoustart-ca`` for So you Start Canada API
* ``kimsufi-eu`` for Kimsufi Europe API
* ``kimsufi-ca`` for Kimsufi Canada API
* Or any arbitrary URL to use in a test for example

The client will successively attempt to locate this configuration file in

1. Current working directory: ``./ovh.conf``
2. Current user's home directory ``~/.ovh.conf``
3. System wide configuration ``/etc/ovh.conf``

When using `NewOVHClient(endpoint, appKey, appSecret, consumerKey, region, serviceName string)` you have to provide all
the values.

### Contributing 🤝

#### Contributing via GitHub

Feel free to join.

#### License

Apache License, Version 2.0
