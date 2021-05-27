module github.com/northcity0406/delayTask

go 1.16

replace github.com/northcity0406/delayTask/tasks => ./tasks

replace github.com/northcity0406/delayTask/DBHandler => ./DBHandler
replace github.com/northcity0406/delayTask/TaskModel => ./TaskModel


require github.com/northcity0406/delayTask/tasks v0.0.0-00010101000000-000000000000
require github.com/northcity0406/delayTask/DBHandler v0.0.0-00010101000000-000000000000
require github.com/northcity0406/delayTask/TaskModel v0.0.0-00010101000000-000000000000
