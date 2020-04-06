chcp 65001
for %%D in (*.dbf) do (
	dbf2csv "%%D" >> "%%D".csv
)	