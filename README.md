
# RSSGator

gator project:

you will need Postgres and Go installed to run the program

to install on system, navigate to RSSGator folder and run "go install"

to set up config file, in home directory create:
~/.gatorconfig.json

containing:

type Config struct {
    DBUrl       string `json:db_url`
    CurrentUserName string`json:"current_user_name"`
}

type "RSSGator" then a command and input

example RSSGator register username - will register user name. 

commands include:
login - name : logs in registered user

register - name : registers user

reset : resets all DB

users : lists all registered users

addfeed - url : adds feed and creates following for logged in user

feeds : list all feeds that have been added by all users

follow - url : adds feed to following for logged in user

following : lists all feeds logged in user is following

unfollow - url : unfollows feed

agg - time(1s or 1m or 1h) : collects all posts from feed continuously at set intervals * ctl + c to stop

browse - num(default 2) : displays latest article links from feeds the logged in user is following. num is number of entries to display. 

