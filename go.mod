module github.com/cristaloleg/gql2yml

go 1.16

require (
	github.com/vektah/gqlparser/v2 v2.2.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

// Fork with a one-liner change - do not marshal schema file per GraphQL definition
replace github.com/vektah/gqlparser/v2 v2.2.0 => github.com/cristaloleg/gqlparser/v2 v2.2.1-0.20210819155019-33200f332744
