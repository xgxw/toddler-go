package services

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/xgxw/toddler-go"
	"github.com/xgxw/toddler-go/internal/tests"
	"github.com/xgxw/toddler-go/internal/tests/mocks"
)

type mockUTServices struct {
	svc *UtilTestService

	demoSvc *mocks.MockDemoService
}

func MockUtilTestService(t *testing.T) (*mockUTServices, sqlmock.Sqlmock, func()) {
	db, mock, teardown_db := tests.MockDB(t)
	demoCtrl := NewController(t)

	demoSvc := mocks.NewMockDemoService(demoCtrl)
	svc := NewUtilTestService(db, demoSvc)

	svcs := &mockUTServices{
		svc:     svc,
		demoSvc: demoSvc,
	}
	teardown := func() {
		// Finish 用于检查是否所有 EXCEPT 的方法都被执行了
		demoCtrl.Finish()
		teardown_db()
	}
	return svcs, mock, teardown
}

func TestUtilTestServiceDoSomething(t *testing.T) {
	/*
		每个Convey定义自己所需的资源. 原因如下:
		1. 各单元测试并发执行, 避免相互干扰.
		2. 某些资源是约定先后顺序的, 如 MockDemoService.
		3. 某些资源限制只能在一个Convey中, 如 Finish(). 如果分散到多个Convey, 则Finish无法正确检测结果.
	*/
	// normal 是单元测试名称/描述, 表示参数符合要求, 正常运行的情况.
	Convey("normal", t, func() {
		svcs, _, teardown := MockUtilTestService(t)
		defer teardown()
		svc := svcs.svc
		respMsg := ""

		// Any 表示任意参数均可, 可以根据需求更换, 如 new(toddler.Request)
		svcs.demoSvc.EXPECT().DoSomething(Any()).Return(&toddler.Response{Msg: respMsg}, nil)

		msg, err := svc.DoSomething()
		So(msg, ShouldEqual, respMsg)
		So(err, ShouldBeNil)
	})
	// 分支错误时的情况.
	SkipConvey("skip demo", t, func() {})
}

func TestUtilTestServiceCreate(t *testing.T) {
	Convey("normal", t, func() {
		svcs, mock, teardown := MockUtilTestService(t)
		defer teardown()
		svc := svcs.svc
		name := "name"
		sql := "^INSERT INTO `util_test_structs` \\(`name`,`created_at`,`updated_at`,`deleted_at`\\) " +
			"VALUES \\(\\?,\\?,\\?,\\?\\)"
		// sqlmock 的顺序必须与调用顺序相同, 否则会报错.
		// NewResult(lastInsertID,rowsAffected), 其他字段是gorm拼接的
		mock.ExpectExec(sql).
			WithArgs(name, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(2, 1))

		demo, err := svc.Create(name)
		t.Logf("create demo: %+v", demo)
		So(demo.ID, ShouldEqual, 2)
		So(err, ShouldBeNil)
	})
}

func TestUtilTestServiceGet(t *testing.T) {
	Convey("normal", t, func() {
		svcs, mock, teardown := MockUtilTestService(t)
		defer teardown()
		svc := svcs.svc
		name := "name"
		sql := "^SELECT \\* FROM `util_test_structs` WHERE `util_test_structs`.`deleted_at` IS NULL AND" +
			" \\(\\(name=\\?\\)\\) ORDER BY `util_test_structs`.`id` ASC LIMIT 1"
		// NewRows(fields).FromCSVString(values), 返回多列则使用[]Rows
		mock.ExpectQuery(sql).WithArgs(name).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).FromCSVString("6"))

		demo, err := svc.Get(name)
		t.Logf("get demo: %+v", demo)
		So(demo.ID, ShouldEqual, 6)
		So(err, ShouldBeNil)
	})
}

func TestUtilTestServiceUpdate(t *testing.T) {
	Convey("normal", t, func() {
		svcs, mock, teardown := MockUtilTestService(t)
		defer teardown()
		svc := svcs.svc
		id := 1
		name := "name"
		sql := "UPDATE `util_test_structs` SET `name` = \\?, `updated_at` = \\? WHERE " +
			"`util_test_structs`.`deleted_at` IS NULL AND \\(\\(id=\\?\\)\\)"
		mock.ExpectExec(sql).
			WithArgs(name, sqlmock.AnyArg(), id).
			WillReturnResult(sqlmock.NewResult(2, 1))

		err := svc.Update(id, name)
		So(err, ShouldBeNil)
	})
}

func TestUtilTestServiceDelete(t *testing.T) {
	Convey("normal", t, func() {
		svcs, mock, teardown := MockUtilTestService(t)
		defer teardown()
		svc := svcs.svc
		name := "name"
		sql := "UPDATE `util_test_structs` SET `deleted_at`=\\? WHERE `util_test_structs`.`deleted_at` IS NULL " +
			"AND \\(\\(name=\\?\\)\\)"
		mock.ExpectExec(sql).
			WithArgs(sqlmock.AnyArg(), name).
			WillReturnResult(sqlmock.NewResult(2, 1))

		err := svc.Delete(name)
		So(err, ShouldBeNil)
	})
}
