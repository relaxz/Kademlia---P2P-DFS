package kdmlib

import (
	"Kademlia---P2P-DFS/github.com/golang/protobuf/proto"
	"Kademlia---P2P-DFS/kdmlib/fileutils"
	pb "Kademlia---P2P-DFS/kdmlib/proto_config"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	Request     = "Request"
	Return      = "Return"
	Ping        = "Ping"
	FindContact = "FindContact"
	FindData    = "FindData"
	Store       = "Store"
)

type Network struct {
	kademlia   *Kademlia
	rt         RoutingTable
	serverConn *net.UDPConn
	mux        *sync.Mutex
	queue      map[string]chan interface{} //<-change?
	timeLimit  int
}

func InitializeNetwork(timeOutLimit int, port int, rt RoutingTable) *Network {
	network := &Network{}
	network.rt = rt
	network.mux = &sync.Mutex{}
	network.queue = make(map[string]chan interface{})
	network.timeLimit = timeOutLimit

	network.UDPConnection(port)
	return network
}
func (network *Network) UDPConnection(Port int) { //TODO: learn how to properly use channels
	ServerAddr, err := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(Port))
	CheckError(err)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	network.serverConn = ServerConn
	buf := make([]byte, 1024)
	network.Listen(buf)

}
func (network *Network) Listen(buf []byte) {
	defer network.serverConn.Close()
	for {
		n, addr, err := network.serverConn.ReadFromUDP(buf)

		container := &pb.Container{}
		err = proto.Unmarshal(buf[0:n], container)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println(addr)
		network.RequestHandler(container, addr)
	}

}
func (network *Network) RequestHandler(container *pb.Container, addr *net.UDPAddr) {

	switch container.REQUEST_ID {
	case Ping:
		fmt.Println("Received Ping_Request")
		network.ReturnPing(addr, container.MSG_ID)
		fmt.Println("Returned Ping_Request")
		break
	case FindContact:
		fmt.Println("Received FindContact_Request")
		network.ReturnContact(addr, container.MSG_ID, container.GetRequestContact().ID)
		fmt.Println("Returned FindContact_Request")
		break
	case FindData:
		fmt.Println("Received FindData_Request")
		//network.ReturnData(addr, container.MSG_ID, container.GetRequestData().KEY)
		fmt.Println("Returned FindData_Request")
		break
	case Store:
		fmt.Println("Received Store_Request")
		network.ReturnStore(addr, container.MSG_ID, container.GetRequestStore().KEY, container.GetRequestStore().VALUE)
		fmt.Println("Returned Store_Request")
		break
	default:
		fmt.Println("Something went horribly wrong! (Request)")
	}
}

//You ask something from someone!
func (network *Network) SendPing(addr *net.UDPAddr, returnChannel chan interface{}) {
	//myID := network.kademlia.nodeId
	msgID := GenerateRandID()
	Info := &pb.REQUEST_PING{ID: "asdasd"}
	Data := &pb.Container_RequestPing{RequestPing: Info}
	Container := &pb.Container{REQUEST_TYPE: Request, REQUEST_ID: Ping, MSG_ID: msgID, Attachment: Data}
	network.putInQueue(msgID, returnChannel)
	network.RequestData(Container, addr)
	go network.checkForTimeOut(msgID, network.timeLimit, returnChannel)
}
func (network *Network) SendFindContact(addr *net.UDPAddr, contactID string, returnChannel chan interface{}) {
	//msgID := GenerateRandID()
	msgID := "12345"
	Info := &pb.REQUEST_CONTACT{ID: contactID}
	Data := &pb.Container_RequestContact{RequestContact: Info}
	Container := &pb.Container{REQUEST_TYPE: Request, REQUEST_ID: FindContact, MSG_ID: msgID, Attachment: Data}
	network.putInQueue(msgID, returnChannel)
	network.RequestData(Container, addr)
	go network.checkForTimeOut(msgID, network.timeLimit, returnChannel)
}
func (network *Network) SendFindData(addr *net.UDPAddr, hash string, returnChannel chan interface{}) {
	msgID := GenerateRandID()
	fmt.Println(msgID)
	Info := &pb.REQUEST_DATA{KEY: hash}
	Data := &pb.Container_RequestData{RequestData: Info}
	Container := &pb.Container{REQUEST_TYPE: Request, REQUEST_ID: FindData, MSG_ID: msgID, Attachment: Data}
	network.putInQueue(msgID, returnChannel)
	network.RequestData(Container, addr)
	go network.checkForTimeOut(msgID, network.timeLimit, returnChannel)
}
func (network *Network) SendStoreData(addr *net.UDPAddr, KEY string, DATA []byte, returnChannel chan interface{}) {
	msgID := GenerateRandID()
	/*	STORE message must contain in addition to the message ID
		at least the data to be stored (including its length) and
		the associated key. As the transport may be UDP, the message needs
		to also contain at least the nodeID of the sender, and the reply the
		nodeID of the recipient. */
	Info := &pb.REQUEST_STORE{KEY: KEY, VALUE: DATA}
	Data := &pb.Container_RequestStore{RequestStore: Info}
	Container := &pb.Container{REQUEST_TYPE: Request, REQUEST_ID: Store, MSG_ID: msgID, Attachment: Data}
	network.putInQueue(msgID, returnChannel)
	network.RequestData(Container, addr)
	go network.checkForTimeOut(msgID, network.timeLimit, returnChannel)
}

//Someone ask something from you and you return!
func (network *Network) ReturnPing(addr *net.UDPAddr, msgID string) {
	//myID := network.kademlia.nodeId
	Info := &pb.RETURN_PING{ID: "qwe"}
	Data := &pb.Container_ReturnPing{ReturnPing: Info}
	Container := &pb.Container{REQUEST_TYPE: Return, REQUEST_ID: Ping, MSG_ID: msgID, Attachment: Data}
	network.ReturnRequestedData(Container, addr)
}
func (network *Network) ReturnContact(addr *net.UDPAddr, msgID string, contactID string) {
	closestContacts := network.rt.FindKClosest(contactID) //mux on this?
	contactListReply := []*pb.RETURN_CONTACTS_CONTACT_INFO{}
	for i := range closestContacts {
		contactReply := &pb.RETURN_CONTACTS_CONTACT_INFO{IP: closestContacts[i].Triple.Ip, PORT: closestContacts[i].Triple.Port,
			ID: closestContacts[i].Triple.Id}
		contactListReply = append(contactListReply, contactReply)
	}
	Info := &pb.RETURN_CONTACTS{ContactInfo: contactListReply}
	Data := &pb.Container_ReturnContacts{ReturnContacts: Info}
	Container := &pb.Container{REQUEST_TYPE: Return, REQUEST_ID: FindContact, MSG_ID: msgID, Attachment: Data}
	network.ReturnRequestedData(Container, addr)
}
func (network *Network) ReturnData(addr *net.UDPAddr, msgID string, DataID string, contactID string) {

	if fileUtilsKademlia.ReadFileFromOS(DataID) != nil {
		Value := fileUtilsKademlia.ReadFileFromOS(DataID)
		Info := &pb.RETURN_DATA{VALUE: Value}
		Data := &pb.Container_ReturnData{ReturnData: Info}
		Container := &pb.Container{REQUEST_TYPE: Return, REQUEST_ID: FindData, MSG_ID: msgID, Attachment: Data}
		network.ReturnRequestedData(Container, addr)
	} else {
		network.ReturnContact(addr, msgID, contactID)
	}
}
func (network *Network) ReturnStore(addr *net.UDPAddr, msgID string, key string, value []byte) { //work in progress

	if fileUtilsKademlia.ReadFileFromOS(key) != nil {
		fmt.Println("Do something? Data exists?")
	} else {
		fmt.Println("Can passed data/value be big? or just 64kb udp package?" +
			"Download file(UDP, TCP), TCP is better? more secure?, add to stored-files-list?") //<- a project in itself
	}
	//check if file already exist, if not, download and reply on the store request.
	//--------------------------------------
	/*
		Info := &pb.RETURN_STORE{VALUE: "Stored"} //i stored the data msg!
		Data := &pb.Container_ReturnStore{ReturnStore: Info}
		Container := &pb.Container{REQUEST_TYPE: Return, REQUEST_ID: Store,MSG_ID:msgID Attachment: Data}
		network.ReturnRequestedData(Container, addr)
	*/
}

//Someone returns something you previously asked for!
func PingReturned(container *pb.Container, returnedRequest chan interface{}) {
	contactID := container.GetReturnPing().ID
	returnedRequest <- contactID
}
func ContactReturned(container *pb.Container, returnedRequest chan interface{}) {
	listOfContacts := []AddressTriple{}
	for i := range container.GetReturnContacts().ContactInfo {
		listOfContacts = append(listOfContacts,
			AddressTriple{Ip: container.GetReturnContacts().ContactInfo[i].IP,
				Port: container.GetReturnContacts().ContactInfo[i].PORT,
				Id:   container.GetReturnContacts().ContactInfo[i].ID})
	}

	returnedRequest <- listOfContacts
}
func DataReturned(container *pb.Container, returnedRequest chan interface{}) {
	Value := container.GetReturnData().VALUE
	returnedRequest <- Value
}
func StoreReturned(container *pb.Container, returnedRequest chan interface{}) {}

//helper functions
func (network *Network) RequestData(container *pb.Container, addr *net.UDPAddr) {

	conn, err := net.Dial("udp", addr.IP.String()+":"+strconv.Itoa(addr.Port))
	CheckError(err)
	buf := []byte(EncodeContainer(container))
	_, err = conn.Write(buf)
	buf = make([]byte, 1024)
	i, err := conn.Read(buf)

	container = &pb.Container{}
	err = proto.Unmarshal(buf[0:i], container)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	network.ReturnHandler(container)

}
func (network *Network) ReturnHandler(container *pb.Container) {
	returnedRequest := network.takeFromQueue(container.MSG_ID)
	if returnedRequest == nil {
		fmt.Println("Timeout")
		return
	}
	switch container.REQUEST_ID {
	case Ping:
		fmt.Println("Ping Returned")
		fmt.Println(container.GetReturnPing())
		PingReturned(container, returnedRequest)
		break
	case FindContact:
		fmt.Println("Contact Returned")
		ContactReturned(container, returnedRequest)
		break
	case FindData:
		fmt.Println("Data Returned")
		DataReturned(container, returnedRequest)
		break
	case Store:
		fmt.Println("Store Returned")
		StoreReturned(container, returnedRequest)
		break
	default:
		fmt.Println("Something went horribly wrong! (Return)")
	}
}
func (network *Network) ReturnRequestedData(container *pb.Container, addr *net.UDPAddr) {

	buf := []byte(EncodeContainer(container))
	_, err := network.serverConn.WriteToUDP(buf, addr)

	if err != nil {
		fmt.Println(EncodeContainer(container), err)
	}
}
func (network *Network) putInQueue(msgID string, returnChannel chan interface{}) {
	network.mux.Lock()
	network.queue[msgID] = returnChannel
	network.mux.Unlock()
}
func (network *Network) takeFromQueue(msgID string) (returnedRequest chan interface{}) {
	network.mux.Lock()
	returnedRequest = network.queue[msgID]
	network.mux.Unlock()
	return returnedRequest
}
func (network *Network) checkForTimeOut(msgID string, timeLimit int, returnChannel chan interface{}) {
	time.Sleep(time.Duration(timeLimit) * time.Second)
	network.mux.Lock()
	network.queue[msgID] = nil
	network.mux.Unlock()
	returnChannel <- false
}
func EncodeContainer(pack *pb.Container) []byte {
	data, err := proto.Marshal(pack)
	if err != nil {
		log.Fatal("marshalling error: ", err)
	}
	return data
}
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

/*	To test this code duplicate this project and insert the following
	before network.Listen(buf) inside UDPConnection in one of the projects
	//---------------------test------------------------
	contactAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:12000") //<-- try this address when testing!
	CheckError(err)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		answerChannel := make(chan interface{})
		network.SendPing(contactAddr, answerChannel)
	}
	//go network.SendFindContact(contactAddr, "00000", answerChannel)
	//go network.SendFindData(contactAddr, GenerateRandID(), answerChannel)
	//go network.SendStoreData(contactAddr, GenerateRandID(), []byte("hello"), answerChannel)

	//-------------------------------------------------
*/
