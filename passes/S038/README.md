# S038

The S038 analyzer reports cases of Schemas which include `ConflictsWith` and have invalid schema attribute references.

NOTE: This pass only works with Terraform resources that are fully defined in a single function.

## Flagged Code

```go
// Attribute reference in multi nested block is not supported
&schema.Resource{
    Schema: map[string]*schema.Schema{
        "x": {
            Type: schema.TypeList,
            Elem: &schema.Resource{
                Schema: map[string]*schema.Schema{
                    "foo": {
                        AtLeastOneOf: []string{"x.0.bar"},
                    },
                    "bar": {
                        AtLeastOneOf: []string{"x.0.foo"},
                    },
                },
            },
        },
    },
}
```

or

```go
// Non-existed attribute reference
&schema.Resource{
    Schema: map[string]*schema.Schema{
        "x": {
            Type: schema.TypeList,
            MaxItems: 1,
            Elem: &schema.Resource{
                Schema: map[string]*schema.Schema{
                    "foo": {
                        AtLeastOneOf: []string{"x.1.bar"},
                    },
                    "bar": {
                        AtLeastOneOf: []string{"x.1.foo"},
                    },
                },
            },
        },
    },
}
```

## Passing Code

```go
&schema.Resource{
    Schema: map[string]*schema.Schema{
        "x": {
            Type: schema.TypeList,
            MaxItems: 1,
            Elem: &schema.Resource{
                Schema: map[string]*schema.Schema{
                    "foo": {
                        AtLeastOneOf: []string{"x.0.bar"},
                    },
                    "bar": {
                        AtLeastOneOf: []string{"x.0.foo"},
                    },
                },
            },
        },
    },
}
```

## Ignoring Reports

Singular reports can be ignored by adding the a `//lintignore:S038` Go code comment at the end of the offending line or on the line immediately proceding, e.g.

```go
//lintignore:S038
&schema.Resource{
    // ...
}
```
