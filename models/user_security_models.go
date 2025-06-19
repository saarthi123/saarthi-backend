package models

type UserSecurity struct {
    UserID           string
    PIN              string
    TwoFactorEnabled  bool
    BiometricEnabled  bool
    Alerts           map[string]bool
}
