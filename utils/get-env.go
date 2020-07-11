package utils

import "os"

// Getenv is a safe function for read enviroment variables
func Getenv(name string) string {
  v := os.Getenv(name)
  if v == "" {
    panic("missing required environment variable " + name)
  }
  return v
}