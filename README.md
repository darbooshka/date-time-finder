# datetimefinder

Usage:

###### main.go

```go
package main

import (
	"fmt"
	"time"

	"github.com/darbooshka/datetimefinder"
)

func main() {
	finder := datetimefinder.NewDateTimeFinder()

	text := `Фізика Тема: Розв'язування задач.
Время: 29 травня 2023 01:30 PM Киев
Інший час: 31 мая 2023 11:30 AM Київ

Подключиться к конференции Zoom`

	results := finder.FindDateTime(text)

	if len(results) > 0 {
		fmt.Println("DateTimes found:")
		for _, dt := range results {
			fmt.Println(dt.Format(time.RFC3339))
		}
	} else {
		fmt.Println("No DateTimes found.")
	}
}
```

Output:

```
DateTimes found:
2023-05-29T13:30:00Z
2023-05-31T11:30:00Z
```
