# XS003

The XS003 analyzer reports cases of schemas where a non-computed nested block property of `TypeList` that only contains optional-only and non-default child properties, which has no `AtLeastOneOf`/`ExactlyOneOf` constraint.

Note that if the schema is not a composite literal, this rule will be skipped so as to make this rule "complete" (less estimation).

## Flagged Code

```go
map[string]*schema.Schema{
  "a": {
    Type: schema.TypeList,
    Optional: true,
    Elem: &schema.Resource{
      Schema: map[string]*schema.Schema{
        "foo": {
          Optional: true,
        },
        "bar": {
          Optional: true,
        },
      },
    },
  },
}
```

## Passing Code

```go
map[string]*schema.Schema{
  "a": {
    Type: schema.TypeList,
    Optional: true,
    Elem: &schema.Resource{
      Schema: map[string]*schema.Schema{
        "foo": {
          Optional: true,
          AtLeastOneOf: []string{"foo", "bar"},
        },
        "bar": {
          Optional: true,
		  AtLeastOneOf: []string{"foo", "bar"},
},
      },
    },
  },
}
```

## Ignoring Reports

Singular reports can be ignored by adding the a `//lintignore:XS003` Go code comment at the end of the offending line or on the line immediately proceding, e.g.

```go
//lintignore:XS003
map[string]*schema.Schema{
  "a": {
    Type: schema.TypeList,
    Optional: true,
    Elem: &schema.Resource{
      Schema: map[string]*schema.Schema{
        "foo": {
          Optional: true,
        },
        "bar": {
          Optional: true,
        },
      },
    },
  },
}
```