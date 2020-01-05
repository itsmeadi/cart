package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/itsme/nuc/src/entities/config"
	sql "github.com/jmoiron/sqlx"
	//"database/sql"
	"log"
)

var (
	createCart     = `insert into NTUC.cart (user_id, status) values(?,?)`
	getCart        = `select id, user_id ,status from NTUC.cart where user_id=?`
	insertItem     = `insert into NTUC.cart_items(cart_id, product_id, qty,status) values(?,?,?,?)`
	getItemsInCart = `select id, cart_id, product_id, qty,status from NTUC.cart_items where cart_id=? and status=? 
							order by updated_at desc limit ?`
	getProductFromCart   = `select id, cart_id, product_id, qty,status from NTUC.cart_items where cart_id=? and product_id=? and status=? `
	updateCartQty        = `update NTUC.cart_items set qty=? where cart_id=? and product_id=? limit 1`
	updateCartItemStatus = `update NTUC.cart_items set status=? where cart_id=? and product_id=? limit 1`

	getCartItemsByUser = `select c.id as cart_id, c.user_id as user_id, i.product_id as product_id, i.qty as qty from cart as c, cart_items as i 
							where c.user_id=? and c.id=i.cart_id and c.status=? and i.status=?`
	//getCartItemsByUserAndProductId = `select c.id as cart_id, c.user_id as user_id, i.product_id as product_id, i.qty as qty from cart as c, cart_items as i left join
	//						where c.user_id=? and c.id=i.cart_id and i.product_id=? and c.status=? and i.status=?`
	getCartItemsByUserAndProductId = `select c.id as cart_id, c.user_id as user_id, i.product_id as product_id, i.qty as qty from cart as c left join cart_items as i
							on c.id=i.cart_id where c.user_id=? and i.product_id=? and c.status=? and i.status=?`

	getUserBySub = `select id from NTUC.user where sub=?`
	addUser      = `insert into NTUC.user (sub) values(?)`
)

type PreparedQueries struct {
	CreateCart *sql.Stmt
	GetCart    *sql.Stmt

	InsertItem           *sql.Stmt
	GetItemsInCart       *sql.Stmt
	GetProductFromCart   *sql.Stmt
	UpdateCartQty        *sql.Stmt
	UpdateCartItemStatus *sql.Stmt

	GetCartItemsByUser             *sql.Stmt
	GetCartItemsByUserAndProductId *sql.Stmt

	GetUserBySub *sql.Stmt
	AddUser      *sql.Stmt
}

type DB struct {
	sqlConn *sql.DB
	queries *PreparedQueries
}

var conn *DB

func GetDb() *DB {
	if conn == nil {
		log.Fatal("Db Uninitialized")
	}
	return conn
}
func InitDb(dbConf config.DBConfig) {
	var err error
	conn = &DB{}

	conn.sqlConn, err = sql.Open("mysql", dbConf.DBStr) //TODO panics
	if err != nil {
		log.Fatal("Cannot init mysql err=", err)
	}
	conn.initQueries()

}
func (Conn *DB) initQueries() {
	Conn.queries = &PreparedQueries{
		CreateCart: Conn.Prepare(createCart),
		GetCart:    Conn.Prepare(getCart),
		InsertItem: Conn.Prepare(insertItem),

		GetItemsInCart:       Conn.Prepare(getItemsInCart),
		GetProductFromCart:   Conn.Prepare(getProductFromCart),
		UpdateCartQty:        Conn.Prepare(updateCartQty),
		UpdateCartItemStatus: Conn.Prepare(updateCartItemStatus),

		GetCartItemsByUser:             Conn.Prepare(getCartItemsByUser),
		GetCartItemsByUserAndProductId: Conn.Prepare(getCartItemsByUserAndProductId),

		GetUserBySub: Conn.Prepare(getUserBySub),
		AddUser:      Conn.Prepare(addUser),
	}

}

func (Conn *DB) Prepare(query string) *sql.Stmt {

	prep, err := Conn.sqlConn.Preparex(query)
	if err != nil {
		log.Fatalf("Cannot Prepare Query=%+v, err=%+v", query, err)
	}
	return prep
}
