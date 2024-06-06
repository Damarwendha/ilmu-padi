package mocking

import "github.com/stretchr/testify/mock"

type MailServMock struct {
	mock.Mock
}

func (m *MailServMock) SendMail(subject, html string, to []string) error {
	args := m.Called(subject, html, to)
	return args.Error(0)
}
