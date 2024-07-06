package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//1
	fmt.Println(removeDuplicateLetter("bananas"))
	fmt.Println(removeDuplicateLetter("lalalamama"))

	//2
	fmt.Println(isAnagram("Otto", "Toto"))
	fmt.Println(isAnagram("Ayam", "Maya"))
	fmt.Println(isAnagram("Tamat", "Kiamat"))

	//3
	words := []string{"this", "is", "a", "kanoha"}
	except := []string{"is", "a"}
	fmt.Println(capitalize(words, except))

	//4
	nums01 := []int{1, 2, 2, 3, 3, 4, 4, 4}
	fmt.Println(removeDuplicate(nums01))

	//5
	nums02 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("oddBeforeEven:", oddBeforeEven(nums02))

	//7
	nums1 := []int{1, 2, 4, 7, 8}
	nums2 := []int{2, 3, 7, 9}
	fmt.Println("sameElement:", sameElement(nums1, nums2))
	nums1 = []int{1, 2, 7, 4, 7, 8}
	nums2 = []int{7, 7, 3, 2, 9}
	fmt.Println("sameElement:", sameElement(nums1, nums2))
	nums1 = []int{2, 4, 6, 8}
	nums2 = []int{1, 3, 5, 7, 9}
	fmt.Println("sameElement:", sameElement(nums1, nums2))

	//9
	nums := []int{1, 2, 3, 4, 4, 4, 3, 3, 2, 4}
	fmt.Println("countFrequentElement:", countFrequentElement(nums))
	nums = []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
	fmt.Println("countFrequentElement:", countFrequentElement(nums))

}

// 1
func removeDuplicateLetter(words string) string {
	seen := make(map[rune]bool)
	result := []rune{}
	for _, char := range words {
		if !seen[char] {
			seen[char] = true
			result = append(result, char)
		}
	}
	return string(result)
}

// 2
func isAnagram(word1, word2 string) bool {

	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)

	r1 := []rune(word1)
	r2 := []rune(word2)

	if len(r1) != len(r2) {
		return false
	}

	sort.Slice(r1, func(i, j int) bool { return r1[i] < r1[j] })
	sort.Slice(r2, func(i, j int) bool { return r2[i] < r2[j] })

	for i := range r1 {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}

// 3
func capitalize(words []string, except []string) []string {
	exceptSet := make(map[string]bool)
	for _, word := range except {
		exceptSet[word] = true
	}

	for i, word := range words {
		if !exceptSet[word] {
			words[i] = strings.Title(word)
		}
	}
	return words
}

// 4
func removeDuplicate(nums []int) []int {
	seen := make(map[int]bool)
	result := []int{}
	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}
	return result
}

// 5
func oddBeforeEven(nums []int) []int {
	var odds []int
	var evens []int
	for _, num := range nums {
		if num%2 == 0 {
			evens = append(evens, num)
		} else {
			odds = append(odds, num)
		}
	}
	return append(odds, evens...)
}

// 7
func sameElement(nums1, nums2 []int) []int {
	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})
	for _, num := range nums1 {
		set1[num] = struct{}{}
	}
	for _, num := range nums2 {
		if _, exists := set1[num]; exists {
			set2[num] = struct{}{}
		}
	}
	var result []int
	for num := range set2 {
		result = append(result, num)
	}
	return result
}

// 9
func countFrequentElement(nums []int) map[int]int {
	sort.Ints(nums)
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}
	return frequencyMap
}
