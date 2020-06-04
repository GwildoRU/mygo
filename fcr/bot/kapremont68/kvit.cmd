cd out\%1\kvit
del *.pdf
set NLS_LANG=russian_russia.CL8MSWIN1251
..\..\..\MakeKv8.exe %2
exit
