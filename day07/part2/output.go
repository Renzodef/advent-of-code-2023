package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type HandRankBid struct {
	Hand string
	Rank int
	Bid  int
}

func cardOrder(card rune) int {
	order := "AKQT98765432J"
	return strings.IndexRune(order, card)
}

func compareHands(hand1, hand2 string) bool {
	for i := 0; i < len(hand1); i++ {
		if cardOrder(rune(hand1[i])) != cardOrder(rune(hand2[i])) {
			return cardOrder(rune(hand1[i])) > cardOrder(rune(hand2[i]))
		}
	}
	return false
}

func calculateRankWithJollies(cardCounts map[rune]int) int {
	switch len(cardCounts) {
	case 1:
		return 7
	case 2:
		for _, count := range cardCounts {
			if count == 4 {
				return 6
			}
			if count == 3 {
				return 5
			}
		}
	case 3:
		twoCount := 0
		for _, count := range cardCounts {
			if count == 2 {
				twoCount++
			}
		}
		if twoCount == 2 {
			return 3
		} else {
			return 4
		}
	case 4:
		return 2
	default:
		return 1
	}
	return 0
}

func calculateInitialHandRank(hand string) int {
	cardCounts := make(map[rune]int)
	jollyCount := 0
	for _, card := range hand {
		if card == 'J' {
			jollyCount++
			continue
		}
		cardCounts[card]++
	}
	for j := 0; j < jollyCount; j++ {
		maxCount := 0
		var maxCard rune
		for card, count := range cardCounts {
			if count > maxCount {
				maxCount = count
				maxCard = card
			}
		}
		cardCounts[maxCard]++
	}
	return calculateRankWithJollies(cardCounts)
}

func initialHandRankBid(line string) HandRankBid {
	parts := strings.Fields(line)
	hand := parts[0]
	bid, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("Error parsing bid:", err)
		return HandRankBid{}
	}
	initialHandRank := calculateInitialHandRank(hand)
	return HandRankBid{
		Hand: hand,
		Rank: initialHandRank,
		Bid:  bid,
	}
}

func processFile(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)
	var initialHandRankBidArray = make([]HandRankBid, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		initialHandRankBidArray = append(initialHandRankBidArray, initialHandRankBid(line))
	}
	groupedHandsBySameRank := make(map[int][]HandRankBid)
	for _, hrb := range initialHandRankBidArray {
		groupedHandsBySameRank[hrb.Rank] = append(groupedHandsBySameRank[hrb.Rank], hrb)
	}
	var ranks []int
	for rank := range groupedHandsBySameRank {
		ranks = append(ranks, rank)
	}
	sort.Ints(ranks)
	var finalHandRankBidArray = make([]HandRankBid, 0)
	finalRank := 1
	for _, rank := range ranks {
		group := groupedHandsBySameRank[rank]
		sort.Slice(group, func(i, j int) bool {
			return compareHands(group[i].Hand, group[j].Hand)
		})
		for _, hrb := range group {
			hrb.Rank = finalRank
			finalHandRankBidArray = append(finalHandRankBidArray, hrb)
			finalRank++
		}
	}
	totalSumOfProductsBidRank := 0
	for _, hrb := range finalHandRankBidArray {
		totalSumOfProductsBidRank += hrb.Bid * hrb.Rank
	}
	return totalSumOfProductsBidRank
}

func main() {
	startTime := time.Now()
	result := processFile("../input.txt")
	elapsedTime := time.Since(startTime)
	fmt.Println("Result:", result)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
