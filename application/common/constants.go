package common

import "time"

const WEEK_DAYS = 7
const MAX_LESSONS = 9
const UserId = "user_id"
const UserRole = "user_role"
//var Domain = "http://localhost:8080/api/v1"
var Domain = "https://10-tka.ru/api/v1"
var CookiesDuration = int((2 * 24 * time.Hour).Seconds())
var PathToSaveStatic = "/home/vlad/10tka/10tka-backend/static"
