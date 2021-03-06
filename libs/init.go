package libs

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"strings"
	"os"
	"log"
	"sync"
	"net"
)


func Init()  {
	log.SetFlags(log.Lshortfile)
	log.SetPrefix(SERVER_TAG)

	LoadArgsModule()
	LoadConfigurationModule()
	LoadLoggerModule()

	initGlobal()

	RunServerApi()


	LoadEntryModule()

}

//global
var (
	GlobalAllocates map[string]*Allocate
	AllocateMutex *sync.Mutex
	RelayPortPool *PortPool
	ServerAddress *net.UDPAddr
)
func initGlobal()  {
	RelayPortPool = NewPortPool(*min_port,*max_port)
	GlobalAllocates = make(map[string]*Allocate)
	AllocateMutex = new(sync.Mutex)
	ServerAddress = new(net.UDPAddr)
	ServerAddress.Port = *server_port
	ServerAddress.IP = getRelayAddress()
}



var (
	App               = kingpin.New("gostun", APP_NAME)
	config            = App.Flag("config", "Configuration file location").PlaceHolder(strings.Join(config_path_array,",")).Short('c').String()
	server_port		  = App.Flag("port","Server port").Default("3478").Short('p').Int()
	external_ip       = App.Flag("external_ip","TURN Server public/private address mapping, if the server is behind NAT.").Short('x').String()
	relay_ip 		  = App.Flag("relay_ip","Relay endpoint ip ").Short('r').String()
	min_port 		  = App.Flag("min_port","Lower bound of the UDP port range for relay endpoints allocation.").Default("49152").Int()
	max_port 		  = App.Flag("max_port","Upper bound of the UDP port range for relay endpoints allocation.").Default("65535").Int()
	secret			  = App.Flag("secret","Secret for generating password").Default("passwordkey").Short('s').String()
)


func LoadArgsModule() {
	App.Version(APP_VERSION)
	App.HelpFlag.Short('h')
	App.VersionFlag.Short('v')
	App.Parse(os.Args[1:])
}