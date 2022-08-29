/*******************************************************
 * @author      : dog head
 * @date        : Created in 2022/8/2
 * @mail        : 396139919@qq.com
 * @project     : get_audio_arr_test
 * @file        : main.go
 * @description : ali 语音sdk 分组测试
 *******************************************************/

package main

import "fmt"

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
	interval := uint64(15000) // 30s，单位ms

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

		// 多段累加，小于统计时长则继续累计
		if Sentences[i].EndTime - begin < interval {
			// 上一段结束时间计算
			if preEnd != 0 {
				// 当前段开始减去上一段结束等于静默时间
				silence = Sentences[i].BeginTime - preEnd + silence
			}

			// 更新结束时间，提供下一个段的开始减去它得到无人声段
			preEnd = Sentences[i].EndTime

		// 累计足够了则停下
		} else {
			// 如果这一段开始前也大于了间隔 interval，那么截断
			if Sentences[i].BeginTime - begin > interval {
				// 剩余的全是静音
				silence += interval - (preEnd - begin)
				// 偏移到截断位置
				preEnd += interval - (preEnd - begin)
			} else {
				// 先计算中间段的静音时长
				silence += Sentences[i].BeginTime - preEnd
				// 偏移尾到当前头
				preEnd = Sentences[i].BeginTime
				// 再计算截断位置
				preEnd += interval - (preEnd - begin)
			}

			temp := DataArr{
				BeginTime:       begin,
				EndTime:         preEnd,
				SilenceDuration: silence,
			}
			res = append(res, temp)

			// 重置计数
			begin = 0
			silence = 0
			preEnd = 0
		}
	}

	var res2 []DataArr
	for _, v := range res {
		if v.EndTime-v.SilenceDuration-v.BeginTime > interval &&
			(1-float64(v.SilenceDuration)/float64(v.EndTime-v.BeginTime)) > 0.7 {
			res2 = append(res2, v)
		}
	}

	return res2
}


func main() {
	var arr []DataArr
	temp1 := DataArr{BeginTime:2870, EndTime:3580, SilenceDuration:2     }
	temp2 := DataArr{BeginTime:24810, EndTime:37010, SilenceDuration:21  }
	temp3 := DataArr{BeginTime:44530, EndTime:55140, SilenceDuration:7   }
	temp4 := DataArr{BeginTime:55450, EndTime:61220, SilenceDuration:0   }
	temp5 := DataArr{BeginTime:62060, EndTime:73677, SilenceDuration:0   }
	temp6 := DataArr{BeginTime:73677, EndTime:88620, SilenceDuration:0   }
	temp7 := DataArr{BeginTime:89550, EndTime:91052, SilenceDuration:0   }
	temp8 := DataArr{BeginTime:91052, EndTime:91720, SilenceDuration:0   }
	temp9 := DataArr{BeginTime:95990, EndTime:106307, SilenceDuration:4  }
	temp10 := DataArr{BeginTime:106307, EndTime:110600, SilenceDuration:0 }
	temp11 := DataArr{BeginTime:110780, EndTime:115460, SilenceDuration:0 }
	temp12 := DataArr{BeginTime:126920, EndTime:128080, SilenceDuration:11}
	temp13 := DataArr{BeginTime:131260, EndTime:134120, SilenceDuration:3 }
	temp14 := DataArr{BeginTime:155630, EndTime:156800, SilenceDuration:21}
	temp15 := DataArr{BeginTime:161660, EndTime:163640, SilenceDuration:4 }
	temp16 := DataArr{BeginTime:177100, EndTime:177790, SilenceDuration:13}
	temp17 := DataArr{BeginTime:206640, EndTime:207290, SilenceDuration:28}
	temp18 := DataArr{BeginTime:221380, EndTime:222340, SilenceDuration:14}
	temp19 := DataArr{BeginTime:259190, EndTime:260230, SilenceDuration:36}
	temp20 := DataArr{BeginTime:299030, EndTime:299680, SilenceDuration:38}
	temp21 := DataArr{BeginTime:325540, EndTime:329540, SilenceDuration:25}
	temp22 := DataArr{BeginTime:335620, EndTime:336410, SilenceDuration:6 }
	temp23 := DataArr{BeginTime:339510, EndTime:341650, SilenceDuration:3 }
	arr = append(arr, temp1)
	arr = append(arr, temp2)
	arr = append(arr, temp3)
	arr = append(arr, temp4)
	arr = append(arr, temp5)
	arr = append(arr, temp6)
	arr = append(arr, temp7)
	arr = append(arr, temp8)
	arr = append(arr, temp9)
	arr = append(arr, temp10)
	arr = append(arr, temp11)
	arr = append(arr, temp12)
	arr = append(arr, temp13)
	arr = append(arr, temp14)
	arr = append(arr, temp15)
	arr = append(arr, temp16)
	arr = append(arr, temp17)
	arr = append(arr, temp18)
	arr = append(arr, temp19)
	arr = append(arr, temp20)
	arr = append(arr, temp21)
	arr = append(arr, temp22)
	arr = append(arr, temp23)

	_ = cutAudioScoreArr(arr)

	sampleCount := 1

	dropCount := uint64(len(arr)/sampleCount)
	var zoomArr  []DataArr
	for i:=0;i<len(arr);i++ {
		zoomArr = append(zoomArr, arr[i])
		for j:=0;j<int(dropCount);j++ {
			i++
		}
	}
	fmt.Println(len(zoomArr))
}
