package event

import (
	"hades-ebpf/user/decoder"

	manager "github.com/ehids/ebpfmanager"
)

var DefaultDoInitModule = &DoInitModule{}

var _ decoder.Event = (*DoInitModule)(nil)

type DoInitModule struct {
	decoder.BasicEvent `json:"-"`
	Exe                string `json:"-"`
	Modname            string `json:"modname"`
	Pidtree            string `json:"pidt_ree"`
	Cwd                string `json:"cwd"`
	PrivEscalation     uint8  `json:"priv_esca"`
}

func (DoInitModule) ID() uint32 {
	return 1026
}

func (DoInitModule) String() string {
	return "do_init_module"
}

func (d *DoInitModule) GetExe() string {
	return d.Exe
}

func (d *DoInitModule) Parse() (err error) {
	if d.Modname, err = decoder.DefaultDecoder.DecodeString(); err != nil {
		return
	}
	if d.Exe, err = decoder.DefaultDecoder.DecodeString(); err != nil {
		return
	}
	if d.Pidtree, err = decoder.DefaultDecoder.DecodePidTree(&d.PrivEscalation); err != nil {
		return
	}
	if d.Cwd, err = decoder.DefaultDecoder.DecodeString(); err != nil {
		return
	}
	return
}

func (d *DoInitModule) GetProbe() []*manager.Probe {
	return []*manager.Probe{
		{
			UID:              "KprobeDoInitModule",
			Section:          "kprobe/do_init_module",
			EbpfFuncName:     "kprobe_do_init_module",
			AttachToFuncName: "do_init_module",
		},
	}
}

func init() {
	decoder.Regist(DefaultDoInitModule)
}
