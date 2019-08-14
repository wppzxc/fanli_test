@echo off
@echo Deleting preItemsFile, proItemsFile, fanli.log
del /F .\preItemsFile
del /F .\proItemsFile
del /F .\fanli.log
@echo Running the program : 
@echo fanli.exe --config .\config.yaml --alsologtostderr --log-file .\fanli.log --log-file-max-size 500 --v 9
start /b fanli.exe --config .\config.yaml --alsologtostderr --log-file .\fanli.log --log-file-max-size 500 --v 9