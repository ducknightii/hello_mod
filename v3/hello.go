package v3

import (
    "rsc.io/quote"
    quoteV3 "rsc.io/quote/v3"
)

func Hello() string {
    return quote.Hello() + "v3"
}

func Proverb() string {
    return quoteV3.Concurrency() + "v3"
}