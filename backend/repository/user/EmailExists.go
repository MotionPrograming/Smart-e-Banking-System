package user

func (r *userRepository) EmailExists(email string) (bool, error) {
	query := `SELECT COUNT(1) FROM users WHERE email = ?`

	var count int64
	if err := r.db.Get(&count, query, email); err != nil {
		return false, err
	}
	return count > 0, nil
}
