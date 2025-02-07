package udp

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/whatap/golib/io"
	"github.com/whatap/golib/util/stringutil"
	"github.com/whatap/golib/util/urlutil"
)

type UdpTxEndPack struct {
	AbstractPack

	// Pack
	Host    string
	Uri     string
	Mtid    int64
	Mdepth  int32
	Mcaller int64

	McallerTxid    int64
	McallerPcode   int64
	McallerSpec    string
	McallerUrl     string
	McallerPoidKey string

	Status int32

	McallerStepId int64
	XTraceId      string

	// Processing data
	ServiceURL     *urlutil.URL
	McallerUrlHash int32
}

func NewUdpTxEndPack() *UdpTxEndPack {
	p := new(UdpTxEndPack)
	p.Ver = UDP_PACK_VERSION
	p.AbstractPack.Flush = true
	return p
}

func NewUdpTxEndPackVer(ver int32) *UdpTxEndPack {
	p := new(UdpTxEndPack)
	p.Ver = ver
	p.AbstractPack.Flush = true
	return p
}

func (this *UdpTxEndPack) GetPackType() uint8 {
	return TX_END
}

func (this *UdpTxEndPack) ToString() string {
	return fmt.Sprint(this.AbstractPack.ToString(), ",host=", this.Host, ",uri=", this.Uri, ",elapsed=", this.Elapsed)
}

func (this *UdpTxEndPack) Clear() {
	this.AbstractPack.Clear()
	this.AbstractPack.Flush = true

	this.Host = ""
	this.Uri = ""
	this.Mtid = 0
	this.Mdepth = 0
	this.Mcaller = 0

	this.McallerTxid = 0
	this.McallerPcode = 0
	this.McallerSpec = ""
	this.McallerUrl = ""
	this.McallerPoidKey = ""

	this.Status = 0

	this.McallerStepId = 0
	this.XTraceId = ""

	// Processing data
	this.ServiceURL = nil
	this.McallerUrlHash = 0
}

func (this *UdpTxEndPack) SetMcallerUrlHash(v int32) {
	this.McallerUrlHash = v
	this.McallerUrl = fmt.Sprintf("%d", v)
}
func (this *UdpTxEndPack) Write(dout *io.DataOutputX) {
	this.AbstractPack.Write(dout)
	if this.Ver > 50000 {
		// Golang
		dout.WriteTextShortLength(this.Host)
		dout.WriteTextShortLength(this.Uri)
		dout.WriteTextShortLength(stringutil.ParseStringZeroToEmpty(int64(this.Mtid)))
		dout.WriteTextShortLength(stringutil.ParseStringZeroToEmpty(int64(this.Mdepth)))
		dout.WriteTextShortLength(stringutil.ParseStringZeroToEmpty(int64(this.McallerTxid)))
		dout.WriteTextShortLength(stringutil.ParseStringZeroToEmpty(int64(this.McallerPcode)))
		dout.WriteTextShortLength(this.McallerSpec)
		dout.WriteTextShortLength(this.McallerUrl)
		dout.WriteTextShortLength(this.McallerPoidKey)
		if this.Ver >= 50100 {
			dout.WriteTextShortLength(stringutil.ParseStringZeroToEmpty(int64(this.Status)))
		}
		if this.Ver >= 50101 {
			dout.WriteTextShortLength(stringutil.ParseStringZeroToEmpty(this.McallerStepId))
			dout.WriteTextShortLength(this.XTraceId)
		}
	} else if this.Ver > 40000 {
		// Batch
	} else if this.Ver > 30000 {
		// Dotnet
		dout.WriteTextShortLength(this.Host)
		dout.WriteTextShortLength(this.Uri)
		dout.WriteTextShortLength(stringutil.ParseStringZeroToEmpty(int64(this.Mtid)))
		dout.WriteTextShortLength(stringutil.ParseStringZeroToEmpty(int64(this.Mdepth)))
		dout.WriteTextShortLength(stringutil.ParseStringZeroToEmpty(int64(this.Mcaller)))
	} else if this.Ver > 20000 {
		// Python
		dout.WriteTextShortLength(this.Host)
		dout.WriteTextShortLength(this.Uri)
		dout.WriteTextShortLength(stringutil.ParseStringZeroToEmpty(int64(this.Mtid)))
		dout.WriteTextShortLength(stringutil.ParseStringZeroToEmpty(int64(this.Mdepth)))
		dout.WriteTextShortLength(stringutil.ParseStringZeroToEmpty(int64(this.McallerTxid)))
		dout.WriteTextShortLength(stringutil.ParseStringZeroToEmpty(int64(this.McallerPcode)))
		dout.WriteTextShortLength(this.McallerSpec)
		dout.WriteTextShortLength(this.McallerUrl)
		dout.WriteTextShortLength(this.McallerPoidKey)
	} else {
		// PHP
		if this.Ver >= 10102 {
			dout.WriteTextShortLength(this.Host)
			dout.WriteTextShortLength(this.Uri)
			dout.WriteTextShortLength(stringutil.ParseStringZeroToEmpty(int64(this.Mtid)))
			dout.WriteTextShortLength(stringutil.ParseStringZeroToEmpty(int64(this.Mdepth)))
			dout.WriteTextShortLength(stringutil.ParseStringZeroToEmpty(int64(this.McallerTxid)))
			dout.WriteTextShortLength(stringutil.ParseStringZeroToEmpty(int64(this.McallerPcode)))
			dout.WriteTextShortLength(this.McallerSpec)
			dout.WriteTextShortLength(this.McallerUrl)
			dout.WriteTextShortLength(this.McallerPoidKey)
		}

		if this.Ver >= 10107 {
			dout.WriteTextShortLength(stringutil.ParseStringZeroToEmpty(int64(this.Status)))
		}

		if this.Ver >= 10108 {
			dout.WriteTextShortLength(stringutil.ParseStringZeroToEmpty(this.McallerStepId))
			dout.WriteTextShortLength(this.XTraceId)
		}
	}
}

func (this *UdpTxEndPack) Read(din *io.DataInputX) {
	this.AbstractPack.Read(din)
	if this.Ver > 50000 {
		// Golang
		this.Host = din.ReadTextShortLength()
		this.Uri = din.ReadTextShortLength()
		this.Mtid = stringutil.ParseInt64(din.ReadTextShortLength())
		this.Mdepth = stringutil.ParseInt32(din.ReadTextShortLength())
		this.McallerTxid = stringutil.ParseInt64(din.ReadTextShortLength())
		this.McallerPcode = stringutil.ParseInt64(din.ReadTextShortLength())
		this.McallerSpec = din.ReadTextShortLength()
		this.McallerUrl = din.ReadTextShortLength()
		this.McallerPoidKey = din.ReadTextShortLength()

		if this.Ver >= 50100 {
			this.Status = stringutil.ParseInt32(din.ReadTextShortLength())
		}
		if this.Ver >= 50101 {
			this.McallerStepId = stringutil.ParseInt64(din.ReadTextShortLength())
			this.XTraceId = din.ReadTextShortLength()
		}

	} else if this.Ver > 40000 {
		// Batch
	} else if this.Ver > 30000 {
		// Dotnet
		this.Host = din.ReadTextShortLength()
		this.Uri = din.ReadTextShortLength()
		this.Mtid = stringutil.ParseInt64(din.ReadTextShortLength())
		this.Mdepth = stringutil.ParseInt32(din.ReadTextShortLength())
		this.Mcaller = stringutil.ParseInt64(din.ReadTextShortLength())
	} else if this.Ver > 20000 {
		// Python
		this.Host = din.ReadTextShortLength()
		this.Uri = din.ReadTextShortLength()
		this.Mtid = stringutil.ParseInt64(din.ReadTextShortLength())
		this.Mdepth = stringutil.ParseInt32(din.ReadTextShortLength())
		this.McallerTxid = stringutil.ParseInt64(din.ReadTextShortLength())
		this.McallerPcode = stringutil.ParseInt64(din.ReadTextShortLength())
		this.McallerSpec = din.ReadTextShortLength()
		this.McallerUrl = din.ReadTextShortLength()
		this.McallerPoidKey = din.ReadTextShortLength()
	} else {
		// PHP
		if this.Ver >= 10102 {
			this.Host = din.ReadTextShortLength()
			this.Uri = din.ReadTextShortLength()
			this.Mtid = stringutil.ParseInt64(din.ReadTextShortLength())
			this.Mdepth = stringutil.ParseInt32(din.ReadTextShortLength())
			this.McallerTxid = stringutil.ParseInt64(din.ReadTextShortLength())
			this.McallerPcode = stringutil.ParseInt64(din.ReadTextShortLength())
			this.McallerSpec = din.ReadTextShortLength()
			this.McallerUrl = din.ReadTextShortLength()
			this.McallerPoidKey = din.ReadTextShortLength()
		}

		if this.Ver >= 10107 {
			// reponse code
			this.Status = stringutil.ParseInt32(din.ReadTextShortLength())
		}

		if this.Ver >= 10108 {
			this.McallerStepId = stringutil.ParseInt64(din.ReadTextShortLength())
			this.XTraceId = din.ReadTextShortLength()
		}
	}
}
func (this *UdpTxEndPack) Process() {
	if this.Host != "" && this.Uri != "" {
		if strings.HasPrefix(this.Uri, "/") {
			this.ServiceURL = urlutil.NewURL(this.Host + this.Uri)
		} else {
			this.ServiceURL = urlutil.NewURL(this.Host + "/" + this.Uri)
		}
	}
	if this.Ver > 50000 {
		// Golang
		if ret, err := strconv.ParseInt(this.McallerUrl, 10, 32); err == nil {
			this.McallerUrlHash = int32(ret)
		}
	} else if this.Ver > 40000 {
		// Batch
	} else if this.Ver > 30000 {
		// Dotnet
	} else if this.Ver > 20000 {
		// Python
		if ret, err := strconv.ParseInt(this.McallerUrl, 10, 32); err == nil {
			this.McallerUrlHash = int32(ret)
		}
	} else {
		// PHP
		if this.Ver >= 10102 {
			if ret, err := strconv.ParseInt(this.McallerUrl, 10, 32); err == nil {
				this.McallerUrlHash = int32(ret)
			}
		}
	}
}
