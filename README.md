# godash

[![Go Reference](https://pkg.go.dev/badge/github.com/taciogt/envtags.svg)](https://pkg.go.dev/github.com/taciogt/godash)
![Version](https://img.shields.io/github/v/release/taciogt/godash)
![Go version](https://img.shields.io/github/go-mod/go-version/taciogt/godash)
[![Tests](https://github.com/taciogt/godash/actions/workflows/tests.yaml/badge.svg)](https://github.com/taciogt/godash/actions/workflows/tests.yaml)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/86a0ff7430d54e0fa614195978c09213)](https://app.codacy.com/gh/taciogt/godash/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/86a0ff7430d54e0fa614195978c09213)](https://app.codacy.com/gh/taciogt/godash/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_coverage)

_godash_ is a Go package that makes working with common data structures like slices and maps easier.
Inspired by lodash and similar libraries, it provides both pure functions and custom data types for
efficient data manipulation.

## Requirements

- Go >= 1.21

## Installation

```shell
go get github.com/taciogt/godash
```

## Documentation

Complete documentation is available in the [GoDoc](https://pkg.go.dev/github.com/taciogt/godash).

## Data Types

### Set

The [`Set`](https://pkg.go.dev/github.com/taciogt/godash#Set) type implements a set data structure using native Go types.
It can store any type that implements the [comparable](https://go.dev/ref/spec#Type_constraints) interface.

#### Set Methods

| Method                                                                         | Description                                |
|--------------------------------------------------------------------------------|--------------------------------------------|
| [`Add(element T)`](https://pkg.go.dev/github.com/taciogt/godash#Set.Add)       | Adds an element to the set                 |
| [`Delete(element T)`](https://pkg.go.dev/github.com/taciogt/godash#Set.Delete) | Removes an element from the set            |
| [`Has(element T)`](https://pkg.go.dev/github.com/taciogt/godash#Set.Has)       | Checks if an element exists in the set     |
| [`Size()`](https://pkg.go.dev/github.com/taciogt/godash#Set.Size)              | Returns the number of elements in the set  |
| [`Values()`](https://pkg.go.dev/github.com/taciogt/godash#Set.Values)          | Returns all elements as a slice            |
| [`String()`](https://pkg.go.dev/github.com/taciogt/godash#Set.String)          | Returns a string representation of the set |

#### Set Operations

| Method                                                                                    | Description                                                            |
|-------------------------------------------------------------------------------------------|------------------------------------------------------------------------|
| [`Union(s Set[T])`](https://pkg.go.dev/github.com/taciogt/godash#Set.Union)               | Returns a new set with elements from both sets                         |
| [`Intersection(s Set[T])`](https://pkg.go.dev/github.com/taciogt/godash#Set.Intersection) | Returns a new set with elements common to both sets                    |
| [`Difference(s Set[T])`](https://pkg.go.dev/github.com/taciogt/godash#Set.Difference)     | Returns a new set with elements in the first set but not in the second |

### Slices

The [`Slice`](https://pkg.go.dev/github.com/taciogt/godash#Slice) type extends the standard slice to enable chainable method calls.

#### Basic Slice Methods

| Method                                                                           | Description                                 |
|----------------------------------------------------------------------------------|---------------------------------------------|
| [`At(index int)`](https://pkg.go.dev/github.com/taciogt/godash#Slice.At)         | Returns the element at the specified index  |
| [`Fill(value T)`](https://pkg.go.dev/github.com/taciogt/godash#Slice.Fill)       | Fills all elements with the specified value |
| [`Pop()`](https://pkg.go.dev/github.com/taciogt/godash#Slice.Pop)                | Removes and returns the last element        |
| [`Push(elements ...T)`](https://pkg.go.dev/github.com/taciogt/godash#Slice.Push) | Adds elements to the end of the slice       |
| [`Reverse()`](https://pkg.go.dev/github.com/taciogt/godash#Slice.Reverse)        | Reverse the elements of the slice in place  |

#### Iteration and Transformation

| Method                                                                                                                                            | Description                                                          |
|---------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------|
| [`ForEach(fn func(T, int))`](https://pkg.go.dev/github.com/taciogt/godash#Slice.ForEach)                                                          | Executes a function for each element                                 |
| [`Map(fn func(T, int) U)`](https://pkg.go.dev/github.com/taciogt/godash#Map)                                                                      | Creates a new slice with the results of the function                 |
| [`Filter(predicate func(T, int) bool)`](https://pkg.go.dev/github.com/taciogt/godash#Slice.Filter)                                                | Returns elements that pass the predicate function                    |
| [`Reduce(s S, reducer func(acc TOut, curr TIn) (TOut, error), initialValue TOut)`](https://pkg.go.dev/github.com/taciogt/godash#Reduce)           | Reduces the slice from left to right, accumulating to a single value |
| [`ReduceRight(s S, reducer func(acc TOut, curr TIn) (TOut, error), initialValue TOut)`](https://pkg.go.dev/github.com/taciogt/godash#ReduceRight) | Reduces the slice from right to left, accumulating to a single value |

#### Search Methods

| Method                                                                                                           | Description                                            |
|------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------|
| [`Every(predicate func(T, int) bool)`](https://pkg.go.dev/github.com/taciogt/godash#Slice.Every)                 | Tests if all elements pass the predicate function      |
| [`Find(predicate func(T, int) bool)`](https://pkg.go.dev/github.com/taciogt/godash#Slice.Find)                   | Returns the first element that satisfies the predicate |
| [`FindIndex(predicate func(T, int) bool)`](https://pkg.go.dev/github.com/taciogt/godash#Slice.FindIndex)         | Returns the index of the first matching element        |
| [`FindLast(predicate func(T, int) bool)`](https://pkg.go.dev/github.com/taciogt/godash#Slice.FindLast)           | Returns the last element that satisfies the predicate  |
| [`FindLastIndex(predicate func(T, int) bool)`](https://pkg.go.dev/github.com/taciogt/godash#Slice.FindLastIndex) | Returns the index of the last matching element         |

### ComparableSlice

The [`ComparableSlice`](https://pkg.go.dev/github.com/taciogt/godash#ComparableSlice) type extends `Slice` with additional functionality for comparable types:

| Method                                                                                       | Description                                                   |
|----------------------------------------------------------------------------------------------|---------------------------------------------------------------|
| [`Includes(value T)`](https://pkg.go.dev/github.com/taciogt/godash#ComparableSlice.Includes) | Determines if the slice includes a certain value              |
| [`IndexOf(value T)`](https://pkg.go.dev/github.com/taciogt/godash#ComparableSlice.IndexOf)   | Returns the first index at which a given element can be found |

## Contributing

See the [Contributing Guide](CONTRIBUTING.md) for details on how to contribute to this project.