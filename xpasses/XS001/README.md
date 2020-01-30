# XS001

The XS001 analyzer reports cases of schemas where `Description` is not configured, which is generally useful for providers that wish to automatically generate documentation based on the schema information.

## Flagged Code

```go
map[string]*schema.Schema{
    "attribute_name": {
        Optional: true,
        Type:     schema.TypeString,
    },
}
```

## Passing Code

```go
map[string]*schema.Schema{
    "attribute_name": {
        Description: "does something useful",
        Optional:    true,
        Type:        schema.TypeString,
    },
}
```
