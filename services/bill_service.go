package services

import (
	"context"
	"database/sql"
	"log"
	"payR/models"
)

func GetBillsByCustomerID(id int, client *sql.DB) (models.Bill, error) {
	var billsofCustomer models.Bill

	sqlQuery :=
		`
	SELECT 
		id,
		customer_id,
		account_id, 
		mobile,
		bill_type,
		equipment_count,
		amount,
		payment_method,
		submit_date
	FROM 
		Bill
	WHERE 
		customer_id = $1
	`

	row := client.QueryRow(sqlQuery, id)

	err := row.Scan(
		&billsofCustomer.ID,
		&billsofCustomer.CustomerID,
		&billsofCustomer.AccountID,
		&billsofCustomer.Mobile,
		&billsofCustomer.BillType,
		&billsofCustomer.EquipmentCount,
		&billsofCustomer.Amount,
		&billsofCustomer.PaymentMethod,
		&billsofCustomer.SubmitDate,
	)

	switch err {
	case sql.ErrNoRows:
		return billsofCustomer, nil
	case nil:
		return billsofCustomer, nil
	default:
		panic(err)
	}
}

func SubmitBill(bill models.Bill, client *sql.DB) {
	ctx := context.Background()

	tx, err := client.BeginTx(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	billPayQuery := `INSERT INTO bill (customer_id, 
		mobile, 
		bill_type, 
		equipment_count, 
		amount, 
		account_id, 
		payment_method, 
		submit_date) 
		VALUES($1,$2,$3,$4,$5,$6,$7,$8)
	`
	_, err = tx.ExecContext(ctx, billPayQuery,
		bill.CustomerID,
		bill.Mobile,
		bill.BillType,
		bill.EquipmentCount,
		bill.Amount,
		bill.AccountID,
		bill.PaymentMethod, bill.SubmitDate)

	if err != nil {
		tx.Rollback()
		return
	}

	totalBillQuery := `SELECT SUM(amount) FROM bill WHERE account_id = 1`

	totalAmount := 0.0

	err = tx.QueryRowContext(ctx, totalBillQuery).Scan(&totalAmount)

	if err != nil {
		tx.Rollback()
		return
	}

	updataAccountBalanceQuery := `UPDATE account SET total_amount = $1 WHERE id = 1`

	_, err = tx.ExecContext(ctx, updataAccountBalanceQuery, totalAmount)

	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}

}
