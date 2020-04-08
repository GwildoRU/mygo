@echo off
rem type qqq.txt | sandbox.exe >> aaa.txt
for %%f in (q*.txt) do (
echo ! > aaa.txt
for /F %%t in (%%f) do (
    echo ===^>^>%%t;%%f >> aaa.txt
)
)