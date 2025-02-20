# S038

The S038 analyzer reports cases of Schemas which has `Default` value declared with incompatible `Type`.

## Flagged Code

```go
&schema.Schema{
    Type: schema.TypeInt,
    Default: true,
}
```

## Passing Code

```go
&schema.Schema{
    Type: schema.TypeBool,
    Default: true,
}
```

## Ignoring Reports

Singular reports can be ignored by adding the a `//lintignore:S038` Go code comment at the end of the offending line or on the line immediately proceding, e.g.

```go
//lintignore:S038
&schema.Schema{
    Type: schema.TypeInt,
    Default: true,
}
```
