# COSMOS

## Tests

### Test cover

First run the command (THE_PACKAGE_NAME is the name of the package ):

```bash
    go test THE_PACKAGE_NAME -coverprofile=cosmos_test_coverage
```

And then you can see the results on the browser using the following comand:

```bash
    go tool cover -html=cosmos_test_coverage
```