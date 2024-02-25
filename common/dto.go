package common

type Response[T any] struct{
  Status int          `json:"status"`
  Data T              `json:"data"`
  Message string      `json:"message"`
}
