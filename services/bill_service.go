package services

import (
	"context"
	"database/sql"
	"log"
	"payR/models"
)

func GetBillsByCustomerID(id int, client *sql.DB) ([]models.Bill, error) {
	billsofCustomer := []models.Bill{}

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
		bill
	WHERE 
		customer_id = $1
	`
	row, err := client.Query(sqlQuery, id)

	for row.Next() {
		var billOfCustomer models.Bill

		err := row.Scan(
			&billOfCustomer.ID,
			&billOfCustomer.CustomerID,
			&billOfCustomer.AccountID,
			&billOfCustomer.Mobile,
			&billOfCustomer.BillType,
			&billOfCustomer.EquipmentCount,
			&billOfCustomer.Amount,
			&billOfCustomer.PaymentMethod,
			&billOfCustomer.SubmitDate,
		)

		if err == nil {
			billsofCustomer = append(billsofCustomer, billOfCustomer)
		}

	}

	return billsofCustomer, err
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

	log.Print(err)

	if err != nil {
		tx.Rollback()
		return
	}

	totalBillQuery := `SELECT SUM(amount) FROM bill WHERE account_id = 1`

	var totalAmount float64

	err = tx.QueryRow(totalBillQuery).Scan(&totalAmount)

	log.Print(err)

	if err != nil {
		tx.Rollback()
		return
	}

	updataAccountBalanceQuery := `UPDATE account SET total_amount = $1 WHERE id = 1`

	_, err = tx.ExecContext(ctx, updataAccountBalanceQuery, totalAmount)

	log.Print(err)

	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}

}
