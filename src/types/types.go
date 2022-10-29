// package types

// type BenInfo struct {
// 	Length int64
// 	Name []byte
//  PieceLength int64
// }

// type BenCoding struct {
// 	Announce []byte
// 	AnnounceList []interface{}
// 	Comment []byte
// 	CreatedBy []byte
// 	CreationDate []byte
// 	Info BenInfo

// 	path
// 	files
// 	pieces

// }

// "announce": []byte("udp://tracker.publicbt.com:80/announce"),
// 	"announce-list": []interface{}{
// 		[]interface{}{[]byte("udp://tracker.publicbt.com:80/announce")},
// 		[]interface{}{[]byte("udp://tracker.openbittorrent.com:80/announce")},
// 	},
// 	"comment": []byte("Debian CD from cdimage.debian.org"),
// 	"info": map[string]interface{}{
// 		"name":         []byte("debian-8.8.0-arm64-netinst.iso"),
// 		"length":       170917888,
// 		"piece length": 262144,
// 	},
