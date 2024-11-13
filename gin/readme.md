import logger first

application.InitLogger()
application.Logger.Info("ping logger")
application.SyncLogger()

CompileDaemon
CompileDaemon --build="go build -o ./bin/main ./src" --command="./bin/main"
