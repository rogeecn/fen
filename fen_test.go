package fen

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func init() {
	DebugMode = true
}

type respUser struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

type tQueryFilter struct {
	Username string `query:"username"`
	Password string `query:"password"`
}

type testApi struct{}

var Err_Test = NewBusError(http.StatusBadRequest, 1001, "TestErr")

func (t *testApi) Func(ctx *fiber.Ctx) error {
	return Err_Test.Wrap(errors.New("TestStack"))
}

func (t *testApi) FuncP1(ctx *fiber.Ctx, uid int) error {
	return Err_Test.Wrap(errors.New("TestStack"))
}

func (t *testApi) Data(ctx *fiber.Ctx) (*respUser, error) {
	return &respUser{Name: "TestName"}, nil
}

func (t *testApi) DataP1(ctx *fiber.Ctx, uid int) (*respUser, error) {
	return &respUser{ID: uid, Name: "TestName"}, nil
}

func (t *testApi) DataP2(ctx *fiber.Ctx, uid int, name string) (*respUser, error) {
	return &respUser{ID: uid, Name: name}, nil
}

func (t *testApi) DataP1Form(ctx *fiber.Ctx, user *respUser) (*respUser, error) {
	return user, nil
}

func (t *testApi) DataP1Query(ctx *fiber.Ctx, query *tQueryFilter) (*tQueryFilter, error) {
	return query, nil
	// Err_Test.Wrap(errors.New("TestStack"))
}

func Test_Func(t *testing.T) {
	ErrProc = defaultErrProc
	DataProc = defaultDataProc
	api := &testApi{}

	app := fiber.New()
	app.Get("/", Func(api.Func))
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	rep, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rep.StatusCode)
	body, _ := ioutil.ReadAll(rep.Body)
	t.Logf("BODY: %s", body)

	var resp JSON
	err = json.Unmarshal(body, &resp)
	assert.NoError(t, err)
	assert.Equal(t, 1001, resp.Code)
	assert.Equal(t, "TestErr", resp.Message)
	assert.NotNil(t, resp.ErrorStack)
}

func Test_Func_P1(t *testing.T) {
	ErrProc = defaultErrProc
	DataProc = defaultDataProc
	api := &testApi{}

	ErrParamNotExist := NewBusError(http.StatusBadRequest, 10001, "%d 参数不存在")

	app := fiber.New()
	app.Get("/step1/:uid", Func1(api.FuncP1, Integer[int]("uidd", ErrParamNotExist)))
	app.Get("/step2/:uid", Func1(api.FuncP1, Integer[int]("uid", ErrParamNotExist)))

	req := httptest.NewRequest(http.MethodGet, "/step1/100", nil)
	rep, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rep.StatusCode)

	req = httptest.NewRequest(http.MethodGet, "/step2/100", nil)
	rep, err = app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rep.StatusCode)
}

func Test_DataFunc(t *testing.T) {
	ErrProc = defaultErrProc
	DataProc = defaultDataProc
	api := &testApi{}

	app := fiber.New()
	app.Get("/", DataFunc(api.Data))
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	rep, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rep.StatusCode)
	body, _ := ioutil.ReadAll(rep.Body)

	t.Logf("BODY: %s", body)
}

func Test_DataFunc_P1(t *testing.T) {
	ErrProc = defaultErrProc
	DataProc = defaultDataProc
	api := &testApi{}

	ErrParamNotExist := NewBusError(http.StatusBadRequest, 10001, "%s 参数不存在")

	app := fiber.New()
	app.Get("/step1/:uid", DataFunc1(api.DataP1, Integer[int]("uidd", ErrParamNotExist)))
	app.Get("/step2/:uid", DataFunc1(api.DataP1, Integer[int]("uid", ErrParamNotExist)))

	req := httptest.NewRequest(http.MethodGet, "/step1/100", nil)
	rep, err := app.Test(req)
	body, _ := ioutil.ReadAll(rep.Body)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rep.StatusCode)
	t.Logf("resp body: %s\n\n", body)

	// step2 test
	req = httptest.NewRequest(http.MethodGet, "/step2/100", nil)
	rep, err = app.Test(req)
	body, _ = ioutil.ReadAll(rep.Body)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rep.StatusCode)
	t.Logf("resp body: %s\n\n", body)

	var resp JSON
	err = json.Unmarshal(body, &resp)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Code)

	b, err := json.Marshal(resp.Error)
	assert.NoError(t, err)

	var data respUser
	err = json.Unmarshal(b, &data)
	assert.NoError(t, err)
	assert.Equal(t, 100, data.ID)
}

func Test_DataFunc_P2(t *testing.T) {
	ErrProc = defaultErrProc
	DataProc = defaultDataProc
	api := &testApi{}

	ErrParamNotExist := NewBusError(http.StatusBadRequest, 10001, "%s 参数不存在")

	app := fiber.New()
	app.Get("/:uid/:name", DataFunc2(api.DataP2, Integer[int]("uid", ErrParamNotExist), String("name", ErrParamNotExist)))

	req := httptest.NewRequest(http.MethodGet, "/100/ZhangSan", nil)
	rep, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rep.StatusCode)
	body, _ := ioutil.ReadAll(rep.Body)

	t.Logf("resp body: %s\n\n", body)
	assert.Equal(t, http.StatusOK, rep.StatusCode)

	var resp JSON
	err = json.Unmarshal(body, &resp)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Code)

	b, err := json.Marshal(resp.Error)
	assert.NoError(t, err)

	var data respUser
	err = json.Unmarshal(b, &data)
	assert.NoError(t, err)
	assert.Equal(t, 100, data.ID)
	assert.Equal(t, "ZhangSan", data.Name)
}

func Test_DataFunc_P1_POST(t *testing.T) {
	ErrProc = defaultErrProc
	DataProc = defaultDataProc
	api := &testApi{}

	ErrBindParam := NewBusError(http.StatusBadRequest, 10001, "参数绑定失败")

	app := fiber.New()
	app.Post("/user", DataFunc1(api.DataP1Form, Body[respUser](ErrBindParam)))
	app.Post("/user2", DataFunc1(api.DataP1Form, Body[respUser](ErrBindParam)))

	form := &url.Values{}
	form.Add("name", "TestName")
	form.Add("id", "100")
	t.Logf("body: %s", form.Encode())

	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rep, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rep.StatusCode)
	body, _ := ioutil.ReadAll(rep.Body)

	var resp JSON
	err = json.Unmarshal(body, &resp)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Code)

	b, err := json.Marshal(resp.Error)
	assert.NoError(t, err)

	var data respUser
	err = json.Unmarshal(b, &data)
	assert.NoError(t, err)
	assert.Equal(t, 100, data.ID)
	assert.Equal(t, "TestName", data.Name)

	t.Log("step 2::::")
	// step2 test
	b, err = json.Marshal(respUser{ID: 100, Name: "TestName"})
	assert.NoError(t, err)

	req = httptest.NewRequest(http.MethodPost, "/user2", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	rep, err = app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rep.StatusCode)
	body, _ = ioutil.ReadAll(rep.Body)

	t.Logf("resp body: %s\n\n", body)

	err = json.Unmarshal(body, &resp)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Code)

	b, err = json.Marshal(resp.Error)
	assert.NoError(t, err)

	err = json.Unmarshal(b, &data)
	assert.NoError(t, err)
	assert.Equal(t, 100, data.ID)
}

func Test_DataFunc_P1_Query(t *testing.T) {
	ErrProc = defaultErrProc
	DataProc = defaultDataProc
	api := &testApi{}

	ErrBindParam := NewBusError(http.StatusBadRequest, 10001, "参数绑定失败")

	app := fiber.New()
	app.Get("/user", DataFunc1(api.DataP1Query, Query[tQueryFilter](ErrBindParam)))

	query := &url.Values{}
	query.Add("username", "TestUsernameStep1")
	query.Add("password", "TestPasswordStep1")
	t.Logf("query: %s", query.Encode())

	req := httptest.NewRequest(http.MethodGet, "/user?"+query.Encode(), nil)

	rep, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rep.StatusCode)
	body, _ := ioutil.ReadAll(rep.Body)

	t.Logf("body: %s", body)

	var resp tQueryFilter
	err = json.Unmarshal(body, &resp)
	assert.NoError(t, err)
	assert.Equal(t, "TestUsernameStep1", resp.Username)
	assert.Equal(t, "TestPasswordStep1", resp.Password)

	query = &url.Values{}
	query.Add("page", "1")
	query.Add("limit", "10")
	req = httptest.NewRequest(http.MethodGet, "/user?"+query.Encode(), nil)

	rep, err = app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rep.StatusCode)
	body, _ = ioutil.ReadAll(rep.Body)

	t.Logf("body: %s", body)

	resp = tQueryFilter{}
	err = json.Unmarshal(body, &resp)
	assert.NoError(t, err)
	assert.Equal(t, "", resp.Username)
	assert.Equal(t, "", resp.Password)
}
