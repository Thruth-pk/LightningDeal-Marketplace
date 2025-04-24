package repositories

import (
	"LightningDeal_Marketplace/common"
	"LightningDeal_Marketplace/datamodels"
	"database/sql"
)

type IProduct interface {
	//链接数据库
	Conn() error
	//crud
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}

type ProductManager struct {
	//表明
	table string
	//数据库链接
	mysqlConn *sql.DB
}

func (m *ProductManager) Conn() error {
	if m.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		m.mysqlConn = mysql
	}
	if m.table == "" {
		m.table = "product"
	}
	return nil
}
func (m *ProductManager) Insert(product *datamodels.Product) (int64, error) {
	//链接失败
	if err := m.Conn(); err != nil {
		return 0, err
	}
	strSql := "INSERT product SET product_name=?,product_num=?,product_img=?,product_url=?"
	stmt, err := m.mysqlConn.Prepare(strSql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
func (m *ProductManager) Delete(productID int64) bool {
	if m.Conn() != nil {
		return false
	}
	strSql := "delete from product where id=?"
	stmt, err := m.mysqlConn.Prepare(strSql)
	if err != nil {
		return false
	}
	_, err = stmt.Exec(productID)
	if err != nil {
		return false
	}
	return true
}
func (m *ProductManager) Update(product *datamodels.Product) error {
	if err := m.Conn(); err != nil {
		return err
	}
	strSql := "update product set product_name=?,product_num=?,product_img=?,product_url=? where id=?"
	stmt, err := m.mysqlConn.Prepare(strSql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl, product.ID)
	if err != nil {
		return err
	}
	return nil

}
func (m *ProductManager) SelectByKey(productID int64) (*datamodels.Product, error) {
	if err := m.Conn(); err != nil {
		return nil, err
	}
	strSql := "select * from product where id=?"
	row, err := m.mysqlConn.Query(strSql)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return nil, nil
	}
	var product *datamodels.Product
	common.DataToStructByTagSql(result, product)
	return product, nil
}
func (m *ProductManager) SelectAll() ([]*datamodels.Product, error) {
	if err := m.Conn(); err != nil {
		return nil, err
	}
	strSql := "select * from product"
	rows, err := m.mysqlConn.Query(strSql)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, nil
	}
	var products []*datamodels.Product
	for _, v := range result {
		product := &datamodels.Product{}
		common.DataToStructByTagSql(v, product)
		products = append(products, product)
	}
	return products, nil
}
func NewProductManager(table string, conn *sql.DB) IProduct {
	return &ProductManager{
		table:     table,
		mysqlConn: conn,
	}
}
