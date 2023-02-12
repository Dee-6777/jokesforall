package ui

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"
	"strconv"

	"github.com/Dee-6777/jokesforall/cmd"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Datas struct {
	firstName string
	number    string
}

var data = make([]Datas, 0)

// Tview
const photo = `/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBxESERYRERERFhYREBYREREWEREREBAWFhYXFxYTFhYZHioiGRsnHBYWIzMjJystMDAwGCE2OzYvOiovMC0BCwsLDw4PGxERGy8hHh4vLy8vLy0tLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vL//AABEIAN8A4gMBIgACEQEDEQH/xAAbAAABBQEBAAAAAAAAAAAAAAAAAQMEBQYCB//EAEAQAAEDAgIGBggEBQQDAQAAAAEAAgMEEQUhBhIxQVFhBxMiMnGRQlJygaGxwdEUQ5LhIzNTYoIWk7LSVKLwFf/EABoBAQACAwEAAAAAAAAAAAAAAAADBAECBQb/xAAzEQACAgECBQIDBgcBAQAAAAAAAQIDEQQSBSExQVETYRQycSJSgaGx8BUjM0KR0eHBBv/aAAwDAQACEQMRAD8A9pQhCAEIQgBCEIAQhCAEIQgBCEWQAhLZFkAiEtkWQCIRZCAEIQgBCEIAQhCAEIQgBCEIAQhCAEIQgBCEIAQhCAEIS2QCJbKNU1rGZXueAzKrp8Se7Zl4bVXu1VVXzMlhTOfRFtJM1u0qJJiY3DzKqS5x4lPR0Tzu+i5NnGG3tqjn6cywtPGPObJLsTdyCPx7+PwCVmGO3kJ5uG8/gtFqNfPpB/oYfooj/jn8fkl//ReN/wAApBw4cfgm34YfWHksu3iEf7H+QTpEZih3gealRVzHb7Kukw144FRXxObtFliPFba3i6DX1WDPo1z+VmkSWWfhq3t2OPhtCsKfFAcni3MZhdSjX029Hghnp5x9ywQhrgRcG44oV0gBCEIAQhCAEIQgBCEIAQhCAEIQgBCE3UTtYLn3DeVhtJZYSyOPeGi7jYDaVS1uKl2TOyOPpH7KJX1xkPIbBuC5pKYvPLiuHrOJvPp09S/Vp1FbrBIwXGwG1WFNhhObjbkFOp4GsFgE7dVK9Gpfauefbt/01s1DfKPIIadjdgCdum7rq66dbjBYisL2Kr59Tq6Lrm6W6m9QC3RdJdF0cwLdcuAO0IuuSVpKzKwxgi1GHMdsyPwVXUUj2brjir665Oa51ukqk8w+y/bp/gnhdOPXmUVNVuYeyfFp7pV3R1jZBlkd7ftxVdXYeO83yVayQtPAg7d4WdPrraJKu7mv30J5VQuWY9TWJFBw/EA/su27jx/dTyu/CcZx3ReUUJRcXhiIQhbmoIQhACEIQAhCEAIQlCA5e8NBc42AFyVmsRrjI7gNgHJSscrrnq2nJve5u4e5VdPEXlcLimsa/lV/idHS0JL1Jfgd0dOXnlvKv4GBosFGp2BosFIaVz6IKpZfzPv/AOIxdY5v2H7pbpsFF1YVhWwOXS3TYKW6kVowd3S3Td0XW/qmuBy6S65ukunqjB3dJdc3SXWjtNsHd0l0l1zdaOwYFcVWYhSX7TdvDirBxTTiorNtkdsiWuTi8ooWSFp3/ZaTDK3rG2Pebt5j1lU19PcXCg0tQ5jg5u1u7cRvaVtodXLT2enN8n+8lm2pXQyupsChcU8wkaHt2EeXEFdr0yeTlAhCFkAhCEAIQhACjYlU9XGSNpyb48VKCzuP1N36vq5e/eoNTb6VbkS0V75pFY91yrKkZqhVtKLm6s2FeQUt0nNnWu+6iU1ycaVGaU4HLfeU5RJIcluo4kHFOByx6hG0PXSXXAKLrZWGMDl0XXF0t1t6pjB1dF1zdJdPVGDu6S64uglau0zg7JXJcuS5NukCx6hlROy5NucjXHFNuKbyRIR5VTWR6rrhWTiotU24WtjyvdFip7WSsAq7O1DsfmOTv3WgKxFPJqnwN1s6eUPYHDeLr0fDL/Uqw+xU1tW2eV3O0IQukUwQhCAEIQEBzNIGtLjuCxdXMXOJ4krSY/Nqx245+SyGvcricXuwlFHT0EOsiwpcgpjCoUJUXGsZjpouseeTWja88AuBHLxFFmayyxr8SigYZJXhrW7XH5DieS88xrpFleS2lZqt2dY4XcfAbAs/jWKy1cmvIchkxg7rByHHmtNoxoGZAJanWYw2LYR2ZHji4+iOW1dGNNdMd1vN+P31IXhczP0OkOKSzBsL5pHk9xo1h/llYDxXsWB9f1LfxGqJLdsNN2grvD6OKBnVwxsjb6rQBfx3k8ypGsqmo1MLOUY4wRPLHg5dByZDkt1V3Gu0fui6ZDkayxvMbR66LpnWRrJvG0cLlyXLi6QlZ3Gdp04rzDpBq8Shm12ukEFhqSRjsA7w47j4r0vWSOIIsbEHIgi4PIhTUXqueXHJsk10PG8L07rIiOsIkbwcACfBwH3XoWj2lUFULNOq8C7ond8cx6w5hQtINBIJgX04bDJmbAfwXnm30fELzasoZqWbVeHMfGbgg58nNO8LoKNGoX2Psv8Af5fQkTTPcS66YkKyOiOlXX2hmt1luy7YJOIPBy1TnLn2RlXLbIljEhSmxWj0eqbjUPiPqs1VlTcDqNWRvtAeeX1V/hd2y3HYauvfV9DYlIlckXqTighCEAJQkShAZ7SiTtAcG/M/ssxEe0r7Sd38Q+AHwVBF3l5fizzYzt6Jfy0WTHZLzHSvEjUVBAN2RksYN2W13vK3mMVBjp5Hja2N1vG2SyeguAGV34iZp1G9xrh/MPHPcPiq2j2wjK2XbkjNr7FvoNo4BaeZm/8AhNI/9yPktu6YDeqjGcZipo9Z5tua0d5x4ALzfF9JqipNmkxs9VpNz7R3rVV26qW58kR7cc31PTavSKCM2dIy/DWCit0wp/6jP1BeTtor7SU5+AHFT/A1LrJmMZ7HstHj8MmTZWE8A4XVkyYFeCmjI2EqywvSOqpiLPL2jax5JHuO0KOfDs/05Z+vIw4ntd0t1n9G9Ioqpl2ZOA7TD3m/cK8uubKMoPbJYZo4nd0ay4uuJZdUXK0GB0vAUWpxGNg7TmjxIC860m04e57oqbIA2Mp2u9kbhzWUkEsh1pHuJO8kldKrh8mt03t/U2UT1ufSunbl1rP1BNx6W05NutYP8gvKBQDeUGgG4qb4Gn7zNtp7ZT1zHi7XAg7wQQoOkGER1MRa4DWA/hvtm0+PBeSUs08DtaKRzbcDkfEbCtxoxpiJiIpgGvOQPoP8OB5KKzSTp+3W8pf5M4T+piaymkppbEFrmOuORByIXqGB4gJ4GS+s3tDg4ZOHmo+lWCtqoeyB1jM43bL/ANp5FZ/o9le3r4HAjq3NdqkWLXG4cPgFJbNX0b+8eptDlLDNXUp2gNio9QU9RHNRaH+oTW/Iby+XuQuIT2G+yPku17NHnwQhCAEoSJQgMppP/MPu+Sz0Z7S0ulbO3fi0fULLg2K8vxaP8xnb0TzWiyjsRmAeRzCkCQDIWyGzgqSbEDrNjiGs95s0fVaCsw8U9OHOOs8ka7uJ4DkttDwW2+tWWPZGXy+X748fUo6zidddjhBbpR6+F/tmM0u0aqKmVskLmlobqlrnauqb7RxRhugJsOumA4hjb/F32VrBpNC12o51jdXcFax4u0gqlO+6uKr6Jexc2t80+pWU+hlG3vCR/tSW+DbKR/pKh/on/cf91Zh3NdBV3db95mrh7spJ9CqN3d61nNsl/wDkCqeu6Owf5VR7ntHzH2W2BSgrMdVbHuY2vszKaK6Gmlk618wc7MBrWkNAPEnatcuLouo7rJWy3S6mEjq6j19KJY3RkkBwIuNovvT10XUaynlGcHnbOjiUPJ6+PVvkbO1reH7q7pdA4G/zJZX8hqsb8ifitTdBKtS1l0u4Sfkpm6I0I/KJ5mR1/muX6JUJ/KcOYkd9SrhISo/Wt+8zOz3Zla3QSAj+FNI3k7Ve36FZqo0EqxINV0ZaHA9YHEFovt1SNq9IlqGtFyQqar0nga7UDgScuSsVaq9dOf4Gyg33J1EDGxrXOuQLF3FdPDbkgC52mwufErvACyoL2O2at1W15fSy9TJm05xPO8cPculdwW30VbS9zxlx6cvbz9Dn18UrVu21YWcJ9s+/18jlQU7RHNRJZAcwpNBtXO0C+2di75Tewfy2+yPku0jRZoHAD5JV7M88CEIQAlCRKSBmdgzJ4ICh0sZ2A7gCPivO6mpc52qwEkmwsCSTuAC1OkuLund1cfcBsBvfzU/RvAGQ2lkAMlsuEfhz5rdcOo3K/ULd4j5fv7FS3iVrTo07x5l4Xt7iaK6OCAdbKLyvHiIx6o58Ux0h1WrCxg9J1/ILU3WC6RpbyRs4Mz95/ZXKW7bcy8P6Ll0XsUL9tVGId2vq8vqzzeDCJKifUY6znAuF9mWZVjLhGJ0/c1yBvadYeW1PMl/DVUEvo64DvA5H4FeoscF5+iuu+Dz5PTW2zqax0wjyaHS+riNpG3ttDmkFWlN0gj8yM+7916NLQwy5SRRu9poKgVGhNBJtha08Wkt+SjnwqqXTH5oytZ5Rn6XTuldtc5vi0/RWMeldKfzm++4XE/RVRv7kk7fexw+IVbP0Qj8up/VGPoVUnwVds/5N1q62XbdI6Y/nx/qCcGP0/wDWi/W1ZR/RHPuqIz7nBNO6J6vdLEf8nf8AVafwZ+WbfEVeTYHH6f8Arxf7jVy7SKmH58X6wsgOier3yxfqP/VOM6JajfNEP1H6LH8Gflj4iryaSTSqkG2ePzuoNRpzSN2PLvZafqoUPRE706hvuYT8yp0PRNTjvzSnkGsb8wVvHgq75/Ix8VWU9V0hx/lsefGw+qqqnTiofkxrW+5b+Do9oY/Re72n/aysYMBpYf5cEY56oJ8yrUOE1x6pfqRvWrsjyiNmJVOzrLHfm1qi4pgE1O9nWkXfmLG+yy9lcANgXnuk0/X4gyIbIW2PibOd82hS3UV0VOS6im6dtiT6Fz0f1BbOGk99pb8L/RbnG8IjqYurfkdrH+kx3EfZeeYNJ1dZHyeB59n6r1MlduOY10zjyzFHnU1Kd0Jc0pM8mq2TU0hikGzfucNzhyWg0fcJHtA3kD4hafGsLiqY9V4zGbHjvNP25LG0bJaSbddjrj1Xj6KG3h1Gpl69S2WL5l2fuvBPRxG7Sr0bHvrfyvvH2Z6aVymaGsbNGJG79o3tO8FPLR9S6nnoCEIQyAVRpPUlsWoD3763sjd8lcBZzSy9x7H1Km08VKxJlfVTcam0Z/CrdZrHjktTHOsPHPquVvT4nYZrq30OXNHnatQo5izTxz5rzzTqfWqPDVH/AN5rTHFmNGsTsCwOKVJllLj6TifDgoq4elGdsuSUWWbJK111R5tyQ9j9B1tLrjbH2udt/wAFI0f0zDI2xzsedUaokbY3A2awJvdWmHWMeqdhFisXiWGdRMYzs7zDuc07CPkvBaHVyrckvLf4Hur6FLCkem4dpBTy9yVt/VN2O8iruGUHevEhTna1xCk01dURdyR3hrGy60OIwfzIpS0b/tZ7bG9PtevIKbTSsj2tjd4h1/MFWcPSRKO9Tj3PIVhayl9yF6W1dj1EOK6DivNmdJo307/1fsnB0nx/+NL5tW3xNX3jHoWeD0XWSFy85d0ns3U8vm1MydKB9Gmd73/YJ8TV5Hw9ng9Ic5NPdzXmM3SRUO7sDB4ucVW1OmFdJuY3wv8AdaPWUruZWlsfY9UnmaNpA96pq/HqePvzNHm4/BeX1FXUyd+Q+ZTH4Vxzc8lQy4hBdCWOjfdm0xHTiFrSIWve62RI1Wjnnn8FUaH0jnukqZMy9xAPE5Fx8/kqKOjL3CNgu551Wjx3+G9b6mpmwQtibsY21+J3nzuuTxHWSnXt8/p3LunojCXIoettUa3B31XqrqjIe5eRVuUpPEra4ZjbXxN1jZzRqkeG9ez0i9fQ0TjzwsHjtS/R1t9cuWeaNEZ1RY85rvELmXFRuVRV1WsVdo07UsnPu1Kw4rmaLRKpLX6m6T/kNh8slrSsLo6T1jPbHzW6K5+sio2cu529BNyqw+wiEIVUvAFWY9R67LjdkfAqzSraEnFpo0sgpxcX3PLq+jc1yiWcvSazB2P2ZctyrHaOZ7F169fHHM4F3DZ56GCfE88Uz+Gsbr0f/TDS09qzrZZZBZOvoHMcWuFiDYrkcc4jKdHo1rlLq/bwdbgfC1Vf61nWPRe/lkehfZScTw5lTHqE6r25xSb2ngeRUFmRU6GVeDlmEsrk0ewklIxb2PjeYpRqvbtG4/3DiE4CtfiVBHUts/Jze5IO837jkspXYfNTnti7d0je77+CuwsjZ7Px/ortOPU4sksuWSApxbA51UaoSoQCaoRqpUIBNVdWSLlz0MilNvfuAJJNgALlxOwAJ2mp5Jnasbb8XbGN8StPg+Dsh7bu3JbvEZM5NG7xWJzjWsy6+DCy+SEwDCOob1kluteLW2iJvqDnxKlVcmScllUGokuqDcrJbpE8Y7UV1RDrFdMgcMxdS4ICStVguj3WR67yWg9wWzd/ceS9h/8AP8QnRF0yWYvnnx/xnm+P8MjqXG6DxNcmvK/2jItD07BTOcVs/wDTVuB+Cm0mBNbmfIL0stfDHI85Dhk89CJo5h5BDj6OfvWkKSNgaLAWAQuTbY7Jbmd6ilVQ2ghCFGTAhCEAIuhCAW6rMawwTNuANYC3tDgrJC1nBTWGZjJxeUeZ11GWEggi27eFFabL0jFMKZOODrZO+h4hYjE8MfE6z22vsPoO8CuBrNA4849Dr6fVKfJ9SNHKpLZARY2IORBzBUDUIXTXrjyrcWXOTIGI6N3JdARnn1ZNvIqknp5IzaRjm+IyPgVsI5rb08Ki4s4Ajgc1JHUyXKSyRurwYYSoEgWyfQ0zu9CzxF2n4Jh2B0h9F48JCpVqK++TXZLwZXrAkMi1TcBpBuefF5TzMNpm7IWHxu75o9RX2yY2S8GQja55sxpceABKtqDRyR1nTdkere7j9lomyNaLMaGjgAAm5JyVHLUt8orBsqvI7ExkbdVgAA3BNSSpl0iaJJUCi28kqSXQ6kkXEcRcU7FATbI5mwFrkngBvWswXRy1nyi28R7zzcfoulpdFKx9CC6+Na5kTAMF1yHOHYG3+7kFsBlkN2SAAMhu3IXo6aY1RwjjW2uyWWLdJdCFMRghCEAIQhACEIQAhCEAIQhAC4qIGSNLHtDgdoK7QgMxiOjJzMR1h6pPaHgd6ztTQuYbOaQeByK9JuuJoGPFnNBHMKjdoa59ORar1c49eZ5c5hCTXst1W6ORuzZly2hUVXo7I3cT4ZrlXcMmuhfr1kJFEJl11yfkw9w2j5hAozwVJ6KZYVsRjrkhlUj8GeCQUR4ItHPwPWiR+sSgEqzpsFkfsa79Jt5q7otGN7zblvVmrh05dSGeqhEyrICVcYdgEsljq6o9Z30G0rW0mGQx91gvxOZ8ypa6lPDoR+YpWa2T+Ug4ZhEUGYGs62bzt93AKeSkQulGKisIpNtvLBCELJgEIQgBCEIAQhCAEIQgBCEIAQhCAEIQgBCEIASgpEIBTZcGFh2tb5BdIWMIDYp4/Ub+kJ0NA2AeSRCYQFui6RCyAQhCAEIQgBCEIAQhCAEIQgBCEID/2Q==`

var pages = tview.NewPages()
var jokeText = tview.NewTextView()
var app = tview.NewApplication()
var form = tview.NewForm()
var jokesList = tview.NewList().ShowSecondaryText(false)
var flex = tview.NewFlex()
var text = tview.NewTextView().
	SetTextColor(tcell.ColorPaleVioletRed).
	SetText("(a) generate a random joke \n(b) get n no of jokes \n(q) quit")

func Greet() {
	jokesList.SetSelectedFunc(func(index int, name string, number string, shortcut rune) {
		setConcatText(&data[index])
	})

	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(jokesList, 0, 1, true).
			AddItem(jokeText, 0, 4, false), 0, 6, false).
		AddItem(text, 0, 1, false)

	flex.SetBorder(true).SetTitle("J[red]O[yellow]K[green]E[darkcyan] B[blue]O[darkmagenta]X[red]").SetTitleAlign(tview.AlignCenter)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 97 {
			form.Clear(true)
			addContactForm_()
			pages.SwitchToPage("Joke of the day")
		} else if event.Rune() == 98 {
			form.Clear(true)
			addContactForm()
			pages.SwitchToPage("Your Jokes")
		} else if event.Rune() == 113 {
			app.Stop()
		}
		return event
	})
	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("Your Jokes", form, true, false)
	pages.AddPage("Joke of the day", form, true, false)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func addJokeList() {
	jokesList.Clear()
	for index, jokes := range data {
		jokesList.AddItem(jokes.firstName+" "+jokes.number, " ", rune(49+index), nil)
	}
}

func addContactForm() *tview.Form {

	jokes_ := Datas{}

	b, _ := base64.StdEncoding.DecodeString(photo)
	img, _ := jpeg.Decode(bytes.NewReader(b))

	form.AddImage("", img, 0, 12, 0)
	form.SetBackgroundColor(tcell.ColorAntiqueWhite)

	form.AddInputField("First Name", "", 20, nil, func(firstName string) {
		jokes_.firstName = firstName
	})

	form.AddInputField("Number", "", 20, nil, func(number string) {
		jokes_.number = number
	})

	form.AddButton("save", func() {
		data = append(data, jokes_)
		addJokeList()
		pages.SwitchToPage("Menu")
	})

	form.AddButton("cancel", func() {
		pages.SwitchToPage("Menu")
	})
	form.SetBorder(true).SetBorderColor(tcell.ColorDarkRed).SetTitle("J[red]O[yellow]K[green]E[darkcyan] B[blue]O[darkmagenta]X[red]").SetTitleAlign(tview.AlignCenter)
	return form
}

func addContactForm_() *tview.Form {

	jokes_ := Datas{}
	b, _ := base64.StdEncoding.DecodeString(photo)
	img, _ := jpeg.Decode(bytes.NewReader(b))

	form.AddImage("", img, 0, 12, 0)
	form.SetBackgroundColor(tcell.ColorAntiqueWhite)

	form.AddInputField("First Name", "", 20, nil, func(firstName string) {
		jokes_.firstName = firstName
	})

	jokes_.number = ""

	form.AddButton("save", func() {
		data = append(data, jokes_)
		addJokeList()
		pages.SwitchToPage("Menu")
	})
	form.AddButton("cancel", func() {
		pages.SwitchToPage("Menu")
	})
	form.SetBorder(true).SetBorderColor(tcell.ColorDarkRed).SetTitle("J[red]O[yellow]K[green]E[darkcyan] B[blue]O[darkmagenta]X[red]").SetTitleAlign(tview.AlignCenter)
	return form
}

func setConcatText(contact *Datas) {
	jokeText.Clear()
	text := "Hey " + contact.firstName + " ready to laugh?\n"
	if contact.number != "" {
		t, _ := strconv.Atoi(contact.number)
		for i := 0; i < t; i++ {
			text += cmd.GetRandomJoke() + "\n"
		}
		if t < 1 {
			text += "You forget to add a number. Enter a number, please."
		}
	} else {
		text += cmd.GetRandomJoke()
	}
	jokeText.SetText(text)
}
