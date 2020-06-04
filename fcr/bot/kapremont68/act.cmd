cd out\%1\act
del *.pdf
set NLS_LANG=russian_russia.CL8MSWIN1251
..\..\..\MakeACT.exe %2
exit