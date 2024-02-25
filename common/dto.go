package common

type Response[T any] struct{
  Status int
  Data T
  Message string
}
