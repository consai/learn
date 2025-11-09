package main

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Account struct {
	ID      uint
	Balance int
}

type Transaction struct {
	ID              uint
	From_account_id uint
	To_account_id   uint
	Amount          int
}

func Trans(db *gorm.DB, fromid uint, toid uint, amount int) error {

	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})
	ctx := context.Background()
	// err := gorm.G[Account](db).Create(ctx, &Account{ID: 1, Balance: 100})
	// if err != nil {
	// 	fmt.Printf("Create err:\n%+v\n", err)
	// }

	// err = gorm.G[Account](db).Create(ctx, &Account{ID: 2, Balance: 0})
	// if err != nil {
	// 	fmt.Printf("Create err:\n%+v\n", err)
	// }

	err := db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		//ctx := context.Background()

		err := gorm.G[Transaction](tx).Create(ctx, &Transaction{From_account_id: fromid, To_account_id: toid, Amount: amount})
		if err != nil {
			return err
		}

		facc, err := gorm.G[Account](tx).Where("ID = ?", fromid).First(ctx)
		if err != nil {
			return err
		}
		if facc.Balance < amount {
			return fmt.Errorf("余额不足")
		}

		tacc, err := gorm.G[Account](tx).Where("ID = ?", toid).First(ctx)
		if err != nil {
			return err
		}

		_, err = gorm.G[Account](tx).Where("ID = ?", fromid).Update(ctx, "Balance", facc.Balance-amount)
		if err != nil {
			return err
		}

		_, err = gorm.G[Account](tx).Where("ID = ?", toid).Update(ctx, "Balance", tacc.Balance+amount)
		if err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
