package cache

import (
	"github.com/bwmarrin/discordgo"
	"github.com/decendgame/bot/model"
)

func init() {
	TATUM_API_KEY = "cedbe61a-44ec-4c8c-b1bb-a481a2534f42"
	ActivePlayers = make(map[string]*model.Player)
	NumberOfHouses = 12
	serverName := "http://teste.com/tokens/"
	SkaleRPCServer = "https://hackathon.skalenodes.com/v1/hoarse-well-made-theemim"
	SkaleNetworkID = 31949730

	GodMasterBot := new(discordgo.User)
	GodMasterBot.Email = "jeffprestes@gmail.com"
	GodMasterBot.Username = "DecendGameBot"
	GodMasterBot.Discriminator = "1424"
	GodMasterBot.ID = "990059486054064128"

	Villa = make(map[int]model.House)

	Villa[1] = model.NewHouse(1, serverName+"1.json", GodMasterBot)
	Villa[2] = model.NewHouse(2, serverName+"2.json", GodMasterBot)
	Villa[3] = model.NewHouse(3, serverName+"3.json", GodMasterBot)
	Villa[4] = model.NewHouse(4, serverName+"4.json", GodMasterBot)
	Villa[5] = model.NewHouse(5, serverName+"5.json", GodMasterBot)
	Villa[6] = model.NewHouse(6, serverName+"6.json", GodMasterBot)
	Villa[7] = model.NewHouse(7, serverName+"7.json", GodMasterBot)
	Villa[8] = model.NewHouse(8, serverName+"8.json", GodMasterBot)
	Villa[9] = model.NewHouse(9, serverName+"9.json", GodMasterBot)
	Villa[10] = model.NewHouse(10, serverName+"10.json", GodMasterBot)
	Villa[11] = model.NewHouse(11, serverName+"11.json", GodMasterBot)
	Villa[12] = model.NewHouse(12, serverName+"12.json", GodMasterBot)
	Villa[13] = model.NewHouse(13, serverName+"13.json", GodMasterBot)
	Villa[14] = model.NewHouse(14, serverName+"14.json", GodMasterBot)
	Villa[15] = model.NewHouse(15, serverName+"15.json", GodMasterBot)
	Villa[16] = model.NewHouse(16, serverName+"16.json", GodMasterBot)
	Villa[17] = model.NewHouse(17, serverName+"17.json", GodMasterBot)
	Villa[18] = model.NewHouse(18, serverName+"18.json", GodMasterBot)
	Villa[19] = model.NewHouse(19, serverName+"19.json", GodMasterBot)
	Villa[20] = model.NewHouse(20, serverName+"20.json", GodMasterBot)
	Villa[21] = model.NewHouse(21, serverName+"21.json", GodMasterBot)
	Villa[22] = model.NewHouse(22, serverName+"22.json", GodMasterBot)
	Villa[23] = model.NewHouse(23, serverName+"23.json", GodMasterBot)
	Villa[24] = model.NewHouse(24, serverName+"24.json", GodMasterBot)
	Villa[25] = model.NewHouse(25, serverName+"25.json", GodMasterBot)
	Villa[26] = model.NewHouse(26, serverName+"26.json", GodMasterBot)
	Villa[27] = model.NewHouse(27, serverName+"27.json", GodMasterBot)
	Villa[28] = model.NewHouse(28, serverName+"28.json", GodMasterBot)

	NFTs = make(map[int]model.HouseNFT)
	NFTs[1] = model.NewHouseNFT("House of Feronia", "abundance is provided in this house by Feronia", "https://bafkreihtx3eq4yqs2tavcdwbdxop3cv4e4dn3ekngeejsyof4uwya2zyry.ipfs.nftstorage.link/")
	NFTs[2] = model.NewHouseNFT("House of Minerva", "Justice reigns in Minerva's house", "https://bafkreigl3j5ehibghbfktuzsdz2vkqf2kcoo6umpecgkpzsejv26h73a2u.ipfs.nftstorage.link/")
	NFTs[3] = model.NewHouseNFT("House of Pales", "The steakshouse", "https://bafkreihgetdvjyflbqmy4iw6ntrmaujhdffcxwr4de5u7qnifhq7x3uys4.ipfs.nftstorage.link/")
	NFTs[4] = model.NewHouseNFT("House of Salus", "The House of Relax and well-being", "https://bafkreihysqydpfvwqbepnl6ajkgxlm4ch55snw7tpww324bpoyloqky26a.ipfs.nftstorage.link/")
	NFTs[5] = model.NewHouseNFT("House of Fortuna", "Here nothing gets wrong and ether is always in ATH", "https://bafkreihh6xbrpriijabbd3ek4sgonwnqebrul5m7srra5hyuin3rbxl67m.ipfs.nftstorage.link/")
	NFTs[6] = model.NewHouseNFT("House of Fides", "The house of the Goddess of the Blockchain", "https://bafkreicvv7xp4crktox7m2sylfb4omog4rundi4arcwj4jvdykbfd2wmmi.ipfs.nftstorage.link/")
	NFTs[7] = model.NewHouseNFT("House of Opis", "In this all bless of nature are given", "https://bafkreifjxzsmjaxqwhxi7fkxfswaqvbc5rovjk7duxe5u63xvsw2276hxq.ipfs.nftstorage.link/")
	NFTs[8] = model.NewHouseNFT("House of Flora", "No other house is this metaverse Villa is more beautiful", "https://bafkreia6uwhf5sbbbjgmhpji5izldo27c2pb3y3w7ehpjacpvrjzjrcfya.ipfs.nftstorage.link/")
	NFTs[9] = model.NewHouseNFT("House of Vejovis", "The house of Hunters", "https://bafkreifzlccym7h3h4xyq6hteopglzy2j6dirnmhhi3t6iux3vrqcmad5q.ipfs.nftstorage.link/")
	NFTs[10] = model.NewHouseNFT("House of Saturn", "In this house decisions are made", "https://bafkreibvkm2paw7e3blyfazn3y4jkfcam52sbohy3xseg2qjpuvnlasnga.ipfs.nftstorage.link/")
	NFTs[11] = model.NewHouseNFT("House of Volcanus", "This house is always on Fire", "https://bafkreigbh3ejnu3fapju6oqvfpfwmuqgpqth2krn3qdx432nk43dabrqpy.ipfs.nftstorage.link/")
	NFTs[12] = model.NewHouseNFT("House of Diana", "Hunting with class in precision here", "https://bafkreidi76en4r5j4zwzh6l7b5juahmycy73lnmq4jehv3eem2xxvrwnrm.ipfs.nftstorage.link/")
	NFTs[13] = model.NewHouseNFT("House of ", "", "")
	NFTs[14] = model.NewHouseNFT("House of ", "", "")
	NFTs[15] = model.NewHouseNFT("House of ", "", "")
	NFTs[16] = model.NewHouseNFT("House of ", "", "")
	NFTs[17] = model.NewHouseNFT("House of ", "", "")
	NFTs[18] = model.NewHouseNFT("House of ", "", "")
	NFTs[19] = model.NewHouseNFT("House of ", "", "")
	NFTs[20] = model.NewHouseNFT("House of ", "", "")
	NFTs[21] = model.NewHouseNFT("House of ", "", "")
	NFTs[22] = model.NewHouseNFT("House of ", "", "")
	NFTs[23] = model.NewHouseNFT("House of ", "", "")
	NFTs[24] = model.NewHouseNFT("House of ", "", "")
	NFTs[25] = model.NewHouseNFT("House of ", "", "")
	NFTs[26] = model.NewHouseNFT("House of ", "", "")
	NFTs[27] = model.NewHouseNFT("House of ", "", "")
	NFTs[28] = model.NewHouseNFT("House of ", "", "")
}
