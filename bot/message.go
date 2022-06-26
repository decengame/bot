package bot

import (
	"strconv"
	"strings"

	"github.com/decendgame/bot/cache"
	"github.com/decendgame/bot/model"
)

func welcome() (msg string) {
	msg = `
Hey, I am DecendGameBot. 
I am not smart. Alexa is smart. I am a Robot Ape. 
I control an modern Villa of Roman Empire. 
My Metaverse is Discord.
And last but not least I AM A CRYPTO DEGEN`
	return
}

func listCommands() (msg string) {
	msg = `
Type any of these commands:
start                   = start the game (work in progress)
who are you             = tell who i am
how to play             = explain the basic game mechanics
story of the game       = tells the tail behind the game
welcome                 = greetings
list commands           = display this message
villa map               = list all houses and their address`
	return
}

func greetings(name string) (msg string) {
	msg = "Salve, DEGEN " + name
	return
}

func howToPlay() (msg string) {
	msg = `
Tell to bot "start" . It will create a wallet if you don't have one and give out some coins.
After this it will read latest ethereum block and see in which house you will be hosted. 
It will offers options about what to do and also give information about the game status and other opponents.
You will see what other players are doing too.
An alfa: meanwhile is not your turn is pointless to write to the Bot`
	return
}

func story() (msg string) {
	msg = `
	DecedGame stands for Decentralized Excelsior Game. A Metaverse Monopoly Play-to-Earn game that mix ancient Roman myths and Gods and crypto operations. The main idea is to bring new people to Metaverse and Ethereum space, having fun and earning tokens.
	
	# How it works #
	
	Players, in Discord, are invited to join a Villa where they navigate thru this Metaverse. They start the game with some amount of coins and in each house in the Villa they stop they can buy it or rent it. Their order of appearance and the numbers of houses they move is based on the latest number of the latest Ethereum block. After the move, if they get in a house that already have an owner they neeed to rent or buy it. 
	
	Who define the houses's style and name, receive the payment for the initial sale and also define houses' initial price are the Godmasters Developers - Core Devs DAO that have built the game node softwares. The house prices are reajusted up 5% each time some player pays a rent. The rent price is 10% of the house's price.
	
	If the player does not have money to pay the rent she/he needs to answer a question showing her/his knowledge about Ethereum. If she/he fails to answer correctly they lost game and leave the board. If she/he does not want to take riskmn to miss the question, they can hire the "Specialists" - Ethereum specialists that put in stake some assets and then can be hired to answer the questions for the players - and if the question correctly answered by the specialist they can still to play.
	
	The game is endless. Players can access different boards. The cheastest one are for new begginers and the expensive ones - accessible only if the player holds NFT proven they have previous experience - are for advanced Ethereum users.
	
	## Game characters
	
	Investor
	Godmasters Developers
	Specialists`
	return
}

func villaMap() (msg []string) {
	var b strings.Builder
	var house model.House
	var nft model.HouseNFT
	for i := 1; i <= cache.NumberOfHouses; i++ {
		nft = cache.NFTs[i]
		house = cache.Villa[i]
		b.WriteString("**")
		b.WriteString(nft.Name)
		b.WriteString("**")
		b.WriteString("\n")
		b.WriteString("Location: ")
		b.WriteString(strconv.Itoa(house.IdOnBoard))
		b.WriteString("\n")
		b.WriteString(nft.Description)
		b.WriteString("\n")
		b.WriteString("Price: ")
		b.WriteString(strconv.Itoa(house.GetPrice()))
		b.WriteString("\n")
		b.WriteString("Rental: ")
		b.WriteString(strconv.Itoa(house.RentPrice()))
		b.WriteString("\n")
		b.WriteString("Token details at: ")
		b.WriteString("https://hoarse-well-made-theemim.explorer.hackathon.skalenodes.com/token/0xe50a4d6167fea0c90bda6dc262d907e750c30d75/instance/")
		b.WriteString(strconv.Itoa(i))
		b.WriteString("\n\n\n")
		msg = append(msg, b.String())
		b.Reset()
	}
	return
}
