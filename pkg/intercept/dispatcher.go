package intercept

import "github.com/blogtitle/wapty/pkg/structures"

func Dispatch(out chan<- structures.Request, in <-chan structures.Request) {}
