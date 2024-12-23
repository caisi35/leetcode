package main

import (
	"fmt"
	"math/bits"
)

func main() {
    fmt.Println(movesToChessboard([][]int{
        {0,1},
        {1,0},
    }))
}

func movesToChessboard(board [][]int) int {
    n := len(board)
    rowMask, colMask := 0, 0
    for i := 0; i < n; i++ {
        rowMask |= board[0][i] << i
        colMask |= board[i][0] << i
    }
    reverseRowMask := 1 << n-1 ^ rowMask
    reverseColMask := 1 << n-1 ^ colMask
    rowCnt, colCnt := 0, 0
    for i := 0; i < n; i++ {
        currRowMask, currColMask := 0, 0
        for j := 0; j < n; j++ {
            currColMask |= board[j][i] << j
            currRowMask |= board[i][j] << j
        }
        if currRowMask != rowMask && currRowMask != reverseRowMask ||
            currColMask != colMask && currColMask != reverseColMask {
                return -1
            }
        if currRowMask == rowMask {
            rowCnt++
        }
        if currColMask == colMask {
            colCnt++
        }
    }
    rowmoves := getMoves(uint(rowMask), rowCnt, n)
    colmoves := getMoves(uint(colMask), colCnt, n)
    if rowmoves == -1 || colmoves == -1 {
        return -1
    }
    return rowmoves + colmoves
}

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}

func getMoves(mask uint, count, n int) int {
    ones := bits.OnesCount(mask)
    if n & 1 > 0 {
        if abs(n-2*ones) != 1 || abs(n-2*count) != 1 {
            return -1
        }
        if ones == n >> 1 {
            return n / 2 - bits.OnesCount(mask&0xAAAAAAAA)
        } else {
            return (n+1)/2 - bits.OnesCount(mask&0x55555555)
        }
    } else {
        if ones != n >> 1 || count != n >> 1 {
            return -1
        }
        count0 := n/2 - bits.OnesCount(mask&0xAAAAAAAA)
        count1 := n/2 - bits.OnesCount(mask&0x55555555)
        return min(count0, count1)
    }
}