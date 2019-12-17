package hello

import (
	"fmt"
	"runtime"

	"rsc.io/quote/v3"
)

func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}

func main() string {
	sdkVersion := "1.0.0"
	runtimeVersion := "2.14.0"
	fmt.Printf("%s vAPI %s %s (%s; %s)", sdkVersion, runtimeVersion, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	return ""
}
