# gql2yml

GraphQL to YAML/JSON - convert your GraphQL schema to YAML/JSON.

```sh
# install gql2yaml
$ go install github.com/cristaloleg/gql2yml

# convert schema
$ gql2yml -schema=server.graphql -result=schema.yaml

# see the result
$ cat schema.yaml

# and also json
$ gql2yml -schema=server.graphql -result=schema.json -json

# or many files
$ gql2yml -schema=schema.graphql -schema=legacy.graphql
```

## License

[MIT License](LICENSE).
