// Code generated by 'freedom new-crud'
package repository

import (
	"github.com/8treenet/freedom"
	"github.com/8treenet/freedom/example/fshop/application/object"
	"github.com/jinzhu/gorm"
	"time"
)

func errorLog(repo freedom.GORMRepository, model, method string, e error, expression ...interface{}) {
	if e == nil || e == gorm.ErrRecordNotFound {
		return
	}
	repo.GetRuntime().Logger().Errorf("Orm error, model: %s, method: %s, expression :%v, reason for error:%v", model, method, expression, e)
}

// findOrderDetailByPrimary .
func findOrderDetailByPrimary(repo freedom.GORMRepository, result interface{}, primary interface{}) (e error) {
	now := time.Now()
	e = repo.DB().Find(result, primary).Error
	freedom.Prometheus().OrmWithLabelValues("OrderDetail", "findOrderDetailByPrimary", e, now)
	errorLog(repo, "OrderDetail", "findOrderDetailByPrimary", e, primary)
	return
}

// findOrderDetailsByPrimarys .
func findOrderDetailsByPrimarys(repo freedom.GORMRepository, results interface{}, primarys ...interface{}) (e error) {
	now := time.Now()
	e = repo.DB().Find(results, primarys).Error
	freedom.Prometheus().OrmWithLabelValues("OrderDetail", "findOrderDetailsByPrimarys", e, now)
	errorLog(repo, "OrderDetail", "findOrderDetailsByPrimarys", e, primarys)
	return
}

// findOrderDetail .
func findOrderDetail(repo freedom.GORMRepository, query object.OrderDetail, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("OrderDetail", "findOrderDetail", e, now)
		errorLog(repo, "OrderDetail", "findOrderDetail", e, query)
	}()
	db := repo.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findOrderDetailByWhere .
func findOrderDetailByWhere(repo freedom.GORMRepository, query string, args []interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("OrderDetail", "findOrderDetailByWhere", e, now)
		errorLog(repo, "OrderDetail", "findOrderDetailByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findOrderDetailByMap .
func findOrderDetailByMap(repo freedom.GORMRepository, query map[string]interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("OrderDetail", "findOrderDetailByMap", e, now)
		errorLog(repo, "OrderDetail", "findOrderDetailByMap", e, query)
	}()

	db := repo.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findOrderDetails .
func findOrderDetails(repo freedom.GORMRepository, query object.OrderDetail, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("OrderDetail", "findOrderDetails", e, now)
		errorLog(repo, "OrderDetail", "findOrderDetails", e, query)
	}()
	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findOrderDetailsByWhere .
func findOrderDetailsByWhere(repo freedom.GORMRepository, query string, args []interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("OrderDetail", "findOrderDetailsByWhere", e, now)
		errorLog(repo, "OrderDetail", "findOrderDetailsByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findOrderDetailsByMap .
func findOrderDetailsByMap(repo freedom.GORMRepository, query map[string]interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("OrderDetail", "findOrderDetailsByMap", e, now)
		errorLog(repo, "OrderDetail", "findOrderDetailsByMap", e, query)
	}()

	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// createOrderDetail .
func createOrderDetail(repo freedom.GORMRepository, object *object.OrderDetail) (rowsAffected int64, e error) {
	now := time.Now()
	db := repo.DB().Create(object)
	rowsAffected = db.RowsAffected
	e = db.Error
	freedom.Prometheus().OrmWithLabelValues("OrderDetail", "createOrderDetail", e, now)
	errorLog(repo, "OrderDetail", "createOrderDetail", e, *object)
	return
}

// saveOrderDetail .
func saveOrderDetail(repo freedom.GORMRepository, object *object.OrderDetail) (affected int64, e error) {
	now := time.Now()
	db := repo.DB().Model(object).Updates(object.TakeChanges())
	e = db.Error
	affected = db.RowsAffected
	freedom.Prometheus().OrmWithLabelValues("OrderDetail", "saveOrderDetail", e, now)
	errorLog(repo, "OrderDetail", "saveOrderDetail", e, *object)
	return
}

// findUserByPrimary .
func findUserByPrimary(repo freedom.GORMRepository, result interface{}, primary interface{}) (e error) {
	now := time.Now()
	e = repo.DB().Find(result, primary).Error
	freedom.Prometheus().OrmWithLabelValues("User", "findUserByPrimary", e, now)
	errorLog(repo, "User", "findUserByPrimary", e, primary)
	return
}

// findUsersByPrimarys .
func findUsersByPrimarys(repo freedom.GORMRepository, results interface{}, primarys ...interface{}) (e error) {
	now := time.Now()
	e = repo.DB().Find(results, primarys).Error
	freedom.Prometheus().OrmWithLabelValues("User", "findUsersByPrimarys", e, now)
	errorLog(repo, "User", "findUsersByPrimarys", e, primarys)
	return
}

// findUser .
func findUser(repo freedom.GORMRepository, query object.User, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("User", "findUser", e, now)
		errorLog(repo, "User", "findUser", e, query)
	}()
	db := repo.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findUserByWhere .
func findUserByWhere(repo freedom.GORMRepository, query string, args []interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("User", "findUserByWhere", e, now)
		errorLog(repo, "User", "findUserByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findUserByMap .
func findUserByMap(repo freedom.GORMRepository, query map[string]interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("User", "findUserByMap", e, now)
		errorLog(repo, "User", "findUserByMap", e, query)
	}()

	db := repo.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findUsers .
func findUsers(repo freedom.GORMRepository, query object.User, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("User", "findUsers", e, now)
		errorLog(repo, "User", "findUsers", e, query)
	}()
	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findUsersByWhere .
func findUsersByWhere(repo freedom.GORMRepository, query string, args []interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("User", "findUsersByWhere", e, now)
		errorLog(repo, "User", "findUsersByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findUsersByMap .
func findUsersByMap(repo freedom.GORMRepository, query map[string]interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("User", "findUsersByMap", e, now)
		errorLog(repo, "User", "findUsersByMap", e, query)
	}()

	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// createUser .
func createUser(repo freedom.GORMRepository, object *object.User) (rowsAffected int64, e error) {
	now := time.Now()
	db := repo.DB().Create(object)
	rowsAffected = db.RowsAffected
	e = db.Error
	freedom.Prometheus().OrmWithLabelValues("User", "createUser", e, now)
	errorLog(repo, "User", "createUser", e, *object)
	return
}

// saveUser .
func saveUser(repo freedom.GORMRepository, object *object.User) (affected int64, e error) {
	now := time.Now()
	db := repo.DB().Model(object).Updates(object.TakeChanges())
	e = db.Error
	affected = db.RowsAffected
	freedom.Prometheus().OrmWithLabelValues("User", "saveUser", e, now)
	errorLog(repo, "User", "saveUser", e, *object)
	return
}

// findAdminByPrimary .
func findAdminByPrimary(repo freedom.GORMRepository, result interface{}, primary interface{}) (e error) {
	now := time.Now()
	e = repo.DB().Find(result, primary).Error
	freedom.Prometheus().OrmWithLabelValues("Admin", "findAdminByPrimary", e, now)
	errorLog(repo, "Admin", "findAdminByPrimary", e, primary)
	return
}

// findAdminsByPrimarys .
func findAdminsByPrimarys(repo freedom.GORMRepository, results interface{}, primarys ...interface{}) (e error) {
	now := time.Now()
	e = repo.DB().Find(results, primarys).Error
	freedom.Prometheus().OrmWithLabelValues("Admin", "findAdminsByPrimarys", e, now)
	errorLog(repo, "Admin", "findAdminsByPrimarys", e, primarys)
	return
}

// findAdmin .
func findAdmin(repo freedom.GORMRepository, query object.Admin, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Admin", "findAdmin", e, now)
		errorLog(repo, "Admin", "findAdmin", e, query)
	}()
	db := repo.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findAdminByWhere .
func findAdminByWhere(repo freedom.GORMRepository, query string, args []interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Admin", "findAdminByWhere", e, now)
		errorLog(repo, "Admin", "findAdminByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findAdminByMap .
func findAdminByMap(repo freedom.GORMRepository, query map[string]interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Admin", "findAdminByMap", e, now)
		errorLog(repo, "Admin", "findAdminByMap", e, query)
	}()

	db := repo.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findAdmins .
func findAdmins(repo freedom.GORMRepository, query object.Admin, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Admin", "findAdmins", e, now)
		errorLog(repo, "Admin", "findAdmins", e, query)
	}()
	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findAdminsByWhere .
func findAdminsByWhere(repo freedom.GORMRepository, query string, args []interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Admin", "findAdminsByWhere", e, now)
		errorLog(repo, "Admin", "findAdminsByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findAdminsByMap .
func findAdminsByMap(repo freedom.GORMRepository, query map[string]interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Admin", "findAdminsByMap", e, now)
		errorLog(repo, "Admin", "findAdminsByMap", e, query)
	}()

	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// createAdmin .
func createAdmin(repo freedom.GORMRepository, object *object.Admin) (rowsAffected int64, e error) {
	now := time.Now()
	db := repo.DB().Create(object)
	rowsAffected = db.RowsAffected
	e = db.Error
	freedom.Prometheus().OrmWithLabelValues("Admin", "createAdmin", e, now)
	errorLog(repo, "Admin", "createAdmin", e, *object)
	return
}

// saveAdmin .
func saveAdmin(repo freedom.GORMRepository, object *object.Admin) (affected int64, e error) {
	now := time.Now()
	db := repo.DB().Model(object).Updates(object.TakeChanges())
	e = db.Error
	affected = db.RowsAffected
	freedom.Prometheus().OrmWithLabelValues("Admin", "saveAdmin", e, now)
	errorLog(repo, "Admin", "saveAdmin", e, *object)
	return
}

// findCartByPrimary .
func findCartByPrimary(repo freedom.GORMRepository, result interface{}, primary interface{}) (e error) {
	now := time.Now()
	e = repo.DB().Find(result, primary).Error
	freedom.Prometheus().OrmWithLabelValues("Cart", "findCartByPrimary", e, now)
	errorLog(repo, "Cart", "findCartByPrimary", e, primary)
	return
}

// findCartsByPrimarys .
func findCartsByPrimarys(repo freedom.GORMRepository, results interface{}, primarys ...interface{}) (e error) {
	now := time.Now()
	e = repo.DB().Find(results, primarys).Error
	freedom.Prometheus().OrmWithLabelValues("Cart", "findCartsByPrimarys", e, now)
	errorLog(repo, "Cart", "findCartsByPrimarys", e, primarys)
	return
}

// findCart .
func findCart(repo freedom.GORMRepository, query object.Cart, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Cart", "findCart", e, now)
		errorLog(repo, "Cart", "findCart", e, query)
	}()
	db := repo.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findCartByWhere .
func findCartByWhere(repo freedom.GORMRepository, query string, args []interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Cart", "findCartByWhere", e, now)
		errorLog(repo, "Cart", "findCartByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findCartByMap .
func findCartByMap(repo freedom.GORMRepository, query map[string]interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Cart", "findCartByMap", e, now)
		errorLog(repo, "Cart", "findCartByMap", e, query)
	}()

	db := repo.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findCarts .
func findCarts(repo freedom.GORMRepository, query object.Cart, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Cart", "findCarts", e, now)
		errorLog(repo, "Cart", "findCarts", e, query)
	}()
	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findCartsByWhere .
func findCartsByWhere(repo freedom.GORMRepository, query string, args []interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Cart", "findCartsByWhere", e, now)
		errorLog(repo, "Cart", "findCartsByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findCartsByMap .
func findCartsByMap(repo freedom.GORMRepository, query map[string]interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Cart", "findCartsByMap", e, now)
		errorLog(repo, "Cart", "findCartsByMap", e, query)
	}()

	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// createCart .
func createCart(repo freedom.GORMRepository, object *object.Cart) (rowsAffected int64, e error) {
	now := time.Now()
	db := repo.DB().Create(object)
	rowsAffected = db.RowsAffected
	e = db.Error
	freedom.Prometheus().OrmWithLabelValues("Cart", "createCart", e, now)
	errorLog(repo, "Cart", "createCart", e, *object)
	return
}

// saveCart .
func saveCart(repo freedom.GORMRepository, object *object.Cart) (affected int64, e error) {
	now := time.Now()
	db := repo.DB().Model(object).Updates(object.TakeChanges())
	e = db.Error
	affected = db.RowsAffected
	freedom.Prometheus().OrmWithLabelValues("Cart", "saveCart", e, now)
	errorLog(repo, "Cart", "saveCart", e, *object)
	return
}

// findDeliveryByPrimary .
func findDeliveryByPrimary(repo freedom.GORMRepository, result interface{}, primary interface{}) (e error) {
	now := time.Now()
	e = repo.DB().Find(result, primary).Error
	freedom.Prometheus().OrmWithLabelValues("Delivery", "findDeliveryByPrimary", e, now)
	errorLog(repo, "Delivery", "findDeliveryByPrimary", e, primary)
	return
}

// findDeliverysByPrimarys .
func findDeliverysByPrimarys(repo freedom.GORMRepository, results interface{}, primarys ...interface{}) (e error) {
	now := time.Now()
	e = repo.DB().Find(results, primarys).Error
	freedom.Prometheus().OrmWithLabelValues("Delivery", "findDeliverysByPrimarys", e, now)
	errorLog(repo, "Delivery", "findDeliverysByPrimarys", e, primarys)
	return
}

// findDelivery .
func findDelivery(repo freedom.GORMRepository, query object.Delivery, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Delivery", "findDelivery", e, now)
		errorLog(repo, "Delivery", "findDelivery", e, query)
	}()
	db := repo.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findDeliveryByWhere .
func findDeliveryByWhere(repo freedom.GORMRepository, query string, args []interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Delivery", "findDeliveryByWhere", e, now)
		errorLog(repo, "Delivery", "findDeliveryByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findDeliveryByMap .
func findDeliveryByMap(repo freedom.GORMRepository, query map[string]interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Delivery", "findDeliveryByMap", e, now)
		errorLog(repo, "Delivery", "findDeliveryByMap", e, query)
	}()

	db := repo.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findDeliverys .
func findDeliverys(repo freedom.GORMRepository, query object.Delivery, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Delivery", "findDeliverys", e, now)
		errorLog(repo, "Delivery", "findDeliverys", e, query)
	}()
	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findDeliverysByWhere .
func findDeliverysByWhere(repo freedom.GORMRepository, query string, args []interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Delivery", "findDeliverysByWhere", e, now)
		errorLog(repo, "Delivery", "findDeliverysByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findDeliverysByMap .
func findDeliverysByMap(repo freedom.GORMRepository, query map[string]interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Delivery", "findDeliverysByMap", e, now)
		errorLog(repo, "Delivery", "findDeliverysByMap", e, query)
	}()

	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// createDelivery .
func createDelivery(repo freedom.GORMRepository, object *object.Delivery) (rowsAffected int64, e error) {
	now := time.Now()
	db := repo.DB().Create(object)
	rowsAffected = db.RowsAffected
	e = db.Error
	freedom.Prometheus().OrmWithLabelValues("Delivery", "createDelivery", e, now)
	errorLog(repo, "Delivery", "createDelivery", e, *object)
	return
}

// saveDelivery .
func saveDelivery(repo freedom.GORMRepository, object *object.Delivery) (affected int64, e error) {
	now := time.Now()
	db := repo.DB().Model(object).Updates(object.TakeChanges())
	e = db.Error
	affected = db.RowsAffected
	freedom.Prometheus().OrmWithLabelValues("Delivery", "saveDelivery", e, now)
	errorLog(repo, "Delivery", "saveDelivery", e, *object)
	return
}

// findGoodsByPrimary .
func findGoodsByPrimary(repo freedom.GORMRepository, result interface{}, primary interface{}) (e error) {
	now := time.Now()
	e = repo.DB().Find(result, primary).Error
	freedom.Prometheus().OrmWithLabelValues("Goods", "findGoodsByPrimary", e, now)
	errorLog(repo, "Goods", "findGoodsByPrimary", e, primary)
	return
}

// findGoodssByPrimarys .
func findGoodssByPrimarys(repo freedom.GORMRepository, results interface{}, primarys ...interface{}) (e error) {
	now := time.Now()
	e = repo.DB().Find(results, primarys).Error
	freedom.Prometheus().OrmWithLabelValues("Goods", "findGoodssByPrimarys", e, now)
	errorLog(repo, "Goods", "findGoodssByPrimarys", e, primarys)
	return
}

// findGoods .
func findGoods(repo freedom.GORMRepository, query object.Goods, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Goods", "findGoods", e, now)
		errorLog(repo, "Goods", "findGoods", e, query)
	}()
	db := repo.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findGoodsByWhere .
func findGoodsByWhere(repo freedom.GORMRepository, query string, args []interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Goods", "findGoodsByWhere", e, now)
		errorLog(repo, "Goods", "findGoodsByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findGoodsByMap .
func findGoodsByMap(repo freedom.GORMRepository, query map[string]interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Goods", "findGoodsByMap", e, now)
		errorLog(repo, "Goods", "findGoodsByMap", e, query)
	}()

	db := repo.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findGoodss .
func findGoodss(repo freedom.GORMRepository, query object.Goods, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Goods", "findGoodss", e, now)
		errorLog(repo, "Goods", "findGoodss", e, query)
	}()
	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findGoodssByWhere .
func findGoodssByWhere(repo freedom.GORMRepository, query string, args []interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Goods", "findGoodssByWhere", e, now)
		errorLog(repo, "Goods", "findGoodssByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findGoodssByMap .
func findGoodssByMap(repo freedom.GORMRepository, query map[string]interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Goods", "findGoodssByMap", e, now)
		errorLog(repo, "Goods", "findGoodssByMap", e, query)
	}()

	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// createGoods .
func createGoods(repo freedom.GORMRepository, object *object.Goods) (rowsAffected int64, e error) {
	now := time.Now()
	db := repo.DB().Create(object)
	rowsAffected = db.RowsAffected
	e = db.Error
	freedom.Prometheus().OrmWithLabelValues("Goods", "createGoods", e, now)
	errorLog(repo, "Goods", "createGoods", e, *object)
	return
}

// saveGoods .
func saveGoods(repo freedom.GORMRepository, object *object.Goods) (affected int64, e error) {
	now := time.Now()
	db := repo.DB().Model(object).Updates(object.TakeChanges())
	e = db.Error
	affected = db.RowsAffected
	freedom.Prometheus().OrmWithLabelValues("Goods", "saveGoods", e, now)
	errorLog(repo, "Goods", "saveGoods", e, *object)
	return
}

// findOrderByPrimary .
func findOrderByPrimary(repo freedom.GORMRepository, result interface{}, primary interface{}) (e error) {
	now := time.Now()
	e = repo.DB().Find(result, primary).Error
	freedom.Prometheus().OrmWithLabelValues("Order", "findOrderByPrimary", e, now)
	errorLog(repo, "Order", "findOrderByPrimary", e, primary)
	return
}

// findOrdersByPrimarys .
func findOrdersByPrimarys(repo freedom.GORMRepository, results interface{}, primarys ...interface{}) (e error) {
	now := time.Now()
	e = repo.DB().Find(results, primarys).Error
	freedom.Prometheus().OrmWithLabelValues("Order", "findOrdersByPrimarys", e, now)
	errorLog(repo, "Order", "findOrdersByPrimarys", e, primarys)
	return
}

// findOrder .
func findOrder(repo freedom.GORMRepository, query object.Order, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Order", "findOrder", e, now)
		errorLog(repo, "Order", "findOrder", e, query)
	}()
	db := repo.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findOrderByWhere .
func findOrderByWhere(repo freedom.GORMRepository, query string, args []interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Order", "findOrderByWhere", e, now)
		errorLog(repo, "Order", "findOrderByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findOrderByMap .
func findOrderByMap(repo freedom.GORMRepository, query map[string]interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Order", "findOrderByMap", e, now)
		errorLog(repo, "Order", "findOrderByMap", e, query)
	}()

	db := repo.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findOrders .
func findOrders(repo freedom.GORMRepository, query object.Order, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Order", "findOrders", e, now)
		errorLog(repo, "Order", "findOrders", e, query)
	}()
	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findOrdersByWhere .
func findOrdersByWhere(repo freedom.GORMRepository, query string, args []interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Order", "findOrdersByWhere", e, now)
		errorLog(repo, "Order", "findOrdersByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findOrdersByMap .
func findOrdersByMap(repo freedom.GORMRepository, query map[string]interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Order", "findOrdersByMap", e, now)
		errorLog(repo, "Order", "findOrdersByMap", e, query)
	}()

	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// createOrder .
func createOrder(repo freedom.GORMRepository, object *object.Order) (rowsAffected int64, e error) {
	now := time.Now()
	db := repo.DB().Create(object)
	rowsAffected = db.RowsAffected
	e = db.Error
	freedom.Prometheus().OrmWithLabelValues("Order", "createOrder", e, now)
	errorLog(repo, "Order", "createOrder", e, *object)
	return
}

// saveOrder .
func saveOrder(repo freedom.GORMRepository, object *object.Order) (affected int64, e error) {
	now := time.Now()
	db := repo.DB().Model(object).Updates(object.TakeChanges())
	e = db.Error
	affected = db.RowsAffected
	freedom.Prometheus().OrmWithLabelValues("Order", "saveOrder", e, now)
	errorLog(repo, "Order", "saveOrder", e, *object)
	return
}
