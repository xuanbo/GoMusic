package util

import (
    "log"
    "strconv"
)

func ParseToInt64(val string) (int64, bool) {
    integer, err := strconv.ParseInt(val, 10, 64)
    if err != nil {
        log.Printf("string parse to int64 error[%s]\n", err.Error())
        return 0, false
    }
    return integer, true
}