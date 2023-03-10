package pipeline

// Of wraps tasks as a single function
func Of(tasks ...func() error) func() error {
	return func() error {
		for _, task := range tasks {
			err := task()
			if err != nil {
				return err
			}
		}

		return nil
	}
}
