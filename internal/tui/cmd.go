package tui

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/acgtools/trace-moe-go/internal/search"

	tea "github.com/charmbracelet/bubbletea"
)

func traceMoe() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(1 * time.Second)

		res, err := unmarshal()
		if err != nil {
			return errMsg{err}
		}

		return successMsg{results: res}
	}
}

func unmarshal() ([]*search.Result, error) {
	var traceResp search.TraceMoeResponse
	err := json.Unmarshal([]byte(j), &traceResp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return traceResp.Result, nil
}

const j = `{
    "frameCount": 3287326,
    "error": "",
    "result": [
        {
            "anilist": {
                "id": 21355,
                "idMal": 31240,
                "title": {
                    "native": "Re:ゼロから始める異世界生活",
                    "romaji": "Re:Zero kara Hajimeru Isekai Seikatsu",
                    "english": "Re:ZERO -Starting Life in Another World-"
                },
                "synonyms": [
                    "Re: Life in a different world from zero",
                    "ReZero",
                    "Re Zero",
                    "Re：从零开始的异世界生活",
                    "Re:Zero รีเซทชีวิต ฝ่าวิกฤตต่างโลก",
                    "Re:Zero — жизнь с нуля в другом мире"
                ],
                "isAdult": false
            },
            "filename": "[Leopard-Raws] Re - Zero kara Hajimeru Isekai Seikatsu - 11 RAW (TX 1280x720 x264 AAC).mp4",
            "episode": 11,
            "from": 1255.67,
            "to": 1261.92,
            "similarity": 0.9195590305126724,
            "video": "https://media.trace.moe/video/21355/%5BLeopard-Raws%5D%20Re%20-%20Zero%20kara%20Hajimeru%20Isekai%20Seikatsu%20-%2011%20RAW%20(TX%201280x720%20x264%20AAC).mp4?t=1258.795&now=1701547200&token=GS3MTJoZ6Tk7Wte0cI26VP4KP8",
            "image": "https://media.trace.moe/image/21355/%5BLeopard-Raws%5D%20Re%20-%20Zero%20kara%20Hajimeru%20Isekai%20Seikatsu%20-%2011%20RAW%20(TX%201280x720%20x264%20AAC).mp4.jpg?t=1258.795&now=1701547200&token=lFC4qxYnXBamdgJwPCBuPfbXpPk"
        },
        {
            "anilist": {
                "id": 21355,
                "idMal": 31240,
                "title": {
                    "native": "Re:ゼロから始める異世界生活",
                    "romaji": "Re:Zero kara Hajimeru Isekai Seikatsu",
                    "english": "Re:ZERO -Starting Life in Another World-"
                },
                "synonyms": [
                    "Re: Life in a different world from zero",
                    "ReZero",
                    "Re Zero",
                    "Re：从零开始的异世界生活",
                    "Re:Zero รีเซทชีวิต ฝ่าวิกฤตต่างโลก",
                    "Re:Zero — жизнь с нуля в другом мире"
                ],
                "isAdult": false
            },
            "filename": "[Ohys-Raws] Re - Zero Kara Hajimeru Isekai Seikatsu - 11 (TX 1280x720 x264 AAC).mp4",
            "episode": 11,
            "from": 1256.08,
            "to": 1262.08,
            "similarity": 0.9177322723758564,
            "video": "https://media.trace.moe/video/21355/%5BOhys-Raws%5D%20Re%20-%20Zero%20Kara%20Hajimeru%20Isekai%20Seikatsu%20-%2011%20(TX%201280x720%20x264%20AAC).mp4?t=1259.08&now=1701547200&token=QjTH4jrv2Lb7A8wvU8pWF8KEgJE",
            "image": "https://media.trace.moe/image/21355/%5BOhys-Raws%5D%20Re%20-%20Zero%20Kara%20Hajimeru%20Isekai%20Seikatsu%20-%2011%20(TX%201280x720%20x264%20AAC).mp4.jpg?t=1259.08&now=1701547200&token=nDmnPYEhuVW6njT2aWfWOjbPfM"
        },
        {
            "anilist": {
                "id": 21355,
                "idMal": 31240,
                "title": {
                    "native": "Re:ゼロから始める異世界生活",
                    "romaji": "Re:Zero kara Hajimeru Isekai Seikatsu",
                    "english": "Re:ZERO -Starting Life in Another World-"
                },
                "synonyms": [
                    "Re: Life in a different world from zero",
                    "ReZero",
                    "Re Zero",
                    "Re：从零开始的异世界生活",
                    "Re:Zero รีเซทชีวิต ฝ่าวิกฤตต่างโลก",
                    "Re:Zero — жизнь с нуля в другом мире"
                ],
                "isAdult": false
            },
            "filename": "[ReinForce] Re - Zero kara Hajimeru Isekai Seikatsu - 11 (BDRip 1920x1080 x264 FLAC).mp4",
            "episode": 11,
            "from": 1245.75,
            "to": 1251.92,
            "similarity": 0.9148126534272799,
            "video": "https://media.trace.moe/video/21355/%5BReinForce%5D%20Re%20-%20Zero%20kara%20Hajimeru%20Isekai%20Seikatsu%20-%2011%20(BDRip%201920x1080%20x264%20FLAC).mp4?t=1248.835&now=1701547200&token=Velj1mXZSNcOVePAac9PdNLMvE",
            "image": "https://media.trace.moe/image/21355/%5BReinForce%5D%20Re%20-%20Zero%20kara%20Hajimeru%20Isekai%20Seikatsu%20-%2011%20(BDRip%201920x1080%20x264%20FLAC).mp4.jpg?t=1248.835&now=1701547200&token=622xLV89Gos1dpKcKRvPfCJU"
        },
        {
            "anilist": {
                "id": 21355,
                "idMal": 31240,
                "title": {
                    "native": "Re:ゼロから始める異世界生活",
                    "romaji": "Re:Zero kara Hajimeru Isekai Seikatsu",
                    "english": "Re:ZERO -Starting Life in Another World-"
                },
                "synonyms": [
                    "Re: Life in a different world from zero",
                    "ReZero",
                    "Re Zero",
                    "Re：从零开始的异世界生活",
                    "Re:Zero รีเซทชีวิต ฝ่าวิกฤตต่างโลก",
                    "Re:Zero — жизнь с нуля в другом мире"
                ],
                "isAdult": false
            },
            "filename": "Re - Zero kara Hajimeru Isekai Seikatsu New edition - 06 (BD 1280x720 x264 AAC).mp4",
            "episode": 6,
            "from": 2631.75,
            "to": 2637.83,
            "similarity": 0.9120919677383382,
            "video": "https://media.trace.moe/video/21355/Re%20-%20Zero%20kara%20Hajimeru%20Isekai%20Seikatsu%20New%20edition%20-%2006%20(BD%201280x720%20x264%20AAC).mp4?t=2634.79&now=1701547200&token=eLU99KljDddnPMhBdbzVpQOD5Qo",
            "image": "https://media.trace.moe/image/21355/Re%20-%20Zero%20kara%20Hajimeru%20Isekai%20Seikatsu%20New%20edition%20-%2006%20(BD%201280x720%20x264%20AAC).mp4.jpg?t=2634.79&now=1701547200&token=RFHFij2xnR8oWc3w63o8k0QuA"
        },
        {
            "anilist": {
                "id": 21355,
                "idMal": 31240,
                "title": {
                    "native": "Re:ゼロから始める異世界生活",
                    "romaji": "Re:Zero kara Hajimeru Isekai Seikatsu",
                    "english": "Re:ZERO -Starting Life in Another World-"
                },
                "synonyms": [
                    "Re: Life in a different world from zero",
                    "ReZero",
                    "Re Zero",
                    "Re：从零开始的异世界生活",
                    "Re:Zero รีเซทชีวิต ฝ่าวิกฤตต่างโลก",
                    "Re:Zero — жизнь с нуля в другом мире"
                ],
                "isAdult": false
            },
            "filename": "Re - Zero Kara Hajimeru Isekai Seikatsu - 11 (BD 1280x720 x264 AACx2).mp4",
            "episode": 11,
            "from": 1246.75,
            "to": 1252.08,
            "similarity": 0.9102652096015222,
            "video": "https://media.trace.moe/video/21355/Re%20-%20Zero%20Kara%20Hajimeru%20Isekai%20Seikatsu%20-%2011%20(BD%201280x720%20x264%20AACx2).mp4?t=1249.415&now=1701547200&token=aavTqMYKQf1z9SHr0jPVQu67Q0",
            "image": "https://media.trace.moe/image/21355/Re%20-%20Zero%20Kara%20Hajimeru%20Isekai%20Seikatsu%20-%2011%20(BD%201280x720%20x264%20AACx2).mp4.jpg?t=1249.415&now=1701547200&token=bOoMEL7FkRugxKgTdJvHFXGm3bs"
        },
        {
            "anilist": {
                "id": 21355,
                "idMal": 31240,
                "title": {
                    "native": "Re:ゼロから始める異世界生活",
                    "romaji": "Re:Zero kara Hajimeru Isekai Seikatsu",
                    "english": "Re:ZERO -Starting Life in Another World-"
                },
                "synonyms": [
                    "Re: Life in a different world from zero",
                    "ReZero",
                    "Re Zero",
                    "Re：从零开始的异世界生活",
                    "Re:Zero รีเซทชีวิต ฝ่าวิกฤตต่างโลก",
                    "Re:Zero — жизнь с нуля в другом мире"
                ],
                "isAdult": false
            },
            "filename": "[Ohys-Raws] Re - Zero kara Hajimeru Isekai Seikatsu New edition - 06 (AT-X 1280x720 x264 AAC).mp4",
            "episode": 6,
            "from": 2642.5,
            "to": 2647.75,
            "similarity": 0.9092413098191069,
            "video": "https://media.trace.moe/video/21355/%5BOhys-Raws%5D%20Re%20-%20Zero%20kara%20Hajimeru%20Isekai%20Seikatsu%20New%20edition%20-%2006%20(AT-X%201280x720%20x264%20AAC).mp4?t=2645.125&now=1701547200&token=kLMJa924yOHB7PO3pp1skycebiY",
            "image": "https://media.trace.moe/image/21355/%5BOhys-Raws%5D%20Re%20-%20Zero%20kara%20Hajimeru%20Isekai%20Seikatsu%20New%20edition%20-%2006%20(AT-X%201280x720%20x264%20AAC).mp4.jpg?t=2645.125&now=1701547200&token=FrXK4BmA8k6CrhKkknUNok1yYE"
        },
        {
            "anilist": {
                "id": 21355,
                "idMal": 31240,
                "title": {
                    "native": "Re:ゼロから始める異世界生活",
                    "romaji": "Re:Zero kara Hajimeru Isekai Seikatsu",
                    "english": "Re:ZERO -Starting Life in Another World-"
                },
                "synonyms": [
                    "Re: Life in a different world from zero",
                    "ReZero",
                    "Re Zero",
                    "Re：从零开始的异世界生活",
                    "Re:Zero รีเซทชีวิต ฝ่าวิกฤตต่างโลก",
                    "Re:Zero — жизнь с нуля в другом мире"
                ],
                "isAdult": false
            },
            "filename": "[Ohys-Raws] Re - Zero Kara Hajimeru Isekai Seikatsu - 11 (TX 1280x720 x264 AAC).mp4",
            "episode": 11,
            "from": 1255.83,
            "to": 1256.33,
            "similarity": 0.9037227076373229,
            "video": "https://media.trace.moe/video/21355/%5BOhys-Raws%5D%20Re%20-%20Zero%20Kara%20Hajimeru%20Isekai%20Seikatsu%20-%2011%20(TX%201280x720%20x264%20AAC).mp4?t=1256.08&now=1701547200&token=tZqC0C8Y2BYIj1yPKSm955tK4A",
            "image": "https://media.trace.moe/image/21355/%5BOhys-Raws%5D%20Re%20-%20Zero%20Kara%20Hajimeru%20Isekai%20Seikatsu%20-%2011%20(TX%201280x720%20x264%20AAC).mp4.jpg?t=1256.08&now=1701547200&token=zcwiMRibIRqpMjJGwtxLrPAsM40"
        },
        {
            "anilist": {
                "id": 21355,
                "idMal": 31240,
                "title": {
                    "native": "Re:ゼロから始める異世界生活",
                    "romaji": "Re:Zero kara Hajimeru Isekai Seikatsu",
                    "english": "Re:ZERO -Starting Life in Another World-"
                },
                "synonyms": [
                    "Re: Life in a different world from zero",
                    "ReZero",
                    "Re Zero",
                    "Re：从零开始的异世界生活",
                    "Re:Zero รีเซทชีวิต ฝ่าวิกฤตต่างโลก",
                    "Re:Zero — жизнь с нуля в другом мире"
                ],
                "isAdult": false
            },
            "filename": "[Leopard-Raws] Re - Zero kara Hajimeru Isekai Seikatsu - 11 RAW (TX 1280x720 x264 AAC).mp4",
            "episode": 11,
            "from": 1255.58,
            "to": 1256.17,
            "similarity": 0.8964110105645932,
            "video": "https://media.trace.moe/video/21355/%5BLeopard-Raws%5D%20Re%20-%20Zero%20kara%20Hajimeru%20Isekai%20Seikatsu%20-%2011%20RAW%20(TX%201280x720%20x264%20AAC).mp4?t=1255.875&now=1701547200&token=UMKfFPZuQwunBFQFc80bfrOYOk",
            "image": "https://media.trace.moe/image/21355/%5BLeopard-Raws%5D%20Re%20-%20Zero%20kara%20Hajimeru%20Isekai%20Seikatsu%20-%2011%20RAW%20(TX%201280x720%20x264%20AAC).mp4.jpg?t=1255.875&now=1701547200&token=6mXvRoTeTcNBHLiKuqwxIeeLcmo"
        },
        {
            "anilist": {
                "id": 21355,
                "idMal": 31240,
                "title": {
                    "native": "Re:ゼロから始める異世界生活",
                    "romaji": "Re:Zero kara Hajimeru Isekai Seikatsu",
                    "english": "Re:ZERO -Starting Life in Another World-"
                },
                "synonyms": [
                    "Re: Life in a different world from zero",
                    "ReZero",
                    "Re Zero",
                    "Re：从零开始的异世界生活",
                    "Re:Zero รีเซทชีวิต ฝ่าวิกฤตต่างโลก",
                    "Re:Zero — жизнь с нуля в другом мире"
                ],
                "isAdult": false
            },
            "filename": "Re - Zero kara Hajimeru Isekai Seikatsu New edition - 06 (BD 1280x720 x264 AAC).mp4",
            "episode": 6,
            "from": 2631.58,
            "to": 2632.25,
            "similarity": 0.8964110105645932,
            "video": "https://media.trace.moe/video/21355/Re%20-%20Zero%20kara%20Hajimeru%20Isekai%20Seikatsu%20New%20edition%20-%2006%20(BD%201280x720%20x264%20AAC).mp4?t=2631.915&now=1701547200&token=jd43FmSYTA4or4TeRjDluSajM",
            "image": "https://media.trace.moe/image/21355/Re%20-%20Zero%20kara%20Hajimeru%20Isekai%20Seikatsu%20New%20edition%20-%2006%20(BD%201280x720%20x264%20AAC).mp4.jpg?t=2631.915&now=1701547200&token=62Vlf58jtJ5CeaWorvneHE4mOZA"
        },
        {
            "anilist": {
                "id": 21355,
                "idMal": 31240,
                "title": {
                    "native": "Re:ゼロから始める異世界生活",
                    "romaji": "Re:Zero kara Hajimeru Isekai Seikatsu",
                    "english": "Re:ZERO -Starting Life in Another World-"
                },
                "synonyms": [
                    "Re: Life in a different world from zero",
                    "ReZero",
                    "Re Zero",
                    "Re：从零开始的异世界生活",
                    "Re:Zero รีเซทชีวิต ฝ่าวิกฤตต่างโลก",
                    "Re:Zero — жизнь с нуля в другом мире"
                ],
                "isAdult": false
            },
            "filename": "Re - Zero Kara Hajimeru Isekai Seikatsu - 11 (BD 1280x720 x264 AACx2).mp4",
            "episode": 11,
            "from": 1245.83,
            "to": 1246.5,
            "similarity": 0.8960165436233183,
            "video": "https://media.trace.moe/video/21355/Re%20-%20Zero%20Kara%20Hajimeru%20Isekai%20Seikatsu%20-%2011%20(BD%201280x720%20x264%20AACx2).mp4?t=1246.165&now=1701547200&token=K6L0lzY9Mc2IAlV4U3R1ocydDY",
            "image": "https://media.trace.moe/image/21355/Re%20-%20Zero%20Kara%20Hajimeru%20Isekai%20Seikatsu%20-%2011%20(BD%201280x720%20x264%20AACx2).mp4.jpg?t=1246.165&now=1701547200&token=T78TSB8mAYv8G9JbXq9sPmAK4u8"
        }
    ]
}`
