# hide

A minimal go library for hiding textual information in plain sight

# What's this

I came across [this webpage](https://invisible-characters.com/) gathering unicode characters that are not (usually)
rendered and thought it'd be fun to find some real-world applications for this

```go
package main

import (
	"fmt"

	"github.com/pieroshka/hide"
)

func main() {
	s := "hello world"

	concealer := hide.New()
	hidden, err := concealer.Hide(s)
	if err != nil {
		panic(err)
	}

	unhidden, err := concealer.Unhide(hidden)
	if err != nil {
		panic(err)
	}

	fmt.Printf(">%s<\n", s)
	fmt.Printf(">%s<\n", hidden)
	fmt.Printf(">%s<\n", unhidden)

    /*
    >hello world<
    >󠁡󠁇󠁖󠁳󠁢󠁇󠀸󠁧󠁤󠀲󠀹󠁹󠁢󠁇󠁑󠀽<
    >hello world<
    */
}
```

> [!TIP]
> As of now, the library takes any string input and base64-encodes it, then maps each base64 character to its equivalent in tag characters,
> starting from [E0041 TAG LATIN CAPITAL LETTER A](https://invisible-characters.com/E0041-TAG-LATIN-CAPITAL-LETTER-A.html)

> [!WARNING]
> This library is created mainly for fun & research purposes and **should not be consider production-ready** at any point.

## How do I use it?

```bash
go get github.com/pieroshka/hide
```

# Use cases?

One thing I immediately thought of was storing conversation metadata for a stateless LLM-powered Discord bot. Since electron doesn't render the tag characters, I could encode additional chat metadata, like tool uses, right within Discord messages
