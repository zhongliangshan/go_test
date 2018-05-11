package parser

import (
	"github.com/zhongliangshan/test/crawler/engine"
<<<<<<< HEAD
	"regexp"
	"github.com/zhongliangshan/test/crawler2/model"
=======
	"github.com/zhongliangshan/test/crawler/model"
	"regexp"
>>>>>>> fd473d8fc6f664aece8abcb043166579d4fd1245
	"strconv"
)

var ageRe = regexp.MustCompile(
	`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var heightRe = regexp.MustCompile(
	`<td><span class="label">身高：</span>(\d+)CM</td>`)
var incomeRe = regexp.MustCompile(
	`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(
	`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var genderRe = regexp.MustCompile(
	`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(
	`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(
	`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(
	`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(
	`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var hokouRe = regexp.MustCompile(
	`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(
	`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(
	`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var guessRe = regexp.MustCompile(
	`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
var idUrlRe = regexp.MustCompile(
	`http://album.zhenai.com/u/([\d]+)`)

func ParserProfile(contents []byte) engine.ParserResult {
	profile := model.Profile{}

	// 数字需要进行转化
	age, err := strconv.Atoi(
		extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(
		extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(
		extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Income = extractString(
		contents, incomeRe)
	profile.Gender = extractString(
		contents, genderRe)
	profile.Car = extractString(
		contents, carRe)
	profile.Education = extractString(
		contents, educationRe)
	profile.Hokou = extractString(
		contents, hokouRe)
	profile.House = extractString(
		contents, houseRe)
	profile.Marriage = extractString(
		contents, marriageRe)
	profile.Occupation = extractString(
		contents, occupationRe)
	profile.Xinzuo = extractString(
		contents, xinzuoRe)

	result := engine.ParserResult{
<<<<<<< HEAD
		Items:[]interface{}{profile},
=======
		Items: []interface{}{profile},
>>>>>>> fd473d8fc6f664aece8abcb043166579d4fd1245
	}

	return result

}

<<<<<<< HEAD
func extractString(contents []byte , re *regexp.Regexp)string {
=======
func extractString(contents []byte, re *regexp.Regexp) string {
>>>>>>> fd473d8fc6f664aece8abcb043166579d4fd1245
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}

	return ""
}
