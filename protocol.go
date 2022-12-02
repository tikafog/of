//go:generate go run github.com/abice/go-enum -f protocol.go

package of

// Protocol version
// ENUM(Answer,Gossip,DataSwap,Message,MessageV2,Instruct,Adapter,Max)
type Protocol int
