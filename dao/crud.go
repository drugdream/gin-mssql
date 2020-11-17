package dao

func Insert(obj interface{}) error {
	return DB.Create(obj).Error
}

func Update(model interface{}, where interface{}, params []interface{}, update interface{}) error {
	err := DB.Model(model).Where(where, params...).Updates(update).Error
	return err
}

func Find(sql string, params []interface{}, out interface{}) error {
	return DB.Where(sql, params...).Find(out).Error
}

func ExecSql(sql string, out interface{}) error {
	return DB.Exec(sql, nil).Scan(out).Error
}
