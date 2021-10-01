# Gqlutil  [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/knesklab/util/blob/master/LICENSE)

A small sets of useful function for [gqlgen.com](http://gqlgen.com) grqphql module.

## Installation
```
go get github.com/ubgo/gqlutil
```

### GetFieldsRequested(ctx context.Context) []string
get the fields requested in a graphql query
Useful for constructing SQL queries from the GraphQL query
```go
gqlutil.GetFieldsRequested(ctx)
```

### GetParentFieldsRequested(ctx context.Context) []string
this get the same as above, but for use within the field resolver. So in the field resolver if we didn't get a field we can grab what was requested above and load it.
- Useful for lazy resolving data
- NB this will panic if there is no parent.
```go
gqlutil.GetParentFieldsRequested(ctx)
```

### GetPreloads(ctx context.Context) []string
get the fields requested as json keys
```go
gqlutil.GetPreloads(ctx)
```

## Contribute

If you would like to contribute to the project, please fork it and send us a pull request.  Please add tests
for any new features or bug fixes.

## Stay in touch

* Author - [Aman Khanakia](https://twitter.com/mrkhanakia)
* Website - [https://khanakia.com](https://khanakia.com/)

## License

goutil is [MIT licensed](LICENSE).
