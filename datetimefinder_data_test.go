package datetimefinder

import (
	"time"
)

func timeParse(s string) time.Time { 
	result, _ := time.Parse("2006-01-02 3:04 PM", s)
	return result
}

var testCases = []testCase[time.Time]{
	{
		text:         "I have a meeting scheduled for 07/03/2023 2:30 PM. Please be on time.",
		expected:     []time.Time{ timeParse("2023-07-03 2:30 PM") },
	},
	{
		text:     "Reminder: Event on 2023-07-04 at 10:00 AM.",
		expected: []time.Time{timeParse("2023-07-04 10:00 AM")},
		ignore:   reason("Demo Ignore, Date format is different."),
	},
	// Add more test cases here...
	{
		text:         `I have a meeting
It is scheduled for 07/03/2023 2:30 PM.
The next is on  07/04/2023 12:30 PM.
Please be on time.`,
		expected:     []time.Time{
			timeParse("2023-07-03 2:30 PM"),
			timeParse("2023-07-04 12:30 PM"),
		},
		//ignore:   reason("Date format is different."),
	},
	{
		text:         `I have a meeting
It is scheduled for 07/03/2023 2:30 PM.
The next is on 2023/07/04 12:30 PM.
Please be on time.`,
		expected:     []time.Time{
			timeParse("2023-07-03 2:30 PM"),
			timeParse("2023-07-04 12:30 PM"),
		},
	},
	{
		text:         ``,
		expected:     []time.Time{}, // nil, //
	},
	{
		text:         `Фізика Тема: Розв'язування задач. Время: 29 May 2023 01:30 PM Киев Подключиться к конференции Zoom`,
		expected:     []time.Time{ timeParse("2023-05-29 01:30 PM") },
	},
	{
		text:         `Фізика Тема: Розв'язування задач. Время: 24 Feb 2023 01:30 PM Киев Подключиться к конференции Zoom`,
		expected:     []time.Time{ timeParse("2023-02-24 01:30 PM") },
	},
	{
		text:         `Фізика Тема: Розв'язування задач. Время: 24 February 2023 01:30 PM Киев Подключиться к конференции Zoom`,
		expected:     []time.Time{ timeParse("2023-02-24 01:30 PM") },
	},
	{
		text:         `Фізика Тема: Розв'язування задач.
Время: 29 мая 2023 01:30 PM Киев

Подключиться к конференции Zoom
https://us04web.zoom.us/j/769268?pwd=FupuVeZ0иукекLJ6Ml3wQuT9S.1

Идентификатор конференции: 765 1993 9268
Код доступа: Byex
`,
		// expected:     []time.Time{ timeParse("29 мая 2023 01:30 PM") },
		expected:     []time.Time{ timeParse("2023-05-29 01:30 PM") },
	},
	{
		text:         `Фізика Тема: Розв'язування задач.
Время: 29 травня 2023 01:30 PM Киев

Подключиться к конференции Zoom
https://us04web.zoom.us/j/76439268?pwd=FupuVк1JGlLJ6Ml3wQuT9S.1

Идентификатор конференции: 765 1993 9268
Код доступа: Byex
`,
		// expected:     []time.Time{ timeParse("29 мая 2023 01:30 PM") },
		expected:     []time.Time{ timeParse("2023-05-29 01:30 PM") },
	},
	{
		text:         `Біологія. Zoom-урок 29.05 о 8:40. Тести.
Катерина Бульба is inviting you to a scheduled Zoom meeting.

Topic: 8 Г
Time: May 29, 2023 08:40 Kyiv

Join Zoom Meeting
https://us04web.zoom.us/j/746765759?pwd=WaQ4dмh25IlLORzk2LfPV.1

Meeting ID: 756 3692 5759
Passcode: Yz0M
`,
		expected:     []time.Time{ timeParse("2023-05-29 08:40 AM") },
	},
	{
		text:         `НОВИЙ МАТЕРІАЛ
Геометрія 19 травня 9.30
Kateryna Bulba is inviting you to a scheduled Zoom meeting.

Topic: Kateryna Bulba's Personal Meeting Room
`,
		expected:     []time.Time{ timeParse("2023-05-29 08:40 AM") },
	},
	// {
	// 	text:         ``,
	// 	expected:     []time.Time{},
	// },
	// {
	// 	text:         ``,
	// 	expected:     []time.Time{},
	// },
	// 		{
	// 	text:         ``,
	// 	expected:     []time.Time{},
	// },
	// {
	// 	text:         ``,
	// 	expected:     []time.Time{},
	// },
}
