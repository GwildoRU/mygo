@echo off
go build
for %%F in (*.zip) do (
	unziprr.exe "%%F" 
)
rem del TMP\*.zip