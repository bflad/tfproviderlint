# v0.4.0

FEATURES

* **New Check:** `S007`: check for `Schema` with both `Required` and `ConflictsWith` configured
* **New Check:** `S008`: check for `Schema` of `TypeList` or `TypeSet` with `Default` configured
* **New Check:** `S009`: check for `Schema` of `TypeList` or `TypeSet` with `ValidateFunc` configured
* **New Check:** `S010`: check for `Schema` of only `Computed` enabled with `ValidateFunc` configured
* **New Check:** `S011`: check for `Schema` of only `Computed` enabled with `DiffSuppressFunc` configured
* **New Check:** `S012`: check for `Schema` that `Type` is configured
* **New Check:** `S013`: check for `map[string]*Schema` that one of `Computed`, `Optional`, or `Required` is configured
* **New Check:** `S014`: check for `Schema` within `Elem` that `Computed`, `Optional`, or `Required` are not configured
* **New Check:** `S015`: check for `map[string]*Schema` that attribute names are valid
* **New Check:** `S016`: check for `Schema` that `Set` is only configured for `TypeSet`
* **New Check:** `S017`: check for `Schema` that `MaxItems` and `MinItems` are only configured for `TypeList`, `TypeMap`, or `TypeSet`

BUG FIXES

* passes/AT001: Ignore file names beginning with `data_source_` [GH-25]

# v0.3.0

ENHANCEMENTS

* Support `map[string]*schema.Schema` in schema checks [GH-24]

# v0.2.0

FEATURES:

* **New Check** `R004` [GH-21]

# v0.1.0

* Initial release with checks:
  * `AT001`, `AT002`, `AT003`, `AT004`
  * `R001`, `R002`, `R003`
  * `S001`, `S002`, `S003`, `S004`, `S005`, `S006`
