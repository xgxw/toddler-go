package helpers

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/xgxw/toddler-go/internal/tests"
)

func MockDemoHelper(t *testing.T) (*DemoHelper, sqlmock.Sqlmock, func()) {
	db, mock, teardown := tests.MockDB(t)
	helper := NewDemoHelper(db)
	return helper, mock, teardown
}

func TestDemoHelperCreate(t *testing.T) {
	svc, mock, teardown := MockDemoHelper(t)
	defer teardown()
	Convey("normal", t, func() {
		name := "name"
		sql := "^INSERT INTO `demos` \\(`name`,`created_at`,`updated_at`,`deleted_at`\\) " +
			"VALUES \\(\\?,\\?,\\?,\\?\\)"
		// NewResult(lastInsertID,rowsAffected), 其他字段是gorm拼接的
		mock.ExpectExec(sql).
			WithArgs(name, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(2, 1))

		demo, err := svc.Create(name)
		tests.Print(t, fmt.Sprintf("create demo: %+v", demo))
		So(demo.ID, ShouldEqual, 2)
		So(err, ShouldBeNil)
	})

	SkipConvey("skip demo", t, func() {})
}

func TestDemoHelperGet(t *testing.T) {
	svc, mock, teardown := MockDemoHelper(t)
	defer teardown()
	Convey("normal", t, func() {
		name := "name"
		sql := "^SELECT \\* FROM `demos` WHERE `demos`.`deleted_at` IS NULL AND" +
			" \\(\\(name=\\?\\)\\) ORDER BY `demos`.`id` ASC LIMIT 1"
		// NewRows(fields).FromCSVString(values), 返回多列则使用[]Rows
		mock.ExpectQuery(sql).WithArgs(name).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).FromCSVString("6"))

		demo, err := svc.Get(name)
		tests.Print(t, fmt.Sprintf("get demo: %+v", demo))
		So(demo.ID, ShouldEqual, 6)
		So(err, ShouldBeNil)
	})
}

func TestDemoHelperUpdate(t *testing.T) {
	svc, mock, teardown := MockDemoHelper(t)
	defer teardown()
	Convey("normal", t, func() {
		id := 1
		name := "name"
		sql := "UPDATE `demos` SET `name` = \\?, `updated_at` = \\? WHERE " +
			"`demos`.`deleted_at` IS NULL AND \\(\\(id=\\?\\)\\)"
		mock.ExpectExec(sql).
			WithArgs(name, sqlmock.AnyArg(), id).
			WillReturnResult(sqlmock.NewResult(2, 1))

		err := svc.Update(id, name)
		So(err, ShouldBeNil)
	})
}

func TestDemoHelperDelete(t *testing.T) {
	svc, mock, teardown := MockDemoHelper(t)
	defer teardown()
	Convey("normal", t, func() {
		name := "name"
		sql := "UPDATE `demos` SET `deleted_at`=\\? WHERE `demos`.`deleted_at` IS NULL " +
			"AND \\(\\(name=\\?\\)\\)"
		mock.ExpectExec(sql).
			WithArgs(sqlmock.AnyArg(), name).
			WillReturnResult(sqlmock.NewResult(2, 1))

		err := svc.Delete(name)
		So(err, ShouldBeNil)
	})
}
