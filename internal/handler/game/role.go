package game

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/internal/db"
	"github.com/lostsnow/keqing/internal/handler"
	"github.com/lostsnow/keqing/pkg/entity"
	"github.com/lostsnow/keqing/pkg/entity/gamerole"
	"github.com/lostsnow/keqing/pkg/entity/gameroleattribute"
	"github.com/lostsnow/keqing/pkg/i18n"
	"github.com/lostsnow/keqing/pkg/object"
)

type RoleCMD struct{}

//nolint:cyclop
func Role(ctx telebot.Context) error {
	userID := ctx.Sender().ID

	vs, err := db.DB.Client.GameRole.
		Query().
		Where(gamerole.UserID(userID)).
		All(context.Background())
	if err != nil {
		_ = ctx.Reply(i18n.T(ctx, "Query game role failed"))

		return fmt.Errorf("query %d game role failed: %w", userID, err)
	}

	if len(vs) == 0 {
		return ctx.Reply(i18n.T(ctx, "No game role found, please send command /game_refresh_role to refresh"))
	}

	cmd := RoleCMD{}
	sel := &telebot.ReplyMarkup{}

	if len(vs) > 1 {
		botBtns := make([]telebot.Btn, 0, len(vs))

		for idx, btn := range vs {
			title := btn.NickName

			if idx == 0 {
				title += handler.CurrentInlineKeywordMark
			}

			botBtn := sel.Data(title, fmt.Sprintf("%d-%s-%s", userID, btn.AccountID, btn.RoleID))
			botBtns = append(botBtns, botBtn)
		}

		chunks := object.ChunkBy(botBtns, 3)
		rows := make([]telebot.Row, 0, len(chunks))

		for _, chunk := range chunks {
			rows = append(rows, chunk)
		}

		sel.Inline(rows...)

		for idx := range botBtns {
			ctx.Bot().Handle(&botBtns[idx], func(c telebot.Context) error {
				parts := strings.Split(c.Callback().Unique, "-")
				if len(parts) != 3 {
					return c.Respond()
				}

				cbUserID, _ := strconv.ParseInt(parts[0], 10, 64)
				cbAccountID := parts[1]
				cbRoleID := parts[2]

				cbV, err := db.DB.Client.GameRole.
					Query().
					Where(gamerole.UserID(cbUserID), gamerole.AccountID(cbAccountID), gamerole.RoleID(cbRoleID)).
					First(context.Background())
				if err != nil {
					_ = c.Edit(i18n.T(c, "Query game role failed"))

					return fmt.Errorf("query %d game role failed: %w", userID, err)
				}

				m := cmd.generateMessage(c, cbV)
				err = handler.UpdateCurrentInlineKeyboard(sel, c.Callback().Unique)
				if err != nil {
					return c.Respond()
				}

				return c.Edit(m, &telebot.SendOptions{ParseMode: telebot.ModeMarkdown, ReplyMarkup: sel})
			})
		}
	}

	m := cmd.generateMessage(ctx, vs[0])

	return ctx.Reply(m, &telebot.SendOptions{ParseMode: telebot.ModeMarkdown, ReplyMarkup: sel})
}

func (RoleCMD) generateMessage(ctx telebot.Context, role *entity.GameRole) any {
	_ = ctx
	format := `*%s*

ID: %s
Region: %s
Level: %d
`
	msg := fmt.Sprintf(format, role.NickName, role.RoleID, role.RegionName, role.Level)
	vs, err := db.DB.Client.GameRoleAttribute.
		Query().
		Where(gameroleattribute.UserID(role.UserID), gameroleattribute.AccountID(role.AccountID),
			gameroleattribute.RoleID(role.RoleID)).
		All(context.Background())

	if err != nil {
		return fmt.Errorf("query %d game role %s attribute failed: %w", role.UserID, role.RoleID, err)
	} else if len(vs) > 0 {
		msg += "\n"
		for _, v := range vs {
			msg += fmt.Sprintf("%s: %s\n", v.Name, v.Value)
		}
	}

	return msg
}
