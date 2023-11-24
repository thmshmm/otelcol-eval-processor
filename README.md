# Demo OTel Collector Processor

A simple demo to bootstrap a custom collector processor. In this example, the processor calls an external 'evaluation' API and sets the `valid` boolean attribute depending of the result of the evaluation. If the mock API is unavailable or returns an error, the `evalError` attribute is set.

## Getting started

```
docker compose up
```

Within the collector container, append log lines to the `/tmp/simple.log` file.

Expected log line format:

```
echo "2023-06-19 05:20:50 ERROR This is a test error message" >> /tmp/simple.log
```
