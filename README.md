# godash

[![Go Reference](https://pkg.go.dev/badge/github.com/taciogt/envtags.svg)](https://pkg.go.dev/github.com/taciogt/godash)
![Version](https://img.shields.io/github/v/release/taciogt/godash)
![Go version](https://img.shields.io/github/go-mod/go-version/taciogt/godash)
[![Tests](https://github.com/taciogt/godash/actions/workflows/tests.yaml/badge.svg)](https://github.com/taciogt/godash/actions/workflows/tests.yaml)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/86a0ff7430d54e0fa614195978c09213)](https://app.codacy.com/gh/taciogt/godash/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/86a0ff7430d54e0fa614195978c09213)](https://app.codacy.com/gh/taciogt/godash/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_coverage)

_godash_ is a Go package that aims to make it easier to deal with common operations on everyday types like slices and maps.
It is inspired by lodash and similar libraries and exposes not only pure functions to interact with common data structures but also custom data types to make
these interactions easier.

## Requirements

Go >= 1.21

## Install

```shell
go get github.com/taciogt/godash
```

## Documentation

Check the [GoDoc here](https://pkg.go.dev/github.com/taciogt/godash).

### Sets

The [Set](https://pkg.go.dev/github.com/taciogt/godash#Set) type represents a set data structure using native Golang data types. It can store any type that
implements the [comparable](https://go.dev/ref/spec#Type_constraints) interface. This custom type provides the following methods:

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

The [Slice](https://pkg.go.dev/github.com/taciogt/godash#Slice) is a custom type over the standard slice to allow for chainable calls.

#### Slice Methods

The `Slice` type provides the following methods:

* [At(index int)](https://pkg.go.dev/github.com/taciogt/godash#Slice.At) - Returns the element at the specified index
* [Every(predicate func(T, int) bool)](https://pkg.go.dev/github.com/taciogt/godash#Slice.Every) - Tests if all elements pass the predicate function
* [Fill(value T)](https://pkg.go.dev/github.com/taciogt/godash#Slice.Fill) - Fills all elements with the specified value
* [Filter(predicate func(T, int) bool)](https://pkg.go.dev/github.com/taciogt/godash#Slice.Filter) - Returns elements that pass the predicate function
* [Find(predicate func(T, int) bool)](https://pkg.go.dev/github.com/taciogt/godash#Slice.Find) - Returns the first element that satisfies the predicate
* [FindIndex(predicate func(T, int) bool)](https://pkg.go.dev/github.com/taciogt/godash#Slice.FindIndex) - Returns the index of the first matching element
* [FindLast(predicate func(T, int) bool)](https://pkg.go.dev/github.com/taciogt/godash#Slice.FindLast) - Returns the last element that satisfies the predicate
* [FindLastIndex(predicate func(T, int) bool)](https://pkg.go.dev/github.com/taciogt/godash#Slice.FindLastIndex) - Returns the index of the last matching
  element
* [ForEach(fn func(T, int))](https://pkg.go.dev/github.com/taciogt/godash#Slice.ForEach) - Executes a function for each element
* [Map(fn func(T, int) U)](https://pkg.go.dev/github.com/taciogt/godash#Map) - Creates a new slice with the results of the function
* [Reduce(fn func(U, T, int) U, initial U)](https://pkg.go.dev/github.com/taciogt/godash#Reduce) - Reduces the slice to a single value
* [Pop()](https://pkg.go.dev/github.com/taciogt/godash#Slice.Pop) - Removes and returns the last element
* [Push(elements ...T)](https://pkg.go.dev/github.com/taciogt/godash#Slice.Push) - Adds elements to the end of the slice

#### ComparableSlice Methods

The [ComparableSlice](https://pkg.go.dev/github.com/taciogt/godash#ComparableSlice) type extends `Slice` with additional functionality for comparable types:

* [Includes(value T)](https://pkg.go.dev/github.com/taciogt/godash#ComparableSlice.Includes) - Determines if the slice includes a certain value
* [IndexOf(value T)](https://pkg.go.dev/github.com/taciogt/godash#ComparableSlice.IndexOf) - Returns the first index at which a given element can be found

## References

[Contributing](CONTRIBUTING.md)