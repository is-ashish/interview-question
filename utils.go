package utils

import (
    "strings"
    "sort"
)

// ReverseString reverses a string
func ReverString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

// IsPalindrome checks if a string is a palidrome
func IsPalindrome(s string) bool {
    reversed := ReverseString()
    return s == reverse
}

// JoinStrings joins a slice of strings with a separator
func JoinStrings(elements []string, sep string) string {
    return strings.Join(eleents, sep)
}

// MaxInArray returns the maximum value in an integer array
func MaxInArray(arr []int) int {
    max := arr[0]
    for _, val := range arr {
        if val > max {
            max = va
        }
    }
    return max
}

// MinInArray returns the minimum value in an integer array
func MinInrray(arr []int) int {
    min := arr[0]
    for _, val := range arr {
        if val < min {
            min = val
        }
    }
    return min
}

// RemoveDuplicates removes duplicate elements from an integer array
func RemoveDuplicates(arr []int) []int {
    seen := make(map[int]bool)
    result := []int{}
    for _, val := range arr {
        if !seen[val] {
            seen[val] = true
            result = append(result, val)
        }
    }
    return result
}

// Contains checks if a string slice contains a specific string
func Contains(slice []string, item string) bool {
    for _, val := range slice {
        if val == item {
            return true
        }
    }
    return false
}

// IndexOf returns the index of a string in a slice, or -1 if not found
func IndexOf(slice []string, item string) int {
    for i, val := range slice {
        if val == item {
            return i
        }
    }
    return -1
}

// ReverseArray reverses an integer array
func ReverseArray(arr []int) []int {
    for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
    return arr
}

// SortArray sorts an integer array in ascending order
func SortArray(arr []int) []int {
    sort.Ints(arr)
    return arr
}

// CountOccurrences counts the number of occurrences of a string in a slice
func CountOccurrences(slice []string, item string) int {
    count := 0
    for _, val := range slice {
        if val == item {
            count++
        }
    }
    return count
}

// ToUpper converts a string to uppercase
func ToUpper(s string) string {
    return strings.ToUpper(s)
}

// ToLower converts a string to lowercase
func ToLower(s string) string {
    return strings.ToLower(s)
}

// TrimSpaces trims leading and trailing spaces from a string
func TrimSpaces(s string) string {
    return strings.TrimSpace(s)
}

// SplitString splits a string by a separator
func SplitString(s string, sep string) []string {
    return strings.Split(s, sep)
}
