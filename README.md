# godash

[![Go Reference](https://pkg.go.dev/badge/github.com/taciogt/envtags.svg)](https://pkg.go.dev/github.com/taciogt/godash)
![Version](https://img.shields.io/github/v/release/taciogt/godash)
![Go version](https://img.shields.io/github/go-mod/go-version/taciogt/godash)

[![Tests](https://github.com/taciogt/godash/actions/workflows/tests.yaml/badge.svg)](https://github.com/taciogt/godash/actions/workflows/tests.yaml)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/86a0ff7430d54e0fa614195978c09213)](https://app.codacy.com/gh/taciogt/godash/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/86a0ff7430d54e0fa614195978c09213)](https://app.codacy.com/gh/taciogt/godash/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_coverage)

_godash_ is a Go package that aims to make it easier to deal with common operations on everyday types like slices and maps. It is inspired by lodash and similar libraries.

## Requirements

Go >= 1.21

## Install

```shell
go get github.com/taciogt/godash
```

## Documentation

Check the [GoDoc here](https://pkg.go.dev/github.com/taciogt/godash).

### Sets

The [Set](https://pkg.go.dev/github.com/taciogt/godash#Set) type represents a set data structure using native Golang data types. It can store any type that implements the [comparable](https://go.dev/ref/spec#Type_constraints) interface. This custom type provides the following methods:
* [Add(element T)](https://pkg.go.dev/github.com/taciogt/godash#Set.Add)
* [Delete(element T)](https://pkg.go.dev/github.com/taciogt/godash#Set.Delete)
* [Difference(s Set[T])](https://pkg.go.dev/github.com/taciogt/godash#Set.Difference)
* [Has(element T)](https://pkg.go.dev/github.com/taciogt/godash#Set.Has)
* [Intersection(s Set[T])](https://pkg.go.dev/github.com/taciogt/godash#Set.Intersection)
* [Size()](https://pkg.go.dev/github.com/taciogt/godash#Set.Size)
* [String()](https://pkg.go.dev/github.com/taciogt/godash#Set.String)
* [Union(s Set[T])](https://pkg.go.dev/github.com/taciogt/godash#Set.Union)
* [Values()](https://pkg.go.dev/github.com/taciogt/godash#Set.Values)


### Slices

The [Slices](https://pkg.go.dev/github.com/taciogt/godash#Slice) type are a custom type over the standard slice to allow for chainable calls. The available methods for slices are:
* [At()](https://pkg.go.dev/github.com/taciogt/godash#Slice.At)
* [Every()](https://pkg.go.dev/github.com/taciogt/godash#Slice.Every)
* [Fill()](https://pkg.go.dev/github.com/taciogt/godash#Slice.Fill)
* [Filter()](https://pkg.go.dev/github.com/taciogt/godash#Slice.Filter)
* [Find()](https://pkg.go.dev/github.com/taciogt/godash#Slice.Find)
* [FindIndex()](https://pkg.go.dev/github.com/taciogt/godash#Slice.FindIndex)
* [FindLast()](https://pkg.go.dev/github.com/taciogt/godash#Slice.FindLast)
* [FindLastIndex()](https://pkg.go.dev/github.com/taciogt/godash#Slice.FindLastIndex)
* [ForEach()](https://pkg.go.dev/github.com/taciogt/godash#Slice.ForEach)
* [Map()](https://pkg.go.dev/github.com/taciogt/godash#Map)
* [Reduce()](https://pkg.go.dev/github.com/taciogt/godash#Reduce)

Some extra functionality is supported by the more specific [ComparableSlice](https://pkg.go.dev/github.com/taciogt/godash#ComparableSlice) type: 
* [Some()](https://pkg.go.dev/github.com/taciogt/godash#CoparableSlice.Some)

## References

[Contributing](CONTRIBUTING.md)