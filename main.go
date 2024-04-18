package main

import (
	"fmt"

	_ "github.com/zaelani23/golib/config"
	_ "github.com/zaelani23/golib/config/conffile"
	_ "github.com/zaelani23/golib/io"
	_ "github.com/zaelani23/golib/lang/pack"
	_ "github.com/zaelani23/golib/lang/pack/udp"
	_ "github.com/zaelani23/golib/lang/ref"
	_ "github.com/zaelani23/golib/lang/service"
	_ "github.com/zaelani23/golib/lang/step"
	_ "github.com/zaelani23/golib/lang/value"
	_ "github.com/zaelani23/golib/lang/variable"
	_ "github.com/zaelani23/golib/logger"
	_ "github.com/zaelani23/golib/logger/logfile"
	_ "github.com/zaelani23/golib/logsink/zip"

	_ "github.com/zaelani23/golib/net"
	_ "github.com/zaelani23/golib/net/oneway"
	_ "github.com/zaelani23/golib/net/udp"
	_ "github.com/zaelani23/golib/util/ansi"
	_ "github.com/zaelani23/golib/util/bitutil"
	_ "github.com/zaelani23/golib/util/compare"
	_ "github.com/zaelani23/golib/util/dateutil"
	_ "github.com/zaelani23/golib/util/hash"
	_ "github.com/zaelani23/golib/util/hexa32"
	_ "github.com/zaelani23/golib/util/hll"
	_ "github.com/zaelani23/golib/util/hmap"
	_ "github.com/zaelani23/golib/util/iputil"
	_ "github.com/zaelani23/golib/util/keygen"
	_ "github.com/zaelani23/golib/util/list"
	_ "github.com/zaelani23/golib/util/mathutil"
	_ "github.com/zaelani23/golib/util/openstack"
	_ "github.com/zaelani23/golib/util/paramtext"
	_ "github.com/zaelani23/golib/util/queue"
	_ "github.com/zaelani23/golib/util/stringutil"
	_ "github.com/zaelani23/golib/util/urlutil"
	_ "github.com/zaelani23/golib/util/uuidutil"
)

func main() {
	fmt.Println("Whatap Golang common library")
}
