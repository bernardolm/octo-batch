package sift

func Start() error {
	config()

	c, err := newClient()
	if err != nil {
		return err
	}
	client = c

	return nil
}
