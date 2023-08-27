package material

import (
	"fmt"
	"strings"
	"time"

	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/internal/handler"
	"github.com/lostsnow/keqing/pkg/i18n"
	_ "github.com/lostsnow/keqing/pkg/i18n/catalog"
)

func Daily(ctx telebot.Context) error {
	var arg string
	if len(ctx.Args()) == 0 {
		arg = ""
	} else {
		arg = strings.Join(ctx.Args(), " ")
	}

	wds := parseWeekDay(arg)
	buttons := make([]handler.PhotoButton, 0)

	for _, wd := range wds {
		var title string

		switch {
		case wd == 1:
			title = fmt.Sprintf("%s/%s", i18n.T(ctx, "Monday"), i18n.T(ctx, "Thursday"))
		case wd == 2:
			title = fmt.Sprintf("%s/%s", i18n.T(ctx, "Tuesday"), i18n.T(ctx, "Friday"))
		default:
			title = fmt.Sprintf("%s/%s", i18n.T(ctx, "Wednesday"), i18n.T(ctx, "Saturday"))
		}

		buttons = append(buttons, handler.PhotoButton{
			Title: title,
			Dir:   "assets/material/daily",
			Name:  fmt.Sprintf("%d.png", wd),
		})
	}

	h := handler.PhotoResponseHandler{
		Buttons:        buttons,
		NoPhotoMessage: handler.NoPhoto,
	}

	return h.Handle(ctx)
}

//nolint:cyclop
func parseWeekDay(s string) []int {
	s = strings.ToLower(strings.TrimSpace(s))

	all := []int{1, 2, 3}
	first := []int{1}
	second := []int{2}
	third := []int{3}

	weekdayAlias := map[string]time.Weekday{
		"sunday":    time.Sunday,
		"sun":       time.Sunday,
		"周日":        time.Sunday,
		"星期天":       time.Sunday,
		"monday":    time.Monday,
		"mon":       time.Monday,
		"周一":        time.Monday,
		"星期一":       time.Monday,
		"tuesday":   time.Tuesday,
		"tue":       time.Tuesday,
		"周二":        time.Tuesday,
		"星期二":       time.Tuesday,
		"wednesday": time.Wednesday,
		"wed":       time.Wednesday,
		"周三":        time.Wednesday,
		"星期三":       time.Wednesday,
		"thursday":  time.Thursday,
		"thu":       time.Thursday,
		"周四":        time.Thursday,
		"星期四":       time.Thursday,
		"friday":    time.Friday,
		"fri":       time.Friday,
		"周五":        time.Friday,
		"星期五":       time.Friday,
		"saturday":  time.Saturday,
		"sat":       time.Saturday,
		"周六":        time.Saturday,
		"星期六":       time.Saturday,
	}

	dailyAlias := map[time.Weekday][]int{
		time.Monday:    first,
		time.Tuesday:   second,
		time.Wednesday: third,
		time.Thursday:  first,
		time.Friday:    second,
		time.Saturday:  third,
		time.Sunday:    all,
	}

	if v, ok := weekdayAlias[s]; ok {
		return dailyAlias[v]
	}

	zone := time.FixedZone("CST", 8*3600)

	var t time.Time

	switch {
	case s == "明天" || s == "明日" || s == "tomorrow":
		t = time.Now().In(zone).AddDate(0, 0, 1)
	case s == "昨天" || s == "昨日" || s == "yesterday":
		t = time.Now().In(zone).AddDate(0, 0, -1)
	case s == "后天" || s == "后日" || s == "after tomorrow":
		t = time.Now().In(zone).AddDate(0, 0, 2)
	case s == "前天" || s == "前日" || s == "before yesterday":
		t = time.Now().In(zone).AddDate(0, 0, -2)
	default:
		t = time.Now().In(zone)
	}

	wd := t.Weekday()

	return dailyAlias[wd]
}
