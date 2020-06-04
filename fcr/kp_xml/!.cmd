chcp 65001

for %%Z in (ZIP\*.zip) do (
    unziprr "%%Z"
)

del TMP\*.zip

del CSV\parsed_kp_xmls.csv

for %%Z in (OUT\*.xml) do (
    parse_kp_xml "%%Z" >> CSV\parsed_kp_xmls.csv
)

rem del OUT\*xml






