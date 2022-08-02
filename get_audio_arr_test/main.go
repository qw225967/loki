/*******************************************************
 * @author      : dog head
 * @date        : Created in 2022/8/2
 * @mail        : 396139919@qq.com
 * @project     : get_audio_arr_test
 * @file        : main.go
 * @description : ali 语音sdk 分组测试
 *******************************************************/

package main

type DataArr struct {
	EndTime         uint64
	SilenceDuration uint64
	BeginTime       uint64
	Text            string
	ChannelId       uint64
	SpeechRate      uint64
	EmotionValue    float64
}

func cutAudioScoreArr(Sentences []DataArr) []DataArr {
	var res     []DataArr

	// 1.init filter params
	interval := uint64(10000) // 30s，单位ms

	// 2.get filter params from apollo

	// 3.start filter
	lens := int(len(Sentences))
	begin := uint64(0)
	silence := uint64(0)
	preEnd := uint64(0)

	for i := 0; i < lens; i++ {
		// 初始开始时间,当第一个段时记录开始时间
		if begin == 0 {
			begin = Sentences[i].BeginTime
		}

		// 上一段结束时间计算
		if preEnd != 0 {
			// 当前段开始减去上一段结束等于静默时间
			silence = Sentences[i].BeginTime - preEnd + silence
		}

		// 多段累加，小于统计时长则继续累计
		if Sentences[i].EndTime - begin < interval {
			// 加上当前段的静音时长
			silence += Sentences[i].SilenceDuration

			// 更新结束时间，提供下一个段的开始减去它得到无人声段
			preEnd = Sentences[i].EndTime

			// 累计足够了则停下
		} else {
			// 之前的静默时间加上当前的静默时间
			silence += Sentences[i].SilenceDuration
			temp := DataArr{
				BeginTime:       begin/1000,
				EndTime:         Sentences[i].EndTime,
				SilenceDuration: silence,
			}
			res = append(res, temp)

			// 重置计数
			begin = 0
			silence = 0
			preEnd = 0
		}

	}

	return res
}


func main() {
	var arr []DataArr
	temp1 := DataArr{
		BeginTime: 1212,
		EndTime: 3222,
		SilenceDuration: 572,
	}
	temp2 := DataArr{
		BeginTime: 5212,
		EndTime: 9222,
		SilenceDuration: 572,
	}
	temp3 := DataArr{
		BeginTime: 12102,
		EndTime: 13222,
		SilenceDuration: 572,
	}
	temp4 := DataArr{
		BeginTime: 21212,
		EndTime: 23222,
		SilenceDuration: 572,
	}
	temp5 := DataArr{
		BeginTime: 26212,
		EndTime: 33222,
		SilenceDuration: 572,
	}
	temp6 := DataArr{
		BeginTime: 36212,
		EndTime: 39222,
		SilenceDuration: 572,
	}
	temp7 := DataArr{
		BeginTime: 41212,
		EndTime: 43222,
		SilenceDuration: 572,
	}
	arr = append(arr, temp1)
	arr = append(arr, temp2)
	arr = append(arr, temp3)
	arr = append(arr, temp4)
	arr = append(arr, temp5)
	arr = append(arr, temp6)
	arr = append(arr, temp7)

	_ = cutAudioScoreArr(arr)

}
