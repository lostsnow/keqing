package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/internal/db"
	"github.com/lostsnow/keqing/internal/handler"
	"github.com/lostsnow/keqing/internal/lab/api/endpoint"
	"github.com/lostsnow/keqing/pkg/entity"
	"github.com/lostsnow/keqing/pkg/entity/gameaccount"
	"github.com/lostsnow/keqing/pkg/entity/gamerole"
	"github.com/lostsnow/keqing/pkg/entity/gameroleattribute"
	"github.com/lostsnow/keqing/pkg/i18n"
)

var ErrQrcodeLoginNotConfirmed = errors.New("QR code login not confirmed")

type QrcodeQueryPayload struct {
	AppID  int    `json:"app_id"`
	Device string `json:"device"`
	Ticket string `json:"ticket"`
}

type QrcodeQueryRequest struct {
	Request
	QrcodeQueryPayload
}

type QrcodeQueryResponse struct {
	Response
	Data QrcodeQueryResponseData `json:"data"`
}

type QrcodeQueryResponseData struct {
	Stat    string                         `json:"stat"`
	Payload QrcodeQueryResponseDataPayload `json:"payload"`
}

type QrcodeQueryResponseDataPayload struct {
	Proto string `json:"proto"`
	Raw   string `json:"raw"`
	Ext   string `json:"ext"`
}

func NewQrcodeQuery(appID int, device string, ticket string) *QrcodeQueryRequest {
	return &QrcodeQueryRequest{
		Request: Request{
			URL:    endpoint.Hk4eSdk + "/hk4e_cn/combo/panda/qrcode/query",
			Method: http.MethodPost,
		},
		QrcodeQueryPayload: QrcodeQueryPayload{
			AppID:  appID,
			Device: device,
			Ticket: ticket,
		},
	}
}

func (r *QrcodeQueryRequest) Do() (*QrcodeQueryResponseData, error) {
	payload := QrcodeQueryPayload{
		AppID:  r.AppID,
		Device: r.Device,
		Ticket: r.Ticket,
	}
	var v QrcodeQueryResponse
	err := SendRequest(r, payload, &v)
	if err != nil {
		return nil, err
	}

	if v.Data.Stat != "Confirmed" {
		return nil, ErrQrcodeLoginNotConfirmed
	}

	return &v.Data, nil
}

func (r *QrcodeQueryRequest) ParseGameToken(raw string) (*GameToken, error) {
	var v GameToken
	err := json.Unmarshal([]byte(raw), &v)
	if err != nil {
		return nil, fmt.Errorf("parse auth token %s failed: %w", raw, err)
	}

	return &v, nil
}

type QrcodeCheckTask struct {
	Context     telebot.Context
	UserID      int64
	QrcodeQuery *QrcodeQueryRequest
}

type QrcodeCheckPool struct {
	Tasks   chan QrcodeCheckTask
	wg      *sync.WaitGroup
	running map[int64]time.Time
}

func NewQrcodeCheckPool() *QrcodeCheckPool {
	return &QrcodeCheckPool{
		Tasks:   make(chan QrcodeCheckTask),
		wg:      &sync.WaitGroup{},
		running: make(map[int64]time.Time, 0),
	}
}

func (p *QrcodeCheckPool) Add(ctx telebot.Context, userID int64, r *QrcodeQueryRequest) {
	if _, ok := p.running[userID]; ok {
		return
	}

	p.running[userID] = time.Now()
	p.Tasks <- QrcodeCheckTask{
		Context:     ctx,
		UserID:      userID,
		QrcodeQuery: r,
	}
}

func (p *QrcodeCheckPool) IsRunning(userID int64) bool {
	if v, ok := p.running[userID]; ok {
		if v.Add(time.Minute * 3).Before(time.Now()) {
			delete(p.running, userID)

			return false
		}

		return true
	}

	return false
}

func (p *QrcodeCheckPool) Worker() {
	for task := range p.Tasks {
		resp, err := task.QrcodeQuery.Do()
		if err != nil {
			if errors.Is(err, ErrQrcodeLoginNotConfirmed) && p.IsRunning(task.UserID) {
				time.Sleep(time.Second * 2)
				go func(t QrcodeCheckTask) {
					p.Tasks <- t
				}(task)
			} else {
				delete(p.running, task.UserID)
				_ = task.Context.Send(i18n.T(task.Context, "Login failed or QR code expired, please try again"))
			}
		} else {
			task.ProcessingToken(resp.Payload.Raw)
			delete(p.running, task.UserID)
		}
	}
}

//nolint:cyclop
func (t *QrcodeCheckTask) ProcessingToken(raw string) {
	senderID := t.Context.Sender().ID
	gameToken, err := t.QrcodeQuery.ParseGameToken(raw)
	if err != nil {
		handler.ReportError(t.Context, "Game token parse for %d failed: %s", senderID, err)
		_ = t.Context.Send(i18n.T(t.Context, "Game token parse failed, please try again"))

		return
	}

	cookieReq := NewCookieTokenReq(gameToken.UID, gameToken.Token)
	cookieToken, err := cookieReq.Do()
	if err != nil {
		handler.ReportError(t.Context, "Cookie token fetch for %d failed: %s", senderID, err)
		_ = t.Context.Send(i18n.T(t.Context, "Cookie token fetch failed, please try again"))

		return
	}

	uid, _ := strconv.ParseInt(gameToken.UID, 10, 64)
	sTokenReq := NewSTokenReq(uid, gameToken.Token)
	sToken, err := sTokenReq.Do()
	if err != nil {
		handler.ReportError(t.Context, "stoken fetch for %d failed: %s", senderID, err)
		_ = t.Context.Send(i18n.T(t.Context, "stoken fetch failed, please try again"))

		return
	}

	err = db.DB.Client.GameAccount.
		Create().
		SetUserID(senderID).
		SetAccountID(gameToken.UID).
		SetGameToken(gameToken.Token).
		SetCookieToken(cookieToken.CookieToken).
		SetStoken(sToken.Token.Token).
		SetMid(sToken.UserInfo.MID).
		OnConflictColumns(gameaccount.FieldUserID, gameaccount.FieldAccountID).
		UpdateNewValues().
		Exec(context.Background())
	if err != nil {
		handler.ReportError(t.Context, "Upsert %d game account %s failed: %s", senderID, gameToken.UID, err)
		_ = t.Context.Send(i18n.T(t.Context, "Upsert game account failed"))

		return
	}

	gameRecordReq := NewGameRecordReq(sToken.UserInfo.AID, cookieToken.CookieToken)
	gameRecord, err := gameRecordReq.Do()
	if err != nil {
		handler.ReportError(t.Context, "fetch %d game record for %s failed: %s", senderID, gameToken.UID, err)
		_ = t.Context.Send(i18n.T(t.Context, "fetch game record failed, please try again"))

		return
	}

	if len(gameRecord.GameRecords) == 0 {
		_ = t.Context.Send(i18n.T(t.Context,
			"The account is not yet bound to any game, please make sure the account of scanning code is correct"))

		return
	}

	bulk := make([]*entity.GameRoleCreate, 0)
	bulkAttr := make([]*entity.GameRoleAttributeCreate, 0)
	for _, role := range gameRecord.GameRecords {
		if role.GameID != 2 {
			continue
		}
		bulk = append(bulk, db.DB.Client.GameRole.
			Create().
			SetUserID(senderID).
			SetAccountID(gameToken.UID).
			SetRoleID(role.GameRoleID).
			SetLevel(role.Level).
			SetRegion(role.Region).
			SetRegionName(role.RegionName).
			SetNickName(role.NickName))

		if len(role.Data) == 0 {
			continue
		}

		for _, attr := range role.Data {
			bulkAttr = append(bulkAttr, db.DB.Client.GameRoleAttribute.
				Create().
				SetUserID(senderID).
				SetAccountID(gameToken.UID).
				SetRoleID(role.GameRoleID).
				SetName(attr.Name).
				SetType(attr.Type).
				SetValue(attr.Value))
		}
	}

	if len(bulk) == 0 {
		_ = t.Context.Send(i18n.T(t.Context,
			"The account has not been bound to Genshin Impact, please confirm that the account is correct"))

		return
	}

	err = db.DB.Client.GameRole.
		CreateBulk(bulk...).
		OnConflictColumns(gamerole.FieldUserID, gamerole.FieldAccountID, gamerole.FieldRoleID).
		UpdateNewValues().
		Exec(context.Background())
	if err != nil {
		handler.ReportError(t.Context, "Upsert %d game roles %s failed: %s", senderID, gameToken.UID, err)
		_ = t.Context.Send(i18n.T(t.Context, "Upsert game roles failed"))

		return
	}

	if len(bulkAttr) > 0 {
		err = db.DB.Client.GameRoleAttribute.
			CreateBulk(bulkAttr...).
			OnConflictColumns(gameroleattribute.FieldUserID, gameroleattribute.FieldAccountID,
				gameroleattribute.FieldRoleID, gameroleattribute.FieldName).
			UpdateNewValues().
			Exec(context.Background())
		if err != nil {
			handler.ReportError(t.Context, "Upsert %d game roles attribute %s failed: %s", senderID, gameToken.UID, err)
			_ = t.Context.Send(i18n.T(t.Context, "Upsert game roles attribute failed"))

			return
		}
	}

	_ = t.Context.Send(i18n.T(t.Context, "Authentication successfully"))
}
